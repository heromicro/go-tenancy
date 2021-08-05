package service

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/chindeo/pkg/file"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	v2 "github.com/go-pay/gopay/wechat"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils"
	"gorm.io/gorm"
)

func PayOrder(req request.PayOrder) (response.PayOrder, error) {
	var res response.PayOrder
	tenancy, err := GetTenancyByID(req.TenancyId)
	if err != nil {
		return res, fmt.Errorf("商户参数错误")
	}
	if tenancy.Status == g.StatusFalse {
		return res, fmt.Errorf("当前商户已被冻结")
	}
	if tenancy.State == g.StatusFalse {
		return res, fmt.Errorf("当前商户已经停业")
	}
	order, err := GetOrderByOrderIdUserIdAndTenancyId(req.OrderId, req.TenancyId, req.UserId, req.OrderType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, fmt.Errorf("当前订单不存在")
	} else if err != nil {
		return res, err
	}
	//已经支付订单不能重复支付
	if order.Paid == g.StatusTrue && !order.PayTime.IsZero() && order.PayType > model.PayTypeUnknown && order.Status > model.OrderStatusNoPay {
		return res, fmt.Errorf("当前订单已经支付，请勿重复支付")
	} else if time.Since(order.CreatedAt).Minutes() >= 15 {
		return res, fmt.Errorf("当前支付订单已超时，请重新下单")
	}

	if strings.Contains(req.UserAgent, "MicroMessenger") {
		return WechatPay(order, tenancy.Name, req.OpenId)
	} else if strings.Contains(req.UserAgent, "Alipay") {
		return Alipay(order, tenancy.Name)
	} else {
		return res, fmt.Errorf("请使用微信或者支付宝扫描支付")
	}
}

func WechatPay(order model.Order, tenancyName, openid string) (response.PayOrder, error) {
	var res response.PayOrder
	siteName, err := GetSeitName()
	if err != nil {
		return res, fmt.Errorf("获取商城名称 %w", err)
	}
	wechatConf, err := GetWechatPayConfig()
	if err != nil {
		return res, err
	}

	pKContent, err := file.ReadString(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, wechatConf.PayWeixinClientKey))
	if err != nil {
		return res, fmt.Errorf("获取支付 %w", err)
	}

	if g.TENANCY_CONFIG.WechatPay.WxPkContent == "" || g.TENANCY_CONFIG.WechatPay.WxPkSerialNo == "" {
		certs, err := wechat.GetPlatformCerts(wechatConf.PayWeixinMchid, wechatConf.PayWeixinKey, wechatConf.PaySerialNo, pKContent)
		if err != nil {
			return res, fmt.Errorf("获取微信支付平台证书错误 %w", err)
		}
		if len(certs.Certs) == 1 {
			g.TENANCY_CONFIG.WechatPay.WxPkContent = certs.Certs[0].PublicKey
			g.TENANCY_CONFIG.WechatPay.WxPkSerialNo = certs.Certs[0].SerialNo
		}
	}

	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：apiV3Key，商户平台获取
	//	pkContent：私钥 apiclient_key.pem 读取后的内容
	client, err := wechat.NewClientV3(wechatConf.PayWeixinAppid, wechatConf.PayWeixinMchid, wechatConf.PaySerialNo, wechatConf.PayWeixinKey, pKContent)
	if err != nil {
		return res, fmt.Errorf("初始化微信支付错误 %w", err)
	}

	// 设置微信平台证书和序列号，并启用自动同步返回验签
	//	注意：请预先通过 wechat.GetPlatformCerts() 获取并维护微信平台证书和证书序列号
	client.SetPlatformCert([]byte(g.TENANCY_CONFIG.WechatPay.WxPkContent), g.TENANCY_CONFIG.WechatPay.WxPkSerialNo).AutoVerifySign()

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOff
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	notifyUrl, err := GetPayNotifyUrl()
	if err != nil {
		return res, err
	}

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("description", fmt.Sprintf("%s-%s", tenancyName, siteName)).
		Set("out_trade_no", order.OrderSn).
		Set("time_expire", expire).
		Set("notify_url", notifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", getOrderPrice(order.PayPrice)*100).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", openid)
		})

	wxRsp, err := client.V3TransactionJsapi(bm)
	if err != nil {
		return res, fmt.Errorf("transaction jsapi 错误 %w", err)
	}
	if wxRsp.Code > 0 {
		return res, fmt.Errorf("%s", wxRsp.Error)
	}
	jsapi, err := client.PaySignOfJSAPI(wxRsp.Response.PrepayId)
	if err != nil {
		return res, fmt.Errorf("微信支付 jsapi 签名错误 %w", err)
	}

	res.JSAPIPayParams = jsapi
	return res, nil
}

