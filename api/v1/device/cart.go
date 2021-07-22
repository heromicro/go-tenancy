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

func CreateCart(ctx *gin.Context) {
	var cart request.CreateCart
	cart.SysTenancyID = multi.GetTenancyId(ctx)
	cart.SysUserID = multi.GetUserId(ctx)
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

func GetCartList(ctx *gin.Context) {
	if list, total, err := service.GetCartList(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", ctx)
	}
}
