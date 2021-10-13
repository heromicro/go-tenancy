/*
migration
基于 https://github.com/go-gormigrate/gormigrate 实现
用于数据库迁移使用，每次使用都需要修改 GormMysql() 方法内的数据库连接信息。
*/

package migration

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/snowlyg/go-tenancy/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrExceptDsn = errors.New("DSN 参数格式错误")
)

// GetMigrate
// - 需要连接数据库，执行迁移或者回滚
func GetMigrate(dsn, dbName string) *gormigrate.Gormigrate {
	dsn = dsn + dbName
	m := gormigrate.New(GormMysql(dsn), gormigrate.DefaultOptions, []*gormigrate.Migration{})
	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			model.SysUser{},
			model.CUser{},
			model.TUser{},
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
			model.FinancialRecord{},
		)
		if err != nil {
			return err
		}
		return nil
	})
	return m

}

// GormMysql 初始化Mysql数据库
func GormMysql(dsn string) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       dsn + "?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		// 禁用外键约束,自动创建。 否则迁移会报错
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(0)
		sqlDB.SetMaxOpenConns(0)
		return db
	}
}

// CreateTable 创建数据库(mysql)
func CreateTable(dsn, tableName string) error {
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", tableName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("open sql error %w", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	if err = db.Ping(); err != nil {
		return fmt.Errorf("db ping error %w", err)
	}
	_, err = db.Exec(createSql)
	if err != nil {
		return fmt.Errorf("create database %s error %w", tableName, err)
	}
	return nil
}

// DorpDB 删除数据表和视图
func DorpDB(db *gorm.DB, uuid string) error {
	// 删除表和视图
	var sqls []string
	if err := db.Raw("select CASE table_type WHEN 'VIEW' THEN concat('drop view ', table_name, ';') ELSE concat('drop table ', table_name, ';') END  from information_schema.tables where table_schema='%s';", uuid).Scan(&sqls).Error; err != nil {
		return err
	}
	for _, sql := range sqls {
		if err := db.Exec(sql).Error; err != nil {
			continue
		}
	}
	if err := db.Exec(fmt.Sprintf("drop database if exists `%s`;", uuid)).Error; err != nil {
		return err
	}
	return nil
}
