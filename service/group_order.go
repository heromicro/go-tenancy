package service

import (
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/utils/param"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetNoPayGroupOrderAutoClose(isRemind bool) ([]uint, error) {
	whereCreatedAt := fmt.Sprintf("now() > SUBDATE(created_at,interval -%d minute)", param.GetOrderAutoCloseTime())
	var orderIds []uint
	db := g.TENANCY_DB.Model(&model.GroupOrder{}).
		Select("id").
		Where("paid = ? and is_cancel=?", g.StatusFalse, g.StatusFalse)

	if isRemind {
		db = db.Where("is_remind = ?", g.StatusFalse)
	}

	err := db.Where(whereCreatedAt).Find(&orderIds).Error
	if err != nil {
		return orderIds, err
	}
	return orderIds, nil
}

func CreateGroupOrder(db *gorm.DB, groupOrder *model.GroupOrder) error {
	return db.Model(&model.GroupOrder{}).Create(groupOrder).Error
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
	orders, err := GetOrdersByGroupOrderId(groupOrderId)
	if err != nil {
		return err
	}
	if len(orders) == 0 {
		g.TENANCY_LOG.Info("自动取消订单", zap.Uint("订单组id", groupOrderId), zap.String("GetOrdersByGroupOrderId()", "没有查询到订单"))
		return nil
	}
	var orderIds []uint
	var orderStatues []model.OrderStatus
	for _, order := range orders {
		orderIds = append(orderIds, order.ID)
		orderStatus := model.OrderStatus{ChangeType: "cancel", ChangeMessage: "取消订单[自动]", ChangeTime: time.Now(), OrderId: order.ID}
		orderStatues = append(orderStatues, orderStatus)
	}
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		data := map[string]interface{}{"is_cancel": g.StatusTrue}
		err := UpdateGroupOrderById(tx, groupOrderId, data)
		if err != nil {
			return err
		}
		UpdateOrderByIds(tx, orderIds, map[string]interface{}{"is_cancel": g.StatusTrue})
		err = tx.Model(&model.OrderStatus{}).Create(&orderStatues).Error
		if err != nil {
			return fmt.Errorf("生成订单操作记录 %w", err)
		}
		return nil
	})
}
