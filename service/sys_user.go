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
		formStr = fmt.Sprintf(`{"rule":[{"type":"select","field":"authorityId","value":["%s"],"title":"身份","props":{"multiple":true,"placeholder":"请选择身份"},"options":[]},{"type":"input","field":"nickName","value":"%s","title":"管理员姓名","props":{"type":"text","placeholder":"请输入管理员姓名"}},{"type":"input","field":"username","value":"%s","title":"账号","props":{"type":"text","placeholder":"请输入账号"},"validate":[{"message":"请输入账号","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"phone","value":"%s","title":" 联系电话","props":{"type":"text","placeholder":"请输入联系电话"}},{"type":"switch","field":"status","value":%d,"title":"是否开启","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"\/sys\/system\/admin\/update\/2.html","method":"PUT","title":"编辑管理员","config":{}}`, user.AuthorityId, user.AdminInfo.NickName, user.Username, user.AdminInfo.Phone, user.Status)
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
	authorityType := multi.NoneAuthority
	if multi.IsAdmin(ctx) {
		authorityType = multi.AdminAuthority
	} else if multi.IsTenancy(ctx) {
		authorityType = multi.TenancyAuthority
	}
	opts, err := GetAuthorityOptions(authorityType)
	if err != nil {
		return form, err
	}
	form.Rule[0].Options = opts
	return form, nil
}

// Register 用户注册
func Register(req request.Register, authorityType int, tenancyId uint) (uint, error) {
	if !errors.Is(g.TENANCY_DB.
		Where("sys_users.username = ?", req.Username).
		Where("sys_authorities.authority_type = ?", authorityType).
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return 0, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	user := model.SysUser{Username: req.Username, Password: utils.MD5V([]byte(req.Password)), AuthorityId: req.AuthorityId[0], Status: req.Status, IsShow: g.StatusTrue, SysTenancyID: tenancyId}
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&user).Error
		if err != nil {
			return err
		}
		adminInfo := model.AdminInfo{NickName: req.NickName, Phone: req.Phone, SysUserID: user.ID}
		err = tx.Create(&adminInfo).Error
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
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
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
	var tenancy response.SysTenancyUser
	var token string
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", multi.TenancyAuthority).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at,sys_tenancies.id  as tenancy_id,sys_tenancies.name as tenancy_name,admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_tenancies on sys_users.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
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
	err := g.TENANCY_DB.Model(&model.AdminInfo{}).Where("sys_user_id = ?", sysUserId).
		Updates(map[string]interface{}{"nick_name": user.NickName, "phone": user.Phone}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdminInfoList 分页获取数据
func GetAdminInfoList(info request.PageInfo, userId uint) ([]response.SysAdminUser, int64, error) {
	var userList []response.SysAdminUser
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
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("sys_users.id,sys_users.status,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Where("sys_users.authority_id IN (?)", adminAuthorityIds).
		Where("sys_users.is_show = ?", g.StatusTrue).
		Not("sys_users.id = ?", userId).
		Find(&userList).Error
	return userList, total, err
}

// GetTenancyByUserIds
func GetTenancyByUserIds(userIds []uint) ([]response.SysTenancyUser, error) {
	var userList []response.SysTenancyUser
	tenancyAuthorityIds, err := GetUserAuthorityIds(multi.TenancyAuthority)
	if err != nil {
		return userList, err
	}
	err = g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.id,sys_users.status,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id,sys_tenancies.name as tenancy_name").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Joins("left join sys_tenancies on sys_users.sys_tenancy_id = sys_tenancies.id").
		Where("sys_users.authority_id IN (?)", tenancyAuthorityIds).
		Where("sys_users.id IN (?)", userIds).
		Find(&userList).Error
	return userList, err
}

// GetAdminByUserIds
func GetAdminByUserIds(userIds []uint) ([]response.SysAdminUser, error) {
	var userList []response.SysAdminUser
	adminAuthorityIds, err := GetUserAuthorityIds(multi.AdminAuthority)
	if err != nil {
		return userList, err
	}
	err = g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.id,sys_users.status,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
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
	user, err := FindUserByStringId(fmt.Sprintf("%d", id))
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id = ?", id).Delete(&model.SysUser{}).Error
		if err != nil {
			return err
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

// UpdateAdminInfo 设置关联信息
func UpdateAdminInfo(userInfo request.UpdateUser, user model.SysUser, tenancyId uint) error {
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&model.SysUser{}).Where("id = ?", user.ID).Updates(map[string]interface{}{"authority_id": userInfo.AuthorityId[0], "username": userInfo.Username, "status": userInfo.Status})

		info := map[string]interface{}{"nick_name": userInfo.NickName, "phone": userInfo.Phone}
		if user.IsAdmin() || user.IsTenancy() {
			if user.AdminInfo.ID > 0 {
				err := tx.Model(&model.AdminInfo{}).Where("id = ?", user.AdminInfo.ID).Updates(info).Error
				if err != nil {
					return err
				}
			} else {
				info["sys_user_id"] = user.ID
				info["sys_tenancy_id"] = tenancyId
				err := tx.Model(&model.AdminInfo{}).Create(info).Error
				if err != nil {
					return err
				}
			}

		} else {
			g.TENANCY_LOG.Error("角色错误", zap.Any("err", user.AuthorityType()))
			return fmt.Errorf("角色错误")
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserByTenancyId(tenanacyId uint) (model.SysUser, error) {
	var u model.SysUser
	adminAuthorityIds, err := GetUserAuthorityIds(multi.TenancyAuthority)
	if err != nil {
		return u, err
	}
	err = g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.*").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Where("sys_users.authority_id IN (?)", adminAuthorityIds).
		Where("sys_users.sys_tenancy_id = ?", tenanacyId).
		First(&u).Error
	return u, err
}

// FindUserByStringId 通过id获取用户信息
func FindUserByStringId(id string) (model.SysUser, error) {
	var u model.SysUser
	err := g.TENANCY_DB.Where("`id` = ?", id).Preload("Authority").Preload("AdminInfo").Preload("GeneralInfo").First(&u).Error
	return u, err
}

// FindUserByTenancyId 通过tenancy_id获取用户信息
func FindUserByTenancyId(tenancyId uint) (model.SysUser, error) {
	var u model.SysUser
	err := g.TENANCY_DB.Where("sys_tenancy_id = ?", tenancyId).Preload("Authority").Preload("AdminInfo").Preload("GeneralInfo").First(&u).Error
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
func CleanToken(userId string) error {
	err := multi.AuthDriver.CleanUserTokenCache(userId)
	if err != nil {
		g.TENANCY_LOG.Error("clean token", zap.Any("err", err))
		return fmt.Errorf("clean token %w", err)
	}
	return nil
}

// GetTenancyInfoList 分页获取数据
func GetTenancyInfoList(info request.PageInfo, userId, tenancyId uint) ([]response.SysTenancyUser, int64, error) {
	var userList []response.SysTenancyUser
	var tenancyAuthorityIds []int
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.TenancyAuthority).Select("authority_id").Find(&tenancyAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	db := g.TENANCY_DB.Model(&model.SysUser{})
	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("sys_users.id,sys_users.status,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id,sys_tenancies.name as tenancy_name").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Joins("left join sys_tenancies on sys_users.sys_tenancy_id = sys_tenancies.id").
		Where("sys_users.authority_id IN (?)", tenancyAuthorityIds).
		Where("sys_users.sys_tenancy_id = ?", tenancyId).
		Where("sys_users.is_show = ?", g.StatusTrue).
		Not("sys_users.id = ?", userId).
		Find(&userList).Error
	return userList, total, err
}
