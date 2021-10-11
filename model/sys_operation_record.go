package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// 如果含有time.Time 请自行import time包
type SysOperationRecord struct {
	BaseOperationRecord
	User SysUser `json:"user"`
}

type BaseOperationRecord struct {
	g.TENANCY_MODEL
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟"`
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`
	ErrorMessage string        `json:"errorMessage" form:"errorMessage" gorm:"column:error_message;comment:错误信息"`
	Body         string        `json:"body" form:"body" gorm:"type:longtext;column:body;comment:请求Body"`
	Resp         string        `json:"resp" form:"resp" gorm:"type:longtext;column:resp;comment:响应Body"`
	UserID       uint          `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id"`
	SysTenancyId uint          `json:"sysTenancyId" form:"sysTenancyId" gorm:"column:sys_tenancy_id;comment:商户id"`
}
