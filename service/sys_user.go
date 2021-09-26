package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ChangeProfileMap(ctx *gin.Context) (Form, error) {
	var form Form
	user, err := GetUserByUserIdAndTenancyId(multi.GetUserId(ctx), multi.GetTenancyId(ctx))
	if err != nil {
		return form, err
	}
	form = Form{Method: "POST", Title: "修改信息"}
	form.AddRule(*NewInput("管理员姓名", "nickName", "请输入管理员姓名", user.NickName)).
		AddRule(*NewInput("联系电话", "phone", "请输入联系电话", user.Phone))
	form.SetAction("/user/changeProfile", ctx)
	return form, nil
}

func ChangePasswordMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := fmt.Sprintf(`{"rule":[{"type":"input","field":"NewPassword","value":"","title":"密码","props":{"type":"password","placeholder":"请输入密码"},"validate":[{"message":"请输入密码","required":true,"type":"string","trigger":"change"}]},
	{"type":"hidden","field":"id","value":%d,"title":"id","props":{"type":"hidden","placeholder":""}},{"type":"input","field":"confirmPassword","value":"","title":"确认密码","props":{"type":"password","placeholder":"请输入确认密码"},"validate":[{"message":"请输入确认密码","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"修改密码","config":{}}`, id)

	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction("/user/changePassword", ctx)
	return form, nil
}

func RegisterAdminMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		user, err := FindUserByStringId(fmt.Sprintf("%d", id))
		if err != nil {
			return Form{}, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"select","field":"authorityId","value":["%s"],"title":"身份","props":{"multiple":true,"placeholder":"请选择身份"},"options":[]},{"type":"input","field":"nickName","value":"%s","title":"管理员姓名","props":{"type":"text","placeholder":"请输入管理员姓名"}},{"type":"input","field":"username","value":"%s","title":"账号","props":{"type":"text","placeholder":"请输入账号"},"validate":[{"message":"请输入账号","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"phone","value":"%s","title":" 联系电话","props":{"type":"text","placeholder":"请输入联系电话"}},{"type":"switch","field":"status","value":%d,"title":"是否开启","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"\/sys\/system\/admin\/update\/2.html","method":"PUT","title":"编辑管理员","config":{}}`, user.AuthorityId, user.NickName, user.Username, user.Phone, user.Status)
	} else {
		formStr = `{"rule":[{"type":"select","field":"authorityId","value":[],"title":"身份","props":{"multiple":true,"placeholder":"请选择身份"},"options":[]},{"type":"input","field":"RealName","value":"","title":"管理员姓名","props":{"type":"text","placeholder":"请输入管理员姓名"}},{"type":"input","field":"username","value":"","title":"账号","props":{"type":"text","placeholder":"请输入账号"},"validate":[{"message":"请输入账号","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"phone","value":"","title":" 联系电话","props":{"type":"text","placeholder":"请输入联系电话"}},{"type":"input","field":"password","value":"","title":"密码","props":{"type":"password","placeholder":"请输入密码"},"validate":[{"message":"请输入密码","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"confirmPassword","value":"","title":"确认密码","props":{"type":"password","placeholder":"请输入确认密码"},"validate":[{"message":"请输入确认密码","required":true,"type":"string","trigger":"change"}]},{"type":"switch","field":"status","value":1,"title":"是否开启","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"","method":"POST","title":"添加管理员","config":{}}`
	}

	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("%s/%d", "/user/setUserInfo", id), ctx)
	} else {
		form.SetAction("/user/registerAdmin", ctx)
	}
	opts, err := GetAuthorityOptions(multi.AdminAuthority)
	if err != nil {
		return form, err
	}
	form.Rule[0].Options = opts
	return form, nil
}

// Register 用户注册
func Register(req request.Register, authorityType int) (uint, error) {
	err := g.TENANCY_DB.
		Where("sys_users.username = ?", req.Username).
		Where("sys_authorities.authority_type = ?", authorityType).
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&model.SysUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return 0, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	user := model.SysUser{Username: req.Username, Password: utils.MD5V([]byte(req.Password)), AuthorityId: req.AuthorityId[0], Status: req.Status, IsShow: g.StatusTrue, NickName: req.NickName, Phone: req.Phone}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&user).Error
		if err != nil {
			return err
		}
		return nil
	})

	return user.ID, err
}

// Login 用户登录
func Login(u *model.SysUser, authorityType int) (response.LoginResponse, error) {
	switch {
	case authorityType == multi.AdminAuthority:
		return adminLogin(u)
	case authorityType == multi.TenancyAuthority:
		return tenancyLogin(u)
	case authorityType == multi.GeneralAuthority:
		return response.LoginResponse{
			User:  nil,
			Token: "",
		}, errors.New("错误用户类型")
	default:
		return response.LoginResponse{
			User:  nil,
			Token: "",
		}, errors.New("用户名或者密码错误")
	}
}

