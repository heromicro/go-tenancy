package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// Login 用户登录
func Login(ctx iris.Context) {
	var L request.Login
	if errs := utils.Verify(ctx.ReadJSON(&L)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if store.Verify(L.CaptchaId, L.Captcha, true) || g.TENANCY_CONFIG.System.Env == "dev" {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if loginResponse, err := service.Login(U, L.AuthorityType); err != nil {
			g.TENANCY_LOG.Error("登陆失败!", zap.Any("err", err))
			response.FailWithMessage(err.Error(), ctx)
		} else {
			response.OkWithDetailed(loginResponse, "登录成功", ctx)
		}
	} else {
		response.FailWithMessage("验证码错误", ctx)
	}
}

// Logout 退出登录
func Logout(ctx iris.Context) {
	token := multi.GetVerifiedToken(ctx)
	if token == nil {
		response.FailWithMessage("授权凭证为空", ctx)
		return
	}
	err := service.DelToken(string(token))
	if err != nil {
		g.TENANCY_LOG.Error("del token", zap.Any("err", err))
		response.FailWithMessage("退出失败", ctx)
		return
	}
	response.OkWithMessage("退出登录", ctx)
}

// Clean 清空 token
func Clean(ctx iris.Context) {
	waitUse := multi.Get(ctx)
	if waitUse == nil {
		response.FailWithMessage("清空TOKEN失败", ctx)
		return
	}
	err := service.CleanToken(waitUse.ID)
	if err != nil {
		g.TENANCY_LOG.Error("清空TOKEN失败", zap.Any("err", err))
		response.FailWithMessage("清空TOKEN失败", ctx)
		return
	}
	response.OkWithMessage("TOKEN清空", ctx)
}

// Register 用户注册账号
func Register(ctx iris.Context) {
	var R request.Register
	if errs := utils.Verify(ctx.ReadJSON(&R)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	user := &model.SysUser{Username: R.Username, Password: R.Password, AuthorityId: R.AuthorityId}
	userReturn, err := service.Register(*user, R.AuthorityType)
	if err != nil {
		g.TENANCY_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithDetailed(iris.Map{"id": userReturn.ID, "userName": userReturn.Username, "authorityId": userReturn.AuthorityId}, "注册成功", ctx)
	}
}

// ChangePassword 用户修改密码
func ChangePassword(ctx iris.Context) {
	var user request.ChangePasswordStruct
	if errs := utils.Verify(ctx.ReadJSON(&user)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	U := &model.SysUser{Username: user.Username, Password: user.Password}
	err := service.ChangePassword(U, user.NewPassword, user.AuthorityType)
	if err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetAdminList 分页获取用户列表
func GetAdminList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAdminInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetTenancyList 分页获取用户列表
func GetTenancyList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetTenancyInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetGeneralList 分页获取用户列表
func GetGeneralList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetGeneralInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// SetUserAuthority 设置用户权限
func SetUserAuthority(ctx iris.Context) {
	var sua request.SetUserAuth
	if errs := utils.Verify(ctx.ReadJSON(&sua)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.SetUserAuthority(sua.Id, sua.AuthorityId); err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage("修改失败", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// DeleteUser 删除用户
func DeleteUser(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	jwtId := ctx.GetID()
	if jwtId == uint(reqId.Id) {
		response.FailWithMessage("删除失败, 自杀失败", ctx)
		return
	}
	if err := service.DeleteUser(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// SetUserInfo 设置用户信息
func SetUserInfo(ctx iris.Context) {
	userId := ctx.Params().GetIntDefault("user_id", 0)
	user, err := service.FindUserById(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if user.IsAdmin() {
		var admin model.SysAdminInfo
		_ = ctx.ReadJSON(&admin)
		admin.ID = user.AdminInfo.ID
		if _, err := service.SetUserAdminInfo(admin, user.AdminInfo.ID > 0); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithMessage("设置成功", ctx)
		}
	} else if user.IsTenancy() {
		var tenancy model.SysTenancyInfo
		_ = ctx.ReadJSON(&tenancy)
		tenancy.ID = user.TenancyInfo.ID
		if _, err := service.SetUserTenancyInfo(tenancy, user.TenancyInfo.ID > 0); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithMessage("设置成功", ctx)
		}
	} else if user.IsGeneral() {
		var general model.SysGeneralInfo
		_ = ctx.ReadJSON(&general)
		general.ID = user.GeneralInfo.ID
		if _, err := service.SetUserGeneralInfo(general, user.GeneralInfo.ID > 0); err != nil {
			g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
			response.FailWithMessage("设置失败", ctx)
		} else {
			response.OkWithMessage("设置成功", ctx)
		}
	} else {
		g.TENANCY_LOG.Error("未知角色", zap.Any("err", user.AuthorityType()))
		response.FailWithMessage("未知角色", ctx)
	}
}

// // getClaims returns the current authorized client claims.
// func getClaims(ctx iris.Context) *multi.CustomClaims {
// 	waitUse := multi.Get(ctx)
// 	if waitUse == nil {
// 		g.TENANCY_LOG.Error("从Context中获取用户ID失败, 请检查路由是否使用multi中间件")
// 	}
// 	return waitUse
// }

// // getUserAuthorityId 从Context中获取用户角色id
// func getUserAuthorityId(ctx iris.Context) string {
// 	if claims := getClaims(ctx); claims == nil {
// 		g.TENANCY_LOG.Error("从Context中获取用户角色id失败, 请检查路由是否使用multi中间件")
// 		return ""
// 	} else {
// 		return claims.AuthorityId
// 	}
// }

// // getUserId 从Context中获取用户id
// func getUserId(ctx iris.Context) int {
// 	if claims := getClaims(ctx); claims == nil {
// 		g.TENANCY_LOG.Error("从Context中获取用户id失败, 请检查路由是否使用multi中间件")
// 		return 0
// 	} else {
// 		id, err := strconv.Atoi(claims.ID)
// 		if err != nil {
// 			g.TENANCY_LOG.Error("strconv atoi ", zap.Any("err", fmt.Errorf("%s strconv atoi %w", claims.ID, err)))
// 			return 0
// 		}
// 		return id
// 	}
// }

// // getTenancyId 从Context中获取商户id
// func getTenancyId(ctx iris.Context) int {
// 	if claims := getClaims(ctx); claims == nil {
// 		g.TENANCY_LOG.Error("从Context中获取商户id失败, 请检查路由是否使用multi中间件")
// 		return 0
// 	} else {
// 		return claims.TenancyId
// 	}
// }

// // getTenancyName 从Context中获取商户id
// func getTenancyName(ctx iris.Context) string {
// 	if claims := getClaims(ctx); claims == nil {
// 		g.TENANCY_LOG.Error("从Context中获取商户名称失败, 请检查路由是否使用multi中间件")
// 		return ""
// 	} else {
// 		return claims.TenancyName
// 	}
// }
