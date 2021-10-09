package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

func UpdateUserMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	user, err := GetGeneralDetail(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"uid","value":%d,"title":"会员 ID","props":{"type":"text","placeholder":"请输入会员 ID","disabled":true},"validate":[{"message":"会员数据类型错误","required":true,"type":"number","trigger":"change"}]},{"type":"input","field":"realName","value":"%s","title":"真实姓名","props":{"type":"text","placeholder":"请输入真实姓名"}},{"type":"input","field":"phone","value":"%s","title":"手机号","props":{"type":"text","placeholder":"请输入手机号"}},{"type":"datePicker","field":"birthday","value":"%s","title":"生日","props":{"type":"date","editable":false,"placeholder":"请选择生日"}},{"type":"input","field":"idCard","value":"%s","title":"身份证","props":{"type":"text","placeholder":"请输入身份证"}},{"type":"input","field":"address","value":"%s","title":"用户地址","props":{"type":"text","placeholder":"请输入用户地址"}},{"type":"input","field":"mark","value":"%s","title":"备注","props":{"type":"textarea","placeholder":"请输入备注"}},{"type":"select","field":"groupId","value":%d,"title":"会员分组","props":{"multiple":false,"placeholder":"请选择会员分组"},"options":[]},{"type":"select","field":"labelId","value":[],"title":"会员标签","props":{"multiple":true,"placeholder":"请选择会员标签"},"options":[]}],"action":"","method":"POST","title":"编辑","config":{}}`, user.Uid, user.RealName, user.Phone, user.Birthday, user.IdCard, user.Address, user.Mark, user.GroupId)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/cuser/editUser", id), ctx)
	groupOpts, err := GetUserGroupOptions()
	if err != nil {
		return form, err
	}
	form.Rule[7].Options = groupOpts
	opts, err := GetUserLabelOptions(multi.GetTenancyId(ctx))
	if err != nil {
		return form, err
	}
	form.Rule[8].Value = user.LabelIds
	form.Rule[8].Options = opts
	return form, nil
}

// UpdateUser
func UpdateUser(id, tenancyId uint, req response.GeneralUserDetail) error {
	update := map[string]interface{}{"address": req.Address, "birthday": req.Birthday, "id_card": req.IdCard, "group_id": req.GroupId, "mark": req.Mark, "phone": req.Phone, "real_name": req.RealName}
	err := g.TENANCY_DB.Model(&model.CUser{}).Where("id = ?", id).Updates(update).Error
	if err != nil {
		return err
	}
	err = SetUserLabel(id, tenancyId, req.LabelIds)
	if err != nil {
		return err
	}
	return nil
}
func SetNowMoneyMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	user, err := GetGeneralDetail(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"radio","field":"type","value":1,"title":"修改余额","props":{},"validate":[{"message":"请选择修改余额","required":true,"type":"string","trigger":"change"}],"options":[{"label":"增加","value":1},{"label":"减少","value":2}]},{"type":"inputNumber","field":"nowMoney","value":%f,"title":"金额","props":{"placeholder":"请输入金额","min":0},"validate":[{"message":"请输入金额","required":true,"type":"number","trigger":"change"}]}],"action":"","method":"POST","title":"修改用户余额","config":{}}`, user.NowMoney)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/cuser/setNowMoney", id), ctx)
	return form, nil
}

// SetNowMoney
func SetNowMoney(id, tenancyId uint, req request.SetNowMoney) error {
	user, err := GetGeneralDetail(id)
	if err != nil {
		return err
	}
	nowMoney := user.NowMoney
	// 增加
	if req.Type == 1 {
		nowMoney = user.NowMoney + req.NowMoney
	} else if req.Type == 2 {
		if user.NowMoney <= req.NowMoney {
			nowMoney = 0
		} else {
			nowMoney = user.NowMoney - req.NowMoney
		}
	}
	if err := g.TENANCY_DB.Model(&model.CUser{}).Where("id = ?", id).Updates(map[string]interface{}{"now_money": nowMoney}).Error; err != nil {
		return err
	}
	return nil
}

func GetUserLabelIdsByUserId(id uint) ([]uint, error) {
	var labelIds []uint
	db := g.TENANCY_DB.Model(&model.UserUserLabel{}).Select("user_label_id").Where("c_user_id = ?", id)
	// db = CheckTenancyId(db, tenancyId, "")
	err := db.Find(&labelIds).Error
	if err != nil {
		return labelIds, fmt.Errorf("get label ids %w", err)
	}
	return labelIds, nil
}

