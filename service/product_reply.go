package service

import (
	"strings"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
)

// GetProductReplyInfoList
func GetProductReplyInfoList(info request.ProductReplyPageInfo, tenancyId uint) ([]response.ProductReplyList, int64, error) {
	var productReplyList []response.ProductReplyList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.ProductReply{}).
		Where("sys_tenancy_id = ?", tenancyId)
	if info.Date != "" {
		db = filterDate(db, info.Date, "")
	}
	if info.IsReply > 0 {
		db = db.Where("is_reply = ?", info.IsReply)
	}
	if info.Nickname != "" {
		userIds, err := GetUserIdsByNickname(info.Nickname, tenancyId)
		if err != nil {
			return productReplyList, total, err
		}
		db = db.Where("sys_user_id in ?", userIds)
	}
	if info.Keyword != "" {
		productIds, err := GetProductIdsByKeyword(info.Keyword, tenancyId)
		if err != nil {
			return productReplyList, total, err
		}
		db = db.Where("product_id in ?", productIds)
	}
	err := db.Count(&total).Error
	if err != nil {
		return productReplyList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&productReplyList).Error
	if err != nil {
		return productReplyList, total, err
	}

	var productIds []uint
	var products []response.ProductForReply
	if len(productReplyList) > 0 {
		for _, productReply := range productReplyList {
			productIds = append(productIds, productReply.ProductID)
		}
		if len(productIds) > 0 {
			products, err = GetProductForReplysByIds(productIds, tenancyId)
			if err != nil {
				return productReplyList, total, err
			}
		}

		for i := 0; i < len(productReplyList); i++ {
			if len(products) > 0 {
				for _, product := range products {
					if productReplyList[i].ProductID == product.ID {
						productReplyList[i].StoreName = product.StoreName
						productReplyList[i].Image = product.Image
					}
				}
			}
			productReplyList[i].Images = strings.Split(productReplyList[i].Pics, ",")
		}

	}

	return productReplyList, total, err
}
