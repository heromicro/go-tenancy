package service

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/chindeo/pkg/file"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	v2 "github.com/go-pay/gopay/wechat"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/go-tenancy/utils/param"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func PayOrder(req request.PayOrder) (response.PayOrder, error) {
	// 清除测试缓存
	defer DeleteTestCache(req.TenancyId, req.UserId)
	var res response.PayOrder
	autoCloseTime := param.GetOrderAutoCloseTime()
	if time.Until(time.Unix(req.Expire, 0)).Minutes() > float64(autoCloseTime) {
		g.TENANCY_LOG.Error("支付二维码已经过期", zap.Float64("sub:", time.Until(time.Unix(req.Expire, 0)).Minutes()), zap.Int64("过期时间（分钟）:", autoCloseTime))
		return res, fmt.Errorf("支付二维码已经过期，请重新下单")
	}
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
	order, err := GetOrderById(req.OrderId)
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

// RefundOrder 退款订单
func RefundOrder(order model.Order, refundOrder model.RefundOrder, refundPrice float64) error {
	if order.Paid != g.StatusTrue || order.Status == model.OrderStatusNoPay {
		return errors.New("订单未支付")
	}

	if order.PayType == model.PayTypeWx {
		return WechatRefund(order.OrderSn, refundOrder.RefundOrderSn, refundOrder.RefundMessage, order.PayPrice, refundPrice)
	} else if order.PayType == model.PayTypeAlipay {
		return AliRefund(order.OrderSn, refundOrder.RefundMessage, refundPrice)
	}

	return nil
}

// getWechatPayClient 微信支付客户端
func getWechatPayClient() (client *wechat.ClientV3, err error) {
	wechatConf, err := param.GetWechatPayConfig()
	if err != nil {
		return nil, err
	}

	pKContent, err := file.ReadString(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, wechatConf.PayWeixinClientKey))
	if err != nil {
		return nil, fmt.Errorf("获取支付 %w", err)
	}

	if g.TENANCY_CONFIG.WechatPay.WxPkContent == "" || g.TENANCY_CONFIG.WechatPay.WxPkSerialNo == "" {
		certs, err := wechat.GetPlatformCerts(wechatConf.PayWeixinMchid, wechatConf.PayWeixinKey, wechatConf.PaySerialNo, pKContent)
		if err != nil {
			return nil, fmt.Errorf("获取微信支付平台证书错误 %w", err)
		}
		if len(certs.Certs) == 1 {
			g.TENANCY_CONFIG.WechatPay.WxPkContent = certs.Certs[0].PublicKey
			g.TENANCY_CONFIG.WechatPay.WxPkSerialNo = certs.Certs[0].SerialNo
		}
	}

	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：apiV3Key，商户平台获取
	//	pkContent：私钥 apiclient_key.pem 读取后的内容
	client, err = wechat.NewClientV3(wechatConf.PayWeixinAppid, wechatConf.PayWeixinMchid, wechatConf.PaySerialNo, wechatConf.PayWeixinKey, pKContent)
	if err != nil {
		return nil, fmt.Errorf("初始化微信支付错误 %w", err)
	}

	// 设置微信平台证书和序列号，并启用自动同步返回验签
	//	注意：请预先通过 wechat.GetPlatformCerts() 获取并维护微信平台证书和证书序列号
	client.SetPlatformCert([]byte(g.TENANCY_CONFIG.WechatPay.WxPkContent), g.TENANCY_CONFIG.WechatPay.WxPkSerialNo).AutoVerifySign()

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOff
	return client, nil
}

