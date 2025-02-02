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
	"github.com/snowlyg/go-tenancy/service/scope"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/go-tenancy/utils/param"
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
	err := g.TENANCY_DB.Model(&model.TUser{}).
		Select("t_users.id,t_users.username,t_users.authority_id,t_users.created_at,t_users.updated_at,sys_tenancies.id as tenancy_id,sys_tenancies.name as tenancy_name,sys_tenancies.status,t_users.email, t_users.phone, t_users.nick_name, t_users.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,t_users.authority_id").
		Joins("left join sys_tenancies on t_users.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
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

	url, err := param.GetSeitURL()
	if err != nil {
		return loginTenancy, err
	}

	loginTenancy.Url = url + g.TENANCY_CONFIG.System.ClientPreix

	return loginTenancy, nil
}

// CreateTenancy
func CreateTenancy(req request.CreateTenancy) (uint, string, string, error) {
	err := g.TENANCY_DB.Where("name = ?", req.Name).First(&model.SysTenancy{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, "", "", errors.New("商户名称已被注冊")
	}

	err = g.TENANCY_DB.
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
		Where("t_users.username = ?", req.Username).
		Where("sys_authorities.authority_type = ?", multi.TenancyAuthority).
		First(&model.TUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, "", "", errors.New("管理员用户名已注册")
	}

	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		req.SysTenancy.UUID = utils.UUIDV5()
		req.SysTenancy.Status = g.StatusTrue
		req.SysTenancy.State = g.StatusTrue
		err = tx.Model(&model.SysTenancy{}).Create(&req.SysTenancy).Error
		if err != nil {
			return err
		}
		defaultPwd, _ := param.GetTenancyDefaultPassword()
		if defaultPwd == "" {
			defaultPwd = "123456"
		}
		user := model.TUser{Username: req.Username, Password: utils.MD5V([]byte(defaultPwd)), AuthorityId: source.TenancyAuthorityId, Status: g.StatusTrue, IsShow: g.StatusFalse, SysTenancyId: req.SysTenancy.ID, NickName: req.Name}
		err = tx.Create(&user).Error
		if err != nil {
			return err
		}
		return nil
	})

	return req.SysTenancy.ID, req.SysTenancy.UUID.String(), req.Username, err
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
	user, err := GetUserByTenancyId(id)
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).Delete(&model.SysTenancy{}).Error
		if err != nil {
			return err
		}
		if user.ID > 0 {
			err = tx.Where("id = ?", user.ID).Delete(&model.TUser{}).Error
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
	db := g.TENANCY_DB.Model(&model.SysTenancy{}).
		Joins("left join t_users on t_users.sys_tenancy_id = sys_tenancies.id").
		Select("sys_tenancies.*,t_users.username as username").
		Where("sys_tenancies.status = ?", info.Status)
	if info.Keyword != "" {
		db = db.Where(g.TENANCY_DB.Where("sys_tenancies.name like ?", info.Keyword+"%").Or("sys_tenancies.tele like ?", info.Keyword+"%"))
	}

	if info.Date != "" {
		db = db.Scopes(scope.FilterDate(info.Date, "created_at", "sys_tenancies"))
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return tenancyList, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy, "sys_tenancies.")
	err = db.Limit(limit).Offset(offset).Find(&tenancyList).Error
	return tenancyList, total, err
}

// GetTenanciesByRegion
func GetTenanciesByRegion(p_code string) ([]response.SysTenancy, error) {
	tenancyList := []response.SysTenancy{}
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).
		Joins("left join t_users on t_users.sys_tenancy_id = sys_tenancies.id").
		Select("sys_tenancies.*,t_users.username as username").
		Where("sys_tenancies.sys_region_code = ?", p_code).Find(&tenancyList).Error
	return tenancyList, err
}

// GetTenancySelect
func GetTenancySelect() ([]response.SelectOption, error) {
	selects := []response.SelectOption{
		{ID: 0, Name: "请选择"},
	}
	tenancySelects := []response.SelectOption{}
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

// GetTenancyNum 获取商户数量
func GetTenancyNum(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var userNum int64
	db := g.TENANCY_DB.Model(&model.SysTenancy{})
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Count(&userNum).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return userNum, nil
}

// GetTenancyNum 商户收入排行
func GetTenancyOrderPayPriceGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]*response.MerchantRateData, error) {
	var rateData []*response.MerchantRateData
	db := g.TENANCY_DB.Model(&model.SysTenancy{}).
		Select("sum(orders.pay_price) as pay_price,sys_tenancies.name as tenancy_name").
		Joins("left join orders on sys_tenancies.id = orders.sys_tenancy_id")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Where("pay_price > ?", 0).Group("sys_tenancy_id").Order("pay_price desc").Limit(4).Find(&rateData).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return rateData, nil
}

func LoginDevice(loginDevice request.LoginDevice) (*response.LoginResponse, error) {
	tenancy, err := GetTenancyByUUID(loginDevice.UUID)
	if err != nil {
		return nil, fmt.Errorf("find tenancy %w", err)
	}
	if tenancy.Status == g.StatusFalse {
		return nil, fmt.Errorf("商户已被冻结")
	}
	cuserId, err := CreateCUserFromDevice(loginDevice, tenancy.ID)
	if err != nil {
		return nil, err
	}

	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(cuserId), 10), // 患者 id
		Username:      loginDevice.HospitalNO,                  // 用户名使用住院号
		TenancyId:     tenancy.ID,
		TenancyName:   tenancy.Name,
		AuthorityId:   source.DeviceAuthorityId,
		AuthorityType: multi.GeneralAuthority,
		LoginType:     multi.LoginTypeDevice,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}
	token, _, err := multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return nil, err
	}
	// user := map[string]interface{}{
	// 	"tenancy": tenancy,
	// 	"patient": patient,
	// }
	loginResponse := &response.LoginResponse{
		User:  loginDevice.HospitalNO,
		Token: token,
	}
	return loginResponse, nil
}
