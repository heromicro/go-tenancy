package client

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/logic"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// GetStatisticsMain 运营数据
func GetStatisticsMain(ctx *gin.Context) {
	if data, err := logic.GetClientStatisticsMaim(multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsOrder 订单金额数据
func GetStatisticsOrder(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetClientStatisticsOrder(dateReq, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsProduct 商品销量排行
func GetStatisticsProduct(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsProduct(dateReq, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsProductVisit 商户访客量排行
func GetStatisticsProductVisit(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsProductVisit(dateReq, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsProductCart 商户销售额占比
func GetStatisticsProductCart(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsProductCart(dateReq, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsUser 用户成交数据
func GetStatisticsUser(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsUser(dateReq, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsUserRate 用户成交占比数据
func GetStatisticsUserRate(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsUserRate(dateReq); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}