func BatchSetUserGroupMap(ids string, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := fmt.Sprintf(`{"rule":[{"type":"hidden","field":"ids","value":[%s]},{"type":"select","field":"group_id","value":"","title":"用户分组","props":{"multiple":false,"placeholder":"请选择用户分组"},"options":[{"label":"不设置","value":"0"},{"value":10,"label":"测试用户"},{"value":11,"label":"普通用户"},{"value":12,"label":"中级用户"},{"value":13,"label":"高级用户"}]}],"action":"","method":"POST","title":"修改用户分组","config":{}}`, ids)

	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction("/cuser/batchSetUserGroup", ctx)
	opts, err := GetUserGroupOptions()
	if err != nil {
		return form, err
	}
	form.Rule[1].Options = opts
	return form, nil
}

// BatchSetUserGroup
func BatchSetUserGroup(req request.SetUserGroup) error {
	if err := g.TENANCY_DB.Model(&model.CUser{}).Where("id in ?", req.Ids).Updates(map[string]interface{}{"group_id": req.GroupId}).Error; err != nil {
		return err
	}
	return nil
}

func BatchSetUserLabelMap(ids string, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := fmt.Sprintf(`{"rule":[{"type":"hidden","field":"ids","value":[%s]},{"type":"select","field":"label_id","value":[],"title":"用户标签","props":{"multiple":true,"placeholder":"请选择用户标签"},"options":[]}],"action":"","method":"POST","title":"修改用户标签","config":{}}`, ids)

	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction("/cuser/batchSetUserLabel", ctx)
	opts, err := GetUserLabelOptions(multi.GetTenancyId(ctx))
	if err != nil {
		return form, err
	}
	form.Rule[1].Options = opts
	return form, nil
}

// BatchSetUserLabel
func BatchSetUserLabel(req request.SetUserLabel, tenancyId uint) error {
	for _, userId := range req.Ids {
		SetUserLabel(userId, tenancyId, req.LabelId)
	}
	return nil
}

func SetUserGroupMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	user, err := GetGeneralDetail(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"hidden","field":"ids","value":[%d]},{"type":"select","field":"group_id","value":%d,"title":"用户分组","props":{"multiple":false,"placeholder":"请选择用户分组"},"options":[]}],"action":"","method":"POST","title":"修改用户分组","config":{}}`, id, user.GroupId)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/cuser/setUserGroup", id), ctx)
	opts, err := GetUserGroupOptions()
	if err != nil {
		return form, err
	}
	form.Rule[1].Options = opts
	return form, nil
}

// SetUserGroup
func SetUserGroup(id uint, req request.SetUserGroup) error {
	if err := g.TENANCY_DB.Model(&model.CUser{}).Where("id = ?", id).Updates(map[string]interface{}{"group_id": req.GroupId}).Error; err != nil {
		return err
	}
	return nil
}

func SetUserLabelMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	user, err := GetGeneralDetail(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"hidden","field":"ids","value":[%d]},{"type":"select","field":"label_id","value":[],"title":"用户标签","props":{"multiple":true,"placeholder":"请选择用户标签"},"options":[]}],"action":"","method":"POST","title":"修改用户标签","config":{}}`, id)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/cuser/setUserLabel", id), ctx)
	opts, err := GetUserLabelOptions(multi.GetTenancyId(ctx))
	if err != nil {
		return form, err
	}
	form.Rule[1].Options = opts
	form.Rule[1].Value = user.LabelIds
	return form, nil
}

