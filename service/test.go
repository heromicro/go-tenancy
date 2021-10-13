package service

import (
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/go-tenancy/utils/param"
)

const PayTestKey = "PAY_TEST_KEY:"

// DeleteTestCache 清除测试缓存
func DeleteTestCache(tenancyId, userId uint) {
	index := fmt.Sprintf("%s%d_%d_%d", PayTestKey, tenancyId, userId)
	cache.DeleteCache(index)
}

// EmailTest 发送邮件测试
func EmailTest() error {
	subject := "test"
	body := "test"
	return utils.EmailTest(subject, body)
}

// PayTest 支付测试
func PayTest(req request.CreateCart) ([]byte, error) {
	cart, err := CreateCart(req)
	if err != nil {
		return nil, err
	}
	res := request.CreateOrder{
		CartIds:   []uint{cart.ID},
		OrderType: 1,
		Remark:    "remark",
	}
	index := fmt.Sprintf("%s%d_%d_%d", PayTestKey, req.SysTenancyId, req.CUserId)
	qrcode, err := cache.GetCacheBytes(index)
	if err != nil || qrcode == nil {
		tenancy, err := GetTenancyByID(req.SysTenancyId)
		if err != nil {
			return nil, err
		}
		qrcode, _, err = CreateOrder(res, req.SysTenancyId, req.CUserId, tenancy.Name)
		if err != nil {
			return nil, err
		}

		autoCloseTime := param.GetOrderAutoCloseTime()
		err = cache.SetCache(index, qrcode, time.Duration(autoCloseTime)*time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return qrcode, nil
}
