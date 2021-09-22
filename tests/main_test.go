package tests

import (
	_ "embed"
	"fmt"
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/core"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
)

//go:embed password.txt
var password string

func TestMain(m *testing.M) {
	g.TENANCY_VP = core.Viper()     // 初始化Viper
	g.TENANCY_LOG = core.Zap()      // 初始化zap日志库
	g.TENANCY_CACHE = cache.Cache() // redis缓存
	// 初始化认证服务
	initialize.Auth()

	uuid := uuid.NewV3(uuid.NewV4(), uuid.NamespaceOID.String()).String()
	mysqlConfig := request.InitDB{
		SqlType: "mysql",
		Sql: request.Sql{
			Host:     "127.0.0.1",
			Port:     "3306",
			UserName: "root",
			Password: password,
			DBName:   uuid,
		},
		CacheType: "redis",
		Cache: request.Cache{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: password,
			PoolSize: 1000,
		},
		Addr:  8089,
		Level: "test",
	}
	err := service.InitDB(mysqlConfig)
	if err != nil {
		fmt.Printf("初始化数据库错误： %v\n", err)
		return
	}

	req := request.CreateTenancy{
		SysTenancy: model.SysTenancy{
			BaseTenancy: model.BaseTenancy{
				Name:          "多商户平台直营医院",
				Tele:          "0755-23568911",
				Address:       "xxx街道666号",
				BusinessTime:  "08:30-17:30",
				Status:        g.StatusTrue,
				SysRegionCode: 1,
				IsAudit:       g.StatusFalse, // 商品无需审核
			},
		},
		Username: "tenancy_hospital",
	}
	tennancyId, tenancyUUID, username, err := service.CreateTenancy(req)
	if err != nil {
		fmt.Printf("初始化商户错误： %v\n", err)
		return
	}
	cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":username", username, 0)
	cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":id", tennancyId, 0)
	cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":uuid", tenancyUUID, 0)

	code := m.Run()

	err = dorpDB(uuid)
	if err != nil {
		fmt.Printf("初始化商户错误： %v\n", err)
		return
	}

	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":username")
	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":id")
	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":uuid")

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

	if err := g.TENANCY_DB.Exec(fmt.Sprintf("drop database if exists `%s`;", uuid)).Error; err != nil {
		return err
	}

	return nil
}
