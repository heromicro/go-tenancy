package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils/param"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// GetUploadConfigMap
func GetUploadConfigMap(tenancyId uint) (Form, error) {
	form := Form{
		Method: "POST",
		Action: "/sys/admin/configValue/saveConfigValue/upload",
	}

	uploadType, err := param.GetConfigValueByKey("upload_type")
	if err != nil {
		return form, err
	}

	rule := NewRadio("上传类型", "upload_type", "文件上传的类型", uploadType)
	cateKeys := []string{"aliyun_oss", "qiniuyun", "tengxun"}
	configs, err := param.GetConfigByCateKeys(cateKeys, tenancyId)
	if err != nil {
		return form, err
	}
	rule.AddControl(GetControl(configs, "qiniuyun", "2"))
	rule.AddControl(GetControl(configs, "aliyun_oss", "3"))
	rule.AddControl(GetControl(configs, "tengxun", "4"))
	rule.AddOption(Option{Value: "1", Label: "本地存储"})
	rule.AddOption(Option{Value: "2", Label: "七牛云存储"})
	rule.AddOption(Option{Value: "3", Label: "阿里云OSS"})
	rule.AddOption(Option{Value: "4", Label: "腾讯COS"})
	rule.Props = map[string]interface{}{}
	form.AddRule(*rule)
	return form, nil
}

func GetControl(configs []response.SysConfig, cateKey string, value interface{}) Control {
	control := Control{Value: value}
	for i := 0; i < len(configs); i++ {
		if configs[i].CateKey != cateKey {
			continue
		}
		controlRule := Rule{
			Title: configs[i].ConfigName,
			Type:  configs[i].ConfigType,
			Field: configs[i].ConfigKey,
			Info:  configs[i].Info,
			Value: configs[i].Value,
		}
		controlRule.TransData(configs[i].ConfigRule, nil)
		control.Rule = append(control.Rule, controlRule)
	}
	return control
}

// GetConfigMapByCate
func GetConfigMapByCate(cate string, ctx *gin.Context) (Form, error) {
	form := Form{
		Method: "POST",
	}

	//TODO: 添加 /sys/ 前缀，兼容前端 form-create/element-ui 组件
	if multi.IsTenancy(ctx) {
		form.Action = "/sys/merchant/configValue/saveConfigValue/" + cate
	} else if multi.IsAdmin(ctx) {
		form.Action = "/sys/admin/configValue/saveConfigValue/" + cate
	}

	configs, err := param.GetConfigByCateKey(cate, multi.GetTenancyId(ctx))
	if err != nil {
		return form, err
	}
	for i := 0; i < len(configs); i++ {
		rule := Rule{
			Title: configs[i].ConfigName,
			Type:  configs[i].ConfigType,
			Field: configs[i].ConfigKey,
			Info:  configs[i].Info,
			Value: configs[i].Value,
		}
		rule.TransData(configs[i].ConfigRule, multi.GetVerifiedToken(ctx))
		form.Rule = append(form.Rule, rule)
		if i == 0 {
			form.Title = GetConfigTypeName(configs[i].ConfigType)
		}
	}

	return form, nil
}

