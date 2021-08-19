package initialize

import (
	"net/http"
	"path/filepath"
	"strings"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/chindeo/pkg/file"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/router/admin"
	"github.com/snowlyg/go-tenancy/router/client"
	"github.com/snowlyg/go-tenancy/router/device"
	"github.com/snowlyg/go-tenancy/router/public"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/thinkerou/favicon"
)

// 初始化总路由
func App() *gin.Engine {
	gin.SetMode(g.TENANCY_CONFIG.System.Level)
	app := gin.Default()
	utils.RegisterValidation() // 注册已定义验证方法
	Routers(app)
	return app
}

// Routers
func Routers(app *gin.Engine) {
	app.Use(limit.MaxAllowed(50))
	// 跨域
	app.Use(middleware.Cors()) // 如需跨域可以打开
	g.TENANCY_LOG.Info("use middleware cors")
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	app.Use(favicon.New("resource/favicon.ico"))
	app.Use(static.Serve("/doc", static.LocalFile("doc/apidoc", true)))
	app.Use(static.Serve("/admin", static.LocalFile("www/admin", true)))
	app.Use(static.Serve("/system", static.LocalFile("www/admin/system", true)))
	app.Use(static.Serve("/merchant", static.LocalFile("www/merchant", true)))
	app.Use(static.Serve("/client", static.LocalFile("www/merchant/client", true)))
	app.Use(static.Serve("/uploads", static.LocalFile("uploads", true)))
	app.Use(static.Serve("/UEditor", static.LocalFile("www/merchant/UEditor", true)))

	app.LoadHTMLGlob(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "resource/template/*"))
	app.StaticFS(g.TENANCY_CONFIG.Local.Path, http.Dir(g.TENANCY_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址

	// 关键点【解决页面刷新404的问题】
	app.NoRoute(func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusOK)
		if strings.Contains(ctx.Request.RequestURI, "admin") {
			file, _ := file.ReadBytes(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "www/admin/index.html"))
			ctx.Writer.Write(file)
		} else {
			file, _ := file.ReadBytes(filepath.Join(g.TENANCY_CONFIG.Casbin.ModelPath, "www/merchant/index.html"))
			ctx.Writer.Write(file)
		}
		ctx.Writer.Header().Add("Accept", "text/html")
		ctx.Writer.Flush()
	})

	app.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/admin")
	})

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := app.Group("/v1")
	{
		public.InitPublicRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		public.InitInitRouter(PublicGroup)   // 自动初始化相关
		public.InitPayRouter(PublicGroup)    // 自动初始化相关
	}

	V1Group := app.Group("/v1", middleware.NeedInit(), middleware.Auth())
	{
		Auth := V1Group.Group("/auth")
		{
			public.InitAuthRouter(Auth) // 注册用户路由
		}

		// 管理员
		AdminGroup := V1Group.Group(g.TENANCY_CONFIG.System.AdminPreix, middleware.IsAdmin(), middleware.CasbinHandler(), middleware.OperationRecord())
		{
			admin.InitApiRouter(AdminGroup)            // 注册功能api路由
			admin.InitUserRouter(AdminGroup)           // 注册用户路由
			admin.InitCUserRouter(AdminGroup)          // 注册c用户路由
			admin.InitTenancyRouter(AdminGroup)        // 注册商户路由
			admin.InitMiniRouter(AdminGroup)           // 注册小程序路由
			admin.InitBrandRouter(AdminGroup)          // 注册品牌路由
			admin.InitBrandCategoryRouter(AdminGroup)  // 注册品牌分类路由
			admin.InitConfigCategoryRouter(AdminGroup) // 注册系统配置分类路由
			admin.InitConfigRouter(AdminGroup)         // 注册系统配置路由
			admin.InitConfigValueRouter(AdminGroup)    // 注册系统配置值路由
			admin.InitMenuRouter(AdminGroup)           // 注册menu路由
			admin.InitTestouter(AdminGroup)            // 邮件相关路由
			admin.InitSystemRouter(AdminGroup)         // system相关路由
			admin.InitCasbinRouter(AdminGroup)         // 权限相关路由
			admin.InitAuthorityRouter(AdminGroup)      // 注册角色路由
			admin.InitMediaRouter(AdminGroup)          // 媒体库路由
			admin.InitCategoryRouter(AdminGroup)       // 商品分类路由
			admin.InitProductRouter(AdminGroup)        // 商品路由
			admin.InitProductReplyRouter(AdminGroup)   // 商品评价路由
			admin.InitOrderRouter(AdminGroup)          // 订单路由
			admin.InitRefundOrderRouter(AdminGroup)    // 退款订单路由
			admin.InitExpressRouter(AdminGroup)        // 物流公司路由
			admin.InitMqttRouter(AdminGroup)           // MQTT路由
			admin.InitJobRouter(AdminGroup)            // 定时任务管理路由
			admin.InitUserGroupRouter(AdminGroup)      // 用户分组路由
			admin.InitUserLabelRouter(AdminGroup)      // 用户标签路由
			admin.InitPatientRouter(AdminGroup)        // 床旁患者路由
		}

		AdminLogGroup := V1Group.Group(g.TENANCY_CONFIG.System.AdminPreix, middleware.IsAdmin(), middleware.CasbinHandler())
		{
			admin.InitSysOperationRecordRouter(AdminLogGroup) // 操作记录
		}

		// 商户员工
		ClientGroup := V1Group.Group(g.TENANCY_CONFIG.System.ClientPreix, middleware.IsTenancy(), middleware.CheckTenancy(), middleware.CasbinHandler(), middleware.OperationRecord())
		{
			client.InitUserRouter(ClientGroup)             // 注册用户路由
			client.InitTenancyRouter(ClientGroup)          // 注册商户路由
			client.InitBrandRouter(ClientGroup)            // 注册品牌路由
			client.InitConfigRouter(ClientGroup)           // 注册系统配置路由
			client.InitConfigValueRouter(ClientGroup)      // 注册系统配置值路由
			client.InitMenuRouter(ClientGroup)             // 注册menu路由
			client.InitMediaRouter(ClientGroup)            // 媒体库路由
			client.InitCategoryRouter(ClientGroup)         // 商品分类路由
			client.InitAttrTemplateRouter(ClientGroup)     // 规格模板路由
			client.InitProductRouter(ClientGroup)          // 商品路由
			client.InitProductReplyRouter(ClientGroup)     // 商品评价路由
			client.InitShippingTemplateRouter(ClientGroup) // 运费模板路由
			client.InitOrderRouter(ClientGroup)            // 订单路由
			client.InitRefundOrderRouter(ClientGroup)      // 退款订单路由
			client.InitExpressRouter(ClientGroup)          // 物流公司路由
			client.InitUserLabelRouter(ClientGroup)        // 用户标签路由
			client.InitCUserRouter(ClientGroup)            // 用户管理路由
			client.InitPatientRouter(ClientGroup)          // 患者管理路由
		}
		ClientLogGroup := V1Group.Group(g.TENANCY_CONFIG.System.ClientPreix, middleware.IsTenancy(), middleware.CasbinHandler())
		{
			client.InitSysOperationRecordRouter(ClientLogGroup) // 操作记录
		}

		// GeneralGroup := V1Group.Group("/user", middleware.IsGeneral())
		// {
		// 	user.InitDeviceRouter(GeneralGroup)
		// }

		DeviceGroup := V1Group.Group("/device", middleware.IsDevice(), middleware.CheckTenancy(), middleware.OperationRecord())
		{
			device.InitDeviceRouter(DeviceGroup)
		}
	}
}
