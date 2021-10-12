package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service/scope"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GettFinancialRecordInfoList
func GettFinancialRecordInfoList(info request.FinancialRecordPageInfo) ([]model.FinancialRecord, int64, error) {
	financialRecord := []model.FinancialRecord{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var total int64
	db := g.TENANCY_DB.Model(&model.FinancialRecord{})
	if info.Keyword != "" {
		db = db.Where(g.TENANCY_DB.Where("order_sn like ?", info.Keyword+"%").Or("user_info like ?", info.Keyword+"%"))
	}
	if info.Date != "" {
		db = db.Scopes(scope.FilterDate(info.Date, "created_at", ""))
	}
	err := db.Count(&total).Error
	if err != nil {
		return financialRecord, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&financialRecord).Error
	return financialRecord, total, err
}

// CreateFinancialRecord 添加流水记录
func CreateFinancialRecord(db *gorm.DB, financialRecord *model.FinancialRecord) error {
	if err := db.Create(financialRecord).Error; err != nil {
		g.TENANCY_LOG.Error("添加流水记录", zap.String("CreateFinancialRecord()", err.Error()))
		return err
	}
	return nil
}
