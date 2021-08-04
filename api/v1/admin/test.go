package admin

import (
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
	req := request.CreateCart{
		BaseCart: model.BaseCart{
			CartNum:           2,
			IsNew:             2,
			ProductAttrUnique: "e2fe28308fd2",
		},
		ProductID:    1,
		SysUserID:    1,
		SysTenancyID: 1,
	}
	if qrcode, err := service.PayTest(req, "宝安中心人民医院"); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败", ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"qrcode": qrcode,
		}, "操作成功", ctx)
	}
}
