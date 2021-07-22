package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

// CreateCart
func CreateCart(req request.CreateCart) (model.Cart, error) {
	cart := model.Cart{BaseCart: req.BaseCart, SysUserID: req.SysUserID, SysTenancyID: req.SysTenancyID, ProductID: req.ProductID}
	err := g.TENANCY_DB.Model(&model.Cart{}).Create(&cart).Error
	return cart, err
}
