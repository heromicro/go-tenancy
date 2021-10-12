package response

import (
	"github.com/snowlyg/go-tenancy/model"
)

type RefundOrderList struct {
	TenancyResponse
	model.BaseRefundOrder
	ActivityType     int32           `json:"activityType" form:"activityType"`
	OrderSn          string          `json:"orderSn" form:"orderSn"`
	UserNickName     string          `json:"userNickName" form:"userNickName"`
	TenancyName      string          `json:"tenancyName" form:"tenancyName"`
	IsTrader         int             `json:"isTrader" form:"isTrader"`
	ReconciliationId uint            `json:"reconciliationId"` // 对账id
	OrderId          uint            `json:"orderId"`          // 订单id
	CUserId          uint            `json:"CUserId" form:"CUserId"`
	SysTenancyId     uint            `json:"sysTenancyId"` // 商户 id
	RefundProduct    []RefundProduct `gorm:"-" json:"refundProduct"`
}

type RefundOrderDetail struct {
	TenancyResponse
	model.BaseRefundOrder
	OrderSn      string  `json:"orderSn" form:"orderSn"`
	PayPrice     float64 `json:"payPrice" form:"payPrice"`
	UserNickName string  `json:"userNickName" form:"userNickName"`
	TenancyName  string  `json:"tenancyName" form:"tenancyName"`
	OrderId      uint    `json:"orderId"` // 订单id
	CUserId      uint    `json:"CUserId" form:"CUserId"`
	SysTenancyId uint    `json:"sysTenancyId"` // 商户 id
}

type RefundProduct struct {
	RefundOrderId  uint  `json:"refundOrderId"`  // 退款单
	OrderProductId uint  `json:"orderProductId"` // 订单产品id
	RefundNum      int64 `json:"refundNum"`      // 退货数
	OrderProduct
}

type CheckRefundOrder struct {
	TotalRefundPrice float64        `json:"totalRefundPrice" form:"totalRefundPrice"`
	PostagePrice     float64        `json:"postagePrice" form:"postagePrice"`
	Product          []OrderProduct `json:"product" form:"product"`
	Status           int            `json:"status" form:"status"`
}
