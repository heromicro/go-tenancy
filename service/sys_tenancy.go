package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

func ChangeTenancyPasswordMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	user, err := GetUserByTenancyId(id)
	if err != nil {
		return form, err
	}
	formStr := fmt.Sprintf(`{"rule":[{"type":"input","field":"NewPassword","value":"","title":"密码","props":{"type":"password","placeholder":"请输入密码"},"validate":[{"message":"请输入密码","required":true,"type":"string","trigger":"change"}]},
	{"type":"hidden","field":"id","value":%d,"title":"id","props":{"type":"hidden","placeholder":""}},{"type":"input","field":"confirmPassword","value":"","title":"确认密码","props":{"type":"password","placeholder":"请输入确认密码"},"validate":[{"message":"请输入确认密码","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"修改密码","config":{}}`, user.ID)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction("/user/changePassword", ctx)
	return form, nil
}

func LoginTenancy(id uint) (response.LoginTenancy, error) {
	var loginTenancy response.LoginTenancy
	var token string
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at,sys_tenancies.id as tenancy_id,sys_tenancies.name as tenancy_name,sys_tenancies.status,tenancy_infos.email, tenancy_infos.phone, tenancy_infos.nick_name, tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join tenancy_infos on tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_tenancies on sys_users.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Where("sys_authorities.authority_type = ?", multi.TenancyAuthority).
		Where("sys_tenancies.id = ?", id).
		First(&loginTenancy.Admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return loginTenancy, errors.New("用户名或者密码错误")
	}
	if err != nil {
		return loginTenancy, err
	}
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(loginTenancy.Admin.ID), 10),
		Username:      loginTenancy.Admin.Username,
		TenancyId:     loginTenancy.Admin.TenancyId,
		TenancyName:   loginTenancy.Admin.TenancyName,
		AuthorityId:   loginTenancy.Admin.AuthorityId,
		AuthorityType: loginTenancy.Admin.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	if loginTenancy.Admin.ID == 0 {
		return loginTenancy, errors.New("用户名或者密码错误")
	}

	if loginTenancy.Admin.Status == g.StatusFalse {
		return loginTenancy, errors.New("商户已被冻结")
	}

	token, exp, err := multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return loginTenancy, err
	}
	loginTenancy.Token = token
	loginTenancy.Exp = exp
	loginTenancy.Url = g.TENANCY_CONFIG.System.ClientURL + g.TENANCY_CONFIG.System.ClientPreix

	return loginTenancy, nil
}

// CreateTenancy
func CreateTenancy(req request.CreateTenancy) (uint, error) {
	err := g.TENANCY_DB.Where("name = ?", req.Name).First(&model.SysTenancy{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("商户名称已被注冊")
	}
	err = g.TENANCY_DB.
		Where("sys_users.username = ?", req.Username).
		Where("sys_authorities.authority_type = ?", multi.TenancyAuthority).
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&model.SysUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return 0, errors.New("管理员用户名已注册")
	}

	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		req.SysTenancy.UUID = uuid.NewV4()
		req.SysTenancy.Status = g.StatusTrue
		req.SysTenancy.State = g.StatusTrue
		err = tx.Model(&model.SysTenancy{}).Create(&req.SysTenancy).Error
		if err != nil {
			return err
		}
		user := model.SysUser{Username: req.Username, Password: utils.MD5V([]byte("123456")), AuthorityId: source.TenancyAuthorityId, Status: g.StatusTrue, SysTenancyID: req.SysTenancy.ID}
		err = tx.Create(&user).Error
		if err != nil {
			return err
		}
		tenancyInfo := model.AdminInfo{NickName: req.Name, SysUserID: user.ID}
		err = tx.Create(&tenancyInfo).Error
		if err != nil {
			return err
		}
		return nil
	})

	return req.SysTenancy.ID, err
}

// GetTenancyByID
func GetTenancyByID(id uint) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("id = ?", id).First(&tenancy).Error
	return tenancy, err
}

// GetTenancyByUUID
func GetTenancyByUUID(uuid string) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("uuid = ?", uuid).First(&tenancy).Error
	return tenancy, err
}

