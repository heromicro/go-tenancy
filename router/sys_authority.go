package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitAuthorityRouter(Router iris.Party) {
	AuthorityRouter := Router.Party("/authority")
	{
		AuthorityRouter.Post("/createAuthority", v1.CreateAuthority)                 // 创建角色
		AuthorityRouter.Post("/deleteAuthority", v1.DeleteAuthority)                 // 删除角色
		AuthorityRouter.Put("/updateAuthority", v1.UpdateAuthority)                  // 更新角色
		AuthorityRouter.Post("/copyAuthority", v1.CopyAuthority)                     // 更新角色
		AuthorityRouter.Post("/getAuthorityList", v1.GetAuthorityList)               // 获取角色列表
		AuthorityRouter.Post("/getAdminAuthorityList", v1.GetAdminAuthorityList)     // 获取员工角色列表
		AuthorityRouter.Post("/getTenancyAuthorityList", v1.GetTenancyAuthorityList) // 获取商户角色列表
		AuthorityRouter.Post("/getGeneralAuthorityList", v1.GetGeneralAuthorityList) // 获取普通用户角色列表
		AuthorityRouter.Post("/setDataAuthority", v1.SetDataAuthority)               // 设置角色资源权限
	}
}
