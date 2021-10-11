package request

import "github.com/snowlyg/go-tenancy/model"

type CreateCart struct {
	model.BaseCart
	ProductId    uint `json:"productId" binding:"required"` // 商品ID
	CUserId      uint `json:"cUserId"`
	PatientId    uint `json:"patientId"`
	SysTenancyId uint `json:"sysTenancyId"` // 商户 id
}

type ChangeCartNum struct {
	CartNum int64 `json:"cartNum" binding:"required"`
}