// SetUserLabel
func SetUserLabel(id, tenancyId uint, reqlabelIds []uint) error {
	labelIds, err := GetUserLabelIdsByUserId(id)
	if err != nil {
		return err
	}

	// 删除
	var delIds []uint
	for _, labelId := range labelIds {
		isDel := true
		for _, reqLabelId := range reqlabelIds {
			if labelId == reqLabelId {
				isDel = false
				break
			}
		}
		if isDel {
			delIds = append(delIds, labelId)
		}
	}

	if len(delIds) > 0 {
		db := g.TENANCY_DB.Where("c_user_id = ?", id)
		db = CheckTenancyId(db, tenancyId, "")
		err := db.Where("user_label_id in ?", delIds).Delete(&model.UserUserLabel{}).Error
		if err != nil {
			return fmt.Errorf("delete user_user_labels %w", err)
		}
	}

	// 增加
	var addIds []uint
	for _, reqLabelId := range reqlabelIds {
		isAdd := true
		for _, labelId := range labelIds {
			if reqLabelId == labelId {
				isAdd = false
				break
			}
		}
		if isAdd {
			addIds = append(addIds, reqLabelId)
		}
	}

	if len(addIds) > 0 {
		labels := []model.UserUserLabel{}
		for _, addId := range addIds {
			labels = append(labels, model.UserUserLabel{UserLabelID: addId, CUserID: id, SysTenancyID: tenancyId})
		}
		if err = g.TENANCY_DB.Model(&model.UserUserLabel{}).Create(&labels).Error; err != nil {
			return fmt.Errorf("create user_user_labels %w", err)
		}
	}

	return nil
}

func GetGeneralDetail(id uint) (response.GeneralUserDetail, error) {
	var user response.GeneralUserDetail
	generalAuthorityIds, err := GetUserAuthorityIds(multi.GeneralAuthority)
	if err != nil {
		return user, err
	}

	err = g.TENANCY_DB.Model(&model.CUser{}).
		Select("c_users.id as uid,c_users.mark,c_users.real_name,c_users.phone,c_users.address,c_users.id_card,c_users.birthday,c_users.avatar_url,c_users.nick_name,c_users.now_money,c_users.pay_count,c_users.pay_price,c_users.group_id").
		Where("c_users.authority_id IN (?)", generalAuthorityIds).
		Where("c_users.id = ?", id).
		First(&user).Error
	if err != nil {
		return user, fmt.Errorf("get general detail %w", err)
	}

	u, err := GetThisMonthOrderPriceByUserId(id)
	if err != nil {
		return user, err
	}
	user.TotalPayCount = u.TotalPayCount
	user.TotalPayPrice = u.TotalPayPrice

	labelIds, err := GetUserLabelIdsByUserId(user.Uid)
	if err != nil {
		return user, err
	}
	user.LabelIds = labelIds
	return user, nil
}

func GetGeneralSelect(tenancyId uint) ([]response.SelectOption, error) {
	selects := []response.SelectOption{
		{ID: 0, Name: "请选择"},
	}
	var userSelects []response.SelectOption
	err := g.TENANCY_DB.Model(&model.CUser{}).
		Select("c_users.id as id,c_users.nick_name as name").
		Where("status = ?", g.StatusTrue).
		Where("is_show = ?", g.StatusTrue).
		Find(&userSelects).Error
	selects = append(selects, userSelects...)
	return selects, err
}

