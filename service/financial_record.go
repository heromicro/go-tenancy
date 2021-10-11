package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

// GetExpressInfoList
func GetExpressInfoList(info request.ExpressPageInfo) ([]model.Express, int64, error) {
	expressList := []model.Express{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Express{})
	if info.Name != "" {
		db = db.Where(g.TENANCY_DB.Where("name like ?", info.Name+"%").Or("code like ?", info.Name+"%"))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return expressList, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&expressList).Error
	return expressList, total, err
}
