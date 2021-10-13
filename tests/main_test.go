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
	"github.com/snowlyg/go-tenancy/migration"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/multi"
)

//go:embed password.txt
var password string

func TestMain(m *testing.M) {
	uuid := uuid.NewV3(uuid.NewV4(), uuid.NamespaceOID.String()).String()

	dsn := "root:" + password + "@tcp(127.0.0.1:3306)/"
	if err := migration.CreateTable(dsn, uuid); err != nil {
		fmt.Printf("新建数据库错误： %v\n", err)
		return
	}
	if err := migration.GetMigrate(dsn, uuid).Migrate(); err != nil {
		fmt.Printf("数据库迁移错误： %v\n", err)
		return
	}
	g.TENANCY_DB = migration.GormMysql(dsn + uuid)
	if err := source.RunSeed(); err != nil {
		fmt.Printf("数据填充错误： %v\n", err)
		return
	}

	g.TENANCY_VP = core.Viper()     // 初始化Viper
	g.TENANCY_LOG = core.Zap()      // 初始化zap日志库
	g.TENANCY_CACHE = cache.Cache() // redis缓存
	// 初始化认证服务
	initialize.Auth()

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

	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":username")
	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":id")
	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":uuid")

	err = migration.DorpDB(g.TENANCY_DB, uuid)
	if err != nil {
		fmt.Printf("初始化商户错误： %v\n", err)
		return
	}

	db, _ := g.TENANCY_DB.DB()
	db.Close()
	multi.AuthDriver.Close()

	os.Exit(code)
}
