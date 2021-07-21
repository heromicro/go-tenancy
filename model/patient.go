package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// Patient 患者
type Patient struct {
	g.TENANCY_MODEL
	Name       string `json:"name" gorm:"type:varchar(32);not null;comment:姓名"`
	Phone      string `json:"phone" gorm:"type:varchar(16);not null;comment:手机号"`
	Sex        int    `json:"sex" form:"sex" gorm:"not null;column:sex;comment:性别 0:女 1:男，2：未知"`
	Age        int    `json:"age" form:"age" gorm:"column:age;comment:年龄"`
	LocName    string `json:"locName" form:"locName" gorm:"type:varchar(50);column:loc_name;comment:科室名称"`
	BedNum     string `json:"bedNum" form:"bedNum" gorm:"type:varchar(10);column:bed_num;comment:床号"`
	HospitalNO string `json:"hospitalNo" form:"hospitalNo" gorm:"type:varchar(20);column:hospital_no;comment:住院号"`
	Disease    string `json:"disease" form:"disease" gorm:"type:varchar(150);column:disease;comment:病种"`

	SysTenancyID uint `json:"sysTenancyId" form:"sysTenancyId" gorm:"column:sys_tenancy_id;comment:关联标记"`
}
