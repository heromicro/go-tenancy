package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestDeviceProductCategoryList(t *testing.T) {
	auth := base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/device/productCategory/getProductCategoryList"
	base.GetList(auth, url, http.StatusOK, "获取成功")
}
