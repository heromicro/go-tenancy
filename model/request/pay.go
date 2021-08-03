package request

type PayOrder struct {
	OrderId   uint `json:"orderId" form:"orderId"`
	TenancyId uint `json:"tenancyId" form:"tenancyId"`
	UserId    uint `json:"userId" form:"userId"`
	OrderType int  `json:"orderType" form:"orderType"`
}
