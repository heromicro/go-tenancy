package request

type PayOrder struct {
	OrderId   uint   `json:"orderId" form:"orderId"`
	TenancyId uint   `json:"tenancyId" form:"tenancyId"`
	UserId    uint   `json:"userId" form:"userId"`
	PatientID uint   `json:"patientId" form:"patientId"`
	OrderType int    `json:"orderType" form:"orderType"`
	OpenId    string `json:"openId" form:"openId"`
	UserAgent string `json:"userAgent" form:"userAgent"`
	Code      string `json:"code" form:"code"`
	State     string `json:"state" form:"state"`
}
