package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

const (
	OrderTypeUnknown int = iota
	OrderTypeGeneral     //普通
	OrderTypeSelf        //自提
)

// 支付类型
const (
	PayTypeUnknown int = iota
	PayTypeWx          //微信
	PayTypeRoutine     //小程序
	PayTypeH5          //h5
	PayTypeBalance     //余额
	PayTypeAlipay      //支付宝
)

// 0:待付款 1:待发货 2：待收货 3：待评价 4：已完成 5：已退款 6:已取消 10:待付尾款 11:尾款过期未付
const (
	OrderStatusNoPay     int = iota //待付款
	OrderStatusNoDeliver            //待发货
	OrderStatusNoReceive            //待收货
	OrderStatusNoComment            //待评价
	OrderStatusFinish               //已完成
	OrderStatusRefund               //已退款
	OrderStatusCancel               //已取消
)

const (
	DeliverTypeUnknown int = iota
	DeliverTypeFH          //发货
	DeliverTypeSH          //送货
	DeliverTypeXN          //虚拟
)

// Order 订单表
type Order struct {
	g.TENANCY_MODEL

	BaseOrder

	SysUserID        uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID     uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
	GroupOrderID     uint `gorm:"column:group_order_id;type:int" json:"groupOrderId"`
	ReconciliationID uint `gorm:"column:reconciliation_id;type:tinyint unsigned;not null;default:0" json:"reconciliationId"` // 对账id
}

type BaseOrder struct {
	OrderSn        string    `gorm:"column:order_sn;type:varchar(36);not null" json:"orderSn"`                                        // 订单号
	RealName       string    `gorm:"column:real_name;type:varchar(32);not null" json:"realName"`                                      // 用户姓名
	UserPhone      string    `gorm:"column:user_phone;type:varchar(18);not null" json:"userPhone"`                                    // 用户电话
	UserAddress    string    `gorm:"column:user_address;type:varchar(128);not null" json:"userAddress"`                               // 详细地址
	TotalNum       int64     `gorm:"column:total_num;type:int unsigned;not null;default:0" json:"totalNum"`                           // 订单商品总数
	TotalPrice     float64   `gorm:"column:total_price;type:decimal(8,2) unsigned;not null;default:0.00" json:"totalPrice"`           // 订单总价
	TotalPostage   float64   `gorm:"column:total_postage;type:decimal(8,2) unsigned;not null;default:0.00" json:"totalPostage"`       // 邮费
	PayPrice       float64   `gorm:"column:pay_price;type:decimal(8,2) unsigned;not null;default:0.00" json:"payPrice"`               // 实际支付金额
	PayPostage     float64   `gorm:"column:pay_postage;type:decimal(8,2) unsigned;not null;default:0.00" json:"payPostage"`           // 支付邮费
	CommissionRate float64   `gorm:"column:commission_rate;type:decimal(6,4) unsigned;not null;default:0.0000" json:"commissionRate"` // 平台手续费
	OrderType      int       `gorm:"column:order_type;type:tinyint unsigned;default:1" json:"orderType"`                              // 1普通 2自提
	Paid           int       `gorm:"column:paid;type:tinyint unsigned;not null;default:2" json:"paid"`                                // 支付状态
	PayTime        time.Time `gorm:"column:pay_time;type:timestamp" json:"payTime"`                                                   // 支付时间
	PayType        int       `gorm:"column:pay_type;type:tinyint(1);not null" json:"payType"`                                         // 支付方式  1=微信 2=小程序 3=h5 4=余额 5=支付宝
	Status         int       `gorm:"column:status;type:tinyint(1);not null;default:0" json:"status"`                                  // 订单状态（0：待付款 1:待发货 2：待收货 3：待评价 4：已完成 5：已退款 6：已取消）
	DeliveryType   int       `gorm:"column:delivery_type;type:varchar(32)" json:"deliveryType"`                                       // 发货类型(1:发货 2: 送货 3: 虚拟)
	DeliveryName   string    `gorm:"column:delivery_name;type:varchar(64)" json:"deliveryName"`                                       // 快递名称/送货人姓名
	DeliveryID     string    `gorm:"column:delivery_id;type:varchar(64)" json:"deliveryId"`                                           // 快递单号/手机号
	Mark           string    `gorm:"column:mark;type:varchar(512);not null" json:"mark"`                                              // 备注
	Remark         string    `gorm:"column:remark;type:varchar(512)" json:"remark"`                                                   // 管理员备注
	AdminMark      string    `gorm:"column:admin_mark;type:varchar(512)" json:"adminMark"`                                            // 总后台备注
	VerifyCode     string    `gorm:"index:verify_code;column:verify_code;type:char(16)" json:"verifyCode"`                            // 核销码
	VerifyTime     time.Time `gorm:"column:verify_time;type:timestamp" json:"verifyTime"`                                             // 核销时间
	ActivityType   int32     `gorm:"column:activity_type;type:tinyint unsigned;not null;default:1" json:"activityType"`               // 1：普通 2:秒杀 3:预售 4:助力
	Cost           float64   `gorm:"column:cost;type:decimal(8,2) unsigned;not null" json:"cost"`                                     // 成本价
	IsDel          int       `gorm:"column:is_del;type:tinyint unsigned;not null;default:2" json:"isDel"`                             // 是否删除
	IsSystemDel    int       `gorm:"column:is_system_del;type:tinyint(1);default:2" json:"isSystemDel"`                               // 后台是否删除
}

