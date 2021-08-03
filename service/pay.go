package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

func PayOrder(req request.PayOrder, userAgent, tenancyName string) (string, error) {
	tenancy, err := GetTenancyByID(req.TenancyId)
	if err != nil {
		return "", fmt.Errorf("商户参数错误")
	}
	if tenancy.Status == g.StatusFalse {
		return "", fmt.Errorf("当前商户已被冻结")
	}
	if tenancy.State == g.StatusFalse {
		return "", fmt.Errorf("当前商户已经停业")
	}
	order, err := GetOrderByOrderIdUserIdAndTenancyId(req.OrderId, req.TenancyId, req.UserId, req.OrderType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("当前订单不存在")
	} else if err != nil {
		return "", err
	}
	//已经支付订单不能重复支付
	if order.Paid == g.StatusTrue && !order.PayTime.IsZero() && order.PayType > model.PayTypeUnknown && order.Status > model.OrderStatusNoPay {
		return "", fmt.Errorf("当前订单已经支付，请勿重复支付")
	} else if time.Since(order.CreatedAt).Minutes() >= 15 {
		return "", fmt.Errorf("当前支付订单已超时，请重新下单")
	}

	if strings.Contains(userAgent, "MicroMessenger") {
		fmt.Println("wechat")
	} else if strings.Contains(userAgent, "Alipay") {
		return Alipay(order, tenancy.Name)
	}

	return "", nil
}

// Alipay
func Alipay(order model.Order, tenancyName string) (string, error) {
	siteName, err := GetSeitName()
	if err != nil {
		return "", err
	}
	//请求参数
	subject := fmt.Sprintf("%s-%s", tenancyName, siteName)
	body := make(gopay.BodyMap)
	body.Set("subject", subject)
	body.Set("out_trade_no", order.OrderSn)
	// body.Set("quit_url", "https://www.fmm.ink") //用户付款中途退出返回商户网站的地址

	// 支付价格测试和开发使用 0.01
	payPrice := order.PayPrice
	if g.TENANCY_CONFIG.System.Env != "pro" {
		payPrice = 0.01
	}
	body.Set("total_amount", payPrice)
	body.Set("product_code", "QUICK_WAP_WAY") //商家和支付宝签约的产品码
	//手机网站支付请求
	client, err := AliPayClient()
	if err != nil {
		return "", err
	}
	payUrl, err := client.TradeWapPay(body)
	if err != nil {
		return "", err
	}
	return payUrl, nil
}

func AliPayClient() (*alipay.Client, error) {
	alipayConf, err := GetAliPay()
	if err != nil {
		return nil, fmt.Errorf("获取支付宝配置错误 %w", err)
	}
	if alipayConf["alipay_open"] == "0" {
		return nil, fmt.Errorf("支付宝支付未开启")
	}
	seitUrl, err := GetSeitURL()
	if err != nil {
		return nil, fmt.Errorf("获取站点url错误 %w", err)
	}
	var isProd bool
	if alipayConf["alipay_env"] == "0" {
		isProd = false
	} else if alipayConf["alipay_env"] == "1" {
		isProd = true
	}
	fmt.Println(alipayConf)
	client := alipay.NewClient(alipayConf["alipay_app_id"], alipayConf["alipay_private_key"], isProd)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetPrivateKeyType(alipay.PKCS1).
		//SetReturnUrl("https://www.fmm.ink").
		SetNotifyUrl(fmt.Sprintf("%s/%s", seitUrl, ""))
		// if !g.TENANCY_CONFIG.Alipay.IsProd {
		// 	client.SetAliPayRootCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayRootCert.crt")).
		// 		SetAppCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "appCertPublicKey_2021000117637854.crt")).
		// 		SetAliPayPublicCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayCertPublicKey_RSA2.crt"))
		// }
	return client, nil
}
