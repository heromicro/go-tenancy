package logic

import (
	"github.com/shopspring/decimal"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/service/scope"
)

// GetStatisticsMain 大盘运营数据
// - 今天，昨天，同比上周
// -  payPrice 支付订单金额
// -  userNum 新增用户 包括小程序，床旁用户
// -  storeNum 店铺数
// -  visitUserNum 访客数
// -  visitNum 浏览量
func GetStatisticsMaim() (response.StaticMain, error) {
	var staticMain response.StaticMain
	today, err := getStaticMainToday()
	if err != nil {
		return staticMain, err
	}
	staticMain.Today = today

	yeaterday, err := getStaticMainYesterday()
	if err != nil {
		return staticMain, err
	}
	staticMain.Yesterday = yeaterday

	lastWeekRate, err := getStaticMainLastWeekRate()
	if err != nil {
		return staticMain, err
	}
	staticMain.LastWeekRate = lastWeekRate

	return staticMain, nil
}

// getStaticMain 大盘运营数据
func getStaticMainToday() (response.StaticMainData, error) {
	var staticMainData response.StaticMainData

	// 订单
	payPrice, err := service.GetOrderPayPrice(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.PayPrice = payPrice

	userNum, err := service.GetCUserNum(scope.FilterToday("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.UserNum = userNum

	storeNum, err := service.GetTenancyNum(scope.FilterToday("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.StoreNum = storeNum

	// 用户浏览记录，查询浏览商品记录的用户
	visitUserNum, err := service.GetVisitUserNum(scope.FilterToday("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitUserNum = visitUserNum

	// 浏览记录，查询浏览商品记录的用户
	visitNum, err := service.GetVisitNum(scope.FilterToday("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitNum = visitNum
	return staticMainData, nil
}

func getStaticMainYesterday() (response.StaticMainData, error) {
	var staticMainData response.StaticMainData

	// 订单
	payPrice, err := service.GetOrderPayPrice(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.PayPrice = payPrice

	userNum, err := service.GetCUserNum(scope.FilterYesterday("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.UserNum = userNum

	storeNum, err := service.GetTenancyNum(scope.FilterYesterday("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.StoreNum = storeNum

	// 用户浏览记录，查询浏览商品记录的用户
	visitUserNum, err := service.GetVisitUserNum(scope.FilterYesterday("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitUserNum = visitUserNum

	// 浏览记录，查询浏览商品记录的用户
	visitNum, err := service.GetVisitNum(scope.FilterYesterday("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitNum = visitNum
	return staticMainData, nil
}

func getStaticMainLastWeekRate() (response.StaticMainDataRate, error) {
	var staticMainData response.StaticMainDataRate

	thisWeek, err := getStaticMainThisWeek()
	if err != nil {
		return staticMainData, err
	}
	lastWeek, err := getStaticMainLastWeek()
	if err != nil {
		return staticMainData, err
	}
	if thisWeek.PayPrice == 0 || lastWeek.PayPrice == 0 {
		staticMainData.PayPrice = 0
	} else {
		decPayPrice, b := decimal.NewFromFloat(thisWeek.PayPrice).Sub(decimal.NewFromFloat(lastWeek.PayPrice)).DivRound(decimal.NewFromFloat(lastWeek.PayPrice), 2).Float64()
		if !b {
			staticMainData.PayPrice = 0
		}
		staticMainData.PayPrice = decPayPrice
	}
	staticMainData.UserNum = getRate(thisWeek.UserNum, lastWeek.UserNum)
	staticMainData.StoreNum = getRate(thisWeek.StoreNum, lastWeek.StoreNum)
	staticMainData.VisitNum = getRate(thisWeek.VisitNum, lastWeek.VisitNum)
	staticMainData.VisitUserNum = getRate(thisWeek.VisitUserNum, lastWeek.VisitUserNum)

	return staticMainData, nil
}

func getRate(thisWeek, lastWeek int64) float64 {
	if lastWeek == 0 || thisWeek == 0 {
		return 0
	}
	decPayPrice, b := decimal.NewFromInt(thisWeek).Sub(decimal.NewFromInt(lastWeek)).DivRound(decimal.NewFromInt(lastWeek), 2).Float64()
	if !b {
		return 0
	}
	return decPayPrice
}

func getStaticMainThisWeek() (response.StaticMainData, error) {
	var staticMainData response.StaticMainData

	// 订单
	payPrice, err := service.GetOrderPayPrice(scope.FilterThisWeek("pay_time", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.PayPrice = payPrice

	userNum, err := service.GetCUserNum(scope.FilterThisWeek("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.UserNum = userNum

	storeNum, err := service.GetTenancyNum(scope.FilterThisWeek("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.StoreNum = storeNum

	// 用户浏览记录，查询浏览商品记录的用户
	visitUserNum, err := service.GetVisitUserNum(scope.FilterThisWeek("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitUserNum = visitUserNum

	// 浏览记录，查询浏览商品记录的用户
	visitNum, err := service.GetVisitNum(scope.FilterThisWeek("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitNum = visitNum
	return staticMainData, nil
}

func getStaticMainLastWeek() (response.StaticMainData, error) {
	var staticMainData response.StaticMainData

	// 订单
	payPrice, err := service.GetOrderPayPrice(scope.FilterLatelyWeek("pay_time", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.PayPrice = payPrice

	userNum, err := service.GetCUserNum(scope.FilterLatelyWeek("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.UserNum = userNum

	storeNum, err := service.GetTenancyNum(scope.FilterLatelyWeek("created_at", ""))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.StoreNum = storeNum

	// 用户浏览记录，查询浏览商品记录的用户
	visitUserNum, err := service.GetVisitUserNum(scope.FilterLatelyWeek("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitUserNum = visitUserNum

	// 浏览记录，查询浏览商品记录的用户
	visitNum, err := service.GetVisitNum(scope.FilterLatelyWeek("created_at", "user_visits"))
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitNum = visitNum
	return staticMainData, nil
}

// GetStatisticsOrder 大盘运营订单金额数据
// - 今天，昨天
func GetStatisticsOrder() (response.StaticOrder, error) {
	var staticOrder response.StaticOrder
	// 订单价格
	todayPrice, err := service.GetOrderPayPrice(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticOrder, err
	}
	staticOrder.TodayPrice = todayPrice

	// 订单价格
	yesterdayPrice, err := service.GetOrderPayPrice(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrder, err
	}
	staticOrder.YesterdayPrice = yesterdayPrice

	// 订单价格
	todayPriceGroup, err := service.GetOrderPayPriceGroup(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticOrder, err
	}

	// 订单价格
	yesterdayPriceGroup, err := service.GetOrderPayPriceGroup(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrder, err
	}
	staticOrder.Order = []response.StaticOrderData{}
	for _, timeStr := range timeStrs {
		order := response.StaticOrderData{Time: timeStr}

		if len(todayPriceGroup) > 0 {
			for _, todayPrice := range todayPriceGroup {
				if todayPrice.Time == timeStr {
					order.Today = todayPrice.Count
				}
			}
		}

		if len(yesterdayPriceGroup) > 0 {
			for _, yesterdayPrice := range yesterdayPriceGroup {
				if yesterdayPrice.Time == timeStr {
					order.Today = yesterdayPrice.Count
				}
			}
		}
		staticOrder.Order = append(staticOrder.Order, order)
	}

	return staticOrder, nil
}

// GetStatisticsOrderNum 大盘运营订单数量数据
// -当日订单数 OrderNum
// -当月订单数 MonthOrderNum
// -当日订单数同比 OrderRate
// -当月订单数同比 MonthRate
func GetStatisticsOrderNum() (response.StaticOrderNum, error) {
	var staticOrderNum response.StaticOrderNum

	orderNum, err := service.GetOrderNum(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.OrderNum = orderNum

	yesterdayOrderNum, err := service.GetOrderNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.OrderRate = getRate(orderNum, yesterdayOrderNum)

	monthOrderNum, err := service.GetOrderNum(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.MonthOrderNum = monthOrderNum

	yesterdayMonthOrderNum, err := service.GetOrderNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.MonthRate = getRate(monthOrderNum, yesterdayMonthOrderNum)

	todayOrderNumGroup, err := service.GetOrderNumGroup(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.Today = []response.StaticOrderNumData{}
	for _, timeStr := range timeStrs {
		order := response.StaticOrderNumData{Time: timeStr}

		if len(todayOrderNumGroup) > 0 {
			for _, todayOrderNum := range todayOrderNumGroup {
				if todayOrderNum.Time == timeStr {
					order.Total = todayOrderNum.Count
				}
			}
		}
	}

	return staticOrderNum, nil
}

// GetStatisticsOrderUser 大盘运营订单用户数量数据 (用户数据未床旁用户加上小程序用户)
// -当日支付人数 OrderNum
// -当月支付人数 MonthOrderNum
// -当日支付人数同比 OrderRate
// -当月支付人数同比 MonthRate
func GetStatisticsOrderUser() (response.StaticOrderNum, error) {
	var staticOrderNum response.StaticOrderNum

	orderUserNum, err := service.GetOrderUserNum(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	orderPatientNum, err := service.GetOrderPatientNum(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.OrderNum = orderUserNum + orderPatientNum

	yesterdayOrderUserNum, err := service.GetOrderUserNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	yesterdayOrderPatientNum, err := service.GetOrderPatientNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.OrderRate = getRate(staticOrderNum.OrderNum, yesterdayOrderUserNum+yesterdayOrderPatientNum)

	monthOrderUserNum, err := service.GetOrderUserNum(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	monthOrderPatientNum, err := service.GetOrderPatientNum(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.MonthOrderNum = monthOrderUserNum + monthOrderPatientNum

	yesterdayMonthOrderUserNum, err := service.GetOrderUserNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	yesterdayMonthOrderPatientNum, err := service.GetOrderPatientNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.MonthRate = getRate(staticOrderNum.MonthOrderNum, yesterdayMonthOrderUserNum+yesterdayMonthOrderPatientNum)

	todayOrderUserNumGroup, err := service.GetOrderUserNumGroup(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	todayOrderPatientNumGroup, err := service.GetOrderPatientNumGroup(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.Today = []response.StaticOrderNumData{}
	for _, timeStr := range timeStrs {
		order := response.StaticOrderNumData{Time: timeStr}

		if len(todayOrderUserNumGroup) > 0 {
			for _, todayOrderUserNum := range todayOrderUserNumGroup {
				if todayOrderUserNum.Time == timeStr {
					order.Total += todayOrderUserNum.Count
				}
			}
		}

		if len(todayOrderPatientNumGroup) > 0 {
			for _, todayOrderPatientNum := range todayOrderPatientNumGroup {
				if todayOrderPatientNum.Time == timeStr {
					order.Total += todayOrderPatientNum.Count
				}
			}
		}
	}

	return staticOrderNum, nil
}

// GetStatisticsMerchantStock 商品销量排行
// - 统计支付订单商品数量
func GetStatisticsMerchantStock(dateReq request.DateReq) (response.MerchantStock, error) {
	var merchantStocks response.MerchantStock
	total, err := service.GetPayOrderProductNum(scope.FilterDate(dateReq.Date, "pay_time", "orders"))
	if err != nil {
		return merchantStocks, err
	}
	merchantData, err := service.GetPayOrderProductNumGroup(scope.FilterDate(dateReq.Date, "pay_time", "orders"))
	if err != nil {
		return merchantStocks, err
	}
	for _, merchant := range merchantData {
		rate, b := decimal.NewFromInt(merchant.Total).DivRound(decimal.NewFromInt(total), 2).Float64()
		if !b {
			rate = 0
		}
		merchant.Rate = rate
	}
	merchantStocks.Total = total
	merchantStocks.List = merchantData
	return merchantStocks, nil
}

// GetStatisticsMerchantVisit 商户访客量排行
func GetStatisticsMerchantVisit(dateReq request.DateReq) (response.MerchantVisit, error) {
	var merchantVisits response.MerchantVisit
	total, err := service.GetVisitUserNum(scope.FilterDate(dateReq.Date, "created_at", "user_visits"))
	if err != nil {
		return merchantVisits, err
	}
	merchantVisitData, err := service.GetVisitUserNumGroup(scope.FilterDate(dateReq.Date, "created_at", "user_visits"))
	if err != nil {
		return merchantVisits, err
	}

	for _, merchantVisit := range merchantVisitData {
		rate, b := decimal.NewFromInt(merchantVisit.Total).DivRound(decimal.NewFromInt(total), 2).Float64()
		if !b {
			rate = 0
		}
		merchantVisit.Rate = rate
	}

	merchantVisits.Total = total
	merchantVisits.List = merchantVisitData
	return merchantVisits, nil

}

// GetStatisticsMerchantRate 商户销售额占比
func GetStatisticsMerchantRate(dateReq request.DateReq) (response.MerchantRate, error) {
	var merchantRates response.MerchantRate
	total, err := service.GetOrderPayPrice(scope.FilterDate(dateReq.Date, "pay_time", "orders"))
	if err != nil {
		return merchantRates, err
	}
	merchantRateData, err := service.GetTenancyOrderPayPriceGroup(scope.FilterDate(dateReq.Date, "pay_time", "orders"))
	if err != nil {
		return merchantRates, err
	}

	for _, merchantRate := range merchantRateData {
		rate, b := decimal.NewFromFloat(merchantRate.Price).DivRound(decimal.NewFromFloat(total), 2).Float64()
		if !b {
			rate = 0
		}
		merchantRate.Rate = rate
	}

	merchantRates.Total = total
	merchantRates.List = merchantRateData
	return merchantRates, nil

}
