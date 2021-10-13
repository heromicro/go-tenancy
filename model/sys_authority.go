package model

import (
	"time"
)

type SysAuthority struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`

	AuthorityId string `json:"authorityId" gorm:"not null;primary_key;type:varchar(90)" binding:"required"`

	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名" binding:"required"`
	AuthorityType   int            `json:"authorityType" gorm:"comment:角色类型"`
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID" binding:"required"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter   string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`
}
