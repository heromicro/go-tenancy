package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

// GetUserBillInfoList
func GetUserBillInfoList(info request.PageInfo, c_user_id uint) ([]model.UserBill, int64, error) {
	userBillList := []model.UserBill{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.UserBill{}).Where("c_user_id = ?", c_user_id)
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return userBillList, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&userBillList).Error
	return userBillList, total, err
}
