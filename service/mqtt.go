package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// GetMqttMap
func GetMqttMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		mqtt, err := GetMqttByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"name","value":"%s","title":"快递公司名称","props":{"type":"text","placeholder":"请输入快递公司名称"},"validate":[{"message":"请输入快递公司名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"code","value":"%s","title":"快递公司编码","props":{"type":"text","placeholder":"请输入快递公司编码"},"validate":[{"message":"请输入快递公司编码","required":true,"type":"string","trigger":"change"}]},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"\/sys\/store\/mqtt\/create.html","method":"PUT","title":"添加快递公司","config":{}}`, mqtt.Name, mqtt.Code, mqtt.Status, mqtt.Sort)

	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"name","value":"%s","title":"快递公司名称","props":{"type":"text","placeholder":"请输入快递公司名称"},"validate":[{"message":"请输入快递公司名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"code","value":"%s","title":"快递公司编码","props":{"type":"text","placeholder":"请输入快递公司编码"},"validate":[{"message":"请输入快递公司编码","required":true,"type":"string","trigger":"change"}]},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"\/sys\/store\/mqtt\/create.html","method":"POST","title":"添加快递公司","config":{}}`, "", "", 1, 0)
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/mqtt/updateMqtt/%d", id), ctx)
	} else {
		form.SetAction("/mqtt/createMqtt", ctx)
	}
	return form, err
}

// GetMqttOptions
func GetMqttOptions() ([]Option, error) {
	var options []Option
	var opts []StringOpt
	err := g.TENANCY_DB.Model(&model.Mqtt{}).Select("code as value,name as label").Where("status = ?", g.StatusTrue).Find(&opts).Error
	if err != nil {
		return options, err
	}
	options = append(options, Option{Label: "请选择", Value: ""})

	for _, opt := range opts {
		options = append(options, Option{Label: opt.Label, Value: opt.Value})
	}

	return options, err
}

// CreateMqtt
func CreateMqtt(mqtt model.Mqtt) (model.Mqtt, error) {
	err := g.TENANCY_DB.Where("code = ?", mqtt.Code).First(&mqtt).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return mqtt, errors.New("物流代码已被注冊")
	}
	err = g.TENANCY_DB.Create(&mqtt).Error
	return mqtt, err
}

// GetMqttByID
func GetMqttByID(id uint) (model.Mqtt, error) {
	var mqtt model.Mqtt
	err := g.TENANCY_DB.Where("id = ?", id).First(&mqtt).Error
	return mqtt, err
}

// GetMqttByCode
// TODO:根据单号获取物流信息，需要对接第三方平台
func GetMqttByCode(code string) (model.Mqtt, error) {
	var mqtt model.Mqtt
	return mqtt, nil
}

// ChangeMqttStatus
func ChangeMqttStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.Mqtt{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateMqtt
func UpdateMqtt(mqtt model.Mqtt, id uint) (model.Mqtt, error) {
	err := g.TENANCY_DB.Where("code = ?", mqtt.Code).Not("id = ?", id).First(&mqtt).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return mqtt, errors.New("物流代码已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Updates(mqtt).Error
	return mqtt, err
}

// DeleteMqtt
func DeleteMqtt(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.Mqtt{}).Error
}

// GetMqttInfoList
func GetMqttInfoList(info request.MqttPageInfo) ([]model.Mqtt, int64, error) {
	var mqttList []model.Mqtt
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Mqtt{})
	if info.Name != "" {
		db = db.Where(g.TENANCY_DB.Where("name like ?", info.Name+"%").Or("code like ?", info.Name+"%"))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return mqttList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&mqttList).Error
	return mqttList, total, err
}
