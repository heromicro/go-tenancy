package model

import "github.com/snowlyg/go-tenancy/g"

// OrderProduct 订单购物详情表
type OrderProduct struct {
	g.TENANCY_MODEL

	BaseOrderProduct

	CartInfo  string `gorm:"column:cart_info;type:text;not null" json:"cartInfo"`                 // 购买东西的详细信息
	OrderID   uint   `gorm:"index:oid;column:order_id;type:int unsigned;not null" json:"orderId"` // 订单id
	SysUserID uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	CartID    uint   `gorm:"column:cart_id;type:int unsigned;not null;default:0" json:"cartId"`                        // 购物车id
	ProductID uint   `gorm:"index:product_id;column:product_id;type:int unsigned;not null;default:0" json:"productId"` // 商品ID
}

type BaseOrderProduct struct {
	ProductSku   string  `gorm:"column:product_sku;type:char(12);not null" json:"productSku"`                   // 商品 sku
	IsRefund     uint8   `gorm:"column:is_refund;type:tinyint unsigned;not null;default:0" json:"isRefund"`     // 是否退款   0:未退款 1:退款中 2:部分退款 3=全退
	ProductNum   int64   `gorm:"column:product_num;type:int unsigned;not null;default:0" json:"productNum"`     // 购买数量
	ProductType  int32   `gorm:"column:product_type;type:int;not null;default:0" json:"productType"`            // 1.普通商品 2.秒杀商品,3.预售商品
	RefundNum    int64    `gorm:"column:refund_num;type:int unsigned;not null;default:0" json:"refundNum"`       // 可申请退货数量
	IsReply      int     `gorm:"column:is_reply;type:tinyint unsigned;not null;default:2" json:"isReply"`       // 是否评价
	ProductPrice float64 `gorm:"column:product_price;type:decimal(10,2) unsigned;not null" json:"productPrice"` // 商品金额

}
