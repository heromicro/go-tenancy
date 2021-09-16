package request

import (
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/datatypes"
)

type ProductPageInfo struct {
	PageInfo
	ProductCategoryId uint   `json:"tenancyCategoryId" form:"tenancyCategoryId"`
	Type              string `json:"type" form:"type"`
	CateId            int    `json:"cateId" form:"cateId"`
	Keyword           string `json:"keyword" form:"keyword"`
}

type UpdateProduct struct {
	Id uint `json:"id"`
	CreateProduct
}
type SetProductFicti struct {
	Ficti  int32  `json:"ficti"`
	Number string `json:"number"`
	Type   int    `json:"type" binding:"required"` // 1:+ ,2:-
}

type ChangeProductStatus struct {
	Id      []uint `json:"id" form:"id" binding:"required,gt=0"`
	Status  int    `json:"status" binding:"required"`
	Refusal string `json:"refusal" `
}
type ChangeProductIsShow struct {
	Id     uint `json:"id" form:"id" binding:"required,gt=0"`
	IsShow int  `json:"isShow" binding:"required"`
}

type CreateProduct struct {
	model.BaseProduct
	SliderImages []string           `json:"sliderImages"`      // 轮播图
	CateId       uint               `json:"cateId"`            // 平台分类id
	CategoryIds  []uint             `json:"tenancyCategoryId"` // 平台分类id
	Content      string             `json:"content"`           // 商品内容
	AttrValue    []ProductAttrValue `json:"attrValue"`         // 商品规格
	Attr         []ProductAttr      `json:"attr"`              // 商品规格
}

type ProductAttrValue struct {
	ProductID uint `json:"productId"` // 商品id
	model.BaseProductAttrValue
	Detail datatypes.JSON `json:"detail"`
	Value0 string         `json:"value"`
}
type ProductAttr struct {
	Detail datatypes.JSON `json:"detail"`
	Value  string         `json:"value"`
}
