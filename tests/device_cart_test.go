package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestDeviceCartList(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	cartList(auth, base.PageRes, base.PageKeys, http.StatusOK, "获取成功")
}

func TestDeviceCartProcess(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)

	create := map[string]interface{}{"cartNum": 2, "isNew": 2, "productAttrUnique": "e2fe28308fd2", "productId": 1, "productType": 1}
	cartId := CreateCart(auth, create, http.StatusOK, "创建成功")
	if cartId == 0 {
		t.Error("添加购物车失败")
		return
	}
	defer DeleteCart(auth, map[string]interface{}{"ids": []uint{cartId}}, http.StatusOK, "操作成功")
	{
		url := fmt.Sprintf("v1/device/cart/changeCartNum/%d", cartId)
		base.Post(auth, url, map[string]interface{}{"cartNum": 2}, http.StatusOK, "操作成功")
	}

}

func TestDeviceGetProductCount(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/device/cart/getProductCount"
	base.Get(auth, url, http.StatusOK, "获取成功")
}

func DeleteCart(auth *httpexpect.Expect, data map[string]interface{}, status int, message string) {
	url := "v1/device/cart/deleteCart"
	base.DeleteMutil(auth, url, data, status, message)
}

func cartList(auth *httpexpect.Expect, pageRes map[string]interface{}, pageKeys base.ResponseKeys, status int, message string) {
	url := "v1/device/cart/getCartList"
	base.GetList(auth, url, 0, nil, status, message)
}

func CreateCart(auth *httpexpect.Expect, create map[string]interface{}, status int, message string) uint {
	url := "v1/device/cart/createCart"
	res := base.IdKeys()
	base.Create(auth, url, create, res, status, message)
	return res.GetId()
}
