package client

import (
	"github.com/gin-gonic/gin"
	client "github.com/snowlyg/go-tenancy/api/v1/client"
)

// 规格参数
func InitAttrTemplateRouter(Router *gin.RouterGroup) {
	AttrTemplateRouter := Router.Group("/attrTemplate")
	{
		AttrTemplateRouter.POST("/getAttrTemplateList", client.GetAttrTemplateList)
		AttrTemplateRouter.POST("/createAttrTemplate", client.CreateAttrTemplate)
		AttrTemplateRouter.GET("/getAttrTemplateById/:id", client.GetAttrTemplateById)
		AttrTemplateRouter.PUT("/updateAttrTemplate/:id", client.UpdateAttrTemplate)
		AttrTemplateRouter.DELETE("/deleteAttrTemplate/:id", client.DeleteAttrTemplate)
	}
}

// 品牌
func InitBrandRouter(Router *gin.RouterGroup) {
	BrandRouter := Router.Group("/brand")
	{
		BrandRouter.GET("/getBrandList", client.GetBrandList)
	}
}

// 系统配置
func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("/config")
	{
		ConfigRouter.GET("/getConfigMap/:category", client.GetConfigMap)
	}
}

// 系统配置数值
func InitConfigValueRouter(Router *gin.RouterGroup) {
	ConfigValueRouter := Router.Group("/configValue")
	{
		ConfigValueRouter.POST("/saveConfigValue/:category", client.SaveConfigValue)
	}
}

// 商品分类
func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("/productCategory")
	{
		CategoryRouter.GET("/getCreateProductCategoryMap", client.GetCreateProductCategoryMap)
		CategoryRouter.GET("/getUpdateProductCategoryMap/:id", client.GetUpdateProductCategoryMap)
		CategoryRouter.GET("/getProductCategorySelect", client.GetProductCategorySelect)
		CategoryRouter.GET("/getAdminProductCategorySelect", client.GetAdminProductCategorySelect)
		CategoryRouter.POST("/createProductCategory", client.CreateProductCategory)
		CategoryRouter.GET("/getProductCategoryList", client.GetProductCategoryList)
		CategoryRouter.GET("/getProductCategoryById/:id", client.GetProductCategoryById)
		CategoryRouter.POST("/changeProductCategoryStatus", client.ChangeProductCategoryStatus)
		CategoryRouter.PUT("/updateProductCategory/:id", client.UpdateProductCategory)
		CategoryRouter.DELETE("/deleteProductCategory/:id", client.DeleteProductCategory)
	}
}

// 多媒体
func InitMediaRouter(Router *gin.RouterGroup) {
	MediaGroup := Router.Group("/media")
	{
		MediaGroup.GET("/getUpdateMediaMap/:id", client.GetUpdateMediaMap) // 修改名称表单
		MediaGroup.POST("/upload", client.UploadFile)                      // 上传文件
		MediaGroup.POST("/getFileList", client.GetFileList)                // 获取上传文件列表
		MediaGroup.POST("/updateMediaName/:id", client.UpdateMediaName)    // 修改名称
		MediaGroup.DELETE("/deleteFile", client.DeleteFile)                // 删除指定文件
	}
}

// 商品
func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("/product")
	{
		ProductRouter.GET("/getProductFilter", client.GetProductFilter)
		ProductRouter.POST("/createProduct", client.CreateProduct)
		ProductRouter.POST("/changeProductIsShow", client.ChangeProductIsShow)
		ProductRouter.POST("/getProductList", client.GetProductList)
		ProductRouter.GET("/getProductById/:id", client.GetProductById)
		ProductRouter.PUT("/updateProduct/:id", client.UpdateProduct)
		ProductRouter.GET("/restoreProduct/:id", client.RestoreProduct)
		ProductRouter.DELETE("/deleteProduct/:id", client.DeleteProduct)
		ProductRouter.DELETE("/destoryProduct/:id", client.DestoryProduct)
	}
}

// 患者管理
func InitProductReplyRouter(Router *gin.RouterGroup) {
	ProductReplyRouter := Router.Group("/productReply")
	{
		ProductReplyRouter.GET("/replyMap/:id", client.GetReplyMap)
		ProductReplyRouter.POST("/reply/:id", client.AddReply)
		ProductReplyRouter.POST("/getProductReplyList", client.GetProductReplyList)
	}
}

// 商户
func InitTenancyRouter(Router *gin.RouterGroup) {
	TenancyRouter := Router.Group("/tenancy")
	{
		TenancyRouter.GET("/getTenancyInfo", client.GetTenancyInfo)           // 登录商户信息
		TenancyRouter.GET("/getUpdateTenancyMap", client.GetUpdateTenancyMap) // 获取商户编辑表单
		TenancyRouter.PUT("/updateTenancy/:id", client.UpdateClientTenancy)   // 获取商户编辑表单
		TenancyRouter.GET("/getTenancyCopyCount", client.GetTenancyCopyCount) // 获取商户复制商品次数
	}
}

// 菜单
func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("/menu")
	{
		MenuRouter.GET("/getMenu", client.GetMenu) // 获取菜单树
	}
}

// 运费模板
func InitShippingTemplateRouter(Router *gin.RouterGroup) {
	ShippingTemplateRouter := Router.Group("/shippingTemplate")
	{
		ShippingTemplateRouter.GET("/getShippingTemplateSelect", client.GetShippingTemplateSelect)
		ShippingTemplateRouter.POST("/getShippingTemplateList", client.GetShippingTemplateList)
		ShippingTemplateRouter.POST("/createShippingTemplate", client.CreateShippingTemplate)
		ShippingTemplateRouter.GET("/getShippingTemplateById/:id", client.GetShippingTemplateById)
		ShippingTemplateRouter.PUT("/updateShippingTemplate/:id", client.UpdateShippingTemplate)
		ShippingTemplateRouter.DELETE("/deleteShippingTemplate/:id", client.DeleteShippingTemplate)
	}
}

