package public

import (
	"github.com/gin-gonic/gin"
	public "github.com/snowlyg/go-tenancy/api/v1/public"
	"github.com/snowlyg/go-tenancy/middleware"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	Router.GET("/logout", public.Logout) // 退出
	Router.GET("/clean", public.Clean)   //清空授权
}

// 数据库初始化检测
func InitInitRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("/init")
	{
		ApiRouter.POST("/initdb", public.InitDB)  // 创建Api
		ApiRouter.GET("/checkdb", public.CheckDB) // 创建Api
	}
}

// 登录授权验证码
func InitPublicRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("/public", middleware.NeedInit())
	{
		BaseRouter.POST("/admin/login", public.AdminLogin)
		BaseRouter.POST("/merchant/login", public.ClientLogin)
		BaseRouter.POST("/device/login", public.LoginDevice)
		BaseRouter.GET("/captcha", public.Captcha)
		BaseRouter.GET("/region/:p_code", public.Region)
		BaseRouter.GET("/getRegionList", public.RegionList)
	}
}
