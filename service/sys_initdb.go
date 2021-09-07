package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	baseMysql = config.Mysql{
		Path:     "",
		Dbname:   "",
		Username: "",
		Password: "",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
	baseSystem = config.System{
		CacheType:   "",
		Level:       "release",
		Env:         "pro",
		Addr:        8089,
		DbType:      "",
		AdminPreix:  "/admin",
		ClientPreix: "/merchant",
	}
	baseCache = config.Redis{
		DB:       0,
		Addr:     "",
		Password: "",
	}
)

// writeConfig 回写配置
func writeConfig() error {
	cs := utils.StructToMap(g.TENANCY_CONFIG)
	for k, v := range cs {
		g.TENANCY_VP.Set(k, v)
	}
	return g.TENANCY_VP.WriteConfig()
}

// refreshConfig 回写配置
func refreshConfig() error {
	g.TENANCY_CONFIG.System = baseSystem
	g.TENANCY_CONFIG.Mysql = baseMysql
	g.TENANCY_CONFIG.Redis = baseCache
	cs := utils.StructToMap(g.TENANCY_CONFIG)
	for k, v := range cs {
		g.TENANCY_VP.Set(k, v)
	}
	g.TENANCY_DB = nil
	g.TENANCY_CACHE = nil
	multi.AuthDriver = nil
	return g.TENANCY_VP.WriteConfig()
}

// createTable 创建数据库(mysql)
func createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func initDB(InitDBFunctions ...model.InitDBFunc) error {
	for _, v := range InitDBFunctions {
		err := v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// InitDB 创建数据库并初始化
func InitDB(conf request.InitDB) error {
	if conf.Level == "" {
		conf.Level = "release"
	}
	if conf.Env == "" {
		conf.Env = "pro"
	}
	if conf.Addr == 0 {
		conf.Addr = 8089
	}

	if conf.CacheType == "redis" {
		g.TENANCY_CONFIG.System.CacheType = conf.CacheType
		g.TENANCY_CONFIG.System.Env = conf.Env
		g.TENANCY_CONFIG.System.Addr = conf.Addr
		g.TENANCY_CONFIG.Redis.Addr = fmt.Sprintf("%s:%s", conf.Cache.Host, conf.Cache.Port)
		g.TENANCY_CONFIG.Redis.Password = conf.Cache.Password
		g.TENANCY_CONFIG.Redis.PoolSize = conf.Cache.PoolSize
		g.TENANCY_CACHE = cache.Cache() // redis缓存
		err := multi.InitDriver(&multi.Config{
			DriverType:      g.TENANCY_CONFIG.System.CacheType,
			UniversalClient: g.TENANCY_CACHE})
		if err != nil {
			g.TENANCY_LOG.Error("初始化缓存驱动:", zap.Any("err", err))
			return fmt.Errorf("初始化缓存驱动失败 %w", err)
		}
		if multi.AuthDriver == nil {
			refreshConfig()
		}
	}

	if conf.Sql.Host == "" {
		conf.Sql.Host = "127.0.0.1"
	}

	if conf.Sql.Port == "" {
		conf.Sql.Port = "3306"
	}

	g.TENANCY_CONFIG.Mysql.Dbname = conf.Sql.DBName
	g.TENANCY_CONFIG.Mysql.Username = conf.Sql.UserName
	g.TENANCY_CONFIG.Mysql.Password = conf.Sql.Password
	g.TENANCY_CONFIG.Mysql.Path = fmt.Sprintf("%s:%s", conf.Sql.Host, conf.Sql.Port)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.Sql.UserName, conf.Sql.Password, conf.Sql.Host, conf.Sql.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.Sql.DBName)

	if err := createTable(dsn, "mysql", createSql); err != nil {
		refreshConfig()
		return err
	}

	m := g.TENANCY_CONFIG.Mysql
	if m.Dbname == "" {
		refreshConfig()
		return errors.New("数据库名称为空")
	}

	mysqlConfig := mysql.Config{
		DSN:                       g.TENANCY_CONFIG.Mysql.Dsn(), // DSN data source name
		DefaultStringSize:         191,                          // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                         // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                        // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		refreshConfig()
		return err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		g.TENANCY_DB = db
	}

	err := MysqlTables(g.TENANCY_DB)
	if err != nil {
		refreshConfig()
		return err
	}
	err = initDB(
		source.Admin,
		source.Api,
		source.AuthorityMenu,
		source.Authority,
		source.AuthoritiesMenus,
		source.Casbin,
		source.DataAuthorities,
		source.BaseMenu,
		source.Region,
		source.Config,
		source.SysConfigCategory,
		source.SysConfigValue,
	)
	if err != nil {
		refreshConfig()
		return err
	}
	writeConfig()
	return nil
}

// MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.SysUser{},
		model.AdminInfo{},
		model.GeneralInfo{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.SysOperationRecord{},
		model.SysTenancy{},
		model.SysRegion{},
		model.SysMini{},
		model.SysConfigCategory{},
		model.SysConfig{},
		model.SysConfigValue{},
		model.SysBrandCategory{},
		model.SysBrand{},
		model.Patient{},

		model.TenancyMedia{},
		model.ProductCategory{},
		model.AttrTemplate{},
		model.Product{},
		model.ProductProductCate{},
		model.ProductContent{},
		model.ProductAttrValue{},
		model.ProductAttr{},
		model.ProductReply{},
		model.ShippingTemplate{},
		model.ShippingTemplateFree{},
		model.ShippingTemplateRegion{},
		model.ShippingTemplateUndelivery{},

		model.Cart{},
		model.Express{},

		model.Order{},
		model.OrderStatus{},
		model.OrderReceipt{},
		model.OrderProduct{},
		model.GroupOrder{},
		model.CartOrder{},

		model.RefundOrder{},
		model.RefundProduct{},
		model.RefundStatus{},

		model.UserAddress{},
		model.UserReceipt{},
		model.UserBill{},
		model.UserExtract{},
		model.UserGroup{},
		model.UserLabel{},
		model.UserUserLabel{},
		model.LabelRule{},
		model.UserMerchant{},
		model.UserRecharge{},
		model.UserRelation{},
		model.UserVisit{},
		model.Mqtt{},
		model.MqttRecord{},
	)
	if err != nil {
		g.TENANCY_LOG.Error("register table failed", zap.Any("err", err))
		return err
	}
	g.TENANCY_LOG.Info("register table success")
	return nil
}
