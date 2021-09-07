package response

import "github.com/snowlyg/go-tenancy/model/request"

type CartList struct {
	SysTenancyID uint          `json:"sysTenancyId"`
	Name         string        `json:"name" form:"name"`
	Avatar       string        `json:"Avatar"`
	Products     []CartProduct `json:"products" gorm:"-"`
}

type CartProduct struct {
	Id                uint                     `json:"id"`
	SysTenancyID      uint                     `json:"sysTenancyId"`
	SpecType          int                      `json:"specType,omitempty" `
	ProductID         uint                     `json:"productId"`
	StoreName         string                   `json:"storeName"`
	Image             string                   `json:"image"`
	CartNum           int64                    `json:"cartNum"`
	IsFail            int                      `json:"isFail"`
	ProductAttrUnique string                   `json:"productAttrUnique"` // 商品属性
	AttrValue         request.ProductAttrValue `gorm:"-" json:"attrValue"`
	Attr              request.ProductAttr      `gorm:"-" json:"attr"`
}
