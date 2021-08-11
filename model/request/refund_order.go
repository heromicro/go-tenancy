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

type CheckRefundOrder struct {
	Id uint `json:"id"  form:"id"  binding:"required"`
}

type CreateRefundOrder struct {
	Id            uint    `json:"id"  form:"id"  binding:"required"`
	RefundMessage string  `json:"refundMessage"  binding:"required"`
	RefundPrice   float64 `json:"refundPrice"  binding:"required"`
	RefundType    int     `json:"refundType"  binding:"required"`
	Num           uint    `json:"num"  binding:"required"`
	Mark          string  `json:"mark"`
}
