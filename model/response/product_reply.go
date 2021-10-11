package response

import "github.com/snowlyg/go-tenancy/model"

type ProductReplyList struct {
	TenancyResponse
	model.BaseProductReply
	ProductId uint     `json:"productId"`
	SysUserId uint     `json:"sysUserId"`
	StoreName string   `gorm:"-" json:"storeName"`
	Image     string   `gorm:"-" json:"image"`
	Images    []string `gorm:"-" json:"images"`
}
