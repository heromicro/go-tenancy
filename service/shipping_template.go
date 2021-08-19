package service

import (
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
)

// CreateShippingTemplate
func CreateShippingTemplate(shippingTem model.ShippingTemplate, tenancyId uint) (uint, error) {
	shippingTem.SysTenancyID = tenancyId
	err := g.TENANCY_DB.Create(&shippingTem).Error
	return shippingTem.ID, err
}

// GetShippingTemplateByID
func GetShippingTemplateByID(id uint) (response.ShippingTemplateDetail, error) {
	var shippingTem response.ShippingTemplateDetail
	err := g.TENANCY_DB.Model(&model.ShippingTemplate{}).
		Where("id = ?", id).
		First(&shippingTem).Error
	return shippingTem, err
}

// UpdateShippingTemplate
func UpdateShippingTemplate(req model.ShippingTemplate, id uint) error {
	shipTemp := map[string]interface{}{
		"name":       req.Name,
		"type":       req.Type,
		"appoint":    req.Appoint,
		"undelivery": req.Undelivery,
		"is_default": req.IsDefault,
		"sort":       req.Sort,
	}
	if err := g.TENANCY_DB.Model(&model.ShippingTemplate{}).Where("id = ?", id).Updates(&shipTemp).Error; err != nil {
		return err
	}
	return nil
}

// DeleteShippingTemplate
func DeleteShippingTemplate(id uint) error {
	var product model.ShippingTemplate
	return g.TENANCY_DB.Where("id = ?", id).Delete(&product).Error
}

// GetShippingTemplateInfoList
func GetShippingTemplateInfoList(info request.ShippingTemplatePageInfo) ([]response.ShippingTemplateList, int64, error) {
	var shippingTemList []response.ShippingTemplateList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.ShippingTemplate{})

	if info.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%s%%", info.Name))
	}

	err := db.Count(&total).Error
	if err != nil {
		return shippingTemList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&shippingTemList).Error

	return shippingTemList, total, err
}

// GetShippingTemplateInfoSelect
func GetShippingTemplateInfoSelect() ([]response.SelectOption, error) {
	var shippingTemList []response.SelectOption
	err := g.TENANCY_DB.Model(&model.ShippingTemplate{}).Select("id,name").Find(&shippingTemList).Error
	if err != nil {
		return nil, err
	}
	return shippingTemList, err
}
