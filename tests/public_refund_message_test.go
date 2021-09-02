package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestRefundMessage(t *testing.T) {
	auth := base.BaseTester(t)
	url := "v1/public/getRefundMessage"
	pageKeys := base.ResponseKeys{
		{Key: "refundMessage", Value: []string{"收货地址填错了", "与描述不符", "信息填错了", "重新拍", "收到商品损坏了", "未按预定时间发货", "其它原因"}},
	}
	base.Get(auth, url, nil, http.StatusOK, "获取成功", pageKeys)
}
