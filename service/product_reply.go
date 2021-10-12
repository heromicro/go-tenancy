package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service/scope"
)

func GetAdminReplyMap(ctx *gin.Context) (Form, error) {
	var form Form
	productIdProps := map[string]interface{}{
		"type":      "image",
		"maxLength": 1,
		"title":     "请选择商品",
		"src":       "/admin/setting/storeProduct?field=productId",
		"width":     "60%",
		"height":    "536px",
		"srcKey":    "src",
		"modal": map[string]interface{}{
			"modal": false,
		},
	}
	avatarProps := map[string]interface{}{
		"type":      "image",
		"maxLength": 1,
		"title":     "请选择用户头像",
		"src":       "/admin/setting/uploadPicture?field=avatar&type=1",
		"width":     "896px",
		"height":    "480px",
		"footer":    false,
		"modal": map[string]interface{}{
			"modal": false,
		},
	}
	picsProps := map[string]interface{}{
		"type":      "image",
		"maxLength": 6,
		"title":     "请选择评价图片",
		"src":       "/admin/setting/uploadPicture?field=pic&type=2",
		"width":     "896px",
		"height":    "480px",
		"spin":      false,
		"modal": map[string]interface{}{
			"modal": false,
		},
	}
	commentProps := map[string]interface{}{
		"type":        "textarea",
		"placeholder": "请输入评价文字",
	}
	form = Form{Method: "POST", Title: "添加虚拟评价"}
	form.AddRule(*NewFrame("商品", "productId", "", "").AddProps(productIdProps)).
		AddRule(*NewInput("用户名称", "nickname", "请输入用户名称", "")).
		AddRule(*NewInput("评价文字", "comment", "", "").AddProps(commentProps)).
		AddRule(*NewRate("商品分数", "productScore", 8, 5)).
		AddRule(*NewRate("物流分数", "postageScore", 8, 5)).
		AddRule(*NewRate("服务分数", "serviceScore", 8, 5)).
		AddRule(*NewFrame("用户头像", "avatar", "", "").AddProps(avatarProps)).
		AddRule(*NewFrame("评价图片", "pic", "", []string{}).AddProps(picsProps))
	form.SetAction("/productReply/reply", ctx)
	return form, nil
}

func GetReplyMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := `{"rule":[{"type":"input","field":"content","value":"","title":"回复内容","props":{"type":"textarea","placeholder":"请输入回复内容"},"validate":[{"message":"请输入回复内容","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"评价回复","config":{}}`
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/productReply/reply", id), ctx)
	return form, err
}

func AddReply(id uint, content string) error {
	err := g.TENANCY_DB.Model(&model.ProductReply{}).Where("id = ?", id).Updates(map[string]interface{}{"merchant_reply_content": content, "merchant_reply_time": time.Now()}).Error
	if err != nil {
		return err
	}
	return nil
}
func AddFictiReply(req request.AddFictiReply) (uint, error) {
	reply := model.ProductReply{BaseProductReply: req.BaseProductReply, ProductId: req.ProductId.Id}
	reply.Pics = strings.Join(req.Pic, ",")
	err := g.TENANCY_DB.Model(&model.ProductReply{}).Create(&reply).Error
	if err != nil {
		return reply.ID, err
	}
	return reply.ID, nil
}

// GetProductReplyInfoList
func GetProductReplyInfoList(info request.ProductReplyPageInfo, tenancyId uint, isAdmin bool) ([]response.ProductReplyList, int64, error) {
	productReplyList := []response.ProductReplyList{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.ProductReply{})
	if isAdmin {
		// 管理員评论，
		db = db.Where("order_product_id = ?", 0).Where("c_user_id = ?", 0).Where("sys_tenancy_id =?", 0)
	}

	if info.Date != "" {
		db = db.Scopes(scope.FilterDate(info.Date, "created_at", ""))
	}
	if info.IsReply > 0 {
		db = db.Where("is_reply = ?", info.IsReply)
	}
	if info.Nickname != "" {
		userIds, err := GetUserIdsByNickname(info.Nickname, tenancyId)
		if err != nil {
			return productReplyList, total, err
		}
		db = db.Where("c_user_id in ?", userIds)
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
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&productReplyList).Error
	if err != nil {
		return productReplyList, total, err
	}

	productIds := []uint{}
	products := []response.ProductForReply{}
	if len(productReplyList) > 0 {
		for _, productReply := range productReplyList {
			productIds = append(productIds, productReply.ProductId)
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
					if productReplyList[i].ProductId == product.ID {
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

func DeleteProductReply(id uint) error {
	err := g.TENANCY_DB.Where("id = ?", id).Delete(&model.ProductReply{}).Error
	if err != nil {
		return err
	}
	return nil
}
