package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

var ConfigTypes = []Option{
	{Value: "input", Label: "文本框"},
	{Value: "number", Label: "数字框"},
	{Value: "textarea", Label: "多行文本框"},
	{Value: "radio", Label: "单选框"},
	{Value: "checkbox", Label: "多选框"},
	{Value: "select", Label: "下拉框"},
	{Value: "file", Label: "文件上传"},
	{Value: "image", Label: "图片上传"},
	{Value: "color", Label: "颜色选择框"},
}

func GetConfigTypeName(value string) string {
	for i := 0; i < len(ConfigTypes); i++ {
		if ConfigTypes[i].Value.(string) == value {
			return ConfigTypes[i].Label
		}
	}
	return ""
}

type Option struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Children []Option    `json:"children"`
}

type Opt struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type StringOpt struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// filterDate
func filterDate(db *gorm.DB, date, perfix string) *gorm.DB {
	field := "created_at"
	if perfix != "" {
		field = fmt.Sprintf("%s.created_at", perfix)
	}
	dates := strings.Split(date, "-")
	if len(dates) == 2 {
		start, _ := time.Parse("2006/01/02", dates[0])
		end, _ := time.Parse("2006/01/02", dates[1])
		return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), start, end)
	}
	if len(dates) == 1 {
		// { text: '今天', val: 'today' },
		// { text: '昨天', val: 'yesterday' },
		// { text: '最近7天', val: 'lately7' },
		// { text: '最近30天', val: 'lately30' },
		// { text: '本月', val: 'month' },
		// { text: '本年', val: 'year' }
		// TODO: 使用内置函数，可能造成索引失效
		switch dates[0] {
		case "today":
			return db.Where(fmt.Sprintf("TO_DAYS(NOW()) - TO_DAYS(%s) < 1", field))
		case "yesterday":
			return db.Where(fmt.Sprintf("TO_DAYS(NOW()) - TO_DAYS(%s) = 1", field))
		case "lately7":
			return db.Where(fmt.Sprintf("DATE_SUB(CURDATE(),INTERVAL 7 DAY) <= DATE(%s)", field))
		case "lately30":
			return db.Where(fmt.Sprintf("DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(%s)", field))
		case "month":
			return db.Where(fmt.Sprintf("DATE_FORMAT( %s, '%%Y%%m' ) = DATE_FORMAT( CURDATE() , '%%Y%%m' )", field))
		case "year":
			return db.Where(fmt.Sprintf("YEAR(%s)=YEAR(NOW())", field))
		}
	}
	return db
}

// 是否是 c用户
func IsCuser(ctx *gin.Context) bool {
	if multi.IsAdmin(ctx) || multi.IsTenancy(ctx) {
		return false
	}
	return true
}

func CheckTenancyId(db *gorm.DB, tenancyId uint, perfix string) *gorm.DB {
	if tenancyId == 0 {
		return db
	}
	return db.Where(perfix+"sys_tenancy_id", tenancyId)
}

func CheckUserId(db *gorm.DB, userId uint, perfix string) *gorm.DB {
	if userId == 0 {
		return db
	}
	return db.Where(perfix+"sys_user_id", userId)
}

// CheckTenancyIdAndUserId
func CheckTenancyIdAndUserId(db *gorm.DB, req request.GetById, perfix string) *gorm.DB {
	if req.TenancyId > 0 {
		db = db.Where(perfix+"sys_tenancy_id", req.TenancyId)
	}
	if req.UserId > 0 {
		db = db.Where(perfix+"sys_user_id", req.UserId)
	}
	if req.PatientId > 0 {
		db = db.Where(perfix+"patient_id", req.PatientId)
	}
	return db
}

//OrderBy 排序
func OrderBy(db *gorm.DB, orderBy, sortBy string, perfixs ...string) *gorm.DB {
	if orderBy == "" {
		orderBy = "created_at"
	}
	if sortBy == "" {
		sortBy = "desc"
	}
	if len(perfixs) == 1 {
		orderBy = perfixs[0] + orderBy
	}
	return db.Order(fmt.Sprintf("%s %s", orderBy, sortBy))
}
