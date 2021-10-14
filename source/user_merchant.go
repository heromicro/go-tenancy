package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var UserMerchant = new(userMerchant)

type userMerchant struct{}

var now = time.Now()
var userMerchants = []model.UserMerchant{
	{CUserId: 7, SysTenancyId: 1, FirstPayTime: &now, LastPayTime: &now, PayCount: 3, PayPrice: 534.00, LastTime: now, Status: g.StatusTrue},
}

var userTenancyUserLabels = []model.UserUserLabel{
	{CUserId: 7, UserLabelID: 3, SysTenancyId: 1},
	{CUserId: 8, UserLabelID: 4, SysTenancyId: 1},
}

func (m *userMerchant) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.UserMerchant{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> user_merchants 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userMerchants).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.UserUserLabel{}).Create(&userTenancyUserLabels).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> user_merchants 表初始数据成功!")
		return nil
	})
}
