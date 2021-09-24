package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/service/scope"
	"gorm.io/gorm"
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
	today, err := getStaticMain(scope.FilterToday("pay_time", ""))
	if err != nil {
		return ginH, err
	}
	ginH["today"] = today

	yeaterday, err := getStaticMain(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return ginH, err
	}
	ginH["yeaterday"] = yeaterday

	lastWeekRate, err := getStaticMain(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return ginH, err
	}
	ginH["lastWeekRate"] = lastWeekRate
	return ginH, nil
}

// getStaticMain 大盘运营数据
// - date 时间类型参数 today ，yesterday ，lastWeekRate
func getStaticMain(scope func(*gorm.DB) *gorm.DB) (gin.H, error) {
	ginH := gin.H{"payPrice": 0, "userNum": 0, "storeNum": 0, "visitUserNum": 0, "visitNum": 0}
	payPrice, err := service.GetOrderPayPrice(scope)
	if err != nil {
		return ginH, err
	}
	ginH["payPrice"] = payPrice
	return ginH, nil
}
