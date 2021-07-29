package response

import "github.com/snowlyg/go-tenancy/model"

type ProductReplyList struct {
	TenancyResponse
	model.BaseProductReply
	ProductID uint     `json:"productId"`
	SysUserID uint     `json:"sysUserId"`
	StoreName string   `gorm:"-" json:"storeName"`
	Image     string   `gorm:"-" json:"image"`
	Images    []string `gorm:"-" json:"images"`
}
