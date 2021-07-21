package user

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/api/v1/user"
)

//收货地址
func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("/address")
	{
		AddressRouter.POST("/createAddress", user.CreateAddress)
		AddressRouter.POST("/getAddressList", user.GetAddressList)
		AddressRouter.GET("/getAddressById/:id", user.GetAddressById)
		AddressRouter.PUT("/updateAddress/:id", user.UpdateAddress)
		AddressRouter.DELETE("/deleteAddress/:id", user.DeleteAddress)
	}
}

// 发票管理
func InitReceiptRouter(Router *gin.RouterGroup) {
	ReceiptRouter := Router.Group("/receipt")
	{
		ReceiptRouter.POST("/createReceipt", user.CreateReceipt)
		ReceiptRouter.POST("/getReceiptList", user.GetReceiptList)
		ReceiptRouter.GET("/getReceiptById/:id", user.GetReceiptById)
		ReceiptRouter.PUT("/updateReceipt/:id", user.UpdateReceipt)
		ReceiptRouter.DELETE("/deleteReceipt/:id", user.DeleteReceipt)
	}
}

// 小程序用户 -> 不同商城（点餐、护工、商城、租赁）-> 医院 -> 病人

// 床旁用户接口
func InitDeviceRouter(Router *gin.RouterGroup) {

}
