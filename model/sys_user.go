package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type SysUser struct {
	g.TENANCY_MODEL

	Username  string `json:"userName" gorm:"not null;type:varchar(32);comment:用户登录名"`
	Password  string `json:"-"  gorm:"not null;type:varchar(128);comment:用户登录密码"`
	Status    int    `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`   // 账号冻结 1为正常，2为禁止
	IsShow    int    `gorm:"column:is_show;type:tinyint(1);not null;default:1" json:"is_show"` // 是否显示 1为正常，2为禁止
	Email     string `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone     string `json:"phone" gorm:"type:char(15);default:'';comment:员工手机号" `
	NickName  string `json:"nickName" gorm:"type:varchar(16);default:'员工姓名';comment:员工姓名" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`

	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"not null;type:varchar(90)"`
}