// adminLogin
func adminLogin(u *model.SysUser) (response.LoginResponse, error) {
	var admin response.SysAdminUser
	var token string
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", multi.AdminAuthority).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, sys_users.email, sys_users.phone, sys_users.nick_name, sys_users.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, errors.New("用户名或者密码错误")
	} else if err != nil {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, err
	} else if admin.ID == 0 {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, errors.New("用户名或者密码错误")
	} else if admin.Status == g.StatusFalse {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, errors.New("账号已被冻结")
	}

	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(admin.ID), 10),
		Username:      admin.Username,
		AuthorityId:   admin.AuthorityId,
		AuthorityType: admin.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}
	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, err
	}

	return response.LoginResponse{
		User:  admin,
		Token: token,
	}, nil
}

// tenancyLogin
func tenancyLogin(u *model.SysUser) (response.LoginResponse, error) {
	var tenancy response.SysAdminUser
	var token string
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.TUser{}).
		Where("t_users.username = ? AND t_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", multi.TenancyAuthority).
		Select("t_users.id,t_users.username,t_users.authority_id,t_users.created_at,t_users.updated_at,sys_tenancies.id  as tenancy_id,sys_tenancies.name as tenancy_name,t_users.email, t_users.phone, t_users.nick_name, t_users.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,t_users.authority_id").
		Joins("left join sys_tenancies on t_users.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
		First(&tenancy).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}
	if err != nil {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, err
	}

	if tenancy.ID == 0 {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}

	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(tenancy.ID), 10),
		Username:      tenancy.Username,
		TenancyId:     tenancy.TenancyId,
		TenancyName:   tenancy.TenancyName,
		AuthorityId:   tenancy.AuthorityId,
		AuthorityType: tenancy.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, err
	}

	return response.LoginResponse{
		User:  tenancy,
		Token: token,
	}, nil
}