// GetGeneralInfoList 分页获取数据
func GetGeneralInfoList(info request.UserPageInfo, ctx *gin.Context) ([]response.GeneralUser, int64, error) {
	tenancyId := multi.GetTenancyId(ctx)
	userList := []response.GeneralUser{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	generalAuthorityIds, err := GetUserAuthorityIds(multi.GeneralAuthority)
	if err != nil {
		return userList, 0, err
	}

	db := g.TENANCY_DB.Model(&model.CUser{})
	if multi.IsTenancy(ctx) {
		db = db.Select("c_user.sex,c_user.nick_name,c_user.avatar_url,c_user.user_type,c_user.id as uid,c_user.username,c_user.authority_id,c_user.created_at,c_user.updated_at,sys_authorities.authority_name,sys_authorities.authority_type,c_user.authority_id,user_groups.group_name,user_merchants.first_pay_time,user_merchants.last_pay_time").
			Joins("left join user_merchants on user_merchants.c_user_id = c_user.id").
			Where("user_merchants.sys_tenancy_id = ?", tenancyId)
	} else {
		db = db.Select("c_user.id as uid,c_user.username,c_user.authority_id,c_user.created_at,c_user.updated_at, c_user.*,sys_authorities.authority_name,sys_authorities.authority_type,c_user.authority_id,user_groups.group_name")
	}
	db = db.Joins("left join sys_authorities on sys_authorities.authority_id = c_user.authority_id").
		Joins("left join user_groups on c_user.group_id = user_groups.id").
		Where("c_user.authority_id IN (?)", generalAuthorityIds)

	if info.UserTimeType != "" && info.UserTime != "" {
		userTimes := strings.Split(info.UserTime, "-")
		start, err := time.Parse("2006/01/02", userTimes[0])
		if err != nil {
			return userList, total, fmt.Errorf("parse time %w", err)
		}
		end, err := time.Parse("2006/01/02", userTimes[1])
		if err != nil {
			return userList, total, fmt.Errorf("parse time %w", err)
		}
		if info.UserTimeType == "add_time" {
			db = db.Where("c_users.created_at BETWEEN ? AND ?", start, end)
		} else if info.UserTimeType == "visit" {
			db = db.Where("c_users.last_time BETWEEN ? AND ?", start, end)
		}
	}

	if info.PayCount != "" {
		if info.PayCount == "0" {
			db = db.Where("c_users.pay_count = ?", info.PayCount)
		} else {
			db = db.Where("c_users.pay_count >= ?", info.PayCount)
		}
	}
	if info.GroupId != "" {
		db = db.Where("c_users.group_id = ?", info.GroupId)
	}
	if info.LabelId != "" {
		userIds, err := GetUserIdsByLabelId(info.LabelId, tenancyId)
		if err != nil {
			return userList, total, err
		}
		db = db.Where("c_users.id in ?", userIds)
	}
	if info.Sex != "" {
		db = db.Where("c_users.sex = ?", info.Sex)
	}
	if info.NickName != "" {
		db = db.Where("c_users.nick_name like ?", info.NickName+"%")
	}
	if info.UserType != "" {
		db = db.Where("c_users.user_type = ?", info.UserType)
	}

	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}

		db = db.Limit(limit).Offset(offset)
	}
	db = OrderBy(db, info.OrderBy, info.SortBy, "c_users.")
	err = db.Find(&userList).Error
	if err != nil {
		return userList, total, err
	}

	if len(userList) > 0 {
		userList, err = getCuserLabels(userList, tenancyId)
		if err != nil {
			return userList, total, err
		}
	}

	return userList, total, err
}

func getCuserLabels(userList []response.GeneralUser, tenancyId uint) ([]response.GeneralUser, error) {
	userIds := []uint{}
	for _, user := range userList {
		userIds = append(userIds, user.Uid)
	}
	userLabels, err := GetUserLabelByUserIds(userIds, tenancyId)
	if err != nil {
		return userList, err
	}
	for i := 0; i < len(userList); i++ {
		userList[i].Label = []string{}
		for _, userLabel := range userLabels {
			if userLabel.SysUserID == userList[i].Uid {
				userList[i].Label = append(userList[i].Label, userLabel.LabelName)
			}
		}
	}
	return userList, nil
}

func GetUserIdsByLabelId(labelId string, tenancyId uint) ([]uint, error) {
	var userIds []uint
	db := g.TENANCY_DB.Model(&model.UserUserLabel{}).Select("c_user_id").Where("user_label_id = ?", labelId)
	db = CheckTenancyId(db, tenancyId, "")
	err := db.Find(&userIds).Error
	if err != nil {
		return userIds, fmt.Errorf("get user ids by label id %w", err)
	}
	return userIds, nil
}

func GetUserIdsByNickname(nickname string, tenancyId uint) ([]uint, error) {
	userIds := []uint{}
	cuserAuthorityIds, err := GetUserAuthorityIds(multi.GeneralAuthority)
	if err != nil {
		return userIds, err
	}
	db := g.TENANCY_DB.Model(&model.CUser{}).
		Select("c_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = c_users.authority_id").
		Joins("left join sys_tenancies on c_users.sys_tenancy_id = sys_tenancies.id").
		Where("c_users.authority_id IN (?)", cuserAuthorityIds).
		Where("c_users.nick_name like ?", nickname+"%")
	if tenancyId > 0 {
		db = db.Joins("left join user_merchants on user_merchants.c_user_id = c_users.id").
			Where("user_merchants.sys_tenancy_id = ?", tenancyId)
	}
	err = db.Find(&userIds).Error
	return userIds, err
}

