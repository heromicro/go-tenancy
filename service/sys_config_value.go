package service

import (
	"errors"
	"fmt"
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

func GetConfigValueByKey(configKey string) (string, error) {
	configValue := model.SysConfigValue{}
	err := g.TENANCY_DB.Where("config_key = ?", configKey).First(&configValue).Error
	if err != nil {
		return "", err
	}
	return configValue.Value, nil
}

func GetSeitURL() (string, error) {
	return GetConfigValueByKey("site_url")
}

func GetSeitName() (string, error) {
	return GetConfigValueByKey("site_name")
}

func GetAliPayConfig() (map[string]string, error) {
	config := map[string]string{}
	alipayConfigs, err := GetConfigByCateKey("alipay", 0)
	if err != nil {
		return config, err
	}

	for _, alipayConfig := range alipayConfigs {
		config[alipayConfig.ConfigKey] = alipayConfig.Value
	}
	if config["alipay_open"] == "0" {
		return config, fmt.Errorf("支付宝支付未开启")
	}
	return config, nil
}

func GetWechatPayConfig() (map[string]string, error) {
	config := map[string]string{}
	wechatConfigs, err := GetConfigByCateKey("wechat_payment", 0)
	if err != nil {
		return config, err
	}

	for _, wechatConfig := range wechatConfigs {
		config[wechatConfig.ConfigKey] = wechatConfig.Value
	}

	if config["pay_weixin_open"] == "0" {
		return config, fmt.Errorf("微信支付未开启")
	}

	return config, nil
}
