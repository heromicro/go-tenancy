package public

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils/param"
	"go.uber.org/zap"
)

func GetRefundMessage(ctx *gin.Context) {
	if refundMessage, err := param.GetConfigValueByKey("refund_message"); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"refundMessage": strings.Split(refundMessage, ";"),
		}, "获取成功", ctx)
	}
}
