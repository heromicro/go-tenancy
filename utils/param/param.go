package param

import (
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

func GetConfigValueByKey(configKey string) (string, error) {
	configValue := model.SysConfigValue{}
	err := g.TENANCY_DB.Where("config_key = ?", configKey).First(&configValue).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	return configValue.Value, nil
}

// GetSeitURL 网站地址
func GetSeitURL() (string, error) {
	return GetConfigValueByKey("site_url")
}

// 支付回调地址
func GetPayNotifyUrl(notifyType string) (string, error) {
	seitUrl, err := GetSeitURL()
	if err != nil {
		return "", fmt.Errorf("获取站点url错误 %w", err)
	}
	return fmt.Sprintf("%s/%s/%s", seitUrl, "/v1/pay/notify", notifyType), nil
}

// GetTenancyDefaultPassword 商户默认密码
func GetTenancyDefaultPassword() (string, error) {
	return GetConfigValueByKey("tenancy_admin_password")
}

// GetSeitName 网站名称
func GetSeitName() (string, error) {
	return GetConfigValueByKey("site_name")
}

// 网站模式
func GetSeitMode() (string, error) {
	return GetConfigValueByKey("site_open")
}

// GetConfigByCateKey
func GetConfigByCateKey(configKey string, tenancyId uint) ([]response.SysConfig, error) {
	var configs []response.SysConfig
	err := g.TENANCY_DB.Model(&model.SysConfig{}).
		Select("sys_configs.*").
		Joins("left join sys_config_categories on sys_configs.sys_config_category_id = sys_config_categories.id").
		Where("sys_config_categories.key = ?", configKey).
		Where("sys_configs.status = ?", g.StatusTrue).
		Find(&configs).Error
	if err != nil {
		return nil, err
	}

	var values []model.SysConfigValue
	err = g.TENANCY_DB.Where("sys_tenancy_id = ?", tenancyId).Find(&values).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(configs); i++ {
		for _, value := range values {
			if configs[i].ConfigKey == value.ConfigKey {
				configs[i].Value = value.Value
				break
			}
		}
	}

	return configs, err
}

// GetConfigByCateKeys
func GetConfigByCateKeys(configKeys []string, tenancyId uint) ([]response.SysConfig, error) {
	var configs []response.SysConfig
	err := g.TENANCY_DB.Model(&model.SysConfig{}).
		Select("sys_configs.*,sys_config_categories.key as cate_key").
		Joins("left join sys_config_categories on sys_configs.sys_config_category_id = sys_config_categories.id").
		Where("sys_config_categories.key in ?", configKeys).
		Where("sys_configs.status = ?", g.StatusTrue).
		Find(&configs).Error
	if err != nil {
		return nil, err
	}

	var values []model.SysConfigValue
	err = g.TENANCY_DB.Where("sys_tenancy_id = ?", tenancyId).Find(&values).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(configs); i++ {
		for _, value := range values {
			if configs[i].ConfigKey == value.ConfigKey {
				configs[i].Value = value.Value
				break
			}
		}
	}

	return configs, err
}