// SetTenancyRegionByID
func SetTenancyRegionByID(regionCode request.SetRegionCode) error {
	return g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", regionCode.Id).Update("sys_region_code", regionCode.SysRegionCode).Error
}

// ChangeTenancyStatus
func ChangeTenancyStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateTenancy
func UpdateTenancy(tenancy model.SysTenancy, id uint) (model.SysTenancy, error) {
	err := g.TENANCY_DB.Where("name = ?", tenancy.Name).Not("id = ?", id).First(&tenancy).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tenancy, errors.New("名称已被注冊")
	}

	err = g.TENANCY_DB.Where("id = ?", id).Omit("uuid").Updates(&tenancy).Error
	return tenancy, err
}

// UpdateClientTenancy
func UpdateClientTenancy(req request.UpdateClientTenancy, id uint) error {
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).
		Where("id = ?", id).Omit("uuid").
		Updates(map[string]interface{}{"avatar": req.Avatar, "banner": req.Banner, "info": req.Info, "tele": req.Tele, "state": req.State}).Error
	return err
}

// DeleteTenancy
func DeleteTenancy(id uint) error {
	user, err := FindUserByTenancyId(id)
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).Delete(&model.SysTenancy{}).Error
		if err != nil {
			return err
		}
		if user.ID > 0 {
			err = tx.Where("id = ?", user.ID).Delete(&model.SysUser{}).Error
			if err != nil {
				return err
			}
		}
		if user.AdminInfo.ID > 0 {
			err = tx.Where("id = ?", user.AdminInfo.ID).Delete(&model.AdminInfo{}).Error
			if err != nil {
				return err
			}
		}

		if user.GeneralInfo.ID > 0 {
			err = tx.Where("id = ?", user.GeneralInfo.ID).Delete(&model.GeneralInfo{}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

// GetTenanciesInfoList
func GetTenanciesInfoList(info request.TenancyPageInfo) ([]response.SysTenancy, int64, error) {
	var tenancyList []response.SysTenancy
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("status = ?", info.Status)
	if info.Keyword != "" {
		db = db.Where(g.TENANCY_DB.Where("name like ?", info.Keyword+"%").Or("tele like ?", info.Keyword+"%"))
	}

	if info.Date != "" {
		db = filterDate(db, info.Date, "")
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return tenancyList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&tenancyList).Error
	return tenancyList, total, err
}

// GetTenanciesByRegion
func GetTenanciesByRegion(p_code string) ([]response.SysTenancy, error) {
	var tenancyList []response.SysTenancy
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("sys_region_code = ?", p_code).Find(&tenancyList).Error
	return tenancyList, err
}

// GetTenancySelect
func GetTenancySelect() ([]response.TenancySelect, error) {
	selects := []response.TenancySelect{
		{ID: 0, Name: "请选择"},
	}
	var tenancySelects []response.TenancySelect
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Select("id,name").Where("status = ?", g.StatusTrue).Where("state = ?", g.StatusTrue).Find(&tenancySelects).Error
	selects = append(selects, tenancySelects...)
	return selects, err
}

type Result struct {
	ID   int
	Name string
	Age  int
}

func ChangeCopyMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	tenancy, err := GetTenancyByID(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"copyNum","value":%d,"title":"复制次数","props":{"type":"text","placeholder":"请输入复制次数","disabled":true,"readonly":true}},{"type":"radio","field":"type","value":1,"title":"修改类型","props":{},"options":[{"value":1,"label":"增加"},{"value":2,"label":"减少"}]},{"type":"inputNumber","field":"num","value":0,"title":"修改数量","props":{"placeholder":"请输入修改数量"},"validate":[{"message":"请输入修改数量","required":true,"type":"number","trigger":"change"}]}],"action":"","method":"POST","title":"修改复制商品次数","config":{}}`, tenancy.CopyProductNum)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/tenancy/setCopyProductNum", id), ctx)
	return form, err
}

// GetTenancyCount
func GetTenancyCount() (gin.H, error) {
	var counts response.Counts
	err := g.TENANCY_DB.Raw("SELECT sum(case when status = ? then 1 else 0 end) as 'valid',sum(case when status = ? then 1 else 0 end) as 'invalid' FROM sys_tenancies WHERE ISNULL(deleted_at)", g.StatusTrue, g.StatusFalse).Scan(&counts).Error
	return gin.H{
		"invalid": counts.Invalid,
		"valid":   counts.Valid,
	}, err
}

// GetTenancyInfo
func GetTenancyInfo(tenancyId uint) (response.TenancyInfo, error) {
	var info response.TenancyInfo
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", tenancyId).Find(&info).Error
	return info, err
}

// GetTenancyCopyCount
func GetTenancyCopyCount(tenancyId uint) (int64, error) {
	var copyProductNum int64
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", tenancyId).Select("copy_product_num").Find(&copyProductNum).Error
	return copyProductNum, err
}

// GetUpdateTenancyMap
func GetUpdateTenancyMap(ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	id := multi.GetTenancyId(ctx)
	tenancy, err := GetTenancyByID(id)
	if err != nil {
		return form, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"info","value":"%s","title":"店铺简介","props":{"type":"textarea","placeholder":"请输入店铺简介"},"validate":[{"message":"请输入店铺简介","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"tele","value":"%s","title":"服务电话","props":{"type":"text","placeholder":"请输入服务电话"},"validate":[{"message":"请输入服务电话","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"banner","value":"%s","title":"店铺Banner(710*200px)","props":{"type":"image","maxLength":1,"title":"请选择店铺Banner(710*200px)","src":"\/merchant\/setting\/uploadPicture?field=banner&type=1","modal":{"modal":false},"width":"896px","height":"480px","footer":false}},{"type":"frame","field":"avatar","value":"%s","title":"店铺头像(120*120px)","props":{"type":"image","maxLength":1,"title":"请选择店铺头像(120*120px)","src":"\/merchant\/setting\/uploadPicture?field=avatar&type=1","modal":{"modal":false},"width":"896px","height":"480px","footer":false}},{"type":"switch","field":"state","value":%d,"title":"是否开启","col":{"span":12},"props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"","method":"PUT","title":"编辑店铺信息","config":{}}`, tenancy.Info, tenancy.Tele, tenancy.Banner, tenancy.Avatar, tenancy.State)
	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("/tenancy/updateTenancy/%d", id), ctx)
	return form, err
}

