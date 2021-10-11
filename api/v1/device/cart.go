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

func GetCartList(ctx *gin.Context) {
	if list, fails, total, err := service.GetCartList(multi.GetTenancyId(ctx), 0, multi.GetUserId(ctx), nil); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		if total > 99 {
			response.OkWithDetailed(gin.H{
				"fails": fails,
				"list":  list,
				"total": "99+",
			}, "获取成功", ctx)
		} else {
			response.OkWithDetailed(gin.H{
				"fails": fails,
				"list":  list,
				"total": total,
			}, "获取成功", ctx)
		}

	}
}

func GetProductCount(ctx *gin.Context) {
	if total, err := service.GetProductCount(multi.GetUserId(ctx), multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		if total > 99 {
			response.OkWithDetailed(gin.H{
				"total": "99+",
			}, "获取成功", ctx)
		} else {
			response.OkWithDetailed(gin.H{
				"total": total,
			}, "获取成功", ctx)
		}
	}
}

func CreateCart(ctx *gin.Context) {
	var cart request.CreateCart
	cart.SysTenancyId = multi.GetTenancyId(ctx)
	cart.PatientId = multi.GetUserId(ctx)
	if errs := ctx.ShouldBindJSON(&cart); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnCart, err := service.CreateCart(cart); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(returnCart, "创建成功", ctx)
	}
}

func ChangeCartNum(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var cartNum request.ChangeCartNum
	if errs := ctx.ShouldBindJSON(&cartNum); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.ChangeCartNum(cartNum.CartNum, req.Id, multi.GetUserId(ctx), multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}
func DeleteCart(ctx *gin.Context) {
	var delCart request.IdsReq
	if errs := ctx.ShouldBindJSON(&delCart); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteCart(delCart.Ids, multi.GetUserId(ctx), multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("操作失败!", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}
