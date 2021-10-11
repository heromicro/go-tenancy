package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var ProductProductCate = new(productProductCate)

type productProductCate struct{}

var productProductCates = []model.ProductProductCate{
	{ProductId: 1, ProductCategoryID: 174, SysTenancyId: 1},
	{ProductId: 2, ProductCategoryID: 174, SysTenancyId: 1},
	{ProductId: 3, ProductCategoryID: 174, SysTenancyId: 1},
	{ProductId: 4, ProductCategoryID: 174, SysTenancyId: 1},
	{ProductId: 5, ProductCategoryID: 174, SysTenancyId: 1},
	{ProductId: 6, ProductCategoryID: 174, SysTenancyId: 1},
	{ProductId: 7, ProductCategoryID: 174, SysTenancyId: 1},
}

func (m *productProductCate) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("product_id IN ?", []int{1}).Find(&[]model.ProductProductCate{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> product_product_cates 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&productProductCates).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> product_product_cates 表初始数据成功!")
		return nil
	})
}
