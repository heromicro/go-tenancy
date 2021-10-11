package request

type FinancialRecordPageInfo struct {
	PageInfo
	Date    string `json:"date" form:"date"`
	Keyword string `json:"keyword" form:"keyword"`
}
