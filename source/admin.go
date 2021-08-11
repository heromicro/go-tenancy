package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"

	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.SysUser{
	{Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", Status: g.StatusTrue, IsShow: g.StatusFalse, AuthorityId: AdminAuthorityId, AdminInfo: model.AdminInfo{Email: "admin@admin.com", Phone: "13800138000", NickName: "超级管理员", SysUserID: 1}},

	// {Username: "a303176530", Password: "e10adc3949ba59abbe56e057f20f883e", Status: g.StatusTrue, IsShow: g.StatusFalse, AuthorityId: TenancyAuthorityId, AdminInfo: model.AdminInfo{Email: "a303176530@admin.com", Phone: "13800138000", NickName: "商户管理员", SysUserID: 2}, SysTenancyID: 1},
	// {Username: "a303176532", Password: "e10adc3949ba59abbe56e057f20f883e", Status: g.StatusTrue, IsShow: g.StatusFalse, AuthorityId: TenancyAuthorityId, AdminInfo: model.AdminInfo{Email: "a303176530@admin.com", Phone: "13800138000", NickName: "商户管理员", SysUserID: 3}, SysTenancyID: 2},
	// {Username: "a303176533", Password: "e10adc3949ba59abbe56e057f20f883e", Status: g.StatusTrue, IsShow: g.StatusFalse, AuthorityId: TenancyAuthorityId, AdminInfo: model.AdminInfo{Email: "a303176530@admin.com", Phone: "13800138000", NickName: "商户管理员", SysUserID: 4}, SysTenancyID: 3},
	// {Username: "a303176534", Password: "e10adc3949ba59abbe56e057f20f883e", Status: g.StatusTrue, IsShow: g.StatusFalse, AuthorityId: TenancyAuthorityId, AdminInfo: model.AdminInfo{Email: "a303176530@admin.com", Phone: "13800138000", NickName: "商户管理员", SysUserID: 5}, SysTenancyID: 4},
	// {Username: "a303176535", Password: "e10adc3949ba59abbe56e057f20f883e", Status: g.StatusTrue, IsShow: g.StatusFalse, AuthorityId: TenancyAuthorityId, AdminInfo: model.AdminInfo{Email: "a303176530@admin.com", Phone: "13800138000", NickName: "商户管理员", SysUserID: 6}, SysTenancyID: 5},
}

var users = []model.SysUser{
	// {Username: "oZM5VwD_PCaPKQZ8zRGt-NUdU2uM", Password: "e10adc3949ba59abbe56e057f20f883e", AuthorityId: GeneralAuthorityId, GeneralInfo: model.GeneralInfo{BaseGeneralInfo: model.BaseGeneralInfo{Email: "a303176530@admin.com", Phone: "13800138000", NickName: "C端用户", AvatarUrl: "https://thirdwx.qlogo.cn/mmopen/vi_32/PEyYoZmTJtaJdeYWWibrnDUadmXKVYyTtyRq2nxtWbBic5jJTLTT4KHmox1tNvOicgIXxspgmxicghpCFob1icAIWFw/132", Sex: model.Female, Subscribe: 1, OpenId: "own1t5TysymNUqcZm-8giuEvT68M", UnionId: "oZM5VwCgvGUZvkrnrGrdJZI4e12k", IdCard: "445281199411285861", IsAuth: 0, Birthday: model.SetBirthday(), RealName: "余思琳", Mark: "mark", Address: "address", LastTime: time.Now(), LastIP: "127.0.0.1", NowMoney: 0.00, UserType: "routine", PayCount: 5, PayPrice: 20.00}, SysUserID: 3, GroupID: 1}},

	// {Username: "oZM5VwD_PCaPKQZ8zRGt-NUdU2uM1", Password: "e10adc3949ba59abbe56e057f20f883e", AuthorityId: GeneralAuthorityId, GeneralInfo: model.GeneralInfo{BaseGeneralInfo: model.BaseGeneralInfo{Email: "a3031765301@admin.com", Phone: "13800138001", NickName: "C端用户1", AvatarUrl: "https://thirdwx.qlogo.cn/mmopen/vi_32/PEyYoZmTJtaJdeYWWibrnDUadmXKVYyTtyRq2nxtWbBic5jJTLTT4KHmox1tNvOicgIXxspgmxicghpCFob1icAIWFw/132", Sex: model.Male, Subscribe: 1, OpenId: "own1t5TysymNUqcZm-8giuEvT68M1", UnionId: "oZM5VwCgvGUZvkrnrGrdJZI4e12k", IdCard: "445281199411285862", IsAuth: 0, Birthday: model.SetBirthday(), RealName: "余思琳1", Mark: "mark", Address: "address", LastTime: time.Now(), LastIP: "127.0.0.1", NowMoney: 0.00, UserType: "wechat", PayCount: 2, PayPrice: 20.00}, SysUserID: 4, GroupID: 2}},
}

var userUserLabels = []model.UserUserLabel{
	// {SysUserID: 7, UserLabelID: 1},
	// {SysUserID: 7, UserLabelID: 2},
	// {SysUserID: 8, UserLabelID: 1},
	// {SysUserID: 8, UserLabelID: 2},
}

//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2, 3}).Find(&[]model.SysUser{}).RowsAffected == 3 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if len(users) > 0 {
			if err := tx.Create(&users).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
		}
		if len(userUserLabels) > 0 {
			if err := tx.Model(&model.UserUserLabel{}).Create(&userUserLabels).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
