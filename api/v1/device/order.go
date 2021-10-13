package device

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/logic"
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
	pageInfo.CUserId = multi.GetUserId(ctx)
	pageInfo.SysTenancyId = multi.GetTenancyId(ctx)
	if ginH, err := service.GetOrderInfoList(pageInfo, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(ginH, "获取成功", ctx)
	}
}

// CheckOrder 结算页面
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
	req.TenancyId = multi.GetTenancyId(ctx)
	req.UserId = multi.GetUserId(ctx)
	if order, err := service.GetOrderDetailById(req.Id); err != nil {
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

// PayOrder 结算订单，生成支付二维码
func PayOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	req.TenancyId = multi.GetTenancyId(ctx)
	req.UserId = multi.GetUserId(ctx)

	if qrcode, err := logic.PayOrder(req); err != nil {
		g.TENANCY_LOG.Error("结算订单失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
		return
	} else {
		response.OkWithDetailed(gin.H{"qrcode": qrcode}, "操作成功", ctx)
	}
}

// CancelOrder 取消订单
func CancelOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	req.TenancyId = multi.GetTenancyId(ctx)
	req.UserId = multi.GetUserId(ctx)
	if err := logic.CancelOrder(req); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// CheckRefundOrder 申请退款结算页面
func CheckRefundOrder(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	// 商品ID
	var idsReq request.IdsReq
	if err := ctx.ShouldBind(&idsReq); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}
	req.UserId = multi.GetUserId(ctx)
	req.TenancyId = multi.GetTenancyId(ctx)
	if refundOrder, err := logic.CheckRefundOrder(req, idsReq.Ids); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(refundOrder, "操作成功", ctx)
	}
}

// RefundOrder 提交退款申请
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
	req.UserId = multi.GetUserId(ctx)
	req.TenancyId = multi.GetTenancyId(ctx)
	if id, err := logic.CreateRefundOrder(req, createRefundOrder); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{"id": id}, "操作成功", ctx)
	}
}
