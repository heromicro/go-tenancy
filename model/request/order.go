package request

import (
	"github.com/snowlyg/go-tenancy/model"
)

type OrderPageInfo struct {
	PageInfo
	SysTenancyId uint   `json:"sysTenancyId" form:"sysTenancyId"`
	CUserId      uint   `json:"cUserId" form:"cUserId"`
	Date         string `json:"date" form:"date"`
	IsTrader     string `json:"isTrader" form:"isTrader"`
	Keywords     string `json:"keywords" form:"keywords"`
	OrderSn      string `json:"orderSn" form:"orderSn"`
	Status       string `json:"status" form:"status"`
	Username     string `json:"username" form:"username"`
	OrderType    string `json:"orderType" form:"orderType"`
	ActivityType string `json:"activityType" form:"activityType"`
}

type CartInfo struct {
	Product     CartInfoProduct     `json:"product" form:"product"`
	ProductAttr CartInfoProductAttr `json:"productAttr" form:"productAttr"`
	ProductType int                 `json:"ProductType,omitempty" form:"ProductType"`
}

type CartInfoProduct struct {
	Image       string                     `json:"image" form:"image"`
	StoreName   string                     `json:"storeName" form:"storeName"`
	ProductId   uint                       `json:"productId,omitempty" form:"productId"`
	IsShow      int                        `json:"isShow,omitempty"` // 商户 状态（2：未上架，1：上架）
	Status      int                        `json:"status,omitempty"`
	UnitName    string                     `json:"unitName,omitempty" `
	Price       float64                    `json:"price,omitempty" `
	TempID      uint                       `json:"tempId,omitempty"`
	ProductType int32                      `json:"productType,omitempty"`
	Temp        model.BaseShippingTemplate `json:"temp,omitempty"`
}

type CartInfoProductAttr struct {
	Price     float64 `json:"price" form:"price"`
	Sku       string  `json:"sku" form:"sku"`
	ProductId uint    `json:"productId,omitempty" form:"productId"`
	Stock     uint    `json:"stock,omitempty" form:"stock"`
	Unique    string  `json:"unique,omitempty" form:"unique"`
	Volume    float64 `json:"volume,omitempty" form:"volume"`
	Weight    float64 `json:"weight,omitempty" form:"weight"`
	OtPrice   float64 `json:"otPrice,omitempty" form:"otPrice"`
	Cost      float64 `json:"cost,omitempty" form:"cost"`
}

type OrderRemarkAndUpdate struct {
	Remark       string  `json:"remark"`
	TotalPrice   float64 `json:"totalPrice"` // 订单总价
	TotalPostage float64 `json:"totalPostage"`
	PayPrice     float64 `json:"payPrice"`
}

type OrderRemark map[string]interface{}
type DeliveryOrder struct {
	DeliveryId   string `json:"deliveryId" form:"deliveryId"`
	DeliveryName string `json:"deliveryName" form:"deliveryName"`
	DeliveryType int    `json:"deliveryType" form:"deliveryType"`
}

type CreateOrder struct {
	CartIds   []uint `json:"cartIds" form:"cartIds"`
	OrderType int    `json:"orderType" form:"orderType"` // 点单类型
	Remark    string `json:"remark" form:"remark"`       //订单备注
}
