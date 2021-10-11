package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Cart = new(cart)

type cart struct{}

var carts = []model.Cart{
	{CUserId: 3, SysTenancyId: 1, BaseCart: model.BaseCart{ProductType: model.GeneralSale, ProductAttrUnique: "167a5a36ded1", CartNum: 1, Source: 0, SourceID: 0, IsPay: g.StatusTrue, IsNew: g.StatusFalse, IsFail: g.StatusFalse}, ProductId: 7},
	{CUserId: 3, SysTenancyId: 1, BaseCart: model.BaseCart{ProductType: model.GeneralSale, ProductAttrUnique: "167a5a36ded2", CartNum: 1, Source: 0, SourceID: 0, IsPay: g.StatusTrue, IsNew: g.StatusFalse, IsFail: g.StatusFalse}, ProductId: 7},
	{CUserId: 3, SysTenancyId: 1, BaseCart: model.BaseCart{ProductType: model.GeneralSale, ProductAttrUnique: "167a5a36ded3", CartNum: 1, Source: 0, SourceID: 0, IsPay: g.StatusTrue, IsNew: g.StatusFalse, IsFail: g.StatusFalse}, ProductId: 7},
	{CUserId: 3, SysTenancyId: 1, BaseCart: model.BaseCart{ProductType: model.GeneralSale, ProductAttrUnique: "167a5a36ded4", CartNum: 4, Source: 0, SourceID: 0, IsPay: g.StatusTrue, IsNew: g.StatusFalse, IsFail: g.StatusFalse}, ProductId: 7},
}

//@description: carts 表数据初始化
func (a *cart) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Cart{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> carts 表的初始数据已存在!")
			return nil
		}

		if err := tx.Create(&carts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> carts 表初始数据成功!")
		return nil
	})
}
