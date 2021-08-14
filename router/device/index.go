package device

import (
	"github.com/gin-gonic/gin"
	device "github.com/snowlyg/go-tenancy/api/v1/device"
)

// 小程序用户 -> 不同商城（点餐、护工、商城、租赁）-> 医院 -> 病人
// 床旁用户接口
func InitDeviceRouter(Router *gin.RouterGroup) {
	// 商品分类
	CategoryRouter := Router.Group("/productCategory")
	{
		CategoryRouter.GET("/getProductCategoryList", device.GetProductCategoryList)
	}
	// 商品
	ProductRouter := Router.Group("/product")
	{
		ProductRouter.POST("/getProductList", device.GetProductList)
		ProductRouter.GET("/getProductById/:id", device.GetProductById)
	}
	// 购物车
	CartRouter := Router.Group("/cart")
	{
		CartRouter.GET("/getCartList", device.GetCartList)
		CartRouter.GET("/getProductCount", device.GetProductCount)
		CartRouter.POST("/createCart", device.CreateCart) // 添加购物车
		CartRouter.POST("/changeCartNum/:id", device.ChangeCartNum)
		CartRouter.DELETE("/deleteCart", device.DeleteCart)
	}
	// 订单
	OrderRouter := Router.Group("/order")
	{
		OrderRouter.POST("/getOrderList", device.GetOrderList)
		OrderRouter.POST("/checkOrder", device.CheckOrder)
		OrderRouter.GET("/getOrderById/:id", device.GetOrderById)
		OrderRouter.POST("/createOrder", device.CreateOrder)
		OrderRouter.GET("/payOrder/:id", device.PayOrder)
		OrderRouter.GET("/cancelOrder/:id", device.CancelOrder)
		OrderRouter.POST("/checkRefundOrder/:id", device.CheckRefundOrder)
		OrderRouter.POST("/refundOrder/:id", device.RefundOrder)
	}
	// 患者管理
	PatientRouter := Router.Group("/patient")
	{
		PatientRouter.GET("/getPatientDetail", device.GetPatientDetail)
		PatientRouter.GET("/getPatientList", device.GetPatientList)
	}
}
