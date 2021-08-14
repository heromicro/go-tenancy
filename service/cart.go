package service

import (
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateCart
func CreateCart(req request.CreateCart) (model.Cart, error) {
	var cart model.Cart
	err := g.TENANCY_DB.Model(&model.Cart{}).
		Where("sys_user_id = ?", req.SysUserID).
		Where("sys_tenancy_id = ?", req.SysTenancyID).
		Where("product_id = ?", req.ProductID).
		Where("is_pay = ?", g.StatusFalse).
		Where("is_fail = ?", g.StatusFalse).
		Where("is_new = ?", g.StatusFalse).
		Where("product_attr_unique = ?", req.ProductAttrUnique).
		First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // 没有商品直接新建
		cart = model.Cart{BaseCart: model.BaseCart{
			ProductType:       req.ProductType,
			ProductAttrUnique: req.ProductAttrUnique,
			CartNum:           req.CartNum,
			IsPay:             g.StatusFalse,
			IsNew:             req.IsNew,
			IsFail:            g.StatusFalse,
		}, SysUserID: req.SysUserID, SysTenancyID: req.SysTenancyID, ProductID: req.ProductID}
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
func ChangeCartNum(cartNum int64, id, userId, tenancyId uint) error {
	return g.TENANCY_DB.Model(&model.Cart{}).
		Where("id = ?", id).
		Where("sys_user_id = ?", userId).
		Where("sys_tenancy_id = ?", tenancyId).
		Update("cart_num", cartNum).Error
}

// ChangeIsPayByIds
func ChangeIsPayByIds(tx *gorm.DB, ids []uint) error {
	return tx.Model(&model.Cart{}).
		Where("id in ?", ids).
		Update("is_pay", g.StatusTrue).Error
}

// DeleteCart
func DeleteCart(ids []uint, userId, tenancyId uint) error {
	return g.TENANCY_DB.Model(&model.Cart{}).
		Where("sys_user_id = ?", userId).
		Where("sys_tenancy_id = ?", tenancyId).
		Where("id in ?", ids).
		Delete(&model.Cart{}).Error
}

// GetProductCount
func GetProductCount(userId, tenancyId uint) (int64, error) {
	var count int64
	err := g.TENANCY_DB.Model(&model.Cart{}).
		Where("sys_user_id = ?", userId).
		Where("sys_tenancy_id = ?", tenancyId).
		Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

// GetCartList
func GetCartList(tenancyId, userId uint, cartIds []uint) ([]response.CartList, []response.CartProduct, int64, error) {
	cartList := []response.CartList{}
	fails := []response.CartProduct{}
	var count int64
	cartProducts, err := GetCartProducts(tenancyId, userId, cartIds)
	if err != nil {
		return cartList, fails, count, fmt.Errorf("get cart %w", err)
	}
	tenancyIds := []uint{}
	if len(cartProducts) > 0 {
		for _, cartProduct := range cartProducts {
			if cartProduct.IsFail == g.StatusFalse {
				tenancyIds = append(tenancyIds, cartProduct.SysTenancyID)
			} else {
				fails = append(fails, cartProduct)
			}
		}
		err := g.TENANCY_DB.Model(&model.SysTenancy{}).
			Select("avatar,name,id as sys_tenancy_id").
			Where("status = ?", g.StatusTrue).
			Where("state = ?", g.StatusTrue).
			Where("id in ?", tenancyIds).
			Find(&cartList).Error
		if err != nil {
			return cartList, fails, count, fmt.Errorf("get cart %w", err)
		}
	}
	if len(cartList) > 0 {
		for i := 0; i < len(cartList); i++ {
			for _, cartProduct := range cartProducts {
				if cartProduct.SysTenancyID == cartList[i].SysTenancyID && cartProduct.IsFail == g.StatusFalse {
					cartList[i].Products = append(cartList[i].Products, cartProduct)
				}
			}
		}
	}
	count = int64(len(cartProducts))

	return cartList, fails, count, err
}
