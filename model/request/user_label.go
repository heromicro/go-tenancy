package request

type UserLabelPageInfo struct {
	PageInfo
	LabelType int    `json:"labelType" form:"labelType"`
	Keyword   string `json:"keyword" form:"keyword"`
}
