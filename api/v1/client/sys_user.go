package client

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// ChangeLoginPasswordMap 修改用户密码
func ChangeLoginPasswordMap(ctx *gin.Context) {
	if detail, err := service.ChangePasswordMap(multi.GetUserId(ctx), ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// ChangeProfileMap 修改用户密码表单
func ChangeProfileMap(ctx *gin.Context) {
	if detail, err := service.ChangeProfileMap(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// ChangePasswordMap 修改用户密码表单
func ChangePasswordMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.ChangePasswordMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// RegisterAdminMap 注册用户表单
func RegisterAdminMap(ctx *gin.Context) {
	if detail, err := service.RegisterClientMap(0, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// UpdateAdminMap 更新用户表单
func UpdateAdminMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.RegisterClientMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// RegisterTenancy 商户注册
func RegisterTenancy(ctx *gin.Context) {
	var req request.Register
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if req.Password != req.ConfirmPassword {
		response.FailWithMessage("两次输入密码不一致", ctx)
		return
	}
	if len(req.AuthorityId) == 0 || req.AuthorityId[0] == "" {
		response.FailWithMessage("用户角色为必选参数", ctx)
		return
	}
	id, err := service.RegisterClient(req, multi.TenancyAuthority, multi.GetTenancyId(ctx))
	if err != nil {
		g.TENANCY_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{"user_id": id}, "注册成功", ctx)
	}
}

// ChangePassword 用户修改密码
func ChangePassword(ctx *gin.Context) {
	var req request.ChangePassword
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		response.FailWithMessage("两次输入密码不一致", ctx)
		return
	}
	err := service.ChangeClientPassword(multi.GetUserId(ctx), req, multi.GetAuthorityType(ctx))
	if err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// ChangeProfile 用户修改信息
func ChangeProfile(ctx *gin.Context) {
	var user request.ChangeProfile
	if errs := ctx.ShouldBindJSON(&user); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	err := service.ChangeClientProfile(user, multi.GetUserId(ctx))
	if err != nil {
		g.TENANCY_LOG.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetAdminList 分页获取用户列表
func GetAdminList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		fmt.Printf("ShouldBindJSON %v\n\n", errs)
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetTenancyInfoList(pageInfo, multi.GetUserId(ctx), multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	jwtId := multi.GetUserId(ctx)
	if jwtId == reqId.Id {
		response.FailWithMessage("删除失败, 自杀失败", ctx)
		return
	}
	if err := service.DeleteClientUser(reqId.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// SetUserInfo 设置用户信息
func SetUserInfo(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	user, err := service.FindClientByStringId(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	var info request.UpdateUser
	_ = ctx.ShouldBindJSON(&info)
	if err := service.UpdateClientInfo(info, user); err != nil {
		g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
		response.FailWithMessage("设置失败", ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}
