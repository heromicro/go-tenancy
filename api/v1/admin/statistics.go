package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/logic"
	"github.com/snowlyg/go-tenancy/model/response"
	"go.uber.org/zap"
)

// GetStatisticsMain 运营数据
func GetStatisticsMain(ctx *gin.Context) {
	if data, err := logic.GetStatisticsMaim(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}
