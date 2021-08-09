package job

import (
	"encoding/json"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// SendMqtt 暂时未使用
type SendMqtt struct {
	Mqtts   []model.Mqtt
	Topic   string
	Qos     byte
	Payload model.Payload
}

func (d SendMqtt) Run() {
	if len(d.Mqtts) > 0 {
		var mqttRecords []model.MqttRecord
		for _, mqtt := range d.Mqtts {
			content, _ := json.Marshal(d.Payload)
			err := mqtt.MqttPublish(d.Topic, string(content), d.Qos)
			if err != nil {
				g.TENANCY_LOG.Error(fmt.Sprintf("主题：%s 消息发送失败", d.Topic), zap.String("错误", err.Error()))
			}
			mqttRecords = append(mqttRecords, model.MqttRecord{Host: mqtt.Host, Port: mqtt.Port, Qos: d.Qos, Topic: d.Topic, Content: string(content)})
		}

		if len(mqttRecords) > 0 {
			err := service.CreateMqttRecords(mqttRecords)
			if err != nil {
				g.TENANCY_LOG.Error("消息记录保持失败", zap.String("错误", err.Error()))
			}
		}
	}
}