// SetCopyProductNum
func SetCopyProductNum(req request.SetCopyProductNum, id uint) error {
	tenancy, err := GetTenancyByID(id)
	if err != nil {
		return err
	}
	copyNum := tenancy.CopyProductNum
	// 增加
	if req.Type == 1 {
		copyNum = copyNum + req.Num
	} else if req.Type == 2 {
		fmt.Println(copyNum, req.Num)
		if copyNum <= req.Num {
			copyNum = 0
		} else {
			copyNum = copyNum - req.Num
		}
	}
	if err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", id).Updates(map[string]interface{}{"copy_product_num": copyNum}).Error; err != nil {
		return err
	}
	return err
}

func LoginDevice(loginDevice request.LoginDevice) (*response.LoginResponse, error) {

	tenancy, err := GetTenancyByUUID(loginDevice.UUID)
	if err != nil {
		return nil, fmt.Errorf("find tenancy %w", err)
	}
	if tenancy.Status == g.StatusFalse {
		return nil, fmt.Errorf("商户已被冻结")
	}

	loginDevice.Patient.SysTenancyID = tenancy.ID
	patient, err := FindOrCreatePatient(loginDevice.Patient)
	if err != nil {
		return nil, err
	}

	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(patient.ID), 10), // 患者 id
		Username:      loginDevice.HospitalNO,
		TenancyId:     tenancy.ID,
		TenancyName:   tenancy.Name,
		AuthorityId:   source.DeviceAuthorityId,
		AuthorityType: model.DeviceAuthority,
		LoginType:     multi.LoginTypeDevice,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}
	token, _, err := multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return nil, err
	}
	user := map[string]interface{}{
		"tenancy": tenancy,
		"patient": patient,
	}
	loginResponse := &response.LoginResponse{
		User:  user,
		Token: token,
	}
	return loginResponse, nil
}
