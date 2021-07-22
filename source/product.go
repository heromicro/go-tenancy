package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Product = new(product)

type product struct{}

// 出售中 1: is_show' => 1, 'status' => 1
// 仓库中 2:'is_show' => 2, 'status' => 1
// 3,4,5 商户才有
// 已售罄 3:'is_show' => 1, 'stock' => 0, 'status' => 1
// 警戒库存 4:'stock' => $stock ? $stock : 0, 'status' => 1
// 回收站 5:'deleted_at' => not null
// 待审核 6:'status' => 2
// 审核未通过 7:'status' => 3
var products = []model.Product{
	{BaseProduct: model.BaseProduct{SysTenancyID: 1, StoreName: "领立裁腰带短袖连衣裙", StoreInfo: "短袖连衣裙", Keyword: "连衣裙", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 40, Rank: 0, Sales: 1, Price: 80.00, Cost: 50.00, OtPrice: 100.00, Stock: 399, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 1, Refusal: "", Rate: 5.0, ReplyCount: 0, IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"}, SliderImage: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"},

	{BaseProduct: model.BaseProduct{SysTenancyID: 1, StoreName: "纯棉珠地撞色领polo裙", StoreInfo: "polo裙", Keyword: "polo裙", BarCode: "", SysBrandID: 2, IsShow: 2, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 40, Rank: 0, Sales: 1, Price: 160.00, Cost: 50.00, OtPrice: 180.00, Stock: 99, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"}, SliderImage: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"},

	{BaseProduct: model.BaseProduct{SysTenancyID: 1, StoreName: "精梳棉修身短袖T恤（圆/V领）", StoreInfo: "精梳", Keyword: "T恤", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 40, Rank: 0, Sales: 2, Price: 40.00, Cost: 20.00, OtPrice: 58.00, Stock: 0, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 102, SpecType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"}, SliderImage: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"},

	{BaseProduct: model.BaseProduct{SysTenancyID: 1, StoreName: "素湃黑科技纯棉疏水抗污短袖T恤", StoreInfo: "黑科技", Keyword: "T恤", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 2, ProductCategoryID: 162, UnitName: "件", Sort: 0, Rank: 0, Sales: 1, Price: 80.00, Cost: 60.00, OtPrice: 100.00, Stock: 99, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"}, SliderImage: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"},

	{BaseProduct: model.BaseProduct{SysTenancyID: 1, StoreName: "智能定制休闲单西 布雷泽海军蓝轻薄斜纹", StoreInfo: "西装定制", Keyword: "西装", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 3, ProductCategoryID: 162, UnitName: "件", Sort: 70, Rank: 0, Sales: 3, Price: 880.00, Cost: 500.00, OtPrice: 1680.00, Stock: 97, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"}, SliderImage: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"},

	{TENANCY_MODEL: g.TENANCY_MODEL{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}}, BaseProduct: model.BaseProduct{SysTenancyID: 1, StoreName: "梅湾街复古雪纺翻领上衣", StoreInfo: "雪纺", Keyword: "上衣", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 56, Rank: 0, Sales: 1, Price: 88.00, Cost: 100.00, OtPrice: 200.00, Stock: 134, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 96, SpecType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, IsGiftBag: 0, CareCount: 4, OldID: 0, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"}, SliderImage: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"},

	{BaseProduct: model.BaseProduct{SysTenancyID: 1, StoreName: "天丝一字领露肩上衣", StoreInfo: "天丝一字领露肩上衣", Keyword: "测试", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 1, UnitName: "件", Sort: 10, Rank: 0, Sales: 1, Price: 50.00, Cost: 20.00, OtPrice: 100.00, Stock: 99, IsHot: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 102, SpecType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, IsGiftBag: 0, CareCount: 1, OldID: 0, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"}, TENANCY_MODEL: g.TENANCY_MODEL{DeletedAt: gorm.DeletedAt{Time: time.Now()}}, SliderImage: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"},
}

var productContents = []model.ProductContent{
	{Content: "<p>好手机</p>", ProductID: 1, Type: 1},
	{Content: "<p>好手机</p>", ProductID: 2, Type: 1},
	{Content: "<p>好手机</p>", ProductID: 3, Type: 1},
	{Content: "<p>好手机</p>", ProductID: 4, Type: 1},
	{Content: "<p>好手机</p>", ProductID: 5, Type: 1},
	{Content: "<p>好手机</p>", ProductID: 6, Type: 1},
	{Content: "<p>好手机</p>", ProductID: 7, Type: 1},
}
var productAttrValues = []model.ProductAttrValue{
	{ProductID: 2, Detail: "", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "", Stock: 99, Sales: 1, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "12444", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd0"}, Type: 1},
	{ProductID: 4, Detail: "", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "", Stock: 99, Sales: 1, Image: "	http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "15454545", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd1"}, Type: 1},
	{ProductID: 1, Detail: "{\"\u5c3a\u5bf8\": \"S\"}", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "S", Stock: 99, Sales: 1, Image: "	http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "123456", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd2"}, Type: 1},
	{ProductID: 1, Detail: "{\"\u5c3a\u5bf8\": \"L\"}", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "L", Stock: 100, Sales: 0, Image: "	http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "123456", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd3"}, Type: 1},
	{ProductID: 1, Detail: "{\"\u5c3a\u5bf8\": \"XL\"}", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "XL", Stock: 100, Sales: 0, Image: "	http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "123456", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd4"}, Type: 1},
	{ProductID: 1, Detail: "{\"\u5c3a\u5bf8\": \"XXL\"}", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "XXL", Stock: 100, Sales: 0, Image: "	http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "123456", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd5"}, Type: 1},
	{ProductID: 6, Detail: "", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "", Stock: 99, Sales: 1, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "1774575", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd6"}, Type: 1},
	{ProductID: 7, Detail: "", BaseProductAttrValue: model.BaseProductAttrValue{Sku: "", Stock: 98, Sales: 2, Image: "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg", BarCode: "10024242", Cost: 50.00, OtPrice: 180.00, Price: 160.00, Volume: 1.00, Weight: 1.00, ExtensionOne: 0.00, ExtensionTwo: 0.00, Unique: "e2fe28308fd7"}, Type: 1},
}

var productAttrs = []model.ProductAttr{
	{AttrName: "S", AttrValues: "35", Type: 1, ProductID: 1},
	{AttrName: "L", AttrValues: "36", Type: 1, ProductID: 1},
	{AttrName: "XL", AttrValues: "37", Type: 1, ProductID: 1},
	{AttrName: "XXL", AttrValues: "38", Type: 1, ProductID: 1},
}

func (m *product) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Product{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_products 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&products).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.ProductContent{}).Create(&productContents).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.ProductAttr{}).Create(&productAttrs).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.ProductAttrValue{}).Create(&productAttrValues).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_products 表初始数据成功!")
		return nil
	})
}
