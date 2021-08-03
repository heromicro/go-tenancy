package core

import (
	"github.com/go-pay/gopay/alipay"
	"github.com/snowlyg/go-tenancy/g"
)

func AliPay() *alipay.Client {
	client := alipay.NewClient(g.TENANCY_CONFIG.Alipay.AppId, g.TENANCY_CONFIG.Alipay.PrivateKey, g.TENANCY_CONFIG.Alipay.IsProd)
	//配置公共参数
	client.SetCharset(g.TENANCY_CONFIG.Alipay.Charset).
		SetSignType(g.TENANCY_CONFIG.Alipay.SignType).
		//SetReturnUrl("https://www.fmm.ink").
		SetNotifyUrl(g.TENANCY_CONFIG.Alipay.NotifyUrl)
	// if !g.TENANCY_CONFIG.Alipay.IsProd {
	// 	client.SetAliPayRootCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayRootCert.crt")).
	// 		SetAppCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "appCertPublicKey_2021000117637854.crt")).
	// 		SetAliPayPublicCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayCertPublicKey_RSA2.crt"))
	// }
	if g.TENANCY_CONFIG.Alipay.PrivateKeyType == "PKCS1" {
		client.SetPrivateKeyType(alipay.PKCS1)
	} else if g.TENANCY_CONFIG.Alipay.PrivateKeyType == "PKCS8" {
		client.SetPrivateKeyType(alipay.PKCS8)
	}
	return client
}
