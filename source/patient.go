package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Patient = new(patient)

type patient struct{}

var patients = []model.Patient{
	{Name: "八两金", Phone: "13845687419", Sex: model.Female, Age: 32, LocName: "泌尿科一区", BedNum: "15", HospitalNO: "88956655", Disease: "不孕不育", SysTenancyId: 1},
}

func (m *patient) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Patient{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_patients 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&patients).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_patients 表初始数据成功!")
		return nil
	})
}
