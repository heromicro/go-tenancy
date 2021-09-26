package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/service/scope"
)

// GetStatisticsMain 大盘运营数据
// - 今天，昨天，同比上周
// -  payPrice 支付订单金额
// -  userNum 新增用户 包括小程序，床旁用户
// -  storeNum 浏览量
// -  visitUserNum 访客
// -  visitNum 店铺数
func GetStatisticsMaim() (gin.H, error) {
	ginH := gin.H{
		"today":        gin.H{"payPrice": 0, "userNum": 0, "storeNum": 0, "visitUserNum": 0, "visitNum": 0}, //今天
		"yesterday":    gin.H{"payPrice": 0, "userNum": 0, "storeNum": 0, "visitUserNum": 0, "visitNum": 0}, //昨日
		"lastWeekRate": gin.H{"payPrice": 0, "userNum": 0, "storeNum": 0, "visitUserNum": 0, "visitNum": 0}, //同比上周
	}
	today, err := getStaticMainToday()
	if err != nil {
		return ginH, err
	}
	ginH["today"] = today

	// yeaterday, err := getStaticMain(scope.FilterYesterday("pay_time", ""))
	// if err != nil {
	// 	return ginH, err
	// }
	// ginH["yesterday"] = yeaterday

	// lastWeekRate, err := getStaticMain(scope.FilterLatelyWeek("pay_time", ""))
	// if err != nil {
	// 	return ginH, err
	// }
	// ginH["lastWeekRate"] = lastWeekRate
	return ginH, nil

}

// getStaticMain 大盘运营数据
// - date 时间类型参数 today ，yesterday ，lastWeekRate
func getStaticMainToday() (gin.H, error) {
	ginH := gin.H{"payPrice": 0, "userNum": 0, "storeNum": 0, "visitUserNum": 0, "visitNum": 0}

	// 订单
	payPrice, err := service.GetOrderPayPrice(scope.FilterToday("pay_time", ""))
	if err != nil {
		return ginH, err
	}
	ginH["payPrice"] = payPrice

	userNum, err := service.GetOrderPayPrice(scope.FilterToday("pay_time", ""))
	if err != nil {
		return ginH, err
	}
	ginH["userNum"] = userNum
	return ginH, nil
}
