package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var FinancialRecord = new(financialRecord)

type financialRecord struct{}

var financialRecords = []model.FinancialRecord{
	{
		RecordSn: "jy1625619291396979921",
		OrderSn:  "wx1625619291308244221",
		UserInfo: "13672286043",

		FinancialType: "order",
		FinancialPm:   model.OutFinancialPm,
		Number:        89.00,
		SysTenancyId:  1,
		CUserId:       1,
		OrderId:       1,
	},
}

func (m *financialRecord) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.FinancialRecord{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> financial_records 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&financialRecords).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> financial_records 表初始数据成功!")
		return nil
	})
}