// OrderStatus 订单操作记录表
type OrderStatus struct {
	g.TENANCY_MODEL

	ChangeType    string    `gorm:"index:change_type;column:change_type;type:varchar(32);not null" json:"changeType"`       // 操作类型
	ChangeMessage string    `gorm:"column:change_message;type:varchar(256);not null" json:"changeMessage"`                  // 操作备注
	ChangeTime    time.Time `gorm:"column:change_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"changeTime"` // 操作时间

	OrderID uint `gorm:"index:order_id;column:order_id;type:int unsigned;not null" json:"orderId"` // 订单id
}

// OrderReceipt 订单发票信息
type OrderReceipt struct {
	g.TENANCY_MODEL

	ReceiptInfo  string    `gorm:"column:receipt_info;type:varchar(500);default:''" json:"receiptInfo"` // 发票类型：1.普通发票，2.增值税发票
	Status       int8      `gorm:"column:status;type:tinyint;default:0" json:"status"`                  // 开票状态：1.已出票,10.已寄出
	ReceiptSn    string    `gorm:"column:receipt_sn;type:varchar(255);default:''" json:"receiptSn"`     // 发票单号
	ReceiptNo    string    `gorm:"column:receipt_no;type:varchar(255)" json:"receiptNo"`                // 发票编号
	DeliveryInfo string    `gorm:"column:delivery_info;type:varchar(255)" json:"deliveryInfo"`          // 收票联系信息
	Mark         string    `gorm:"column:mark;type:varchar(255)" json:"mark"`                           // 用户备注
	ReceiptPrice float64   `gorm:"column:receipt_price;type:decimal(10,2)" json:"receiptPrice"`         // 开票金额
	OrderPrice   float64   `gorm:"column:order_price;type:decimal(10,2)" json:"orderPrice"`             // 订单金额
	StatusTime   time.Time `gorm:"column:status_time;type:datetime;not null" json:"statusTime"`         // 状态变更时间
	MerMark      string    `gorm:"column:mer_mark;type:varchar(255)" json:"merMark"`                    // 备注

	OrderID      string `gorm:"column:order_id;type:varchar(255);not null;default:0" json:"orderId"` // 订单ID
	SysUserID    uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID uint   `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}

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
	RefundNum    int64   `gorm:"column:refund_num;type:int unsigned;not null;default:0" json:"refundNum"`       // 可申请退货数量
	IsReply      int     `gorm:"column:is_reply;type:tinyint unsigned;not null;default:2" json:"isReply"`       // 是否评价
	ProductPrice float64 `gorm:"column:product_price;type:decimal(10,2) unsigned;not null" json:"productPrice"` // 商品金额

}
