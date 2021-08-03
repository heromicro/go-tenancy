package job

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CheckOrderPayStatus struct {
	OrderId   uint `json:"orderId" form:"orderId"`
	TenancyId uint `json:"tenancyId" form:"tenancyId"`
	UserId    uint `json:"userId" form:"userId"`
	OrderType int  `json:"orderType" form:"orderType"`
	CreatedAt time.Time
}

func (d CheckOrderPayStatus) Run() {
	if time.Since(d.CreatedAt).Minutes() >= 15 {
		err := ChangeOrderStatusByOrderId(d.OrderId, d.TenancyId, d.UserId, d.OrderType, model.OrderStatusCancel, "cancel", "取消订单[自动]")
		if err != nil {
			g.TENANCY_LOG.Error("定时自动取消任务错误", zap.String("自动取消订单任务", err.Error()))
		}
	}
}

func ChangeOrderStatusByOrderId(orderId, tenancyId, userId uint, orderType, status int, changeType, changeMessage string) error {
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Order{}).Where("sys_tenancy_id = ?", tenancyId).
			Where("id = ?", orderId).
			Where("sys_user_id = ?", userId).
			Where("order_type = ?", orderType).
			Where("is_system_del = ?", g.StatusFalse).
			Where("is_del = ?", g.StatusFalse).
			Update("status", status).Error
		if err != nil {
			return err
		}
		orderStatus := model.OrderStatus{ChangeType: changeType, ChangeMessage: changeMessage, ChangeTime: time.Now(), OrderID: orderId}
		err = tx.Create(&orderStatus).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
