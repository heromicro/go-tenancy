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
		err = g.TENANCY_DB.Model(&model.Cart{}).Create(&cart).Error
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

// ChangeCartNum
func ChangeCartNum(cartNum uint16, id, sysUserID, sysTenancyID uint) error {
	return g.TENANCY_DB.Model(&model.Cart{}).
		Where("sys_user_id = ?", sysUserID).
		Where("sys_tenancy_id = ?", sysTenancyID).
		Where("product_id = ?", id).
		Update("cart_num", cartNum).Error
}

// DeleteCart
func DeleteCart(ids []int, sysUserID, sysTenancyID uint) error {
	return g.TENANCY_DB.Model(&model.Cart{}).
		Where("sys_user_id = ?", sysUserID).
		Where("sys_tenancy_id = ?", sysTenancyID).
		Where("product_id in ?", ids).
		Delete(&model.Cart{}).Error
}

// GetProductCount
func GetProductCount(sysUserID, sysTenancyID uint) (int64, error) {
	var count int64
	err := g.TENANCY_DB.Model(&model.Cart{}).
		Where("sys_user_id = ?", sysUserID).
		Where("sys_tenancy_id = ?", sysTenancyID).
		Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

// GetCartList
func GetCartList(ctx *gin.Context) ([]response.CartList, int64, error) {
	cartList := []response.CartList{}
	var count int64
	cartProducts, err := GetCartProducts(multi.GetTenancyId(ctx), multi.GetUserId(ctx))
	if err != nil {
		return cartList, count, fmt.Errorf("get cart %w", err)
	}
	tenancyIds := []uint{}
	if len(cartProducts) > 0 {
		for _, cartProduct := range cartProducts {
			tenancyIds = append(tenancyIds, cartProduct.SysTenancyID)
		}
		err := g.TENANCY_DB.Model(&model.SysTenancy{}).
			Select("avatar,name,id as sys_tenancy_id").
			Where("status = ?", g.StatusTrue).
			Where("state = ?", g.StatusTrue).
			Where("id in ?", tenancyIds).
			Find(&cartList).Error
		if err != nil {
			return cartList, count, fmt.Errorf("get cart %w", err)
		}
	}
	if len(cartList) > 0 {
		for i := 0; i < len(cartList); i++ {
			for _, cartProduct := range cartProducts {
				if cartProduct.SysTenancyID == cartList[i].SysTenancyID {
					cartList[i].Products = append(cartList[i].Products, cartProduct)
				}
			}
		}
	}
	count = int64(len(cartProducts))

	return cartList, count, err
}
