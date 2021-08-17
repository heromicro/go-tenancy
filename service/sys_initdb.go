package service

import (
	"database/sql"
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

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// WriteConfig 回写配置
func WriteConfig(viper *viper.Viper, mysql config.Mysql) error {
	g.TENANCY_CONFIG.Mysql = mysql
	cs := utils.StructToMap(g.TENANCY_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

// WriteCacheTypeConfig 回写配置
func WriteCacheTypeConfig(viper *viper.Viper, system config.System) error {
	g.TENANCY_CONFIG.System = system
	cs := utils.StructToMap(g.TENANCY_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

// WriteRedisConfig 回写配置
func WriteRedisConfig(viper *viper.Viper, redis config.Redis) error {
	g.TENANCY_CONFIG.Redis = redis
	cs := utils.StructToMap(g.TENANCY_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
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
	level := conf.Level
	if level == "" {
		level = "release"
	}
	env := conf.Level
	if env == "" {
		env = "pro"
	}
	addr := conf.Addr
	if addr == 0 {
		addr = 80
	}
	BaseSystem := config.System{
		CacheType:   conf.CacheType,
		Level:       level,
		Env:         env,
		Addr:        addr,
		OssType:     "local",
		DbType:      conf.SqlType,
		AdminPreix:  "/admin",
		ClientPreix: "/merchant",
	}
	if err := WriteCacheTypeConfig(g.TENANCY_VP, BaseSystem); err != nil {
		return err
	}
	if BaseSystem.CacheType == "redis" {
		BaseCache := config.Redis{
			DB:       0,
			Addr:     fmt.Sprintf("%s:%s", conf.Cache.Host, conf.Cache.Port),
			Password: conf.Cache.Password,
		}
		if err := WriteRedisConfig(g.TENANCY_VP, BaseCache); err != nil {
			return err
		}
		g.TENANCY_CACHE = cache.Cache() // redis缓存
		err := multi.InitDriver(&multi.Config{
			DriverType:      g.TENANCY_CONFIG.System.CacheType,
			UniversalClient: g.TENANCY_CACHE})
		if err != nil {
			g.TENANCY_LOG.Error("初始化缓存驱动:", zap.Any("err", err))
		}
	}

	BaseMysql := config.Mysql{
		Path:     "",
		Dbname:   "",
		Username: "",
		Password: "",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if conf.Sql.Host == "" {
		conf.Sql.Host = "127.0.0.1"
	}

	if conf.Sql.Port == "" {
		conf.Sql.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.Sql.UserName, conf.Sql.Password, conf.Sql.Host, conf.Sql.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.Sql.DBName)

	if err := createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", conf.Sql.Host, conf.Sql.Port),
		Dbname:   conf.Sql.DBName,
		Username: conf.Sql.UserName,
		Password: conf.Sql.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if err := WriteConfig(g.TENANCY_VP, MysqlConfig); err != nil {
		return err
	}
	m := g.TENANCY_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	linkDns := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		_ = WriteConfig(g.TENANCY_VP, BaseMysql)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		g.TENANCY_DB = db
	}

	err := g.TENANCY_DB.AutoMigrate(
		model.SysUser{},
		model.AdminInfo{},
		model.GeneralInfo{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.SysRegion{},
		model.SysOperationRecord{},
		model.SysTenancy{},
		model.SysMini{},
		model.SysConfig{},
		model.SysConfigCategory{},
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
		_ = WriteConfig(g.TENANCY_VP, BaseMysql)
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
		_ = WriteConfig(g.TENANCY_VP, BaseMysql)
		return err
	}
	return nil
}
