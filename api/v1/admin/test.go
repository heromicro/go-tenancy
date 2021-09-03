package admin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
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

// PayTest 发送测试邮件
func PayTest(ctx *gin.Context) {
	var req request.CreateCart
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if qrcode, err := service.PayTest(req, "宝安中心人民医院"); err != nil {
		g.TENANCY_LOG.Error("测试失败!", zap.Any("err", err))
		response.FailWithMessage("测试失败", ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"qrcode": qrcode,
		}, "测试成功", ctx)
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