// WechatPay 微信支付
func WechatPay(order model.Order, tenancyName, openid string) (response.PayOrder, error) {
	var res response.PayOrder
	siteName, err := param.GetSeitName()
	if err != nil {
		return res, fmt.Errorf("获取商城名称 %w", err)
	}

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	notifyUrl, err := param.GetPayNotifyUrl("wechat")
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

	client, err := getWechatPayClient()
	if err != nil {
		return res, err
	}

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

// WechatRefund 微信
func WechatRefund(orderSn, refundOrderSn, refundMessage string, totalPrice, refundPrice float64) error {
	client, err := getWechatPayClient()
	if err != nil {
		return err
	}
	notifyUrl, err := param.GetPayNotifyUrl("wechat_return")
	if err != nil {
		return err
	}
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderSn).
		Set("reason", refundMessage).
		Set("out_refund_no", refundOrderSn).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("refund", getOrderPrice(refundPrice)*100).
				Set("total", getOrderPrice(totalPrice)*100).
				Set("currency", "CNY")
		}).
		Set("notify_url", notifyUrl)

	//请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	wxRsp, err := client.V3Refund(bm)
	if err != nil {
		return err
	}

	if wxRsp.Code != 0 {
		return fmt.Errorf("微信支付退款失败 %s", wxRsp.Error)

	}
	g.TENANCY_LOG.Info("微信支付退款", zap.String("aliRsp", fmt.Sprintf("%+v", wxRsp.Response)))
	return nil
}

// getOrderPrice 计算订单价格
func getOrderPrice(price float64) float64 {
	seitMode, err := param.GetSeitMode()
	if err != nil {
		return price
	}

	if seitMode == "0" {
		return 0.01
	}

	return price
}

// Alipay 支付宝支付
func Alipay(order model.Order, tenancyName string) (response.PayOrder, error) {
	var res response.PayOrder
	siteName, err := param.GetSeitName()
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
	client, err := AliPayClient()
	if err != nil {
		return res, err
	}
	//手机网站支付请求
	payUrl, err := client.TradeWapPay(body)
	if err != nil {
		return res, err
	}
	res.AliPayUrl = payUrl
	return res, nil
}

//  AliRefund 支付宝退款
func AliRefund(orderSn, refundMessage string, refundPrice float64) error {
	client, err := AliPayClient()
	if err != nil {
		g.TENANCY_LOG.Error("支付宝退款错误", zap.String("AliPayClient()", err.Error()))
		return err
	}
	//请求参数
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", orderSn)
	body.Set("refund_amount", getOrderPrice(refundPrice))
	body.Set("refund_reason", refundMessage)

	aliRsp, err := client.TradeRefund(body)
	if err != nil {
		g.TENANCY_LOG.Error("支付宝退款错误", zap.String("client.TradeRefund()", err.Error()))
		return err
	}
	g.TENANCY_LOG.Debug("支付宝退款", zap.String("aliRsp", fmt.Sprintf("%+v", aliRsp.Response)))
	return nil
}

// AliPayClient 支付客户端
func AliPayClient() (*alipay.Client, error) {
	alipayConf, err := param.GetAliPayConfig()
	if err != nil {
		return nil, fmt.Errorf("获取支付宝配置错误 %w", err)
	}
	notifyUrl, err := param.GetPayNotifyUrl("ali")
	if err != nil {
		return nil, err
	}

	// 测试采用 PKCS1，正式使用 PKCS8
	pkcs := alipay.PKCS8
	if !alipayConf.AlipayEnv {
		pkcs = alipay.PKCS1
	}

	client := alipay.NewClient(alipayConf.AlipayAppId, alipayConf.AlipayPrivateKey, alipayConf.AlipayEnv)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetPrivateKeyType(pkcs).
		SetNotifyUrl(notifyUrl)
		//SetReturnUrl("https://www.fmm.ink").
		// if !g.TENANCY_CONFIG.Alipay.IsProd {
		// 	client.SetAliPayRootCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayRootCert.crt")).
		// 		SetAppCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "appCertPublicKey_2021000117637854.crt")).
		// 		SetAliPayPublicCertSN(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "alipayCertPublicKey_RSA2.crt"))
		// }
	return client, nil
}

