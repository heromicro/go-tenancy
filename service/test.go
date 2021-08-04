package service

import (
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
)

// EmailTest 发送邮件测试
func EmailTest() error {
	subject := "test"
	body := "test"
	return utils.EmailTest(subject, body)
}

// PayTest 支付测试
func PayTest(req request.CreateCart, tenancyName string) ([]byte, error) {
	cart, err := CreateCart(req)
	if err != nil {
		return nil, err
	}
	return CreateOrder(request.CreateOrder{
		CartIds:   []uint{cart.ID},
		OrderType: 1,
		Remark:    "remark",
	}, req.SysTenancyID, req.SysUserID, tenancyName)
}
