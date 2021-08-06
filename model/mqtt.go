package model

import "github.com/snowlyg/go-tenancy/g"

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
	Qos     int    `gorm:"column:qos;type:int;not null" json:"qos"`             // mqtt qos
	Topic   string `gorm:"column:topic;type:varchar(30);not null" json:"topic"` // mqtt topic
	Content string `gorm:"column:content;type:text;not null" json:"content"`    // content
}
