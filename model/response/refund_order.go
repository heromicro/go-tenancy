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
	ReconciliationID uint            `json:"reconciliationId"` // 对账id
	OrderID          uint            `json:"orderId"`          // 订单id
	SysUserID        uint            `json:"sysUserId" form:"sysUserId"`
	SysTenancyID     uint            `json:"sysTenancyId"` // 商户 id
	RefundProduct    []RefundProduct `gorm:"-" json:"refundProduct"`
}

type RefundProduct struct {
	RefundOrderID  uint  `json:"refundOrderId"`  // 退款单
	OrderProductID uint  `json:"orderProductId"` // 订单产品id
	RefundNum      int64 `json:"refundNum"`      // 退货数
	OrderProduct
}
type CheckRefundOrder struct {
}
