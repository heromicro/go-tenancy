package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var ProductReply = new(productReply)

type productReply struct{}

var productReplys = []model.ProductReply{
	{BaseProductReply: model.BaseProductReply{PostageScore: 5, ServiceScore: 5, ProductScore: 5, Rate: 5.0, Comment: "jfskdfksdjgkjgksfksdfkds", Pics: "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg", MerchantReplyContent: "sdfsd", MerchantReplyTime: time.Now(), IsReply: g.StatusTrue, IsVirtual: g.StatusFalse, Nickname: "sfsdfsd", Avatar: ""}, SysUserId: 7, SysTenancyId: 1, ProductId: 1, OrderProductId: 1, Unique: "e2fe28308fd2", ProductType: model.GeneralSale},
	{BaseProductReply: model.BaseProductReply{PostageScore: 5, ServiceScore: 5, ProductScore: 5, Rate: 5.0, Comment: "jfskdfksdjgkjgksfksdfkds", Pics: "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg", MerchantReplyContent: "sdfsd", MerchantReplyTime: time.Now(), IsReply: g.StatusTrue, IsVirtual: g.StatusFalse, Nickname: "sfsdfsd", Avatar: ""}, SysUserId: 7, SysTenancyId: 1, ProductId: 1, OrderProductId: 1, Unique: "e2fe28308fd3", ProductType: model.GeneralSale},
	{BaseProductReply: model.BaseProductReply{PostageScore: 5, ServiceScore: 5, ProductScore: 5, Rate: 5.0, Comment: "jfskdfksdjgkjgksfksdfkds", Pics: "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg", MerchantReplyContent: "sdfsd", MerchantReplyTime: time.Now(), IsReply: g.StatusTrue, IsVirtual: g.StatusFalse, Nickname: "sfsdfsd", Avatar: ""}, SysUserId: 7, SysTenancyId: 1, ProductId: 1, OrderProductId: 1, Unique: "e2fe28308fd4", ProductType: model.GeneralSale},
	{BaseProductReply: model.BaseProductReply{PostageScore: 5, ServiceScore: 5, ProductScore: 5, Rate: 5.0, Comment: "jfskdfksdjgkjgksfksdfkds", Pics: "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg", MerchantReplyContent: "sdfsd", MerchantReplyTime: time.Now(), IsReply: g.StatusTrue, IsVirtual: g.StatusFalse, Nickname: "sfsdfsd", Avatar: ""}, SysUserId: 7, SysTenancyId: 1, ProductId: 1, OrderProductId: 1, Unique: "e2fe28308fd5", ProductType: model.GeneralSale},
}

func (m *productReply) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.ProductReply{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_product_replys 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&productReplys).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_product_replys 表初始数据成功!")
		return nil
	})
}
