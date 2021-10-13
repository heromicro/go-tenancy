package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/multi"

	"gorm.io/gorm"
)

var (
	Authority          = new(authority)
	AdminAuthorityId   = "999"
	TenancyAuthorityId = "998"
	LiteAuthorityId    = "997" // 小程序用户
	DeviceAuthorityId  = "996" // 床旁设备用户
)

type authority struct{}

var authorities = []model.SysAuthority{
	{AuthorityId: AdminAuthorityId, AuthorityType: multi.AdminAuthority, AuthorityName: "超级管理员", ParentId: "0", DefaultRouter: "dashboard"},
	{AuthorityId: TenancyAuthorityId, AuthorityType: multi.TenancyAuthority, AuthorityName: "商户管理员", ParentId: "0", DefaultRouter: "dashboard"},
	{AuthorityId: LiteAuthorityId, AuthorityType: multi.GeneralAuthority, AuthorityName: "C端用户", ParentId: "0", DefaultRouter: "dashboard"},
	{AuthorityId: DeviceAuthorityId, AuthorityType: multi.GeneralAuthority, AuthorityName: "设备用户", ParentId: "0", DefaultRouter: "dashboard"},
	{AuthorityId: "8881", AuthorityName: "普通用户子角色", ParentId: AdminAuthorityId, AuthorityType: multi.AdminAuthority, DefaultRouter: "dashboard"},
	{AuthorityId: "9528", AuthorityType: multi.GeneralAuthority, AuthorityName: "测试角色", ParentId: "0", DefaultRouter: "dashboard"},
}

//@description: sys_authorities 表数据初始化
func (a *authority) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{AdminAuthorityId, TenancyAuthorityId}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authorities 表初始数据成功!")
		return nil
	})
}
