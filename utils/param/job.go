package param

// GetOrderAutoCloseTime 订单自动关闭时间(分钟)
func GetOrderAutoCloseTime() string {
	autoCloseTime, _ := GetConfigValueByKey("auto_close_order_timer")
	if autoCloseTime == "" {
		return "15"
	}
	return autoCloseTime
}

// GetRefundOrderAutoAgreeTime 退款单自动确认时间(天)
func GetRefundOrderAutoAgreeTime() string {
	autoAgreeTime, _ := GetConfigValueByKey("mer_refund_order_agree")
	if autoAgreeTime == "" {
		return "7"
	}
	return autoAgreeTime
}

// GetOrderAutoTakeOrderTime 订单自动收货时间(天)
func GetOrderAutoTakeOrderTime() string {
	autoTakeOrderTime, _ := GetConfigValueByKey("auto_take_order_timer")
	if autoTakeOrderTime == "" {
		return "15"
	}
	return autoTakeOrderTime
}