//GetAutoCode 获取微信网页授权
func GetAutoCode(redirectUri string) (string, error) {
	seitUrl, err := param.GetSeitURL()
	if err != nil {
		return "", fmt.Errorf("获取站点url错误 %w", err)
	}
	wechatConf, err := param.GetWechatPayConfig()
	if err != nil {
		return "", err
	}
	url := utils.GetAutoCode(wechatConf.PayWeixinAppid, seitUrl+redirectUri, "snsapi_base", g.TENANCY_CONFIG.WechatPay.State)
	if err != nil {
		return "", fmt.Errorf("微信网页授权失败 %w", err)
	}
	return url, nil
}

// GetOpenId 获取用户openid
func GetOpenId(code string) (string, error) {
	wechatConf, err := param.GetWechatPayConfig()
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

// NotifyAliPay
// 支付宝异步通知回调
// total_amount=2.00&buyer_id=20****7&body=大乐透2.1&trade_no=2016071921001003030200089909&refund_fee=0.00&notify_time=2016-07-19 14:10:49&subject=大乐透2.1&sign_type=RSA2&charset=utf-8&notify_type=trade_status_sync&out_trade_no=0719141034-6418&gmt_close=2016-07-19 14:10:46&gmt_payment=2016-07-19 14:10:47&trade_status=TRADE_SUCCESS&version=1.0&sign=kPbQIjX+xQc8F0/A6/AocELIjhhZnGbcBN6G4MM/HmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo+lHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l/KL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM=&gmt_create=2016-07-19 14:10:44&app_id=20151*****3&seller_id=20881021****8&notify_id=4a91b7a78a503640467525113fb7d8bg8e
func NotifyAliPay(ctx *gin.Context) error {
	notifyReq, err := alipay.ParseNotifyToBodyMap(ctx.Request)
	if err != nil {
		return err
	}

	alipayConf, err := param.GetAliPayConfig()
	if err != nil {
		return err
	}
	// 支付宝异步通知验签（公钥模式）
	ok, err := alipay.VerifySign(alipayConf.AlipayPublicKey, notifyReq)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("支付宝异步通知回调验签失败 %w", err)
	}

	var orderSn, outBizNo, gmtRefund, tradeStatus, refundFee string
	if notifyReq["out_trade_no"] != nil {
		orderSn = notifyReq["out_trade_no"].(string) //商户订单号。原支付请求的商户订单号。
	}
	if notifyReq["out_biz_no"] != nil {
		outBizNo = notifyReq["out_biz_no"].(string) //商户业务号。商户业务 ID，主要是退款通知中返回退款申请的流水号。
	}
	if notifyReq["refund_fee"] != nil {
		refundFee = notifyReq["refund_fee"].(string) //总退款金额
	}
	if notifyReq["gmt_refund"] != nil {
		gmtRefund = notifyReq["gmt_refund"].(string) //交易退款时间
	}
	if notifyReq["trade_status"] != nil {
		tradeStatus = notifyReq["trade_status"].(string) //交易状态
	}
	g.TENANCY_LOG.Info("支付异步通知: 支付宝支付异步通知回调", zap.String("订单号", orderSn), zap.String("流水号", outBizNo), zap.String("总退款金额", refundFee), zap.String("交易退款时间", gmtRefund), zap.String("交易状态", tradeStatus))
	// 退款
	if outBizNo != "" && gmtRefund != "" {
		if tradeStatus != "TRADE_SUCCESS" && tradeStatus != "TRADE_CLOSED" {
			return fmt.Errorf("退款异步通知: %s 支付宝异步通知回调返回状态: %s", orderSn, tradeStatus)
		}

		// 部分退款
		if tradeStatus == "TRADE_SUCCESS" {
			err := ChangeReturnOrderStatusByReturnOrderSn(model.PayTypeAlipay, model.RefundStatusEnd, outBizNo, model.RefundChangeTypeSuccess)
			if err != nil {
				g.TENANCY_LOG.Error("退款异步通知: 支付宝支付异步通知回调错误", zap.String(orderSn, err.Error()))
			}
		} else if tradeStatus == "TRADE_CLOSED" { // 全部退款
			err := ChangeReturnOrderStatusByOrderSn(model.PayTypeAlipay, model.RefundStatusEnd, orderSn, model.RefundChangeTypeSuccess)
			if err != nil {
				g.TENANCY_LOG.Error("退款异步通知: 支付宝支付异步通知回调错误", zap.String(orderSn, err.Error()))
			}
		}
	} else { // 支付
		if tradeStatus != "TRADE_SUCCESS" && tradeStatus != "TRADE_FINISHED" {
			return fmt.Errorf("支付异步通知: %s 支付宝异步通知回调返回状态: %s", orderSn, tradeStatus)
		}
		// 发送 mqtt
		changeData := map[string]interface{}{
			"status":   model.OrderStatusNoDeliver,
			"pay_type": model.PayTypeAlipay,
			"pay_time": time.Now(),
			"paid":     g.StatusTrue,
		}
		payload, err := ChangeOrderPayNotifyByOrderSn(changeData, orderSn, model.ChangeTypePaySuccess)
		if err != nil {
			g.TENANCY_LOG.Error("支付异步通知: 支付宝支付异步通知回调错误", zap.String(orderSn, err.Error()))
		}
		payload.PayType = model.PayTypeAlipay
		payload.PayNotifyType = "pay"
		// 异步发送 mqtt
		go func() {
			SendMqttMsgs(model.TOPIC, payload, model.QOS)
		}()
	}

	return nil
}

