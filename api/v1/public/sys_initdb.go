package public

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// InitDB 初始化用户项目
func InitDB(ctx *gin.Context) {
	if g.IsInit() {
		g.TENANCY_LOG.Error("项目已经初始化")
		response.FailWithMessage("项目已经初始化", ctx)
		return
	}
	var dbInfo request.InitDB
	if err := ctx.ShouldBindJSON(&dbInfo); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}
	if err := service.InitDB(dbInfo); err != nil {
		g.TENANCY_LOG.Error("初始化项目失败", zap.Any("err", err))
		response.FailWithMessage("初始化项目失败，请查看后台日志", ctx)
		return
	}
	response.OkWithData("初始化项目成功", ctx)
}

// CheckDB 初始化用户项目
func CheckDB(ctx *gin.Context) {
	if !g.IsInit() {
		g.TENANCY_LOG.Info("前往初始化项目")
		response.OkWithDetailed(gin.H{
			"needInit": true,
		}, "前往初始化项目", ctx)
		return
	} else {
		g.TENANCY_LOG.Info("项目无需初始化")
		response.OkWithDetailed(gin.H{
			"needInit": false,
		}, "项目无需初始化", ctx)
		return
	}
}
