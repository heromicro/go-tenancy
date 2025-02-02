package admin

import (
	"github.com/gin-gonic/gin"

	"github.com/snowlyg/go-tenancy/model/response"

	"github.com/snowlyg/go-tenancy/service"

	"github.com/snowlyg/go-tenancy/model/request"

	"github.com/snowlyg/go-tenancy/model"

	"github.com/snowlyg/go-tenancy/g"
	"go.uber.org/zap"
)

// CreateApi 创建基础api
func CreateApi(ctx *gin.Context) {
	var api model.SysApi
	if errs := ctx.ShouldBindJSON(&api); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if api, err := service.CreateApi(api); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(gin.H{"id": api.ID, "path": api.Path, "method": api.Method}, "创建成功", ctx)
	}
}

// DeleteApi 删除api
func DeleteApi(ctx *gin.Context) {
	var req request.DeleteApi
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteApi(req); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// GetApiList 分页获取API列表
func GetApiList(ctx *gin.Context) {
	var pageInfo request.SearchApiParams
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if list, total, err := service.GetAPIInfoList(pageInfo); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
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

// GetApiById 根据id获取api
func GetApiById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	api, err := service.GetApiById(req.Id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(api, ctx)
	}
}

// UpdateApi 更新基础api
func UpdateApi(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var api model.SysApi
	if errs := ctx.ShouldBindJSON(&api); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.UpdateApi(api, req.Id); err != nil {
		g.TENANCY_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", ctx)
	} else {
		response.OkWithMessage("修改成功", ctx)
	}
}

// GetAllApis 获取所有的Api不分页
func GetAllApis(ctx *gin.Context) {
	var req request.AuthorityType
	if err := ctx.ShouldBind(&req); err != nil {
		g.TENANCY_LOG.Error("参数校验不通过", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", ctx)
		return
	}
	if apis, err := service.GetAllApis(req); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(apis, "获取成功", ctx)
	}
}

// DeleteApisByIds 删除选中Api
func DeleteApisByIds(ctx *gin.Context) {
	var ids request.IdsReq
	_ = ctx.ShouldBindJSON(&ids)
	if err := service.DeleteApisByIds(ids); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