// GetConfigMap
func GetConfigMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		config, err := GetConfigByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"sysConfigCategoryId","value":%d,"title":"配置分类","props":{"type":"other","options":[],"placeholder":"请选择分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"select","field":"userType","value":%d,"title":"后台类型","props":{"multiple":false,"placeholder":"请选择后台类型"},"validate":[{"message":"请选择后台类型","required":true,"type":"number","trigger":"change"}],"options":[{"label":"总后台配置","value":2},{"label":"商户后台配置","value":1}]},{"type":"input","field":"configName","value":"%s","title":"配置名称","props":{"type":"text","placeholder":"请输入配置名称"},"validate":[{"message":"请输入配置名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"configKey","value":"%s","title":"配置key","props":{"type":"text","placeholder":"请输入配置key"},"validate":[{"message":"请输入配置key","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"info","value":"%s","title":"说明","props":{"type":"text","placeholder":"请输入说明"}},{"type":"select","field":"configType","value":"%s","title":"配置类型","props":{"multiple":false,"placeholder":"请选择配置类型"},"validate":[{"message":"请选择配置类型","required":true,"type":"string","trigger":"change"}],"options":[]},{"type":"input","field":"configRule","value":"","title":"规则","props":{"type":"textarea","placeholder":"请输入规则"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}},{"type":"switch","field":"required","value":%d,"title":"必填","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"","method":"PUT","title":"添加配置","config":{}}`, config.SysConfigCategoryID, config.UserType, config.ConfigKey, config.ConfigName, config.Info, config.ConfigType, config.Sort, config.Required, config.Status)
	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"sysConfigCategoryId","value":%d,"title":"配置分类","props":{"type":"other","options":[],"placeholder":"请选择分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"select","field":"userType","value":%d,"title":"后台类型","props":{"multiple":false,"placeholder":"请选择后台类型"},"validate":[{"message":"请选择后台类型","required":true,"type":"number","trigger":"change"}],"options":[{"label":"总后台配置","value":2},{"label":"商户后台配置","value":1}]},{"type":"input","field":"configName","value":"%s","title":"配置名称","props":{"type":"text","placeholder":"请输入配置名称"},"validate":[{"message":"请输入配置名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"configKey","value":"%s","title":"配置key","props":{"type":"text","placeholder":"请输入配置key"},"validate":[{"message":"请输入配置key","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"info","value":"%s","title":"说明","props":{"type":"text","placeholder":"请输入说明"}},{"type":"select","field":"configType","value":"%s","title":"配置类型","props":{"multiple":false,"placeholder":"请选择配置类型"},"validate":[{"message":"请选择配置类型","required":true,"type":"string","trigger":"change"}],"options":[]},{"type":"input","field":"configRule","value":"","title":"规则","props":{"type":"textarea","placeholder":"请输入规则"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}},{"type":"switch","field":"required","value":%d,"title":"必填","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"","method":"POST","title":"添加配置","config":{}}`, 0, 2, "", "", "", "", 0, 2, 1)
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	opts, err := GetConfigCategoriesOptions()
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/config/updateConfig/%d", id), ctx)
	} else {
		form.SetAction("/config/createConfig", ctx)
	}
	form.Rule[0].Props["options"] = opts
	form.Rule[5].Options = ConfigTypes
	return form, err
}

// CreateConfig
func CreateConfig(m model.SysConfig) (model.SysConfig, error) {
	err := g.TENANCY_DB.Where("config_key = ?", m.ConfigKey).First(&model.SysConfig{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return m, fmt.Errorf("设置key:%s已经使用", m.ConfigKey)
	}
	err = g.TENANCY_DB.Create(&m).Error
	return m, err
}

// GetTenancyConfigValue
func GetTenancyConfigValue(config_key string, sys_tenancy_id uint) (response.SysConfig, error) {
	var config response.SysConfig
	err := g.TENANCY_DB.
		Select("sys_config_values.value").
		Joins("left join sys_config_categories on sys_configs.sys_config_category_id = sys_config_categories.id").
		Joins("left join sys_config_values on sys_configs.config_key = sys_config_values.config_key").
		Where("sys_config_categories.key = ?", config_key).
		Where("sys_config_values.sys_tenancy_id = ?", sys_tenancy_id).
		First(&config).Error

	return config, err
}

// GetConfigByKey
func GetConfigByKey(config_key string) (model.SysConfig, error) {
	var config model.SysConfig
	err := g.TENANCY_DB.Where("config_key = ?", config_key).First(&config).Error
	return config, err
}

// GetConfigByID
func GetConfigByID(id uint) (model.SysConfig, error) {
	var config model.SysConfig
	err := g.TENANCY_DB.Where("id = ?", id).First(&config).Error
	return config, err
}

// GetConfigInfoList
func GetConfigInfoList(info request.PageInfo) ([]model.SysConfig, int64, error) {
	configList := []model.SysConfig{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysConfig{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return configList, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&configList).Error
	if err != nil {
		return configList, total, err
	}
	// 获取类型名称
	for i := 0; i < len(configList); i++ {
		configList[i].TypeName = GetConfigTypeName(configList[i].ConfigType)
	}
	return configList, total, err
}

// ChangeConfigStatus
func ChangeConfigStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.SysConfig{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateConfig
func UpdateConfig(m model.SysConfig, id uint) (model.SysConfig, error) {
	err := g.TENANCY_DB.Where("config_key = ?", m.ConfigKey).Not("id = ?", id).First(&model.SysConfig{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return m, fmt.Errorf("设置key:%s已经使用", m.ConfigKey)
	}
	err = g.TENANCY_DB.Where("id= ?", id).Updates(&m).Error
	return m, err
}

// DeleteConfig
func DeleteConfig(id uint) error {
	var config model.SysConfig
	return g.TENANCY_DB.Unscoped().Where("id = ?", id).Delete(&config).Error
}
