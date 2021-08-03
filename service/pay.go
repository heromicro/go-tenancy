package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-pay/gopay"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

func PayOrder(req request.PayOrder, userAgent, tenancyName string) (string, error) {
	order, err := GetOrderByOrderIdUserIdAndTenancyId(req.OrderId, req.TenancyId, req.UserId, req.OrderType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("当前订单不存在")
	} else if err != nil {
		return "", err
	}

	//
	if order.Paid == g.StatusTrue && !order.PayTime.IsZero() && order.PayType > model.PayTypeUnknown && order.Status > model.OrderStatusUnknown {
		return "", fmt.Errorf("当前订单已经支付，请勿重复支付")
	}

	if strings.Contains(userAgent, "MicroMessenger") {
		fmt.Println("wechat")
	} else if strings.Contains(userAgent, "Alipay") {
		return Alipay(order, tenancyName)
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
	payUrl, err := g.TENANCY_ALIAPY.TradeWapPay(body)
	if err != nil {
		return "", err
	}
	return payUrl, nil
}
