package response

type CartList struct {
	SysTenancyID uint            `json:"sysTenancyId"`
	Name         string          `json:"name" form:"name"`
	Avatar       string          `json:"Avatar"`
	ProductID    uint            `json:"productId"`
	Products     []ProductDetail `json:"products" gorm:"-"`
}
