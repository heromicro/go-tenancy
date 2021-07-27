package service

import (
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CreateAuthority 创建一个角色
func CreateAuthority(auth model.SysAuthority) (model.SysAuthority, error) {
	var authorityBox model.SysAuthority
	err := g.TENANCY_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, errors.New("存在相同角色id")
	}
	err = g.TENANCY_DB.Create(&auth).Error
	return auth, err
}

// CopyAuthority 复制一个角色
func CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (model.SysAuthority, error) {
	var authorityBox model.SysAuthority
	err := g.TENANCY_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return authorityBox, errors.New("存在相同角色id")
	}
	copyInfo.Authority.Children = []model.SysAuthority{}
	menus, err := GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return copyInfo.Authority, err
	}
	var baseMenu []model.SysBaseMenu
	for _, v := range menus {
		v.SysBaseMenu.ID = v.MenuId
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = g.TENANCY_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return copyInfo.Authority, err
	}

	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		var authority request.DeleteAuthority
		authority.AuthorityId = copyInfo.Authority.AuthorityId
		_ = DeleteAuthority(&authority)
	}
	return copyInfo.Authority, err
}

// UpdateAuthority 更改一个角色
func UpdateAuthority(auth model.SysAuthority) (model.SysAuthority, error) {
	err := g.TENANCY_DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Updates(&auth).Error
	return auth, err
}

// DeleteAuthority 删除角色
func DeleteAuthority(request *request.DeleteAuthority) error {
	err := g.TENANCY_DB.Where("authority_id = ?", request.AuthorityId).First(&model.SysUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	err = g.TENANCY_DB.Where("parent_id = ?", request.AuthorityId).First(&model.SysAuthority{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	var auth model.SysAuthority
	err = g.TENANCY_DB.Preload("SysBaseMenus").Where("authority_id = ?", request.AuthorityId).First(&auth).Error
	if err != nil {
		return fmt.Errorf("fond authority  %w", err)
	}
	err = g.TENANCY_DB.Unscoped().Delete(auth).Error
	if err != nil {
		return fmt.Errorf("delete authority %w", err)
	}

	if len(auth.SysBaseMenus) > 0 {
		err = g.TENANCY_DB.Model(&model.SysAuthority{}).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			g.TENANCY_LOG.Error("association delete sys_base_menus ", zap.Any("err", err))
		}
		//err = db.Association("SysBaseMenus").DELETE(&auth)
	}
	ClearCasbin(0, auth.AuthorityId)
	return nil
}

// GetAuthorityInfoList 分页获取数据
func GetAuthorityInfoList(info request.PageInfo, authorityType int) ([]model.SysAuthority, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("parent_id = 0")
	if authorityType > 0 {
		db = db.Where("authority_type = ?", authorityType)
	}
	var total int64
	db.Count(&total)
	var authority []model.SysAuthority
	err := db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return authority, total, err
}

// GetAuthorityInfo 获取角色信息
func GetAuthorityInfo(auth model.SysAuthority) (model.SysAuthority, error) {
	var sa model.SysAuthority
	err := g.TENANCY_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

// SetDataAuthority 设置角色资源权限
func SetDataAuthority(auth request.SetDataAuthority) error {
	var s model.SysAuthority
	g.TENANCY_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := g.TENANCY_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

// SetMenuAuthority 菜单与角色绑定
func SetMenuAuthority(auth *model.SysAuthority) error {
	var s model.SysAuthority
	g.TENANCY_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := g.TENANCY_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// findChildrenAuthority 查询子角色
func findChildrenAuthority(authority *model.SysAuthority) error {
	err := g.TENANCY_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

func GetUserAuthorityIds() ([]int, error) {
	var generalAuthorityIds []int
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.GeneralAuthority).Select("authority_id").Find(&generalAuthorityIds).Error
	if err != nil {
		return generalAuthorityIds, fmt.Errorf("get authority ids %w", err)
	}
	return generalAuthorityIds, nil
}

// getAuthorityMap
func getAuthorityMap(authorityType int) (map[string][]model.SysAuthority, error) {
	var authority []model.SysAuthority
	treeMap := make(map[string][]model.SysAuthority)
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", authorityType).Find(&authority).Error
	for _, v := range authority {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// GetAuthorityOptions
func GetAuthorityOptions(authorityType int) ([]Option, error) {
	var options []Option
	options = append(options, Option{Label: "请选择", Value: 0})
	treeMap, err := getAuthorityMap(authorityType)

	for _, opt := range treeMap["0"] {
		options = append(options, Option{Label: opt.AuthorityName, Value: opt.AuthorityId})
	}
	for i := 0; i < len(options); i++ {
		getAuthorityOption(&options[i], treeMap)
	}

	return options, err
}

// getAuthorityOption
func getAuthorityOption(op *Option, treeMap map[string][]model.SysAuthority) {
	id, ok := op.Value.(string)
	if ok {
		for _, opt := range treeMap[id] {
			op.Children = append(op.Children, Option{Label: opt.AuthorityName, Value: opt.AuthorityId})
		}
		for i := 0; i < len(op.Children); i++ {
			getAuthorityOption(&op.Children[i], treeMap)
		}
	}
}
