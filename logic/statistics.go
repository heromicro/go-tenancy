package logic

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/service/scope"
	"github.com/snowlyg/go-tenancy/utils"
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
	payPrice, err := service.GetOrderPayPrice(scope.FilterToday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
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
	payPrice, err := service.GetOrderPayPrice(scope.FilterYesterday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
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

	staticMainData.PayPrice = getRateFloat(thisWeek.PayPrice, lastWeek.PayPrice)
	staticMainData.UserNum = getRateInt(thisWeek.UserNum, lastWeek.UserNum)
	staticMainData.StoreNum = getRateInt(thisWeek.StoreNum, lastWeek.StoreNum)
	staticMainData.VisitNum = getRateInt(thisWeek.VisitNum, lastWeek.VisitNum)
	staticMainData.VisitUserNum = getRateInt(thisWeek.VisitUserNum, lastWeek.VisitUserNum)

	return staticMainData, nil
}

func getRateInt(thisWeek, lastWeek int64) float64 {
	if lastWeek == 0 || thisWeek == 0 {
		return 0
	}
	decPayPrice, b := decimal.NewFromInt(thisWeek).Sub(decimal.NewFromInt(lastWeek)).DivRound(decimal.NewFromInt(lastWeek), 2).Float64()
	if !b {
		return 0
	}
	return decPayPrice
}

func getRateFloat(thisWeek, lastWeek float64) float64 {
	if lastWeek == 0 || thisWeek == 0 {
		return 0
	}
	decPayPrice, b := decimal.NewFromFloat(thisWeek).Sub(decimal.NewFromFloat(lastWeek)).DivRound(decimal.NewFromFloat(lastWeek), 2).Float64()
	if !b {
		return 0
	}
	return decPayPrice
}

func getStaticMainThisWeek() (response.StaticMainData, error) {
	var staticMainData response.StaticMainData

	// 订单
	payPrice, err := service.GetOrderPayPrice(scope.FilterThisWeek("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
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
	payPrice, err := service.GetOrderPayPrice(scope.FilterLatelyWeek("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
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
	todayPrice, err := service.GetOrderPayPrice(scope.FilterToday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
	if err != nil {
		return staticOrder, err
	}
	staticOrder.TodayPrice = todayPrice

	// 订单价格
	yesterdayPrice, err := service.GetOrderPayPrice(scope.FilterYesterday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
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

	staticOrderNum.OrderRate = getRateInt(orderNum, yesterdayOrderNum)

	monthOrderNum, err := service.GetOrderNum(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.MonthOrderNum = monthOrderNum

	yesterdayMonthOrderNum, err := service.GetOrderNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.MonthRate = getRateInt(monthOrderNum, yesterdayMonthOrderNum)

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

	orderUserNum, err := service.GetOrderUserNum(scope.FilterToday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
	if err != nil {
		return staticOrderNum, err
	}
	orderPatientNum, err := service.GetOrderPatientNum(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.OrderNum = orderUserNum + orderPatientNum

	yesterdayOrderUserNum, err := service.GetOrderUserNum(scope.FilterYesterday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
	if err != nil {
		return staticOrderNum, err
	}
	yesterdayOrderPatientNum, err := service.GetOrderPatientNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.OrderRate = getRateInt(staticOrderNum.OrderNum, yesterdayOrderUserNum+yesterdayOrderPatientNum)

	monthOrderUserNum, err := service.GetOrderUserNum(scope.FilterMonth("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
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
	staticOrderNum.MonthRate = getRateInt(staticOrderNum.MonthOrderNum, yesterdayMonthOrderUserNum+yesterdayMonthOrderPatientNum)

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
	total, err := service.GetOrderPayPrice(scope.FilterDate(dateReq.Date, "pay_time", "orders"), scope.FilterBase("paid", "=", "", g.StatusTrue))
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

// GetStatisticsUserData 用户数据
// -new 新增用户
// -visit 新增访客
// -total 用户总数
func GetStatisticsUserData(dateReq request.DateReq) ([]*response.UserData, error) {
	var userDatas []*response.UserData

	scopeFilterBase := scope.FilterBase("created_at", "<", "", utils.GetStartTime(dateReq.Date).Format("2006-01-02"))
	baseUserNum, err := service.GetCUserNum(scopeFilterBase)
	if err != nil {
		return userDatas, err
	}
	userDataGroups, err := service.GetUserNumGroup(scope.FilterDate(dateReq.Date, "created_at", ""))
	if err != nil {
		return userDatas, err
	}
	visitDataGroups, err := service.GetVisitNumGroup(scope.FilterDate(dateReq.Date, "created_at", ""))
	if err != nil {
		return userDatas, err
	}
	dates := utils.GetDatesBetweenTwoDays(utils.GetStartTime(dateReq.Date), time.Now().AddDate(0, 0, -1))
	for _, date := range dates {
		userData := &response.UserData{
			Day: date,
		}
		for _, userDataGroup := range userDataGroups {
			if userDataGroup.Day == date {
				userData.New = userDataGroup.New
				userData.Total = userDataGroup.New + baseUserNum
			}
		}
		for _, visitDataGroup := range visitDataGroups {
			if visitDataGroup.Day == date {
				userData.Visit = visitDataGroup.Visit
			}
		}
		userDatas = append(userDatas, userData)
	}

	return userDatas, nil

}

// GetStatisticsUser 用户成交数据
// - OrderUser 下单人数
// - PayOrderUser 支付人数
// - OrderPrice 订单金额
// - PayOrderPrice 支付订单金额
// - VisitUser 访问人数
// - PayOrderRate 订单支付率 OrderPrice/PayOrderPrice
// - OrderRate 用户支付率  PayOrderUser/OrderUser
// - UserRate 用户下单率 OrderUser/VisitUser
func GetStatisticsUser(dateReq request.DateReq) (response.User, error) {
	var user response.User
	orderUser, err := service.GetOrderUserNum(scope.FilterToday("pay_time", ""))
	if err != nil {
		return user, err
	}
	user.OrderUser = orderUser

	payOrderUser, err := service.GetOrderUserNum(scope.FilterToday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
	if err != nil {
		return user, err
	}
	user.PayOrderUser = payOrderUser

	orderPrice, err := service.GetOrderPayPrice(scope.FilterDate(dateReq.Date, "pay_time", "orders"))
	if err != nil {
		return user, err
	}
	user.OrderPrice = orderPrice

	payOrderPrice, err := service.GetOrderPayPrice(scope.FilterDate(dateReq.Date, "pay_time", "orders"), scope.FilterBase("paid", "=", "", g.StatusTrue))
	if err != nil {
		return user, err
	}
	user.PayOrderPrice = payOrderPrice

	visitUser, err := service.GetVisitUserNum(scope.FilterDate(dateReq.Date, "created_at", "user_visits"))
	if err != nil {
		return user, err
	}
	user.VisitUser = visitUser

	user.PayOrderRate = getRateFloat(payOrderPrice, orderPrice)
	user.OrderRate = getRateInt(payOrderUser, orderUser)
	user.UserRate = getRateInt(orderUser, visitUser)

	return user, nil
}

// GetStatisticsUserRate 用户成交占比数据
// - newTotalPrice 所选日期内消费金额
// - newUser  所选日期内新增用户
// - oldTotalPrice 所选日期前消费金额
// - oldUser 所选日期前用户数
// - totalPrice 总消费金额
// - user 总用户数
func GetStatisticsUserRate(dateReq request.DateReq) (response.UserRate, error) {
	var userRate response.UserRate

	return userRate, nil
}
