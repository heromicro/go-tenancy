package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

const (
	UnknownSale int32 = iota
	GeneralSale       // 普通商品
	FlashSale         //秒杀商品
	PreSale           // 预售商品
	AssistSale        // 助力商品
)

const (
	UnknownSpec int = iota
	SingleSpec      // 单规格
	DoubleSpec      // 多规格
)

const (
	UnknownProductStatus int = iota
	SuccessProductStatus     // 审核通过
	AuditProductStatus       //审核中
	FailProductStatus        // 未通过
)

type Product struct {
	g.TENANCY_MODEL
	BaseProduct
	SliderImage string `gorm:"column:slider_image;type:varchar(2000);not null" json:"sliderImage"` // 轮播图
}

type BaseProduct struct {
	StoreName   string  `gorm:"column:store_name;type:varchar(128);not null" json:"storeName" binding:"required"` // 商品名称
	StoreInfo   string  `gorm:"column:store_info;type:varchar(256);not null" json:"storeInfo" `                   // 商品简介
	Keyword     string  `gorm:"column:keyword;type:varchar(128);not null" json:"keyword"`                         // 关键字
	BarCode     string  `gorm:"column:bar_code;type:varchar(15);not null;default:''" json:"barCode"`              // 产品条码（一维码）
	IsShow      int     `gorm:"column:is_show;type:tinyint;not null;default:2" json:"isShow"`                     // 商户 状态（2：未上架，1：上架）
	Status      int     `gorm:"column:status;type:tinyint;not null;default:2" json:"status"`                      // 管理员 状态（1：审核通过,2：审核中 3: 未通过）
	UnitName    string  `gorm:"column:unit_name;type:varchar(16);not null" json:"unitName" `                      // 单位名
	Sort        int16   `gorm:"index;column:sort;type:smallint;not null;default:0" json:"sort"`                   // 排序
	Rank        int16   `gorm:"column:rank;type:smallint;not null;default:0" json:"rank"`                         // 总后台排序
	Sales       int64   `gorm:"index:sales;column:sales;type:mediumint unsigned;not null;default:0" json:"sales"` // 销量
	Price       float64 `gorm:"column:price;type:decimal(10,2) unsigned;default:0.00" json:"price" `              // 最低价格
	Cost        float64 `gorm:"column:cost;type:decimal(10,2);default:0.00" json:"cost" `                         // 成本价
	OtPrice     float64 `gorm:"column:ot_price;type:decimal(10,2);default:0.00" json:"otPrice" `                  // 原价
	Stock       int64   `gorm:"column:stock;type:int unsigned;default:0" json:"stock" `                           // 总库存
	IsHot       int     `gorm:"column:is_hot;type:tinyint unsigned;not null;default:0" json:"isHot"`              // 是否热卖
	IsBenefit   int     `gorm:"column:is_benefit;type:tinyint unsigned;not null;default:0" json:"isBenefit"`      // 促销单品
	IsBest      int     `gorm:"column:is_best;type:tinyint unsigned;not null;default:0" json:"isBest"`            // 是否精品
	IsNew       int     `gorm:"column:is_new;type:tinyint unsigned;not null;default:0" json:"isNew"`              // 是否新品
	IsGood      int     `gorm:"column:is_good;type:tinyint;not null;default:2" json:"isGood"`                     // 是否优品推荐
	ProductType int32   `gorm:"column:product_type;type:tinyint unsigned;not null;default:0" json:"productType" ` // 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
	Ficti       int32   `gorm:"column:ficti;type:mediumint;default:0" json:"ficti"`                               // 虚拟销量
	Browse      int     `gorm:"column:browse;type:int;default:0" json:"browse"`                                   // 浏览量
	CodePath    string  `gorm:"column:code_path;type:varchar(64);not null;default:''" json:"codePath"`            // 产品二维码地址(用户小程序海报)
	VideoLink   string  `gorm:"column:video_link;type:varchar(200);not null;default:''" json:"videoLink"`         // 主图视频链接
	SpecType    int     `gorm:"column:spec_type;type:tinyint;not null" json:"specType" `                          // 规格 1单 2多
	Refusal     string  `gorm:"column:refusal;type:varchar(255)" json:"refusal"`                                  // 审核拒绝理由
	Rate        float64 `gorm:"column:rate;type:decimal(2,1);default:5.0" json:"rate"`                            // 评价分数
	ReplyCount  uint    `gorm:"column:reply_count;type:int unsigned;default:0" json:"replyCount"`                 // 评论数

	CareCount int    `gorm:"column:care_count;type:int;not null;default:0" json:"careCount"` // 收藏数
	Image     string `gorm:"column:image;type:varchar(256);not null" json:"image"`           // 商品图片

	OldID        uint `gorm:"column:old_id;type:int;default:0" json:"oldId"`
	TempID       uint `gorm:"column:temp_id;type:int;not null;default:0" json:"tempId"`                         // 运费模板ID
	SysTenancyId uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
	SysBrandID   uint `gorm:"column:sys_brand_id;type:int" json:"sysBrandId"`                                   // 品牌 id

	ProductCategoryID uint `gorm:"index:product_category_id;column:product_category_id;type:int;not null" json:"productCategoryId"` // 平台分类
}

