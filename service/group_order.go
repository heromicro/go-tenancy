package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/utils/param"
	"gorm.io/gorm"
)

func GetNoPayGroupOrderAutoClose(isRemind bool) ([]uint, error) {
	whereCreatedAt := fmt.Sprintf("now() > SUBDATE(created_at,interval -%s minute)", param.GetOrderAutoCloseTime())
	orderIds := []uint{}
	db := g.TENANCY_DB.Model(&model.GroupOrder{}).
		Where("paid = ?", g.StatusFalse)

	if isRemind {
		db = db.Where("is_remind = ?", g.StatusFalse)
	}

	err := db.Where(whereCreatedAt).Find(&orderIds).Error
	if err != nil {
		return orderIds, err
	}
	return orderIds, nil
}

func CreateGroupOrder(db *gorm.DB, groupOrder model.GroupOrder) error {
	return db.Model(&model.GroupOrder{}).Create(&groupOrder).Error
}

func GetGroupOrderById(id uint) (model.GroupOrder, error) {
	var grouopOrder model.GroupOrder
	err := g.TENANCY_DB.Model(&model.GroupOrder{}).Where("id = ?", id).First(&grouopOrder).Error
	if err != nil {
		return grouopOrder, fmt.Errorf("获取订单组错误 %w", err)
	}
	return grouopOrder, nil
}

func UpdateGroupOrderById(db *gorm.DB, id uint, data map[string]interface{}) error {
	err := db.Model(&model.GroupOrder{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return fmt.Errorf("更新订单组错误 %w", err)
	}
	return nil
}

// CancelNoPayGroupOrders 取消订单
// - 取消订单组
// - 取消订单
// - 更新订单状态
// - 回退库存
func CancelNoPayGroupOrders(groupOrderId uint) error {
	orderGroup, err := GetGroupOrderById(groupOrderId)
	if err != nil {
		return err
	}
	if orderGroup.Paid != 0 {
		return errors.New("订单组状态错误")
	}
	orders, err := GetOrdersByGroupOrderId(groupOrderId)
	if err != nil {
		return err
	}
	if len(orders) == 0 {
		return errors.New("订单数据错误")
	}
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		data := map[string]interface{}{"is_cancel": g.StatusTrue}
		err := UpdateGroupOrderById(tx, groupOrderId, data)
		if err != nil {
			return err
		}
		var orderStatues []model.OrderStatus
		for _, order := range orders {
			err := UpdateOrderById(tx, order.ID, map[string]interface{}{"is_cancel": g.StatusTrue})
			if err != nil {
				return err
			}
			orderStatus := model.OrderStatus{ChangeType: "cancel", ChangeMessage: "取消订单[自动]", ChangeTime: time.Now(), OrderID: order.ID}
			orderStatues = append(orderStatues, orderStatus)
		}

		err = tx.Model(&model.OrderStatus{}).Create(&orderStatues).Error
		if err != nil {
			return fmt.Errorf("生成订单操作记录 %w", err)
		}
		return nil
	})
}
