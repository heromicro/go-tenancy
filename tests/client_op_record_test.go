package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientOperationRecode(t *testing.T) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)

	url := "v1/merchant/sysOperationRecord/getSysOperationRecordList"
	base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", base.PageKeys)
}
