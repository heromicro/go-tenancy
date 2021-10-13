/*
migration
基于 https://github.com/go-gormigrate/gormigrate 实现
用于数据库迁移使用，每次使用都需要修改 GormMysql() 方法内的数据库连接信息。
*/

package main

import (
	"github.com/snowlyg/go-tenancy/migration"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/spf13/cobra"
)

var MigrationId, DSN, DbName string
var Seed bool

func main() {
	var cmdSeed = &cobra.Command{
		Use:   "seed",
		Short: "exec seed datas",
		Long:  `exec seed  datas which are you defined in source`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return source.RunSeed()
		},
	}

	var cmdRun = &cobra.Command{
		Use:   "migrate",
		Short: "exec run migration",
		Long:  `exec run  migrations which are you writed in  migrate.go file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := migration.CreateTable(DSN, DbName); err != nil {
				return err
			}
			return migration.GetMigrate(DSN, DbName).Migrate()
		},
	}

	var cmdRollback = &cobra.Command{
		Use:   "rollback",
		Short: "exec rollback",
		Long:  `exec rollback migrate command which are you execed`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if MigrationId == "" {
				return migration.GetMigrate(DSN, DbName).RollbackLast()
			}
			return migration.GetMigrate(DSN, DbName).RollbackTo(MigrationId)
		},
	}
	cmdRollback.PersistentFlags().StringVarP(&MigrationId, "to", "t", "", "Rollback to migration id")
	cmdRun.PersistentFlags().BoolVarP(&Seed, "seed", "s", true, "Seed data to database")

	var rootCmd = &cobra.Command{Use: "go-tenancy"}
	rootCmd.AddCommand(cmdRun, cmdRollback, cmdSeed)
	rootCmd.PersistentFlags().StringVarP(&DSN, "dsn", "d", "", "Dsn example [xxx:xxx@tcp(127.0.0.1:3306)/]")
	rootCmd.PersistentFlags().StringVarP(&DbName, "db-name", "db", "", "Database name")
	rootCmd.Execute()
}
