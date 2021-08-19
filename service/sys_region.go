package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
)

// GetRegion 获取行政区域
func GetRegion(pCode string) ([]response.SysRegion, error) {
	regions := []response.SysRegion{}
	err := g.TENANCY_DB.Where("p_code", pCode).Preload("SubRegions.SubRegions").Find(&regions).Error
	return regions, err
}

// GetRegionList 获取行政区域
func GetRegionList() ([]response.SysRegion, error) {
	regions := []response.SysRegion{}
	err := g.TENANCY_DB.Where("p_code = ?", 0).Preload("SubRegions.SubRegions").Find(&regions).Error
	return regions, err
}
