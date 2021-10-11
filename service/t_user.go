package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// ChangeClientPassword 修改用户密码
func ChangeClientPassword(id uint, req request.ChangePassword, authorityType int) error {
	if req.Id > 0 {
		id = req.Id
	} else {
		err := g.TENANCY_DB.Model(&model.TUser{}).
			Where("t_users.id = ? AND t_users.password = ?", id, utils.MD5V([]byte(req.Password))).
			Where("sys_authorities.authority_type = ?", authorityType).
			Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
			First(&model.TUser{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("修改失败，原密码与当前账户不符")
		}
	}
	err := g.TENANCY_DB.Model(&model.TUser{}).Where("id = ?", id).Update("password", utils.MD5V([]byte(req.NewPassword))).Error
	if err != nil {
		return err
	}
	return nil
}

// ChangeClientProfile 修改用户信息
func ChangeClientProfile(user request.ChangeProfile, sysUserId uint) error {
	err := g.TENANCY_DB.Model(&model.TUser{}).Where("id = ?", sysUserId).
		Updates(map[string]interface{}{"nick_name": user.NickName, "phone": user.Phone}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteClientUser 删除用户
func DeleteClientUser(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.TUser{}).Error
}

// UpdateClientInfo 设置关联信息
func UpdateClientInfo(userInfo request.UpdateUser, user model.TUser) error {
	err := g.TENANCY_DB.Model(&model.TUser{}).Where("id = ?", user.ID).Updates(map[string]interface{}{"authority_id": userInfo.AuthorityId[0], "username": userInfo.Username, "status": userInfo.Status, "nick_name": userInfo.NickName, "phone": userInfo.Phone}).Error
	if err != nil {
		return err
	}
	return nil
}

// FindClientByStringId 通过id获取用户信息
func FindClientByStringId(id string) (model.TUser, error) {
	var u model.TUser
	err := g.TENANCY_DB.Where("`id` = ?", id).Preload("Authority").First(&u).Error
	return u, err
}

// RegisterClient 用户注册
func RegisterClient(req request.Register, authorityType int, tenancyId uint) (uint, error) {
	err := g.TENANCY_DB.
		Where("t_users.username = ?", req.Username).
		Where("sys_authorities.authority_type = ?", authorityType).
		Joins("left join sys_authorities on sys_authorities.authority_id = t_users.authority_id").
		First(&model.TUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return 0, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	user := model.TUser{Username: req.Username, Password: utils.MD5V([]byte(req.Password)), AuthorityId: req.AuthorityId[0], Status: req.Status, IsShow: g.StatusTrue, NickName: req.NickName, Phone: req.Phone, SysTenancyId: tenancyId}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&user).Error
		if err != nil {
			return err
		}
		return nil
	})

	return user.ID, err
}

func RegisterClientMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		user, err := FindClientByStringId(fmt.Sprintf("%d", id))
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
	opts, err := GetAuthorityOptions(multi.TenancyAuthority)
	if err != nil {
		return form, err
	}
	form.Rule[0].Options = opts
	return form, nil
}
