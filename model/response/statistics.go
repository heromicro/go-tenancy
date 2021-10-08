package response

type MerchantStock struct {
	Total int64                `json:"total"`
	List  []*MerchantStockData `json:"list"`
}

type MerchantStockData struct {
	Total     int64   `json:"total"`
	Rate      float64 `json:"rate"`
	ProductId int64   `json:"product_id"`
	StoreName string  `json:"store_name"`
}

type MerchantVisit struct {
	Total int64                `json:"total"`
	List  []*MerchantVisitData `json:"list"`
}

type MerchantVisitData struct {
	Total       int64   `json:"total"`
	Rate        float64 `json:"rate"`
	TenancyId   int64   `json:"tenancy_id"`
	TenancyName string  `json:"tenancy_name"`
}

type MerchantRate struct {
	Total float64             `json:"total"`
	List  []*MerchantRateData `json:"list"`
}

type MerchantRateData struct {
	Price       float64 `json:"price"`
	Rate        float64 `json:"rate"`
	TenancyName string  `json:"tenancy_name"`
}

type StaticOrderNum struct {
	MonthOrderNum int64                `json:"monthOrderNum"`
	MonthRate     float64              `json:"monthRate"`
	OrderNum      int64                `json:"orderNum"`
	OrderRate     float64              `json:"orderRate"`
	Today         []StaticOrderNumData `json:"today"`
}

type StaticOrderNumData struct {
	Time  string  `json:"time"`
	Total float64 `json:"total"`
}

type StaticOrder struct {
	TodayPrice     float64           `json:"todayPrice"`
	YesterdayPrice float64           `json:"yesterdayPrice"`
	Order          []StaticOrderData `json:"order"`
}

type StaticOrderData struct {
	Time      string  `json:"time"`
	Today     float64 `json:"today"`
	Yesterday float64 `json:"yesterday"`
}

type StaticMain struct {
	Today        StaticMainData     `json:"today"`
	Yesterday    StaticMainData     `json:"yesterday"`
	LastWeekRate StaticMainDataRate `json:"lastWeekRate"`
}

type StaticMainDataRate struct {
	PayPrice     float64 `json:"payPrice"`
	UserNum      float64 `json:"userNum"`
	StoreNum     float64 `json:"storeNum"`
	VisitUserNum float64 `json:"visitUserNum"`
	VisitNum     float64 `json:"visitNum"`
}
type StaticMainData struct {
	PayPrice     float64 `json:"payPrice"`
	UserNum      int64   `json:"userNum"`
	StoreNum     int64   `json:"storeNum"`
	VisitUserNum int64   `json:"visitUserNum"`
	VisitNum     int64   `json:"visitNum"`
}

type UserData struct {
	Day   string `json:"day"`
	New   int64  `json:"new"`
	Total int64  `json:"total"`
	Visit int64  `json:"visit"`
}

type User struct {
	OrderPrice    float64 `json:"orderPrice"`
	OrderRate     float64 `json:"orderRate"`
	OrderUser     int64   `json:"orderUser"`
	PayOrderPrice float64 `json:"payOrderPrice"`
	PayOrderRate  float64 `json:"payOrderRate"`
	PayOrderUser  int64   `json:"payOrderUser"`
	UserRate      float64 `json:"userRate"`
	VisitUser     int64   `json:"visitUser"`
}

type UserRate struct {
	NewTotalPrice float64 `json:"newTotalPrice"`
	NewUser       int64   `json:"newUser"`
	OldTotalPrice float64 `json:"oldTotalPrice"`
	OldUser       int64   `json:"oldUser"`
	TotalPrice    float64 `json:"totalPrice"`
	User          int64   `json:"user"`
}
