package request

type RefundOrderPageInfo struct {
	Page          int    `json:"page" form:"page" binding:"required"`
	PageSize      int    `json:"pageSize" form:"pageSize" binding:"required"`
	Date          string `json:"date" form:"date"`
	IsTrader      string `json:"isTrader" form:"isTrader"`
	OrderSn       string `json:"orderSn" form:"orderSn"`
	RefundOrderSn string `json:"refundOrderSn" form:"refundOrderSn"`
	Status        string `json:"status" form:"status"`
}

type OrderAudit struct {
	Status      int    `json:"status"  binding:"required"`
	FailMessage string `json:"failMessage"`
}

type CreateRefundOrder struct {
	Ids           []uint  `json:"ids"  form:"ids"  binding:"required"` // 退款商品id
	RefundMessage string  `json:"refundMessage"  binding:"required"`   // 退款原因
	RefundPrice   float64 `json:"refundPrice"  binding:"required"`     // 退款金额
	RefundType    int     `json:"refundType"  binding:"required"`      // 退款类型
	Num           uint    `json:"num"  binding:"required"`             // 退款数量
	Mark          string  `json:"mark"`
}