// ProductReply 商品评论表
type ProductReply struct {
	g.TENANCY_MODEL
	BaseProductReply
	CUserId        uint   `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`                            // 用户ID
	SysTenancyId   uint   `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"`       // 商户 id
	ProductId      uint   `gorm:"index:product_id;column:product_id;type:int;not null" json:"productId"`                  // 商品id  // 商品id
	OrderProductId uint   `gorm:"index:order_product_id;column:order_product_id;type:int;not null" json:"orderProductId"` // 订单商品ID
	Unique         string `gorm:"index:order_product_id;column:unique;type:char(12)" json:"unique"`                       // 商品 sku
	ProductType    int32  `gorm:"column:product_type;type:tinyint;not null;default:1" json:"productType"`                 // 1=普通商品
}

type BaseProductReply struct {
	ProductScore         int       `gorm:"column:product_score;type:tinyint(1);not null" json:"productScore"`           // 商品分数
	ServiceScore         int       `gorm:"column:service_score;type:tinyint(1);not null" json:"serviceScore"`           // 服务分数
	PostageScore         int       `gorm:"column:postage_score;type:tinyint(1);not null" json:"postageScore"`           // 物流分数
	Rate                 float64   `gorm:"column:rate;type:float(2,1);default:5.0" json:"rate"`                         // 平均值
	Comment              string    `gorm:"column:comment;type:varchar(512);not null" json:"comment"`                    // 评论内容
	Pics                 string    `gorm:"column:pics;type:text;not null" json:"pics"`                                  // 评论图片
	MerchantReplyContent string    `gorm:"column:merchant_reply_content;type:varchar(300)" json:"merchantReplyContent"` // 管理员回复内容
	MerchantReplyTime    time.Time `gorm:"column:merchant_reply_time;type:timestamp;not nul" json:"merchantReplyTime"`  // 管理员回复时间
	IsReply              int       `gorm:"column:is_reply;type:tinyint(1);not null;default:1" json:"isReply"`           // 2未回复1已回复
	IsVirtual            int       `gorm:"column:is_virtual;type:tinyint(1);not null;default:1" json:"isVirtual"`       // 2不是虚拟评价1是虚拟评价
	Nickname             string    `gorm:"column:nickname;type:varchar(64);not null" json:"nickname"`                   // 用户名称
	Avatar               string    `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`                      // 用户头像
}

// ProductCate 商品商户分类关联表
type ProductProductCate struct {
	ProductId         uint `gorm:"column:product_id;type:int" json:"productId"`
	ProductCategoryID uint `gorm:"index:product_category_id;column:product_category_id;type:int;not null" json:"productCategoryId"` // 分类id
	SysTenancyId      uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"`                // 商户 id
}

// ProductContent 商品详情表
type ProductContent struct {
	Content string `gorm:"column:content;type:longtext;not null" json:"content"`       // 商品详情
	Type    int32  `gorm:"column:type;type:tinyint(1);not null;default:1" json:"type"` // 商品类型 1=普通

	ProductId uint `gorm:"product_contents:product_id;column:product_id;type:int;not null" json:"productId"` // 商品id
}

// ProductAttrValue 商品属性值表
type ProductAttrValue struct {
	g.TENANCY_MODEL

	Detail string `gorm:"column:detail;type:varchar(1000);not null;default:''" json:"detail"`
	BaseProductAttrValue
	Type      int32 `gorm:"column:type;type:tinyint(1);default:1" json:"type"`                     // 活动类型 1=商品
	ProductId uint  `gorm:"index:product_id;column:product_id;type:int;not null" json:"productId"` // 商品id
}

// ProductAttr 商品属性值表
type ProductAttr struct {
	g.TENANCY_MODEL
	AttrName   string `gorm:"column:attr_name;type:varchar(32);not null" json:"attrName"`            // 属性名
	AttrValues string `gorm:"column:attr_values;type:varchar(2000);not null" json:"attrValues"`      // 属性值
	Type       int32  `gorm:"column:type;type:tinyint(1);default:1" json:"type"`                     // 活动类型 1=商品
	ProductId  uint   `gorm:"index:product_id;column:product_id;type:int;not null" json:"productId"` // 商品id
}

type BaseProductAttrValue struct {
	Sku     string  `gorm:"index:sku;column:sku;type:varchar(128);not null" json:"sku"`             // 商品属性索引值 (attr_value|attr_value[|....])
	Stock   int64   `gorm:"column:stock;type:int unsigned;not null" json:"stock"`                   // 属性对应的库存
	Sales   uint    `gorm:"column:sales;type:int unsigned;not null;default:0" json:"sales"`         // 销量
	Image   string  `gorm:"column:image;type:varchar(128)" json:"image"`                            // 图片
	BarCode string  `gorm:"column:bar_code;type:varchar(50);not null;default:''" json:"barCode"`    // 产品条码
	Cost    float64 `gorm:"column:cost;type:decimal(8,2) unsigned;not null" json:"cost"`            // 成本价
	OtPrice float64 `gorm:"column:ot_price;type:decimal(8,2);not null;default:0.00" json:"otPrice"` // 原价
	Price   float64 `gorm:"column:price;type:decimal(8,2) unsigned;not null" json:"price"`          // 价格
	Volume  float64 `gorm:"column:volume;type:decimal(8,2);not null;default:0.00" json:"volume"`    // 体积
	Weight  float64 `gorm:"column:weight;type:decimal(8,2);not null;default:0.00" json:"weight"`    // 重量
	Unique  string  `gorm:"index;column:unique;type:char(12);not null;default:''" json:"unique"`    // 唯一值
}
