package service

import (
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/utils"
)

// GetSystemConfig 读取配置文件
func GetSystemConfig() config.Server {
	return g.TENANCY_CONFIG
}

// SetSystemConfig 设置配置文件
func SetSystemConfig(system model.System) error {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		g.TENANCY_VP.Set(k, v)
	}
	return g.TENANCY_VP.WriteConfig()
}
