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

// GetOrderList
func GetOrderList(ctx *gin.Context) {
	var pageInfo request.OrderPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	pageInfo.SysUserId = multi.GetUserId(ctx)
	pageInfo.SysTenancyId = multi.GetTenancyId(ctx)
	if list, stat, total, err := service.GetOrderInfoList(pageInfo, ctx); err != nil {
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

func CheckOrder(ctx *gin.Context) {
	var req request.IdsReq
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if order, err := service.CheckOrder(req.Ids, multi.GetTenancyId(ctx), multi.GetUserId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(order, "获取成功", ctx)
	}
}

// GetOrderById
func GetOrderById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if order, err := service.GetOrderDetailById(req.Id, multi.GetTenancyId(ctx), multi.GetUserId(ctx), service.GetIsDelField(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(order, ctx)
	}
}

// CreateOrder
func CreateOrder(ctx *gin.Context) {
	var req request.CreateOrder
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if qrcode, orderId, err := service.CreateOrder(req, multi.GetTenancyId(ctx), multi.GetUserId(ctx), multi.GetTenancyName(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{"qrcode": qrcode, "id": orderId}, "获取成功", ctx)
	}
}

// PayOrder
func PayOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	var payOrder request.PayOrder
	if err := ctx.ShouldBind(&payOrder); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}

	err := service.CheckOrderStatusBeforePay(req.Id, multi.GetTenancyId(ctx), multi.GetUserId(ctx))
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
		return
	}

	if qrcode, err := service.GetQrCode(req.Id, multi.GetTenancyId(ctx), multi.GetUserId(ctx), payOrder.OrderType); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
		return
	} else {
		response.OkWithDetailed(gin.H{"qrcode": qrcode}, "获取成功", ctx)
	}
}

// CancelOrder 
func CancelOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.CancelOrder(req.Id, multi.GetTenancyId(ctx), multi.GetUserId(ctx)); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// CheckRefundOrder
func CheckRefundOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var idsReq request.IdsReq
	if err := ctx.ShouldBind(&idsReq); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}

	if refundOrder, err := service.CheckRefundOrder(req.Id, multi.GetTenancyId(ctx), multi.GetUserId(ctx), idsReq.Ids); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(refundOrder, "操作成功", ctx)
	}
}

// RefundOrder
func RefundOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var createRefundOrder request.CreateRefundOrder
	if err := ctx.ShouldBind(&createRefundOrder); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}
	if id, err := service.CreateRefundOrder(req.Id, multi.GetTenancyId(ctx), multi.GetUserId(ctx), createRefundOrder); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{"id": id}, "操作成功", ctx)
	}
}
