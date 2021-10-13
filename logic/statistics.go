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
	"gorm.io/gorm"
)

// GetClientStatisticsMaim 大盘运营数据
// - 今天，昨天，同比上周
// - LikeStore 关注店铺
// - OrderNum 支付订单
// - PayPrice 支付金额
// - PayUser 支付人数
// - VisitNum 访客人数
func GetClientStatisticsMaim(tenancyId uint) (response.ClientStaticMain, error) {
	staticMain := response.ClientStaticMain{Day: time.Now().Format("2006-01-02")}
	today, err := getClientStaticMain(
		scope.FilterToday("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterToday("created_at", ""),
		scope.FilterToday("created_at", "user_visits"),
		scope.FilterBase("sys_tenancy_id", "=", "", tenancyId),
		scope.FilterBase("type_id", "=", "", tenancyId),
		scope.FilterBase("sys_tenancy_id", "=", "products", tenancyId),
	)
	if err != nil {
		return staticMain, err
	}
	staticMain.Today = today

	yeaterday, err := getClientStaticMain(
		scope.FilterYesterday("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterYesterday("created_at", ""),
		scope.FilterYesterday("created_at", "user_visits"),
		scope.FilterBase("sys_tenancy_id", "=", "", tenancyId),
		scope.FilterBase("type_id", "=", "", tenancyId),
		scope.FilterBase("sys_tenancy_id", "=", "products", tenancyId),
	)
	if err != nil {
		return staticMain, err
	}
	staticMain.Yesterday = yeaterday

	lastWeekRate, err := getClientStaticMainLastWeekRate(tenancyId)
	if err != nil {
		return staticMain, err
	}
	staticMain.LastWeekRate = lastWeekRate

	return staticMain, nil
}

func getClientStaticMainLastWeekRate(tenancyId uint) (response.ClientStaticMainDataRate, error) {
	var staticMainData response.ClientStaticMainDataRate
	thisWeek, err := getClientStaticMain(
		scope.FilterThisWeek("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterThisWeek("created_at", ""),
		scope.FilterThisWeek("created_at", "user_visits"),
		scope.FilterBase("sys_tenancy_id", "=", "", tenancyId),
		scope.FilterBase("type_id", "=", "", tenancyId),
		scope.FilterBase("sys_tenancy_id", "=", "products", tenancyId),
	)
	if err != nil {
		return staticMainData, err
	}
	lastWeek, err := getClientStaticMain(
		scope.FilterLatelyWeek("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterLatelyWeek("created_at", ""),
		scope.FilterLatelyWeek("created_at", "user_visits"),
		scope.FilterBase("sys_tenancy_id", "=", "", tenancyId),
		scope.FilterBase("type_id", "=", "", tenancyId),
		scope.FilterBase("sys_tenancy_id", "=", "products", tenancyId),
	)
	if err != nil {
		return staticMainData, err
	}

	staticMainData.PayPrice = getRateFloat(thisWeek.PayPrice, lastWeek.PayPrice)
	staticMainData.PayUser = getRateInt(thisWeek.PayUser, lastWeek.PayUser)
	staticMainData.LikeStore = getRateInt(thisWeek.LikeStore, lastWeek.LikeStore)
	staticMainData.VisitNum = getRateInt(thisWeek.VisitNum, lastWeek.VisitNum)
	staticMainData.OrderNum = getRateInt(thisWeek.OrderNum, lastWeek.OrderNum)

	return staticMainData, nil
}

// getClientStaticMain 大盘运营数据
func getClientStaticMain(payTime, paid, createdAt, pCreatedAt, tenanacyId, typeId, pTenanacyId func(db *gorm.DB) *gorm.DB) (response.ClientStaticMainData, error) {
	var staticMainData response.ClientStaticMainData

	// 订单
	payPrice, err := service.GetOrderPayPrice(payTime, paid, tenanacyId)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.PayPrice = payPrice

	payUser, err := service.GetOrderUserNum(payTime, paid, tenanacyId)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.PayUser = payUser

	orderNum, err := service.GetPayOrderNum(payTime, tenanacyId)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.OrderNum = orderNum

	// 用户浏览记录，查询浏览商品记录的用户
	likeStore, err := service.GetLikeStore(createdAt, typeId)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.LikeStore = likeStore

	// 浏览记录，查询浏览商品记录的用户
	visitNum, err := service.GetVisitNum(pCreatedAt, pTenanacyId)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitNum = visitNum
	return staticMainData, nil
}

// GetStatisticsMain 大盘运营数据
// - 今天，昨天，同比上周
// -  payPrice 支付订单金额
// -  userNum 新增用户 包括小程序，床旁用户
// -  storeNum 店铺数
// -  visitUserNum 访客数
// -  visitNum 浏览量
func GetStatisticsMaim() (response.StaticMain, error) {
	var staticMain response.StaticMain
	today, err := getStaticMain(
		scope.FilterToday("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterToday("created_at", ""),
		scope.FilterToday("created_at", "user_visits"),
	)
	if err != nil {
		return staticMain, err
	}
	staticMain.Today = today

	yeaterday, err := getStaticMain(
		scope.FilterYesterday("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterYesterday("created_at", ""),
		scope.FilterYesterday("created_at", "user_visits"),
	)
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
func getStaticMain(payTime, paid, createdAt, pCreatedAt func(db *gorm.DB) *gorm.DB) (response.StaticMainData, error) {
	var staticMainData response.StaticMainData

	// 订单
	payPrice, err := service.GetOrderPayPrice(payTime, paid)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.PayPrice = payPrice

	userNum, err := service.GetCUserNum(createdAt)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.UserNum = userNum

	storeNum, err := service.GetTenancyNum(createdAt)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.StoreNum = storeNum

	// 用户浏览记录，查询浏览商品记录的用户
	visitUserNum, err := service.GetVisitUserNum(pCreatedAt)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitUserNum = visitUserNum

	// 浏览记录，查询浏览商品记录的用户
	visitNum, err := service.GetVisitNum(pCreatedAt)
	if err != nil {
		return staticMainData, err
	}
	staticMainData.VisitNum = visitNum
	return staticMainData, nil
}

func getStaticMainLastWeekRate() (response.StaticMainDataRate, error) {
	var staticMainData response.StaticMainDataRate

	thisWeek, err := getStaticMain(
		scope.FilterThisWeek("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterThisWeek("created_at", ""),
		scope.FilterThisWeek("created_at", "user_visits"),
	)
	if err != nil {
		return staticMainData, err
	}
	lastWeek, err := getStaticMain(
		scope.FilterLatelyWeek("pay_time", ""),
		scope.FilterBase("paid", "=", "", g.StatusTrue),
		scope.FilterLatelyWeek("created_at", ""),
		scope.FilterLatelyWeek("created_at", "user_visits"),
	)
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

// GetStatisticsOrder 大盘运营订单金额数据
// - 今天，昨天
func GetClientStatisticsOrder(dateReq request.DateReq, tenancyId uint) ([]response.ClientStaticOrder, error) {
	var orders []response.ClientStaticOrder
	// 订单价格
	staticOrders, err := service.GetOrderGroup(scope.FilterBase("sys_tenancy_id", "=", "", tenancyId))
	if err != nil {
		return orders, err
	}
	dates := utils.GetDatesBetweenTwoDays(utils.GetStartTime(dateReq.Date), time.Now().AddDate(0, 0, -1))
	for _, date := range dates {
		order := response.ClientStaticOrder{
			Day: date,
		}
		for _, staticOrder := range staticOrders {
			if staticOrder.Day == date {
				order.PayPrice = staticOrder.PayPrice
				order.User = staticOrder.User
				order.Total = staticOrder.Total
			}
		}
		orders = append(orders, order)
	}

	return orders, nil
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

	orderNum, err := service.GetPayOrderNum(scope.FilterToday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.OrderNum = orderNum

	yesterdayOrderNum, err := service.GetPayOrderNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.OrderRate = getRateInt(orderNum, yesterdayOrderNum)

	monthOrderNum, err := service.GetPayOrderNum(scope.FilterMonth("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}
	staticOrderNum.MonthOrderNum = monthOrderNum

	yesterdayMonthOrderNum, err := service.GetPayOrderNum(scope.FilterYesterday("pay_time", ""))
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

	staticOrderNum.OrderNum = orderUserNum

	yesterdayOrderUserNum, err := service.GetOrderUserNum(scope.FilterYesterday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.OrderRate = getRateInt(staticOrderNum.OrderNum, yesterdayOrderUserNum)

	monthOrderUserNum, err := service.GetOrderUserNum(scope.FilterMonth("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.MonthOrderNum = monthOrderUserNum

	yesterdayMonthOrderUserNum, err := service.GetOrderUserNum(scope.FilterYesterday("pay_time", ""))
	if err != nil {
		return staticOrderNum, err
	}

	staticOrderNum.MonthRate = getRateInt(staticOrderNum.MonthOrderNum, yesterdayMonthOrderUserNum)

	todayOrderUserNumGroup, err := service.GetOrderUserNumGroup(scope.FilterMonth("pay_time", ""))
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
func GetStatisticsUser(dateReq request.DateReq, tenancyId uint) (response.User, error) {
	var user response.User

	orderUserNumScope := []func(*gorm.DB) *gorm.DB{scope.FilterToday("pay_time", "")}
	if tenancyId > 0 {
		orderUserNumScope = append(orderUserNumScope, scope.FilterBase("sys_tenancy_id", "=", "", tenancyId))
	}
	orderUser, err := service.GetOrderUserNum(orderUserNumScope...)
	if err != nil {
		return user, err
	}
	user.OrderUser = orderUser

	payOrderUserNumScope := []func(*gorm.DB) *gorm.DB{scope.FilterToday("pay_time", ""), scope.FilterBase("paid", "=", "", g.StatusTrue)}
	if tenancyId > 0 {
		payOrderUserNumScope = append(payOrderUserNumScope, scope.FilterBase("sys_tenancy_id", "=", "", tenancyId))
	}
	payOrderUser, err := service.GetOrderUserNum(payOrderUserNumScope...)
	if err != nil {
		return user, err
	}
	user.PayOrderUser = payOrderUser

	orderPayPriceScope := []func(*gorm.DB) *gorm.DB{scope.FilterDate(dateReq.Date, "pay_time", "orders")}
	if tenancyId > 0 {
		orderPayPriceScope = append(orderPayPriceScope, scope.FilterBase("sys_tenancy_id", "=", "orders", tenancyId))
	}
	orderPrice, err := service.GetOrderPayPrice(orderPayPriceScope...)
	if err != nil {
		return user, err
	}
	user.OrderPrice = orderPrice

	payPriceScope := []func(*gorm.DB) *gorm.DB{scope.FilterDate(dateReq.Date, "pay_time", "orders"), scope.FilterBase("paid", "=", "", g.StatusTrue)}
	if tenancyId > 0 {
		payPriceScope = append(payPriceScope, scope.FilterBase("sys_tenancy_id", "=", "orders", tenancyId))
	}
	payOrderPrice, err := service.GetOrderPayPrice(payPriceScope...)
	if err != nil {
		return user, err
	}
	user.PayOrderPrice = payOrderPrice

	visitUserScope := []func(*gorm.DB) *gorm.DB{scope.FilterDate(dateReq.Date, "created_at", "user_visits")}
	if tenancyId > 0 {
		visitUserScope = append(visitUserScope, scope.FilterBase("sys_tenancy_id", "=", "products", tenancyId))
	}
	visitUser, err := service.GetVisitUserNum(visitUserScope...)
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

// GetStatisticsProduct
func GetStatisticsProduct(dateReq request.DateReq, tenancyId uint) ([]*response.MerchantStockData, error) {
	productData, err := service.GetPayOrderProductNumGroup(scope.FilterDate(dateReq.Date, "pay_time", "orders"), scope.FilterBase("sys_tenancy_id", "=", "orders", tenancyId))
	if err != nil {
		return productData, err
	}
	return productData, nil
}

// GetStatisticsProductVisit
func GetStatisticsProductVisit(dateReq request.DateReq, tenancyId uint) ([]*response.ProductVisitData, error) {
	productVisitData, err := service.GetVisitProductNumGroup(scope.FilterDate(dateReq.Date, "created_at", "user_visits"), scope.FilterBase("sys_tenancy_id", "=", "products", tenancyId))
	if err != nil {
		return productVisitData, err
	}
	return productVisitData, nil
}

// GetStatisticsProductCart
func GetStatisticsProductCart(dateReq request.DateReq, tenancyId uint) ([]*response.ProductVisitData, error) {
	productCartData, err := service.GetCartProductNumGroup(scope.FilterDate(dateReq.Date, "created_at", "carts"), scope.FilterBase("sys_tenancy_id", "=", "products", tenancyId))
	if err != nil {
		return productCartData, err
	}
	return productCartData, nil
}
