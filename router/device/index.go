package device

import (
	"github.com/gin-gonic/gin"
	device "github.com/snowlyg/go-tenancy/api/v1/device"
)

// //收货地址
// func InitAddressRouter(Router *gin.RouterGroup) {
// 	AddressRouter := Router.Group("/address")
// 	{
// 		AddressRouter.POST("/createAddress", device.CreateAddress)
// 		AddressRouter.POST("/getAddressList", device.GetAddressList)
// 		AddressRouter.GET("/getAddressById/:id", device.GetAddressById)
// 		AddressRouter.PUT("/updateAddress/:id", device.UpdateAddress)
// 		AddressRouter.DELETE("/deleteAddress/:id", device.DeleteAddress)
// 	}
// }

// // 发票管理
// func InitReceiptRouter(Router *gin.RouterGroup) {
// 	ReceiptRouter := Router.Group("/receipt")
// 	{
// 		ReceiptRouter.POST("/createReceipt", device.CreateReceipt)
// 		ReceiptRouter.POST("/getReceiptList", device.GetReceiptList)
// 		ReceiptRouter.GET("/getReceiptById/:id", device.GetReceiptById)
// 		ReceiptRouter.PUT("/updateReceipt/:id", device.UpdateReceipt)
// 		ReceiptRouter.DELETE("/deleteReceipt/:id", device.DeleteReceipt)
// 	}
// }

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
		OrderRouter.POST("/checkOrder", device.CheckOrder)
	}
	// 患者管理
	PatientRouter := Router.Group("/patient")
	{
		PatientRouter.GET("/getPatientList", device.GetPatientList)
	}
}
