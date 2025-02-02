package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// GetGeneralSelect 获取c用户下拉选项
func GetGeneralSelect(ctx *gin.Context) {
	var req request.GetByTenancyId
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if users, err := service.GetGeneralSelect(req.TenancyId); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(users, "获取成功", ctx)
	}
}

// GetGeneralList 分页获取c用户列表
func GetGeneralList(ctx *gin.Context) {
	var pageInfo request.UserPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetGeneralInfoList(pageInfo, ctx); err != nil {
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

// UpdateUserMap 设置余额表单
func UpdateUserMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.UpdateUserMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// UpdateUser 设置余额
func UpdateUser(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var update response.GeneralUserDetail
	if errs := ctx.ShouldBindJSON(&update); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.UpdateUser(req.Id, multi.GetTenancyId(ctx), update); err != nil {
		g.TENANCY_LOG.Error("操作失败", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// SetNowMoneyMap 设置余额表单
func SetNowMoneyMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.SetNowMoneyMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// SetNowMoney 设置余额
func SetNowMoney(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var setNowMoney request.SetNowMoney
	if errs := ctx.ShouldBindJSON(&setNowMoney); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.SetNowMoney(req.Id, multi.GetTenancyId(ctx), setNowMoney); err != nil {
		g.TENANCY_LOG.Error("操作失败", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// BatchSetUserGroupMap 设置用户分组表单
func BatchSetUserGroupMap(ctx *gin.Context) {
	var setUserGroup request.BatchSetUserGroup
	if errs := ctx.ShouldBindJSON(&setUserGroup); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.BatchSetUserGroupMap(setUserGroup.Ids, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// BatchSetUserGroup 设置用户分组
func BatchSetUserGroup(ctx *gin.Context) {
	var setUserGroup request.SetUserGroup
	if errs := ctx.ShouldBindJSON(&setUserGroup); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.BatchSetUserGroup(setUserGroup); err != nil {
		g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
		response.FailWithMessage("设置失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// BatchSetUserLabelMap 设置用户分组表单
func BatchSetUserLabelMap(ctx *gin.Context) {
	var setUserGroup request.BatchSetUserLabel
	if errs := ctx.ShouldBindJSON(&setUserGroup); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.BatchSetUserLabelMap(setUserGroup.Ids, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// BatchSetUserLabel 设置用户分组
func BatchSetUserLabel(ctx *gin.Context) {
	var setUserGroup request.SetUserLabel
	if errs := ctx.ShouldBindJSON(&setUserGroup); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.BatchSetUserLabel(setUserGroup, multi.GetTenancyId(ctx)); err != nil {
		g.TENANCY_LOG.Error("设置失败", zap.Any("err", err))
		response.FailWithMessage("设置失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("设置成功", ctx)
	}
}

// SetUserGroupMap 设置用户分组表单
func SetUserGroupMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.SetUserGroupMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// SetUserGroup 设置用户分组
func SetUserGroup(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var setUserGroup request.SetUserGroup
	if errs := ctx.ShouldBindJSON(&setUserGroup); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.SetUserGroup(req.Id, setUserGroup); err != nil {
		g.TENANCY_LOG.Error("操作失败", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// SetUserLabelMap 设置用户分组表单
func SetUserLabelMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.SetUserLabelMap(req.Id, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// SetUserLabel 设置用户分组
func SetUserLabel(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var setUserLabel request.SetUserLabel
	if errs := ctx.ShouldBindJSON(&setUserLabel); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.SetUserLabel(req.Id, multi.GetTenancyId(ctx), setUserLabel.LabelId); err != nil {
		g.TENANCY_LOG.Error("操作失败", zap.Any("err", err))
		response.FailWithMessage("操作失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}

// GetGeneralDetail 用户详情
func GetGeneralDetail(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if detail, err := service.GetGeneralDetail(req.Id); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(detail, "获取成功", ctx)
	}
}

// GetUserOrderList
func GetUserOrderList(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var pageInfo request.OrderPageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	pageInfo.CUserId = req.Id
	if ginH, err := service.GetOrderInfoList(pageInfo, ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(ginH, "获取成功", ctx)
	}
}

func GetBillList(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetUserBillInfoList(pageInfo, req.Id); err != nil {
		g.TENANCY_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}
