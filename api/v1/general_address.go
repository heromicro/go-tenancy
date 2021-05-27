package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// CreateAddress
func CreateAddress(ctx iris.Context) {
	var address request.CreateAddress
	if errs := utils.Verify(ctx.ReadJSON(&address)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if returnAddress, err := service.CreateAddress(address, user_id); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getAddressMap(returnAddress), "创建成功", ctx)
	}
}

// UpdateAddress
func UpdateAddress(ctx iris.Context) {
	var address request.UpdateAddress
	if errs := utils.Verify(ctx.ReadJSON(&address)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnAddress, err := service.UpdateAddress(address); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithDetailed(getAddressMap(returnAddress), "更新成功", ctx)
	}
}

// getAddressMap
func getAddressMap(returnAddress model.GeneralAddress) context.Map {
	return iris.Map{"id": returnAddress.ID, "name": returnAddress.Name, "phone": returnAddress.Phone, "sex": returnAddress.Sex, "country": returnAddress.Country, "province": returnAddress.Province, "city": returnAddress.City, "district": returnAddress.District, "isDefault": returnAddress.IsDefault, "detail": returnAddress.Detail, "postcode": returnAddress.Postcode, "age": returnAddress.Age, "hospitalName": returnAddress.HospitalName, "locName": returnAddress.LocName, "bedNum": returnAddress.BedNum, "hospitalNo": returnAddress.HospitalNO, "disease": returnAddress.Disease}
}

// GetAddressList
func GetAddressList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if list, total, err := service.GetAddressInfoList(pageInfo, user_id); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
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

// GetAddressById
func GetAddressById(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	address, err := service.GetAddressByID(reqId.Id, user_id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(address, ctx)
	}
}

// DeleteAddress
func DeleteAddress(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if err := service.DeleteAddress(reqId.Id, user_id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
