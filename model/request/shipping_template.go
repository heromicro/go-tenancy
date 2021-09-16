package request

type ShippingTemplatePageInfo struct {
	PageInfo
	Name string `json:"name" form:"name"`
}
