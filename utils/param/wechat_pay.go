package param

import "fmt"

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
