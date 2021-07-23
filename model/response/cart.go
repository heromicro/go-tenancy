package response

type CartList struct {
	SysTenancyID uint          `json:"sysTenancyId"`
	Name         string        `json:"name" form:"name"`
	Avatar       string        `json:"Avatar"`
	Products     []CartProduct `json:"products" gorm:"-"`
}

type CartProduct struct {
	SysTenancyID uint   `json:"sysTenancyId"`
	ProductID    uint   `json:"productId"`
	StoreName    string `json:"storeName"`
	Image        string `json:"image"`
	Price        string `json:"price"`
	CartNum      uint16 `json:"cartNum"`
}