// 操作日志
func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	SysOperationRecordRouter := Router.Group("/sysOperationRecord")
	{
		SysOperationRecordRouter.POST("/getSysOperationRecordList", client.GetSysOperationRecordList) // 获取SysOperationRecord列表
	}
}

// 订单
func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("/order")
	{
		OrderRouter.GET("/deliveryOrderMap/:id", client.DeliveryOrderMap)
		OrderRouter.GET("/getOrderRemarkMap/:id", client.GetOrderRemarkMap)
		OrderRouter.GET("/getEditOrderMap/:id", client.GetEditOrderMap)
		OrderRouter.POST("/getOrderList", client.GetOrderList)
		OrderRouter.GET("/getOrderChart", client.GetOrderChart)
		OrderRouter.GET("/getOrderFilter", client.GetOrderFilter)
		OrderRouter.GET("/getOrderById/:id", client.GetOrderById)
		OrderRouter.POST("/getOrderRecord/:id", client.GetOrderRecord)
		OrderRouter.POST("/deliveryOrder/:id", client.DeliveryOrder)
		OrderRouter.POST("/remarkOrder/:id", client.RemarkOrder)
		OrderRouter.POST("/updateOrder/:id", client.UpdateOrder)
		OrderRouter.DELETE("/deleteOrder/:id", client.DeleteOrder)
	}
}

// 管理员
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.GET("/changePasswordMap/:id", client.ChangePasswordMap) // 修改密码表单
		UserRouter.POST("/changePassword", client.ChangePassword)          // 修改密码
		UserRouter.GET("/registerAdminMap", client.RegisterAdminMap)       // 注册管理员表单
		UserRouter.GET("/updateAdminMap/:id", client.UpdateAdminMap)       // 编辑管理员表单
		UserRouter.POST("/registerAdmin", client.RegisterTenancy)          // 注册管理员
		UserRouter.POST("/changeProfile", client.ChangeProfile)            // 修改个人信息
		UserRouter.POST("/getAdminList", client.GetAdminList)              // 分页获取管理员列表
		UserRouter.DELETE("/deleteUser", client.DeleteUser)                // 删除用户
		UserRouter.PUT("/setUserInfo/:user_id", client.SetUserInfo)        // 设置用户信息

	}
}

// 退款单
func InitRefundOrderRouter(Router *gin.RouterGroup) {
	RefundOrderRouter := Router.Group("/refundOrder")
	{
		RefundOrderRouter.GET("/getRefundOrderRemarkMap/:id", client.GetRefundOrderRemarkMap)
		RefundOrderRouter.POST("/remarkRefundOrder/:id", client.RemarkRefundOrder)
		RefundOrderRouter.GET("/getRefundOrderMap/:id", client.GetRefundOrderMap)
		RefundOrderRouter.POST("/auditRefundOrder/:id", client.AuditRefundOrder)
		RefundOrderRouter.POST("/getRefundOrderList", client.GetRefundOrderList)
		RefundOrderRouter.POST("/getRefundOrderRecord/:id", client.GetRefundOrderRecord)
		RefundOrderRouter.DELETE("/deleteRefundOrder/:id", client.DeleteRefundOrder)
	}
}

// 物流信息
func InitExpressRouter(Router *gin.RouterGroup) {
	ExpressRouter := Router.Group("/express")
	{
		ExpressRouter.GET("/getExpressByCode/:code", client.GetExpressByCode)
	}
}

// 用户标签
func InitUserLabelRouter(Router *gin.RouterGroup) {
	UserLabelRouter := Router.Group("/userLabel")
	{
		UserLabelRouter.POST("/getLabelList", client.GetLabelList)
		UserLabelRouter.GET("/getCreateUserLabelMap", client.GetCreateUserLabelMap)
		UserLabelRouter.GET("/getUpdateUserLabelMap/:id", client.GetUpdateUserLabelMap)
		UserLabelRouter.POST("/createUserLabel", client.CreateUserLabel)
		UserLabelRouter.PUT("/updateUserLabel/:id", client.UpdateUserLabel)
		UserLabelRouter.DELETE("/deleteUserLabel/:id", client.DeleteUserLabel)

		// 自动标签
		AutoLabelRouter := UserLabelRouter.Group("/auto")
		{
			AutoLabelRouter.POST("/getLabelList", client.GetAutoLabelList)
			AutoLabelRouter.POST("/createUserLabel", client.CreateAutoLabel)
			AutoLabelRouter.PUT("/updateUserLabel/:id", client.UpdateAutoLabel)
			AutoLabelRouter.DELETE("/deleteUserLabel/:id", client.DeleteAutoLabel)
		}
	}
}

// 用户管理
func InitCUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/cuser")
	{
		UserRouter.GET("/setUserLabelMap/:id", client.SetUserLabelMap) // 设置标签表单
		UserRouter.POST("/setUserLabel/:id", client.SetUserLabel)      // 设置标签
		UserRouter.POST("/getOrderList/:id", client.GetUserOrderList)  // 用户订单列表
		UserRouter.POST("/getGeneralList", client.GetGeneralList)      // 分页获取c用户列表
	}
}

// 患者管理
func InitPatientRouter(Router *gin.RouterGroup) {
	PatientRouter := Router.Group("/patient")
	{
		PatientRouter.POST("/getPatientList", client.GetPatientList)
	}
}
