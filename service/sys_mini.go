package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils"
	"gorm.io/gorm"
)

// CreateMini
func CreateMini(m request.CreateSysMini) (uint, error) {
	var mini model.SysMini
	err := g.TENANCY_DB.Where("name = ?", m.Name).First(&mini).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return mini.ID, errors.New("商户名称已被注冊")
	}
	mini.UUID = utils.UUIDV5()
	mini.Name = m.Name
	mini.AppID = m.AppID
	mini.AppSecret = m.AppSecret
	mini.Remark = m.Remark
	err = g.TENANCY_DB.Create(&mini).Error
	return mini.ID, err
}

// GetMiniByID
func GetMiniByID(id uint) (model.SysMini, error) {
	var mini model.SysMini
	err := g.TENANCY_DB.Where("id = ?", id).First(&mini).Error
	return mini, err
}

// UpdateMini
func UpdateMini(id uint, m request.UpdateSysMini) error {
	var mini model.SysMini
	err := g.TENANCY_DB.Where("name = ?", m.Name).Not("id = ?", id).First(&mini).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("商户名称已被注冊")
	}

	data := map[string]interface{}{
		"name":       m.Name,
		"app_id":     m.AppID,
		"app_secret": m.AppSecret,
		"remark":     m.Remark,
	}
	err = g.TENANCY_DB.Model(&model.SysMini{}).Where("id =?", id).Updates(&data).Error
	return err
}

// DeleteMini
func DeleteMini(id uint) error {
	var mini model.SysMini
	return g.TENANCY_DB.Where("id = ?", id).Delete(&mini).Error
}

// GetMiniInfoList
func GetMiniInfoList(info request.PageInfo) ([]response.SysMini, int64, error) {
	miniList := []response.SysMini{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysMini{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return miniList, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&miniList).Error
	return miniList, total, err
}
