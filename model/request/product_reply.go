package request

import "github.com/snowlyg/go-tenancy/model"

type ProductReplyPageInfo struct {
	PageInfo
	IsReply  int    `json:"isReply" form:"isReply"`
	Keyword  string `json:"keyword" form:"keyword"`
	Nickname string `json:"nickname" form:"nickname"`
	Status   string `json:"status" form:"status"`
	Date     string `json:"date" form:"date"`
}

type ProductReply struct {
	Content string `json:"content" form:"content"`
}

type AddFictiReply struct {
	model.BaseProductReply
	ProductID ProductID `json:"productId"  form:"productId"`
	Pic       []string  `json:"pic"  form:"pic"`
}
type ProductID struct {
	Id  uint   `json:"id"  form:"id"`
	Src string `json:"src"  form:"src"`
}