func getOrderPrice(price float64) float64 {
	seitMode, err := GetSeitMode()
	if err != nil {
		return price
	}

	if seitMode == "0" {
		return 0.01
	}

	return price
}

// Alipay
func Alipay(order model.Order, tenancyName string) (response.PayOrder, error) {
	var res response.PayOrder
	siteName, err := GetSeitName()
	if err != nil {
		return res, fmt.Errorf("获取商城名称 %w", err)
	}
	//请求参数
	body := make(gopay.BodyMap)
	body.Set("subject", fmt.Sprintf("%s-%s", tenancyName, siteName))
	body.Set("out_trade_no", order.OrderSn)
	// body.Set("quit_url", "https://www.fmm.ink") //用户付款中途退出返回商户网站的地址

	body.Set("total_amount", getOrderPrice(order.PayPrice))
	body.Set("product_code", "QUICK_WAP_WAY") //商家和支付宝签约的产品码
	//手机网站支付请求
	client, err := AliPayClient()
	if err != nil {
		return res, err
	}
	payUrl, err := client.TradeWapPay(body)
	if err != nil {
		return res, err
	}
	res.AliPayUrl = payUrl
	return res, nil
}

func AliPayClient() (*alipay.Client, error) {
	alipayConf, err := GetAliPayConfig()
	if err != nil {
		return nil, fmt.Errorf("获取支付宝配置错误 %w", err)
	}
	notifyUrl, err := GetPayNotifyUrl()
	if err != nil {
		return nil, err
	}

	client := alipay.NewClient(alipayConf.AlipayAppId, alipayConf.AlipayPrivateKey, alipayConf.AlipayEnv)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetPrivateKeyType(alipay.PKCS1).
		SetNotifyUrl(notifyUrl)
		//SetReturnUrl("https://www.fmm.ink").
		// if !g.TENANCY_CONFIG.Alipay.IsProd {
		// 	client.SetAliPayRootCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayRootCert.crt")).
		// 		SetAppCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "appCertPublicKey_2021000117637854.crt")).
		// 		SetAliPayPublicCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayCertPublicKey_RSA2.crt"))
		// }
	return client, nil
}

func GetAutoCode(redirectUri string) (string, error) {
	seitUrl, err := GetSeitURL()
	if err != nil {
		return "", fmt.Errorf("获取站点url错误 %w", err)
	}
	wechatConf, err := GetWechatPayConfig()
	if err != nil {
		return "", err
	}
	url := utils.GetAutoCode(wechatConf.PayWeixinAppid, seitUrl+redirectUri, "snsapi_base", g.TENANCY_CONFIG.WechatPay.State)
	if err != nil {
		return "", fmt.Errorf("微信网页授权失败 %w", err)
	}
	return url, nil
}

func GetOpenId(code string) (string, error) {
	wechatConf, err := GetWechatPayConfig()
	if err != nil {
		return "", err
	}
	accessToken, err := v2.GetOauth2AccessToken(wechatConf.PayWeixinAppid, wechatConf.PayWeixinAppsecret, code)
	if err != nil {
		return "", fmt.Errorf("微信获取openid失败 %w", err)
	}
	if accessToken.Errcode > 0 {
		return "", fmt.Errorf("%s", accessToken.Errmsg)
	}
	return accessToken.Openid, nil
}
