package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Receipt = new(receipt)

type receipt struct{}

var receipts = []model.UserReceipt{
	{ReceiptType: 1, ReceiptTitle: "title", ReceiptTitleType: 1, DutyGaragraph: "garagraph", Email: "email", BankName: "bank_name", BankCode: "bank_code", Address: "松山湖阿里产业园", Tel: "413514", IsDefault: g.StatusTrue, CUserID: 3},
}

func (m *receipt) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.UserReceipt{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_receipts 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&receipts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_receipts 表初始数据成功!")
		return nil
	})
}
