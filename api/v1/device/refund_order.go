package device

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// GetRefundOrderList
func GetRefundOrderList(ctx *gin.Context) {
	var pageInfo request.RefundOrderPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	pageInfo.SysUserId = multi.GetUserId(ctx)
	pageInfo.SysTenancyId = multi.GetTenancyId(ctx)
	if list, stat, total, err := service.GetRefundOrderInfoList(pageInfo, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"stat":     stat,
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetOrderById
func GetRefundOrderById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if order, err := service.GetRefundOrderByID(req.Id, service.GetIsDelField(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(order, ctx)
	}
}
