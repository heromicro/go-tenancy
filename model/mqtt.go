package model

import "github.com/snowlyg/go-tenancy/g"

type Mqtt struct {
	g.TENANCY_MODEL
	Host     string `gorm:"column:host;type:varchar(30);not null"`             // mqtt host
	Port     int    `gorm:"column:port;type:int;not null"`                     // mqtt port
	Username string `gorm:"column:username;type:varchar(20);not null"`         // mqtt username
	Password string `gorm:"column:password;type:varchar(30);not null"`         // mqtt password
	ClientID string `gorm:"unique;column:client_id;type:varchar(50);not null"` // mqtt client_id
	Status   int    `gorm:"column:status;type:tinyint(1);not null;default:2"`  // 是否有效
}

type MqttRecord struct {
	g.TENANCY_MODEL
	Host    string `gorm:"column:host;type:varchar(30);not null"`      // mqtt host
	Port    int    `gorm:"column:port;type:int;not null;default:1883"` // mqtt port
	Qos     int    `gorm:"column:qos;type:int;not null;default:1883"`  // mqtt qos
	Topic   string `gorm:"column:topic;type:varchar(30);not null"`     // mqtt topic
	Content string `gorm:"column:content;type:text;not null"`          // content
}
