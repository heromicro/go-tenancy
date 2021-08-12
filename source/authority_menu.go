package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

//@description: authority_menu 视图数据初始化
func (a *authorityMenu) Init() error {
	db := g.TENANCY_DB.Find(&[]model.SysMenu{})
	err := db.Error
	if err == nil && db.RowsAffected > 0 {
		color.Danger.Println("\n[Mysql] --> authority_menu 视图已存在!")
		return nil
	}
	if err := g.TENANCY_DB.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`, `sys_base_menus`.`updated_at` AS `updated_at`, `sys_base_menus`.`deleted_at` AS `deleted_at`, `sys_base_menus`.`pid` AS `pid`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`menu_name` AS `menu_name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`route` AS `route`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_base_menus`.`is_tenancy` AS `is_tenancy`,`sys_base_menus`.`is_menu` AS `is_menu`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id` from (`sys_authority_menus` join `sys_base_menus` on ((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)))").Error; err != nil {
		return err
	}
	color.Info.Println("\n[Mysql] --> authority_menu 视图创建成功!")
	return nil
}
