package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

const (
	RefundTypeUnknown int = iota
	RefundTypeTK          //退款
	RefundTypeAll         //退款退货
)

// 1:待审核 2:待退货 3:待收货 4:已退款 5:审核未通过
const (
	RefundStatusUnknown  int = iota //
	RefundStatusAudit               //待审核
	RefundStatusAgree               //待退货，审核通过
	RefundStatusBackgood            //待收货
	RefundStatusEnd                 //已退款
	RefundStatusRefuse              //审核未通过
)

// RefundOrder 订单退款表
type RefundOrder struct {
	g.TENANCY_MODEL
	BaseRefundOrder
	ReconciliationId uint `gorm:"column:reconciliation_id;type:int unsigned;default:0" json:"reconciliationId"` // 对账id
	PatientId        uint `json:"patientId" form:"patientId" gorm:"column:patient_id;comment:患者"`
	OrderId          uint `gorm:"index:oid;column:order_id;type:int unsigned;not null" json:"orderId"` // 订单id
	CUserId          uint `json:"cUserId" form:"cUserId" gorm:"column:c_user_id;comment:关联标记"`
	SysTenancyId     uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}

type BaseRefundOrder struct {
	RefundOrderSn      string    `gorm:"unique;column:refund_order_sn;type:varchar(32);not null" json:"refundOrderSn"`           // 退款单号
	DeliveryType       string    `gorm:"column:delivery_type;type:varchar(32)" json:"deliveryType"`                              // 快递公司
	DeliveryID         string    `gorm:"column:delivery_id;type:varchar(32)" json:"deliveryId"`                                  // 快递单号
	DeliveryMark       string    `gorm:"column:delivery_mark;type:varchar(200)" json:"deliveryMark"`                             // 快递备注
	DeliveryPics       string    `gorm:"column:delivery_pics;type:varchar(255)" json:"deliveryPics"`                             // 快递凭证
	DeliveryPhone      string    `gorm:"column:delivery_phone;type:varchar(18)" json:"deliveryPhone"`                            // 联系电话
	MerDeliveryUser    string    `gorm:"column:mer_delivery_user;type:varchar(32)" json:"merDeliveryUser"`                       // 收货人
	MerDeliveryAddress string    `gorm:"column:mer_delivery_address;type:varchar(32)" json:"merDeliveryAddress"`                 // 收货地址
	Phone              string    `gorm:"column:phone;type:varchar(18)" json:"phone"`                                             // 联系电话
	Mark               string    `gorm:"column:mark;type:varchar(200)" json:"mark"`                                              // 备注
	MerMark            string    `gorm:"column:mer_mark;type:varchar(255)" json:"merMark"`                                       // 商户备注
	AdminMark          string    `gorm:"column:admin_mark;type:varchar(255)" json:"adminMark"`                                   // 平台备注
	Pics               string    `gorm:"column:pics;type:varchar(255)" json:"pics"`                                              // 图片
	RefundType         int       `gorm:"column:refund_type;type:tinyint(1);not null" json:"refundType"`                          // 退款类型 1:退款 2:退款退货
	RefundMessage      string    `gorm:"column:refund_message;type:varchar(128);not null" json:"refundMessage"`                  // 退款原因
	RefundPrice        float64   `gorm:"column:refund_price;type:decimal(8,2);not null;default:0.00" json:"refundPrice"`         // 退款金额
	RefundNum          uint      `gorm:"column:refund_num;type:int unsigned;not null;default:0" json:"refundNum"`                // 退款数
	FailMessage        string    `gorm:"column:fail_message;type:varchar(200)" json:"failMessage"`                               // 未通过原因
	Status             int       `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`                         // 状态 1:待审核 2:待退货 3:待收货 4:已退款 5:审核未通过
	StatusTime         time.Time `gorm:"column:status_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"statusTime"` // 状态改变时间

	IsCancel    int `gorm:"column:is_cancel;type:tinyint unsigned;not null;default:2" json:"isCancel"`
	IsSystemDel int `gorm:"column:is_system_del;type:tinyint(1);default:2" json:"isSystemDel"` // 商户删除
}

// RefundProduct 退款单产品表
type RefundProduct struct {
	g.TENANCY_MODEL

	RefundOrderId  uint  `gorm:"index:refund_order_id;column:refund_order_id;type:int unsigned;not null" json:"refundOrderId"` // 退款单
	OrderProductId uint  `gorm:"column:order_product_id;type:int unsigned;not null" json:"orderProductId"`                     // 订单产品id
	RefundNum      int64 `gorm:"column:refund_num;type:int unsigned;not null;default:0" json:"refundNum"`                      // 退货数
}

// RefundStatus 订单操作记录表
type RefundStatus struct {
	g.TENANCY_MODEL

	RefundOrderId uint      `gorm:"index:refund_order_id;column:refund_order_id;type:int unsigned;not null" json:"refundOrderId"` // 退款单订单id
	ChangeType    string    `gorm:"index:change_type;column:change_type;type:varchar(32);not null" json:"changeType"`             // 操作类型
	ChangeMessage string    `gorm:"column:change_message;type:varchar(256);not null" json:"changeMessage"`                        // 操作备注
	ChangeTime    time.Time `gorm:"column:change_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"changeTime"`       // 操作时间
}
