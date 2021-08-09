package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"go.uber.org/zap"
)

// StartJob
func StartJob(ctx *gin.Context) {
	var req request.Task
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := g.TENANCY_Timer.StartTask(req.Name); err != nil {
		g.TENANCY_LOG.Error("启动失败!", zap.Any("err", err))
		response.FailWithMessage("启动失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("启动成功", ctx)
	}
}

// StopJob
func StopJob(ctx *gin.Context) {
	var req request.Task
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := g.TENANCY_Timer.StopTask(req.Name); err != nil {
		g.TENANCY_LOG.Error("启动失败!", zap.Any("err", err))
		response.FailWithMessage("启动失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("启动成功", ctx)
	}
}

// GetJobList
func GetJobList(ctx *gin.Context) {
	tasks := g.TENANCY_Timer.GetTasks()
	response.OkWithDetailed(tasks, "获取成功", ctx)

}

// DeleteJob
func DeleteJob(ctx *gin.Context) {
	var req request.Task
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := g.TENANCY_Timer.Clear(req.Name); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
