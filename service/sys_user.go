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

func RegisterAdminMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		user, err := FindUserById(fmt.Sprintf("%d", id))
		if err != nil {
			return Form{}, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"select","field":"authorityId","value":["%s"],"title":"身份","props":{"multiple":true,"placeholder":"请选择身份"},"options":[]},{"type":"input","field":"nickName","value":"%s","title":"管理员姓名","props":{"type":"text","placeholder":"请输入管理员姓名"}},{"type":"input","field":"username","value":"%s","title":"账号","props":{"type":"text","placeholder":"请输入账号"},"validate":[{"message":"请输入账号","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"phone","value":"%s","title":" 联系电话","props":{"type":"text","placeholder":"请输入联系电话"}},{"type":"switch","field":"status","value":%d,"title":"是否开启","props":{"activeValue":1,"inactiveValue":0,"inactiveText":"关闭","activeText":"开启"}}],"action":"\/sys\/system\/admin\/update\/2.html","method":"PUT","title":"编辑管理员","config":{}}`, user.AuthorityId, user.AdminInfo.NickName, user.Username, user.AdminInfo.Phone, user.Status)
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
	if !errors.Is(g.TENANCY_DB.
		Where("sys_users.username = ?", req.Username).
		Where("sys_authorities.authority_type = ?", authorityType).
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return 0, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	user := model.SysUser{Username: req.Username, Password: utils.MD5V([]byte(req.Password)), AuthorityId: req.AuthorityId[0], Status: req.Status}
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&user).Error
		if err != nil {
			return err
		}
		adminInfo := model.AdminInfo{BaseUserInfo: model.BaseUserInfo{NickName: req.NickName, Phone: req.Phone, SysUserID: user.ID}}
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
	}
	if err != nil {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, err
	}

	if admin.ID == 0 {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, errors.New("用户名或者密码错误")
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
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at,sys_tenancies.id  as tenancy_id,sys_tenancies.name as tenancy_name,tenancy_infos.email, tenancy_infos.phone, tenancy_infos.nick_name, tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join tenancy_infos on tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_tenancies on tenancy_infos.sys_tenancy_id = sys_tenancies.id").
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
func ChangePassword(u *model.SysUser, newPassword string, authorityType int) error {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", authorityType).
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("修改失败，原密码与当前账户不符")
	}
	if user.ID == 0 {
		return errors.New("修改失败，原密码与当前账户不符")
	}
	err = g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", user.ID).Update("password", utils.MD5V([]byte(newPassword))).Error
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
func GetAdminInfoList(info request.PageInfo) ([]response.SysAdminUser, int64, error) {
	var userList []response.SysAdminUser
	var adminAuthorityIds []int
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.AdminAuthority).Select("authority_id").Find(&adminAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("sys_users.authority_id IN (?)", adminAuthorityIds)
	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Find(&userList).Error
	return userList, total, err
}

// GetTenancyInfoList 分页获取数据
func GetTenancyInfoList(info request.PageInfo) ([]response.SysTenancyUser, int64, error) {
	var userList []response.SysTenancyUser
	var tenancyAuthorityIds []int
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.TenancyAuthority).Select("authority_id").Find(&tenancyAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("sys_users.authority_id IN (?)", tenancyAuthorityIds)
	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, tenancy_infos.email, tenancy_infos.phone, tenancy_infos.nick_name, tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id,sys_tenancies.name as tenancy_name").
		Joins("left join tenancy_infos on tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Joins("left join sys_tenancies on tenancy_infos.sys_tenancy_id = sys_tenancies.id").
		Find(&userList).Error
	return userList, total, err
}

// SetUserAuthority  设置一个用户的权限
func SetUserAuthority(id uint, authorityId string) error {
	return g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
}

// DeleteUser 删除用户
func DeleteUser(id uint) (err error) {
	var user model.SysUser
	return g.TENANCY_DB.Where("id = ?", id).Delete(&user).Error
}

// UpdateAdminInfo 设置关联信息
func UpdateAdminInfo(userInfo request.UpdateUser, user model.SysUser, tenancyId uint) error {
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&model.SysUser{}).Where("id = ?", user.ID).Updates(map[string]interface{}{"authority_id": userInfo.AuthorityId[0], "username": userInfo.Username, "status": userInfo.Status})

		info := map[string]interface{}{"nick_name": userInfo.NickName, "phone": userInfo.Phone}
		if user.IsAdmin() {
			if user.AdminInfo.ID > 0 {
				err := tx.Model(&model.AdminInfo{}).Where("id = ?", user.AdminInfo.ID).Updates(info).Error
				if err != nil {
					return err
				}
			} else {
				info["sys_user_id"] = user.ID
				err := tx.Model(&model.AdminInfo{}).Create(info).Error
				if err != nil {
					return err
				}
			}

		} else if user.IsTenancy() {
			if user.TenancyInfo.ID > 0 {
				err := tx.Model(&model.TenancyInfo{}).Where("id = ?", user.TenancyInfo.ID).Updates(info).Error
				if err != nil {
					return err
				}
			} else {
				info["sys_user_id"] = user.ID
				info["sys_tenancy_id"] = tenancyId
				err := tx.Model(&model.TenancyInfo{}).Create(info).Error
				if err != nil {
					return err
				}
			}
		} else {
			g.TENANCY_LOG.Error("未知角色", zap.Any("err", user.AuthorityType()))
			return fmt.Errorf("未知角色")
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// SetUserGeneralInfo 设置普通用户信息
func SetUserGeneralInfo(reqUser model.GeneralInfo, infoId uint, userId string) (model.GeneralInfo, error) {
	if infoId > 0 {
		reqUser.ID = infoId
		err := g.TENANCY_DB.Updates(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	} else {
		id, err := strconv.Atoi(userId)
		if err != nil {
			return reqUser, err
		}
		reqUser.SysUserID = uint(id)
		err = g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// FindUserById 通过id获取用户信息
func FindUserById(id string) (model.SysUser, error) {
	var u model.SysUser
	err := g.TENANCY_DB.Where("`id` = ?", id).Preload("Authority").Preload("AdminInfo").Preload("TenancyInfo").Preload("GeneralInfo").First(&u).Error
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
