package request

import "github.com/snowlyg/go-tenancy/model"

type CreateCart struct {
	model.BaseCart
	ProductID    uint `json:"productId" binding:"required"` // 商品ID
	SysUserID    uint `json:"sysUserId"`
	SysTenancyID uint `json:"sysTenancyId"` // 商户 id
}
