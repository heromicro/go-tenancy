package main

import (
	"github.com/snowlyg/go-tenancy/core"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/go-tenancy/utils/job"
	"github.com/snowlyg/multi"
)

func main() {
	g.TENANCY_VP = core.Viper()      // 初始化Viper
	g.TENANCY_LOG = core.Zap()       // 初始化zap日志库
	g.TENANCY_DB = initialize.Gorm() // gorm连接数据库
	if g.TENANCY_DB != nil {
		job.Timer()
		// 程序结束前关闭数据库链接
		db, _ := g.TENANCY_DB.DB()
		defer db.Close()
	}
	// 初始化认证服务
	initialize.Auth()
	if multi.AuthDriver != nil {
		defer multi.AuthDriver.Close()
	}

	core.RunServer()
}
