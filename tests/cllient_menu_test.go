package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientMenu(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	url := "v1/merchant/menu/getMenu"
	base.GetList(auth, url, http.StatusOK, "获取成功")
}
