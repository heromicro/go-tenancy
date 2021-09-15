package service

import (
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
)

const payTestKey = "PAY_TEST_KEY:"

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

	qrcode, err := cache.GetCacheBytes(fmt.Sprintf("%s%d", payTestKey, cart.ID))
	if err != nil || qrcode == nil {
		tenancy, err := GetTenancyByID(req.SysTenancyID)
		if err != nil {
			return nil, err
		}
		qrcode, _, err = CreateOrder(res, req.SysTenancyID, req.SysUserID, req.PatientID, tenancy.Name)
		if err != nil {
			return nil, err
		}
		err = cache.SetCache(payTestKey, qrcode, 15*time.Minute)
		if err != nil {
			return nil, err
		}
	}
	return qrcode, nil
}
