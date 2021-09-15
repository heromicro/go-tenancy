package admin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils/param"
	"go.uber.org/zap"
)

// EmailTest 发送测试邮件
func EmailTest(ctx *gin.Context) {
	if err := service.EmailTest(); err != nil {
		g.TENANCY_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", ctx)
	} else {
		response.OkWithData("发送成功", ctx)
	}
}

// PayTest 支付测试
func PayTest(ctx *gin.Context) {

	wechatConf, err := param.GetWechatPayConfig()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if !wechatConf.PayWeixinOpen {
		response.FailWithMessage("微信支付未开启", ctx)
		return
	}

	alipayConfi, err := param.GetAliPayConfig()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if !alipayConfi.AlipayOpen {
		response.FailWithMessage("支付宝支付未开启", ctx)
		return
	}

	var req request.CreateCart
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if qrcode, err := service.PayTest(req); err != nil {
		g.TENANCY_LOG.Error("生成支付二维码失败!", zap.Any("err", err))
		response.FailWithMessage("生成支付二维码失败", ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"qrcode": qrcode,
		}, "生成支付二维码成功", ctx)
	}
}

// MqttTest 发送测试消息
func MqttTest(ctx *gin.Context) {
	payload := model.Payload{
		OrderId:       1,
		TenancyId:     1,
		UserId:        1,
		PayType:       1,
		PayNotifyType: "test",
		CreatedAt:     time.Now(),
	}
	if err := service.SendMqttMsgs("tenancy_notify_test", payload, 2); err != nil {
		g.TENANCY_LOG.Error("测试失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithMessage("测试成功", ctx)
	}
}
