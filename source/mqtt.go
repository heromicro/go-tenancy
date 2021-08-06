package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/utils"
	"gorm.io/gorm"
)

var Mqtt = new(mqtt)

type mqtt struct{}

// uuid.NewV3(ns uuid.UUID, name string)
var mqtts = []model.Mqtt{
	{
		Host:     "10.0.0.27",
		Port:     1883,
		Username: "Chindeo",
		Password: "P@ssw0rd",
		ClientID: utils.UUIDV5().String(),
		Status:   g.StatusTrue,
	},
	{
		Host:     "10.0.0.23",
		Port:     1883,
		Username: "Chindeo",
		Password: "P@ssw0rd",
		ClientID: utils.UUIDV5().String(),
		Status:   g.StatusTrue,
	},
}
var mqttRecords = []model.MqttRecord{
	{
		Host:    "10.0.0.27",
		Port:    1883,
		Qos:     2,
		Topic:   "chindeo/php/dinning",
		Content: "订单号：I202011132300302543962744 操作：cancel",
	},
	{
		Host:    "10.0.0.27",
		Port:    1883,
		Qos:     2,
		Topic:   "chindeo/php/dinning",
		Content: "订单号：I202011132300302543962744 操作：cancel",
	},
	{
		Host:    "10.0.0.27",
		Port:    1883,
		Qos:     2,
		Topic:   "chindeo/php/dinning",
		Content: "订单号：I202011132300302543962744 操作：cancel",
	},
}

func (m *mqtt) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Mqtt{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_mqtts 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&mqtts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Create(&mqttRecords).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_mqtts 表初始数据成功!")
		return nil
	})
}