// NotifyWechatPay 微信支付异步回调通知
func NotifyWechatPay(ctx *gin.Context) error {
	// ========异步通知验签========
	notifyReq, err := wechat.V3ParseNotify(ctx.Request)
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 微信支付异异步通知解析错误", zap.String("V3ParseNotify()", err.Error()))
		return fmt.Errorf("异步回调: 微信支付异异步通知解析错误: %w", err)
	}
	// WxPkContent 是通过 wechat.GetPlatformCerts() 接口向微信获取的微信平台公钥证书内容
	err = notifyReq.VerifySign(g.TENANCY_CONFIG.WechatPay.WxPkContent)
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 微信支付异异步通知验签错误", zap.String("VerifySign()", err.Error()))
		return fmt.Errorf("支付异步回调: 微信支付异异步通知验签错误: %w", err)
	}

	wechatConf, err := param.GetWechatPayConfig()
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 获取微信支付配置错误", zap.String("GetWechatPayConfig()", err.Error()))
		return fmt.Errorf("支付异步回调: 获取微信支付配置错误: %w", err)
	}

	// ========异步通知敏感信息解密========
	// 普通支付通知解密
	result, err := notifyReq.DecryptCipherText(wechatConf.PayWeixinKey)
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 解密支付信息错误", zap.String("DecryptCipherText()", err.Error()))
		return fmt.Errorf("支付异步回调: 解密支付信息错误: %w", err)
	}

	if result != nil {
		return wechatPayNotifyForPay(result)
	}

	return nil
}

// NotifyWechatPayReturn 微信支付异步回调通知
func NotifyWechatPayReturn(ctx *gin.Context) error {
	// ========异步通知验签========
	notifyReq, err := wechat.V3ParseNotify(ctx.Request)
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 微信支付异异步通知解析错误", zap.String("V3ParseNotify()", err.Error()))
		return fmt.Errorf("异步回调: 微信支付异异步通知解析错误: %w", err)
	}
	// WxPkContent 是通过 wechat.GetPlatformCerts() 接口向微信获取的微信平台公钥证书内容
	err = notifyReq.VerifySign(g.TENANCY_CONFIG.WechatPay.WxPkContent)
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 微信支付异异步通知验签错误", zap.String("VerifySign()", err.Error()))
		return fmt.Errorf("支付异步回调: 微信支付异异步通知验签错误: %w", err)
	}

	wechatConf, err := param.GetWechatPayConfig()
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 获取微信支付配置错误", zap.String("GetWechatPayConfig()", err.Error()))
		return fmt.Errorf("支付异步回调: 获取微信支付配置错误: %w", err)
	}

	// 退款通知解密
	resultRefund, err := notifyReq.DecryptRefundCipherText(wechatConf.PayWeixinKey)
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 解密支付信息错误", zap.String("DecryptCipherText()", err.Error()))
		return fmt.Errorf("支付异步回调: 解密支付信息错误: %w", err)
	}
	if resultRefund != nil {
		return wechatPayNotifyForRefund(resultRefund)
	}
	return nil
}