// GetCuserByUserIds
func GetCuserByUserIds(userIds []uint, tenancyId uint) ([]response.SysGeneralUser, error) {
	userList := []response.SysGeneralUser{}
	cuserAuthorityIds, err := GetUserAuthorityIds(multi.GeneralAuthority)
	if err != nil {
		return userList, err
	}
	db := g.TENANCY_DB.Model(&model.CUser{}).
		Select("c_users.id,c_users.status,c_users.username,c_users.authority_id,c_users.created_at,c_users.updated_at, c_users.email, c_users.phone, c_users.nick_name, c_users.avatar_url,sys_authorities.authority_name,sys_authorities.authority_type,c_users.authority_id,sys_tenancies.name as tenancy_name").
		Joins("left join sys_authorities on sys_authorities.authority_id = c_users.authority_id").
		Joins("left join sys_tenancies on c_users.sys_tenancy_id = sys_tenancies.id").
		Where("c_users.authority_id IN (?)", cuserAuthorityIds).
		Where("c_users.id IN (?)", userIds)
	if tenancyId > 0 {
		db = db.Joins("left join user_merchants on user_merchants.c_user_id = c_users.id").Where("user_merchants.sys_tenancy_id = ?", tenancyId)
	}
	err = db.Find(&userList).Error
	return userList, err
}

// GetCUserNum 获取c用户数量
func GetCUserNum(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var userNum int64
	db := g.TENANCY_DB.Model(&model.CUser{})
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Count(&userNum).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return userNum, nil
}

// GetVisitUserNum 用户浏览记录，查询用户浏览商品记录
func GetVisitUserNum(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var userNum int64
	db := g.TENANCY_DB.Model(&model.UserVisit{}).
		Joins("left join products on products.id = user_visits.type_id").
		Where("user_visits.type = ?", "product")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Count(&userNum).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return userNum, nil
}

// GetVisitUserNumGroup 用户浏览记录，查询用户浏览商品记录
func GetVisitUserNumGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]*response.MerchantVisitData, error) {
	var visitData []*response.MerchantVisitData
	db := g.TENANCY_DB.Model(&model.UserVisit{}).
		Select("count(user_visits.type) as total,products.sys_tenancy_id as tenancy_id,sys_tenancies.name as tenancy_name").
		Joins("left join products on products.id = user_visits.type_id").
		Joins("left join sys_tenancies on products.sys_tenancy_id = sys_tenancies.id").
		Where("user_visits.type = ?", "product")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Limit(7).Group("tenancy_id").Order("total DESC").Find(&visitData).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return visitData, nil
}

// GetVisitProductNumGroup 用户浏览记录，查询用户浏览商品记录
func GetVisitProductNumGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]*response.ProductVisitData, error) {
	var visitData []*response.ProductVisitData
	db := g.TENANCY_DB.Model(&model.UserVisit{}).
		Select("count(user_visits.type) as total,products.image as image,products.store_name as store_name").
		Joins("left join products on products.id = user_visits.type_id").
		Where("user_visits.type = ?", "product")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Limit(7).Group("type_id").Order("total DESC").Find(&visitData).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return visitData, nil
}

// GetVisitNum 浏览记录，查询用户浏览产品和小程序记录
func GetVisitNum(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var userNum int64
	db := g.TENANCY_DB.Model(&model.UserVisit{}).
		Joins("left join products on products.id = user_visits.type_id").
		Where("user_visits.type in ?", []string{"page", "smallProgram"})
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Count(&userNum).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return userNum, nil
}

// GetUserNumGroup 用户数量
func GetUserNumGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]*response.UserData, error) {
	var userData []*response.UserData
	db := g.TENANCY_DB.Model(&model.CUser{}).
		Select("from_unixtime(unix_timestamp(created_at),'%m-%d') as time, count(id) as new")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Group("time").Order("time asc").Find(&userData).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return userData, nil
}

// GetVisitNumGroup 访客用户数量
func GetVisitNumGroup(scopes ...func(*gorm.DB) *gorm.DB) ([]*response.UserData, error) {
	var userData []*response.UserData
	db := g.TENANCY_DB.Model(&model.UserVisit{}).
		Select("from_unixtime(unix_timestamp(created_at),'%m-%d') as time, count(id) as visit")
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Group("time").Order("time asc").Find(&userData).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return userData, nil
}

// GetLikeStore 商户关注用户数
func GetLikeStore(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := g.TENANCY_DB.Model(&model.UserRelation{}).Where("type = ?", 10)
	if len(scopes) > 0 {
		db = db.Scopes(scopes...)
	}
	err := db.Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return count, err
	}
	return count, nil
}
