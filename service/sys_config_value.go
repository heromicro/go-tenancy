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

func GetPayNotifyUrl(notifyType string) (string, error) {
	seitUrl, err := GetSeitURL()
	if err != nil {
		return "", fmt.Errorf("获取站点url错误 %w", err)
	}
	return fmt.Sprintf("%s/%s/%s", seitUrl, "v1/pay/notify", notifyType), nil
}

func GetSeitName() (string, error) {
	return GetConfigValueByKey("site_name")
}

func GetSeitMode() (string, error) {
	return GetConfigValueByKey("site_open")
}

type AlipayConfig struct {
	AlipayOpen       bool
	AlipayEnv        bool
	AlipayAppId      string
	AlipayPrivateKey string
	AlipayPublicKey  string
}

func GetAliPayConfig() (AlipayConfig, error) {
	config := AlipayConfig{}
	alipayConfigs, err := GetConfigByCateKey("alipay", 0)
	if err != nil {
		return config, err
	}

	for _, alipayConfig := range alipayConfigs {
		if alipayConfig.ConfigKey == "alipay_open" {
			if alipayConfig.Value == "1" {
				config.AlipayOpen = true
			}
		}
		if alipayConfig.ConfigKey == "alipay_env" {
			if alipayConfig.Value == "1" {
				config.AlipayEnv = true
			}
		}
		if alipayConfig.ConfigKey == "alipay_app_id" {
			config.AlipayAppId = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "alipay_private_key" {
			config.AlipayPrivateKey = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "alipay_public_key" {
			config.AlipayPublicKey = alipayConfig.Value
		}
	}
	if !config.AlipayOpen {
		return config, fmt.Errorf("支付宝支付未开启")
	}
	return config, nil
}

type WechatConfig struct {
	PayWeixinAppid      string
	PayWeixinAppsecret  string
	PayWeixinMchid      string
	PayWeixinClientCert string
	PayWeixinClientKey  string
	PaySerialNo         string
	PayWeixinKey        string
	PayWeixinOpen       bool
}

func GetWechatPayConfig() (WechatConfig, error) {
	config := WechatConfig{}
	wechatConfigs, err := GetConfigByCateKey("wechat_payment", 0)
	if err != nil {
		return config, fmt.Errorf("获取微信支付配置错误 %w", err)
	}

	for _, wechatConfig := range wechatConfigs {
		if wechatConfig.ConfigKey == "pay_weixin_open" {
			if wechatConfig.Value == "1" {
				config.PayWeixinOpen = true
			}
		}
		if wechatConfig.ConfigKey == "pay_weixin_key" {
			config.PayWeixinKey = wechatConfig.Value
		}
		if wechatConfig.ConfigKey == "pay_weixin_appid" {
			config.PayWeixinAppid = wechatConfig.Value
		}
		if wechatConfig.ConfigKey == "pay_weixin_appsecret" {
			config.PayWeixinAppsecret = wechatConfig.Value
		}
		if wechatConfig.ConfigKey == "pay_weixin_mchid" {
			config.PayWeixinMchid = wechatConfig.Value
		}
		if wechatConfig.ConfigKey == "pay_weixin_client_cert" {
			config.PayWeixinClientCert = wechatConfig.Value
		}
		if wechatConfig.ConfigKey == "pay_weixin_client_key" {
			config.PayWeixinClientKey = wechatConfig.Value
		}
		if wechatConfig.ConfigKey == "pay_serial_no" {
			config.PaySerialNo = wechatConfig.Value
		}
	}

	if !config.PayWeixinOpen {
		return config, fmt.Errorf("微信支付未开启")
	}

	return config, nil
}
