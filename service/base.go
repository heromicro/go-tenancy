package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
