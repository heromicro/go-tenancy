package request

import "github.com/snowlyg/go-tenancy/model"

type CreateCart struct {
	model.BaseCart
	ProductID    uint `json:"productId" binding:"required"` // 商品ID
	SysUserID    uint `json:"sysUserId"`
	PatientID    uint `json:"patientId"`
	SysTenancyID uint `json:"sysTenancyId"` // 商户 id
}

type ChangeCartNum struct {
	CartNum int64 `json:"cartNum" binding:"required"`
}