// ChangePassword 修改用户密码
func ChangePassword(id uint, req request.ChangePassword, authorityType int) error {
	if req.Id > 0 {
		id = req.Id
	} else {
		err := g.TENANCY_DB.Model(&model.SysUser{}).
			Where("sys_users.id = ? AND sys_users.password = ?", id, utils.MD5V([]byte(req.Password))).
			Where("sys_authorities.authority_type = ?", authorityType).
			Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
			First(&model.SysUser{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("修改失败，原密码与当前账户不符")
		}
	}
	err := g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", id).Update("password", utils.MD5V([]byte(req.NewPassword))).Error
	if err != nil {
		return err
	}
	return nil
}

// ChangeUserStatus 修改用户密码
func ChangeUserStatus(changeStatus request.ChangeStatus) error {
	err := g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
	if err != nil {
		return err
	}
	return nil
}

// ChangeProfile 修改用户信息
func ChangeProfile(user request.ChangeProfile, sysUserId uint) error {
	err := g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", sysUserId).
		Updates(map[string]interface{}{"nick_name": user.NickName, "phone": user.Phone}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdminInfoList 分页获取数据
func GetAdminInfoList(info request.PageInfo, userId uint) ([]response.SysAdminUser, int64, error) {
	userList := []response.SysAdminUser{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	adminAuthorityIds, err := GetUserAuthorityIds(multi.AdminAuthority)
	if err != nil {
		return userList, total, err
	}
	db := g.TENANCY_DB.Model(&model.SysUser{})

	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = OrderBy(db, info.OrderBy, info.SortBy, "sys_users.")
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("sys_users.*,sys_authorities.authority_name,sys_authorities.authority_type").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Where("sys_users.authority_id IN (?)", adminAuthorityIds).
		Where("sys_users.is_show = ?", g.StatusTrue).
		Not("sys_users.id = ?", userId).
		Find(&userList).Error
	return userList, total, err
}

// GetTenancyByUserIds
func GetTenancyByUserIds(userIds []uint, tenancyId uint) ([]response.SysAdminUser, error) {
	userList := []response.SysAdminUser{}
	tenancyAuthorityIds, err := GetUserAuthorityIds(multi.TenancyAuthority)
	if err != nil {
		return userList, err
	}
	db := g.TENANCY_DB.Model(&model.TUser{}).
		Select("t_users.*,sys_authorities.authority_name,sys_authorities.authority_type,sys_tenancies.name as tenancy_name").
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
		Joins("left join sys_tenancies on t_users.sys_tenancy_id = sys_tenancies.id").
		Where("t_users.authority_id IN (?)", tenancyAuthorityIds).
		Where("t_users.id IN (?)", userIds)
	if tenancyId > 0 {
		db = db.Where("t_users.sys_tenancy_id = ?", tenancyId)
	}
	err = db.Find(&userList).Error
	return userList, err
}

// GetAdminByUserIds
func GetAdminByUserIds(userIds []uint) ([]response.SysAdminUser, error) {
	userList := []response.SysAdminUser{}
	adminAuthorityIds, err := GetUserAuthorityIds(multi.AdminAuthority)
	if err != nil {
		return userList, err
	}
	err = g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.*,sys_authorities.authority_name,sys_authorities.authority_type").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Where("sys_users.authority_id IN (?)", adminAuthorityIds).
		Where("sys_users.id IN (?)", userIds).
		Find(&userList).Error
	return userList, err
}

// SetUserAuthority  设置一个用户的权限
func SetUserAuthority(id uint, authorityId string) error {
	return g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
}

// DeleteUser 删除用户
func DeleteUser(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.SysUser{}).Error
}

// UpdateAdminInfo 设置关联信息
func UpdateAdminInfo(userInfo request.UpdateUser, user model.SysUser) error {
	err := g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", user.ID).Updates(map[string]interface{}{"authority_id": userInfo.AuthorityId[0], "username": userInfo.Username, "status": userInfo.Status, "nick_name": userInfo.NickName, "phone": userInfo.Phone}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByTenancyId(tenanacyId uint) (model.SysUser, error) {
	var u model.SysUser
	tenancyAuthorityIds, err := GetUserAuthorityIds(multi.TenancyAuthority)
	if err != nil {
		return u, err
	}
	err = g.TENANCY_DB.Model(&model.TUser{}).
		Select("t_users.*").
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
		Where("t_users.authority_id IN (?)", tenancyAuthorityIds).
		Where("t_users.sys_tenancy_id = ?", tenanacyId).
		Preload("Authority").
		First(&u).Error
	return u, err
}

func GetUserByUserIdAndTenancyId(userId, tenanacyId uint) (model.SysUser, error) {
	var u model.SysUser
	tenancyAuthorityIds, err := GetUserAuthorityIds(multi.TenancyAuthority)
	if err != nil {
		return u, err
	}
	err = g.TENANCY_DB.Model(&model.TUser{}).
		Select("t_users.*").
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
		Where("t_users.authority_id IN (?)", tenancyAuthorityIds).
		Where("t_users.sys_tenancy_id = ?", tenanacyId).
		Where("t_users.id = ?", userId).
		Preload("Authority").
		First(&u).Error
	return u, err
}

// FindUserByStringId 通过id获取用户信息
func FindUserByStringId(id string) (model.SysUser, error) {
	var u model.SysUser
	err := g.TENANCY_DB.Where("`id` = ?", id).Preload("Authority").First(&u).Error
	return u, err
}

// DelToken 删除token
func DelToken(token string) error {
	err := multi.AuthDriver.DelUserTokenCache(token)
	if err != nil {
		g.TENANCY_LOG.Error("del token", zap.Any("err", err))
		return fmt.Errorf("del token %w", err)
	}
	return nil
}

// CleanToken 清空 token
func CleanToken(authorityType int, userId string) error {
	err := multi.AuthDriver.CleanUserTokenCache(authorityType, userId)
	if err != nil {
		g.TENANCY_LOG.Error("clean token", zap.Any("err", err))
		return fmt.Errorf("clean token %w", err)
	}
	return nil
}

// GetTenancyInfoList 分页获取数据
func GetTenancyInfoList(info request.PageInfo, userId, tenancyId uint) ([]response.SysAdminUser, int64, error) {
	userList := []response.SysAdminUser{}
	var tenancyAuthorityIds []int
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.TenancyAuthority).Select("authority_id").Find(&tenancyAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	db := g.TENANCY_DB.Model(&model.TUser{})
	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = OrderBy(db, info.OrderBy, info.SortBy, "t_users.")
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("t_users.id,t_users.status,t_users.username,t_users.authority_id,t_users.created_at,t_users.updated_at, t_users.email, t_users.phone, t_users.nick_name, t_users.header_img,sys_authorities.authority_name,sys_authorities.authority_type,t_users.authority_id,sys_tenancies.name as tenancy_name").
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
		Joins("left join sys_tenancies on t_users.sys_tenancy_id = sys_tenancies.id").
		Where("t_users.authority_id IN (?)", tenancyAuthorityIds).
		Where("t_users.sys_tenancy_id = ?", tenancyId).
		Where("t_users.is_show = ?", g.StatusTrue).
		Not("t_users.id = ?", userId).
		Find(&userList).Error
	return userList, total, err
}
