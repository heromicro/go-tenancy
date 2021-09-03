package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetMqttMap
func GetMqttMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	validatorHost := map[string]interface{}{"message": "请输入host地址", "required": true, "type": "string", "trigger": "change"}
	validatorPort := map[string]interface{}{"message": "请输入端口", "required": true, "type": "number", "trigger": "change"}
	validatorUsername := map[string]interface{}{"message": "请输入用户名", "required": true, "type": "string", "trigger": "change"}
	validatorPassword := map[string]interface{}{"message": "请输入密码", "required": true, "type": "string", "trigger": "change"}
	if id > 0 {
		mqtt, err := GetMqttByID(id)
		if err != nil {
			return form, err
		}

		form = Form{Method: "PUT", Title: "修改信息"}
		form.AddRule(*NewInput("HOST", "host", "请输入host地址", mqtt.Host).AddValidator(validatorHost)).
			AddRule(*NewInput("PORT", "port", "请输入端口", mqtt.Port).AddValidator(validatorPort)).
			AddRule(*NewInput("用户名", "username", "请输入用户名", mqtt.Username).AddValidator(validatorUsername)).
			AddRule(*NewInput("密码", "password", "请输入密码", mqtt.Password).AddValidator(validatorPassword)).
			AddRule(*NewSwitch("是否显示", "status", mqtt.Status))
	} else {
		form = Form{Method: "POST", Title: "添加信息"}
		form.AddRule(*NewInput("HOST", "host", "请输入host地址", "").AddValidator(validatorHost)).
			AddRule(*NewInput("PORT", "port", "请输入端口", 1883).AddValidator(validatorPort)).
			AddRule(*NewInput("用户名", "username", "请输入用户名", "").AddValidator(validatorUsername)).
			AddRule(*NewInput("密码", "password", "请输入密码", "").AddValidator(validatorPassword)).
			AddRule(*NewSwitch("是否显示", "status", 0))
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/mqtt/updateMqtt/%d", id), ctx)
	} else {
		form.SetAction("/mqtt/createMqtt", ctx)
	}
	return form, nil
}

// CreateMqtt
func CreateMqtt(mqtt model.Mqtt) (uint, error) {
	mqtt.ClientID = utils.UUIDV5().String()
	err := g.TENANCY_DB.Where("host = ?", mqtt.Host).First(&mqtt).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.New("mqtt地址已被注冊")
	}
	err = g.TENANCY_DB.Create(&mqtt).Error
	return mqtt.ID, err
}

// GetMqttByID
func GetMqttByID(id uint) (model.Mqtt, error) {
	var mqtt model.Mqtt
	err := g.TENANCY_DB.Where("id = ?", id).First(&mqtt).Error
	return mqtt, err
}

// ChangeMqttStatus
func ChangeMqttStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.Mqtt{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateMqtt
func UpdateMqtt(mqtt model.Mqtt, id uint) error {
	err := g.TENANCY_DB.Where("host = ?", mqtt.Host).Not("id = ?", id).First(&mqtt).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("mqtt地址已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Omit("client_id").Updates(mqtt).Error
	return err
}

// DeleteMqtt
func DeleteMqtt(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.Mqtt{}).Error
}

// GetMqttInfoList
func GetMqttInfoList(info request.PageInfo) ([]model.Mqtt, int64, error) {
	mqttList := []model.Mqtt{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Mqtt{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return mqttList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&mqttList).Error
	return mqttList, total, err
}

// GetMqttRecordList
func GetMqttRecordList(info request.PageInfo) ([]model.MqttRecord, int64, error) {
	mqttList := []model.MqttRecord{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.MqttRecord{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return mqttList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&mqttList).Error
	return mqttList, total, err
}

func GetStatusMqtts() ([]model.Mqtt, error) {
	mqtts := []model.Mqtt{}
	err := g.TENANCY_DB.Model(&model.Mqtt{}).Where("status = ?", g.StatusTrue).Find(&mqtts).Error
	if err != nil {
		return mqtts, err
	}
	return mqtts, nil
}

// CreateMqttRecords
func CreateMqttRecords(mqttRecords []model.MqttRecord) error {
	err := g.TENANCY_DB.Create(&mqttRecords).Error
	if err != nil {
		return err
	}
	return nil
}

func SendMqttMsgs(topic string, payload model.Payload, qos byte) error {
	mqttRecords := []model.MqttRecord{}
	mqtts, err := GetStatusMqtts()
	if err != nil {
		return err
	}
	if len(mqtts) == 0 {
		return errors.New("请添加mqtt客户端")
	}
	for _, mqtt := range mqtts {
		content, _ := json.Marshal(payload)
		err := mqtt.MqttPublish(topic, string(content), qos)
		if err != nil {
			g.TENANCY_LOG.Error(fmt.Sprintf("主题：%s 消息发送失败", topic), zap.String("错误", err.Error()))
		}

		mqttRecords = append(mqttRecords, model.MqttRecord{Host: mqtt.Host, Port: mqtt.Port, Qos: qos, Topic: topic, Content: string(content)})
	}

	if len(mqttRecords) > 0 {
		err := CreateMqttRecords(mqttRecords)
		if err != nil {
			g.TENANCY_LOG.Error("消息记录保持失败", zap.String("错误", err.Error()))
			return err
		}
	}
	return nil
}
