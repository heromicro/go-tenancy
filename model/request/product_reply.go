package request

type ProductReplyPageInfo struct {
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
	IsReply  int    `json:"isReply" form:"isReply"`
	Keyword  string `json:"keyword" form:"keyword"`
	Nickname string `json:"nickname" form:"nickname"`
	Status   string `json:"status" form:"status"`
	Date     string `json:"date" form:"date"`
}
