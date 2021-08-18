package param

import "fmt"

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
