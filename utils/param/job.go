package param

import (
	"github.com/snowlyg/go-tenancy/g"
	"go.uber.org/zap"
)

// GetOrderAutoCloseTime 订单自动关闭时间(分钟)
func GetOrderAutoCloseTime() string {
	autoCloseTime, err := GetConfigValueByKey("auto_close_order_timer")
	g.TENANCY_LOG.Error("获取订单自动关闭时间错误", zap.String("错误", err.Error()))
	if autoCloseTime == "" {
		return "15"
	}
	return autoCloseTime
}

// GetRefundOrderAutoAgreeTime 退款单自动确认时间(天)
func GetRefundOrderAutoAgreeTime() string {
	autoAgreeTime, err := GetConfigValueByKey("mer_refund_order_agree")
	g.TENANCY_LOG.Error("退款单自动确认时间", zap.String("错误", err.Error()))
	if autoAgreeTime == "" {
		return "7"
	}
	return autoAgreeTime
}

// GetRefundOrderAutoTakeOrderTime 订单自动收货时间(天)
func GetRefundOrderAutoTakeOrderTime() string {
	autoTakeOrderTime, err := GetConfigValueByKey("auto_take_order_timer")
	g.TENANCY_LOG.Error("订单自动收货时间", zap.String("错误", err.Error()))
	if autoTakeOrderTime == "" {
		return "15"
	}
	return autoTakeOrderTime
}
