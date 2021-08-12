package tests

import (
	"os"
	"testing"

	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/core"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
)

func TestMain(m *testing.M) {
	g.TENANCY_VP = core.Viper() // 初始化Viper
	MysqlConfig := config.Mysql{
		Path:     "127.0.0.1:3306",
		Dbname:   "tenancy_test",
		Username: "root",
		Password: "Chindeo",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if err := service.WriteConfig(g.TENANCY_VP, MysqlConfig); err != nil {
		return
	}
	g.TENANCY_LOG = core.Zap()       // 初始化zap日志库
	g.TENANCY_DB = initialize.Gorm() // gorm连接数据库
	g.TENANCY_CACHE = cache.Cache()  // redis缓存
	// initialize.Timer()
	if g.TENANCY_DB != nil {
		initialize.MysqlTables(g.TENANCY_DB) // 初始化表
	}
	// 初始化认证服务
	initialize.Auth()

	// call flag.Parse() here if TestMain uses flags
	// 如果 TestMain 使用了 flags，这里应该加上 flag.Parse()
	os.Exit(m.Run())

	db, _ := g.TENANCY_DB.DB()
	db.Close()
	multi.AuthDriver.Close()
}
