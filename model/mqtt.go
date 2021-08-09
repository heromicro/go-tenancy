package model

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/snowlyg/go-tenancy/g"
)

type Mqtt struct {
	g.TENANCY_MODEL
	Host     string `gorm:"unique;column:host;type:varchar(30);not null" json:"host" binding:"ip"`         // mqtt host
	Port     int    `gorm:"column:port;type:int;not null" json:"port"  binding:"required"`                 // mqtt port
	Username string `gorm:"column:username;type:varchar(20);not null" json:"username"  binding:"required"` // mqtt username
	Password string `gorm:"column:password;type:varchar(30);not null" json:"password"  binding:"required"` // mqtt password
	ClientID string `gorm:"unique;column:client_id;type:varchar(50);not null"  json:"clientId"`            // mqtt client_id
	Status   int    `gorm:"column:status;type:tinyint(1);not null;default:2"  json:"status"`               // 是否有效
}

type MqttRecord struct {
	g.TENANCY_MODEL
	Host    string `gorm:"column:host;type:varchar(30);not null" json:"host"`   // mqtt host
	Port    int    `gorm:"column:port;type:int;not null" json:"port"`           // mqtt port
	Qos     byte   `gorm:"column:qos;type:int;not null" json:"qos"`             // mqtt qos
	Topic   string `gorm:"column:topic;type:varchar(30);not null" json:"topic"` // mqtt topic
	Content string `gorm:"column:content;type:text;not null" json:"content"`    // content
}

type Payload struct {
	OrderId       uint   `json:"orderId"`
	TenancyId     uint   `json:"tenancyId"`
	UserId        uint   `json:"userId"`
	OrderType     int    `json:"orderType"`
	PayType       int    `json:"payType"`
	PayNotifyType string `json:"payNotifyType"` // pay
	CreatedAt     time.Time
}

func NewMqttClient(mq Mqtt) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mq.Host, mq.Port))
	opts.SetClientID(mq.ClientID)
	opts.SetUsername(mq.Username)
	opts.SetPassword(mq.Password)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	return mqtt.NewClient(opts)

}

func (mq *Mqtt) MqttPublish(topic string, payload interface{}, qos byte) error {
	c := NewMqttClient(*mq)
	token := c.Connect()
	if !token.Wait() {
		return fmt.Errorf("mqtt token wait %w", token.Error())
	}
	if token.Error() != nil {
		return fmt.Errorf("mqtt 连接失败 %w", token.Error())
	}
	defer c.Disconnect(250)

	token = c.Publish(topic, 2, false, payload)
	if !token.Wait() {
		return fmt.Errorf("mqtt token wait %w", token.Error())
	}
	return nil
}
