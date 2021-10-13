package request

type PayOrder struct {
	OrderId   uint   `json:"orderId" form:"orderId"`
	TenancyId uint   `json:"tenancyId" form:"tenancyId"`
	UserId    uint   `json:"userId" form:"userId"`
	OrderType int    `json:"orderType" form:"orderType"`
	Expire    int64  `json:"expire" form:"expire"`
	OpenId    string `json:"openId" form:"openId"`
	UserAgent string `json:"userAgent" form:"userAgent"`
	Code      string `json:"code" form:"code"`
	State     string `json:"state" form:"state"`
}
