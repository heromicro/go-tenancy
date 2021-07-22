package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// CreateCart
func CreateCart(req request.CreateCart) (model.Cart, error) {
	cart := model.Cart{BaseCart: req.BaseCart, SysUserID: req.SysUserID, SysTenancyID: req.SysTenancyID, ProductID: req.ProductID}
	err := g.TENANCY_DB.Model(&model.Cart{}).Where("sys_user_id = ?", req.SysUserID).Where("sys_tenancy_id = ?", req.SysTenancyID).Where("product_id = ?", req.ProductID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // 没有商品直接新建
		err = g.TENANCY_DB.Model(&model.Cart{}).FirstOrCreate(&cart).Error
		if err != nil {
			return cart, err
		}
	} else if err != nil {
		return cart, err
	} else { // 商品存在增加数量
		cartNum := cart.CartNum + req.CartNum
		err = g.TENANCY_DB.Model(&model.Cart{}).Where("id = ?", cart.ID).Update("cart_num", cartNum).Error
		if err != nil {
			return cart, err
		}
	}

	return cart, nil
}

// GetCartList
func GetCartList(ctx *gin.Context) ([]response.CartList, int64, error) {
	cartList := []response.CartList{}
	var count int64
	err := g.TENANCY_DB.Model(&model.Cart{}).
		Select("carts.product_id,sys_tenancies.avatar,sys_tenancies.name,sys_tenancies.id as sys_tenancy_id").
		Joins("left join sys_tenancies on sys_tenancies.id = carts.sys_tenancy_id").
		Where("carts.sys_user_id", multi.GetUserId(ctx)).
		Where("carts.is_pay", g.StatusFalse).
		Where("carts.is_fail", g.StatusFalse).
		Where("carts.sys_tenancy_id", multi.GetTenancyId(ctx)).
		Find(&cartList).Error
	if err != nil {
		return cartList, count, fmt.Errorf("get cart %w", err)
	}
	var productIds []uint
	for _, cart := range cartList {
		productIds = append(productIds, cart.ProductID)
	}
	if len(productIds) > 0 {
		products, err := GetProductByIDs(productIds, IsCuser(ctx))
		if err != nil {
			return cartList, count, fmt.Errorf("get cart %w", err)
		}
		if len(products) > 0 {
			for i := 0; i < len(cartList); i++ {
				for _, product := range products {
					if cartList[i].ProductID == product.ID {
						cartList[i].Products = append(cartList[i].Products, product)
					}
				}
			}
		}
		count = int64(len(products))
	}

	return cartList, count, err
}
