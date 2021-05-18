package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
)

type SysUser struct {
	g.TENANCY_MODEL
	Username    string       `json:"userName" gorm:"comment:用户登录名"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`

	AdminInfo   SysAdminInfo   `json:"adminInfo" gorm:"foreignKey:SysUserID;references:ID;comment:管理员信息"`
	TenancyInfo SysTenancyInfo `json:"tenancyInfo" gorm:"foreignKey:SysUserID;references:ID;comment:商户信息"`
	GeneralInfo SysGeneralInfo `json:"generalInfo" gorm:"foreignKey:SysUserID;references:ID;comment:普通用户信息"`
}

// IsAdmin
func (su *SysUser) IsAdmin() bool {
	return su.Authority.AuthorityType == multi.AdminAuthority
}

// IsTenancy
func (su *SysUser) IsTenancy() bool {
	return su.Authority.AuthorityType == multi.TenancyAuthority
}

// IsGeneral
func (su *SysUser) IsGeneral() bool {
	return su.Authority.AuthorityType == multi.GeneralAuthority
}

// AuthorityType
func (su *SysUser) AuthorityType() int {
	return su.Authority.AuthorityType
}

type SysAdminInfo struct {
	g.TENANCY_MODEL
	Email     string `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone     string `json:"phone" gorm:"default:'';comment:员工手机号" `
	NickName  string `json:"nickName" gorm:"default:'员工姓名';comment:员工姓名" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	SysUserID int    `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}

type SysTenancyInfo struct {
	g.TENANCY_MODEL
	Email        string `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone        string `json:"phone" gorm:"default:'';comment:员工手机号" `
	NickName     string `json:"nickName" gorm:"default:'员工姓名';comment:员工姓名" `
	HeaderImg    string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	SysUserID    int    `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID int    `json:"sysTenancyId" form:"sysTenancyId" gorm:"column:sys_tenancy_id;comment:关联标记"`
}

type Sex int

const (
	UnknownSex = iota
	Male
	Female
)

type SysGeneralInfo struct {
	g.TENANCY_MODEL
	Email     string    `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone     string    `json:"phone" gorm:"default:'';comment:员工手机号" `
	NickName  string    `json:"nickName" gorm:"default:'员工姓名';comment:员工姓名" `
	AvatarUrl string    `json:"avatarUrl" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	SysUserID int       `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	Sex       Sex       `json:"sex" form:"sex" gorm:"column:sex;comment:性别 1:男，2：女"`
	Subscribe int       `json:"subscribe" form:"subscribe" gorm:"column:subscribe;comment:是否订阅"`
	OpenId    string    `json:"openId" form:"openId" gorm:"column:open_id;comment:openid"`
	UnionId   string    `json:"unionId" form:"unionId" gorm:"column:union_id;comment:unionId"`
	Country   string    `json:"country" form:"country" gorm:"column:country;comment:国家"`
	Province  string    `json:"provice" form:"provice" gorm:"column:provice;comment:省份"`
	City      string    `json:"city" form:"city" gorm:"column:city;comment:城市"`
	IdCard    string    `json:"idCard" form:"idCard" gorm:"column:id_card;comment:身份证号"`
	IsAuth    int       `json:"isAuth" form:"isAuth" gorm:"column:is_auth;comment:是否实名认证"`
	RealName  string    `json:"realName" form:"realName" gorm:"column:real_name;comment:真实IP"`
	Birthday  time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日"`
}
