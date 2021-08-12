package tests

import (
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/core"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
)

func TestMain(m *testing.M) {
	g.TENANCY_VP = core.Viper()     // 初始化Viper
	g.TENANCY_LOG = core.Zap()      // 初始化zap日志库
	g.TENANCY_CACHE = cache.Cache() // redis缓存
	// 初始化认证服务
	initialize.Auth()

	uuid := uuid.NewV3(uuid.NewV4(), uuid.NamespaceOID.String()).String()
	mysqlConfig := request.InitDB{
		Host:     "127.0.0.1",
		Port:     "3306",
		UserName: "root",
		Password: "Chindeo",
		DBName:   uuid,
	}
	service.InitDB(mysqlConfig)

	// call flag.Parse() here if TestMain uses flags
	// 如果 TestMain 使用了 flags，这里应该加上 flag.Parse()
	code := m.Run()

	err := dorpDB(uuid)
	if err != nil {
		return
	}

	db, _ := g.TENANCY_DB.DB()
	db.Close()
	multi.AuthDriver.Close()

	os.Exit(code)
}

func dorpDB(uuid string) error {
	// 删除表和视图
	var sqls []string
	if err := g.TENANCY_DB.Raw("select CASE table_type WHEN 'VIEW' THEN concat('drop view ', table_name, ';') ELSE concat('drop table ', table_name, ';') END  from information_schema.tables where table_schema='%s';", uuid).Scan(&sqls).Error; err != nil {
		return err
	}

	for _, sql := range sqls {
		if err := g.TENANCY_DB.Exec(sql).Error; err != nil {
			continue
		}
	}

	if err := g.TENANCY_DB.Exec("drop database if exists %s;", uuid).Error; err != nil {
		return err
	}

	return nil
}
