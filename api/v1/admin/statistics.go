package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/logic"
	"github.com/snowlyg/go-tenancy/model/request"
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

// GetStatisticsOrder 订单金额数据
func GetStatisticsOrder(ctx *gin.Context) {
	if data, err := logic.GetStatisticsOrder(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsOrderNum 订单数据
func GetStatisticsOrderNum(ctx *gin.Context) {
	if data, err := logic.GetStatisticsOrderNum(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsOrderUser 订单用户数据
func GetStatisticsOrderUser(ctx *gin.Context) {
	if data, err := logic.GetStatisticsOrderUser(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsMerchantStock 商品销量排行
func GetStatisticsMerchantStock(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsMerchantStock(dateReq); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsMerchantVisit 商户访客量排行
func GetStatisticsMerchantVisit(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsMerchantVisit(dateReq); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}

// GetStatisticsMerchantRate 商户销售额占比
func GetStatisticsMerchantRate(ctx *gin.Context) {
	var dateReq request.DateReq
	if errs := ctx.ShouldBind(&dateReq); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if data, err := logic.GetStatisticsMerchantRate(dateReq); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(data, "获取成功", ctx)
	}
}
