package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
)

// IsAdmin
func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !multi.IsAdmin(ctx) {
			response.ForbiddenFailWithMessage("无此操作权限", ctx)
			ctx.Abort()
		}
		ctx.Next()
	}
}

// IsTenancy
func IsTenancy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !multi.IsTenancy(ctx) {
			response.ForbiddenFailWithMessage("无此操作权限", ctx)
			ctx.Abort()
		}
		ctx.Next()
	}
}

// IsGeneral
func IsGeneral() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !multi.IsGeneral(ctx) {
			response.ForbiddenFailWithMessage("无此操作权限", ctx)
			ctx.Abort()
		}
		ctx.Next()
	}
}

// CheckTenancy
func CheckTenancy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tenancy, err := service.GetTenancyByID(multi.GetTenancyId(ctx))
		if err != nil {
			response.ForbiddenFailWithMessage("商户参数错误", ctx)
			ctx.Abort()
		}
		if tenancy.Status == g.StatusFalse {
			response.ForbiddenFailWithMessage("当前商户已被冻结", ctx)
			ctx.Abort()
		}
		if tenancy.State == g.StatusFalse {
			response.ForbiddenFailWithMessage("当前商户已经停业", ctx)
			ctx.Abort()
		}
		ctx.Next()
	}
}
