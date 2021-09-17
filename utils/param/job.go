package param

import "strconv"

// GetOrderAutoCloseTime 订单自动关闭时间(分钟)
func GetOrderAutoCloseTime() int64 {
	autoCloseTime, err := GetConfigValueByKey("auto_close_order_timer")
	if autoCloseTime == "" || err != nil {
		return 15
	}
	iact, err := strconv.ParseInt(autoCloseTime, 10, 64)
	if err != nil {
		return 15
	}
	return iact
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
