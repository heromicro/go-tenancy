package service

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

// SaveConfigValue
// TODO: configKey 没有使用，可用于过滤参数
func SaveConfigValue(values map[string]interface{}, configKey string, tenancyId uint) error {
	for key, value := range values {
		var val string
		typeName := reflect.TypeOf(value).Name()
		if typeName == "string" {
			val = value.(string)
		} else if typeName == "int" {
			val = strconv.FormatInt(int64(value.(int)), 10)
		} else if typeName == "float64" {
			val = strconv.FormatFloat(value.(float64), 'E', -1, 64)
		}
		configValue := model.SysConfigValue{}
		err := g.TENANCY_DB.Where("config_key = ?", key).Where("sys_tenancy_id = ?", tenancyId).First(&configValue).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			configValue.Value = val
			g.TENANCY_DB.Save(configValue)
		} else {
			g.TENANCY_DB.Create(&model.SysConfigValue{ConfigKey: key, Value: val, SysTenancyID: tenancyId})
		}
	}
	return nil
}