// wechatPayNotifyForPay 微信支付异步回调
func wechatPayNotifyForPay(result *wechat.V3DecryptResult) error {
	g.TENANCY_LOG.Info("支付异步回调: 微信支付异支付异步通知回调",
		zap.String("订单号", result.OutTradeNo),
		zap.String("通知状态", result.TradeState),
		zap.String("通知状态", result.TradeStateDesc),
		zap.String("通知类型", result.TradeType),
		zap.String("时间", result.SuccessTime),
		zap.String("流水号", result.TransactionId),
		zap.Int("用户支付金额，单位为分", result.Amount.PayerTotal),
		zap.String("人民币", result.Amount.Currency),
		zap.String("用户支付币种", result.Amount.PayerCurrency),
		zap.Int("订单总金额，单位为分", result.Amount.Total))

	// 支付
	if result.TradeState != "SUCCESS" {
		return fmt.Errorf("支付异步回调: %s 微信支付异异步通知回调返回状态: %s", result.OutTradeNo, result.TradeState)
	}
	// 发送 mqtt
	changeData := map[string]interface{}{
		"status":   model.OrderStatusNoDeliver,
		"pay_type": model.PayTypeWx,
		"pay_time": time.Now(),
		"paid":     g.StatusTrue,
	}
	payload, err := ChangeOrderPayNotifyByOrderSn(changeData, result.OutTradeNo, model.ChangeTypePaySuccess)
	if err != nil {
		g.TENANCY_LOG.Error("支付异步回调: 微信支付异支付异步通知回调错误", zap.String(result.OutTradeNo, err.Error()))
	}
	payload.PayType = model.PayTypeWx
	payload.PayNotifyType = "pay"
	// 异步发送 mqtt
	go func() {
		SendMqttMsgs(model.TOPIC, payload, model.QOS)
	}()

	return nil
}

// wechatPayNotifyForRefund 微信退款异步回调
func wechatPayNotifyForRefund(result *wechat.V3DecryptRefundResult) error {
	g.TENANCY_LOG.Info("退款异步回调: 微信支付异支付异步通知回调",
		zap.String("订单号", result.OutTradeNo),
		zap.String("通知状态", result.RefundStatus),
		zap.String("退款单号", result.OutRefundNo),
		zap.String("时间", result.SuccessTime),
		zap.String("流水号", result.TransactionId),
		zap.Int("用户支付金额，单位为分", result.Amount.PayerTotal),
		zap.Int("退款金额", result.Amount.Refund),
		zap.Int("退款给用户的金额", result.Amount.PayerRefund),
		zap.Int("订单总金额，单位为分", result.Amount.Total))

	g.TENANCY_LOG.Info("支付异步回调 ", zap.String("wechatPayNotifyForRefund()", fmt.Sprintf("%+v", result)))
	if result.RefundStatus != "SUCCESS" {
		return fmt.Errorf("退款异步回调返回状态错误")
	}

	err := ChangeReturnOrderStatusByReturnOrderSn(model.PayTypeWx, model.RefundStatusEnd, result.OutRefundNo, model.RefundChangeTypeSuccess)
	if err != nil {
		g.TENANCY_LOG.Error("退款异步回调: 微信支付异支付异步通知回调错误", zap.String(result.OutTradeNo, err.Error()))
	}

	return nil
}
