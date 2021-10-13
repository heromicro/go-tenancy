package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/multi"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

func BaseApisLen() int {
	return len(baseApis)
}

var baseApis = []model.SysApi{

	// 授权
	{Path: "/v1/auth/logout", Description: "退出", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/auth/clean", Description: "清空 token", ApiGroup: "user", Method: "GET"},

	// 管理员管理
	{Path: "/v1/admin/user/changePasswordMap/:id", AuthorityType: multi.AdminAuthority, Description: "管理员修改密码表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/admin/user/updateAdminMap/:id", AuthorityType: multi.AdminAuthority, Description: "管理员编辑表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/admin/user/registerAdminMap", AuthorityType: multi.AdminAuthority, Description: "管理员注册表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/admin/user/registerAdmin", AuthorityType: multi.AdminAuthority, Description: "管理员注册", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/changeUserStatus", AuthorityType: multi.AdminAuthority, Description: "修改状态", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/changePassword", AuthorityType: multi.AdminAuthority, Description: "修改密码", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/changeProfile", AuthorityType: multi.AdminAuthority, Description: "修改个人信息", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/getAdminList", AuthorityType: multi.AdminAuthority, Description: "管理员列表", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/setUserAuthority", AuthorityType: multi.AdminAuthority, Description: "修改用户角色", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/admin/user/setUserInfo/:user_id", AuthorityType: multi.AdminAuthority, Description: "设置用户信息", ApiGroup: "user", Method: "PUT"},
	{Path: "/v1/admin/user/deleteUser", AuthorityType: multi.AdminAuthority, Description: "删除用户", ApiGroup: "user", Method: "DELETE"},

	// 用户管理
	{Path: "/v1/admin/cuser/getGeneralSelect/:tenancy_id", AuthorityType: multi.AdminAuthority, Description: "c用户下拉选项", ApiGroup: "cuser", Method: "GET"},
	{Path: "/v1/admin/cuser/getGeneralList", AuthorityType: multi.AdminAuthority, Description: "c用户列表", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/getGeneralDetail/:id", AuthorityType: multi.AdminAuthority, Description: "c用户列表", ApiGroup: "cuser", Method: "GET"},
	{Path: "/v1/admin/cuser/getOrderList/:id", AuthorityType: multi.AdminAuthority, Description: "消费列表", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/getBillList/:id", AuthorityType: multi.AdminAuthority, Description: "余额变动", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/setNowMoney/:id", AuthorityType: multi.AdminAuthority, Description: "设置余额", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/setNowMoneyMap/:id", AuthorityType: multi.AdminAuthority, Description: "设置余额表单", ApiGroup: "cuser", Method: "GET"},
	{Path: "/v1/admin/cuser/setUserGroup/:id", AuthorityType: multi.AdminAuthority, Description: "设置用户分组", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/setUserGroupMap/:id", AuthorityType: multi.AdminAuthority, Description: "设置用户分组表单", ApiGroup: "cuser", Method: "GET"},
	{Path: "/v1/admin/cuser/setUserLabel/:id", AuthorityType: multi.AdminAuthority, Description: "设置用户标签", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/setUserLabelMap/:id", AuthorityType: multi.AdminAuthority, Description: "设置用户标签表单", ApiGroup: "cuser", Method: "GET"},
	{Path: "/v1/admin/cuser/editUser/:id", AuthorityType: multi.AdminAuthority, Description: "编辑用户", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/editUserMap/:id", AuthorityType: multi.AdminAuthority, Description: "编辑用户表单", ApiGroup: "cuser", Method: "GET"},
	{Path: "/v1/admin/cuser/batchSetUserGroupMap", AuthorityType: multi.AdminAuthority, Description: "批量设置用户分组", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/batchSetUserGroup", AuthorityType: multi.AdminAuthority, Description: "批量设置用户分组表单", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/batchSetUserLabelMap", AuthorityType: multi.AdminAuthority, Description: "批量设置用户标签", ApiGroup: "cuser", Method: "POST"},
	{Path: "/v1/admin/cuser/batchSetUserLabel", AuthorityType: multi.AdminAuthority, Description: "批量设置用户标签表单", ApiGroup: "cuser", Method: "POST"},

	// api 管理
	{Path: "/v1/admin/api/createApi", AuthorityType: multi.AdminAuthority, Description: "创建api", ApiGroup: "api", Method: "POST"},
	{Path: "/v1/admin/api/getApiList", AuthorityType: multi.AdminAuthority, Description: "api列表", ApiGroup: "api", Method: "POST"},
	{Path: "/v1/admin/api/getApiById/:id", AuthorityType: multi.AdminAuthority, Description: "api详细信息", ApiGroup: "api", Method: "GET"},
	{Path: "/v1/admin/api/deleteApi", AuthorityType: multi.AdminAuthority, Description: "删除Api", ApiGroup: "api", Method: "DELETE"},
	{Path: "/v1/admin/api/updateApi/:id", AuthorityType: multi.AdminAuthority, Description: "更新Api", ApiGroup: "api", Method: "PUT"},
	{Path: "/v1/admin/api/getAllApis", AuthorityType: multi.AdminAuthority, Description: "所有api", ApiGroup: "api", Method: "GET"},
	{Path: "/v1/admin/api/deleteApisByIds", AuthorityType: multi.AdminAuthority, Description: "批量删除api", ApiGroup: "api", Method: "DELETE"},

	// 角色管理
	{Path: "/v1/admin/authority/createAuthority", AuthorityType: multi.AdminAuthority, Description: "创建角色", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/deleteAuthority", AuthorityType: multi.AdminAuthority, Description: "删除角色", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/getAuthorityList", AuthorityType: multi.AdminAuthority, Description: "角色列表", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/setDataAuthority", AuthorityType: multi.AdminAuthority, Description: "设置角色资源权限", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/updateAuthority", AuthorityType: multi.AdminAuthority, Description: "更新角色信息", ApiGroup: "authority", Method: "PUT"},
	{Path: "/v1/admin/authority/copyAuthority", AuthorityType: multi.AdminAuthority, Description: "拷贝角色", ApiGroup: "authority", Method: "POST"},

	// 菜单管理
	{Path: "/v1/admin/menu/getMenu", AuthorityType: multi.AdminAuthority, Description: "菜单树", ApiGroup: "menu", Method: "GET"},
	{Path: "/v1/admin/menu/getAddMenuMap", AuthorityType: multi.AdminAuthority, Description: "添加表单", ApiGroup: "menu", Method: "GET"},
	{Path: "/v1/admin/menu/getAddTenancyMenuMap", AuthorityType: multi.AdminAuthority, Description: "添加是商户菜单表单", ApiGroup: "menu", Method: "GET"},
	{Path: "/v1/admin/menu/getEditMenuMap/:id", AuthorityType: multi.AdminAuthority, Description: "编辑表单", ApiGroup: "menu", Method: "GET"},
	{Path: "/v1/admin/menu/getMenuList", AuthorityType: multi.AdminAuthority, Description: "分页基础menu列表", ApiGroup: "menu", Method: "GET"},
	{Path: "/v1/admin/menu/addBaseMenu", AuthorityType: multi.AdminAuthority, Description: "新增菜单", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/addTenancyBaseMenu", AuthorityType: multi.AdminAuthority, Description: "新增商户菜单", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/getBaseMenuTree", AuthorityType: multi.AdminAuthority, Description: "用户动态路由", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/addMenuAuthority", AuthorityType: multi.AdminAuthority, Description: "增加menu和角色关联关系", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/getMenuAuthority", AuthorityType: multi.AdminAuthority, Description: "指定角色menu", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/deleteBaseMenu/:id", AuthorityType: multi.AdminAuthority, Description: "删除菜单", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/updateBaseMenu/:id", AuthorityType: multi.AdminAuthority, Description: "更新菜单", ApiGroup: "menu", Method: "POST"},
	{Path: "/v1/admin/menu/getBaseMenuById/:id", AuthorityType: multi.AdminAuthority, Description: "根据id菜单", ApiGroup: "menu", Method: "POST"},
	// 商户菜单
	{Path: "/v1/admin/menu/merchant/getClientMenuList", AuthorityType: multi.AdminAuthority, Description: "商户菜单", ApiGroup: "menu", Method: "GET"},

	// 多媒体文件管理
	{Path: "/v1/admin/media/getUpdateMediaMap/:id", AuthorityType: multi.AdminAuthority, Description: "媒体文件表单", ApiGroup: "media", Method: "GET"},
	{Path: "/v1/admin/media/upload", AuthorityType: multi.AdminAuthority, Description: "上传媒体文件", ApiGroup: "media", Method: "POST"},
	{Path: "/v1/admin/media/getFileList", AuthorityType: multi.AdminAuthority, Description: "媒体文件列表", ApiGroup: "media", Method: "POST"},
	{Path: "/v1/admin/media/updateMediaName/:id", AuthorityType: multi.AdminAuthority, Description: "修改媒体文件名称", ApiGroup: "media", Method: "POST"},
	{Path: "/v1/admin/media/deleteFile", AuthorityType: multi.AdminAuthority, Description: "删除媒体文件", ApiGroup: "media", Method: "DELETE"},

	// casbin 管理
	{Path: "/v1/admin/casbin/updateCasbin", AuthorityType: multi.AdminAuthority, Description: "更改角色api权限", ApiGroup: "casbin", Method: "POST"},
	{Path: "/v1/admin/casbin/getPolicyPathByAuthorityId", AuthorityType: multi.AdminAuthority, Description: "权限列表", ApiGroup: "casbin", Method: "POST"},

	// 系统配置管理
	{Path: "/v1/admin/system/getSystemConfig", AuthorityType: multi.AdminAuthority, Description: "配置文件内容", ApiGroup: "system", Method: "POST"},
	{Path: "/v1/admin/system/setSystemConfig", AuthorityType: multi.AdminAuthority, Description: "设置配置文件内容", ApiGroup: "system", Method: "POST"},
	{Path: "/v1/admin/system/getServerInfo", AuthorityType: multi.AdminAuthority, Description: "服务器信息", ApiGroup: "system", Method: "POST"},
	// 系统配置值管理
	{Path: "/v1/admin/configValue/saveConfigValue/:category", AuthorityType: multi.AdminAuthority, Description: "保持配置表单", ApiGroup: "configValue", Method: "POST"},

	// 配置
	{Path: "/v1/admin/config/getUploadConfigMap", AuthorityType: multi.AdminAuthority, Description: "配置表单", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getConfigMap/:category", AuthorityType: multi.AdminAuthority, Description: "配置表单", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getCreateConfigMap", AuthorityType: multi.AdminAuthority, Description: "配置创建表单", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getUpdateConfigMap/:id", AuthorityType: multi.AdminAuthority, Description: "配置编辑表单", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getConfigList", AuthorityType: multi.AdminAuthority, Description: "配置项列表", ApiGroup: "config", Method: "POST"},
	{Path: "/v1/admin/config/createConfig", AuthorityType: multi.AdminAuthority, Description: "添加配置项", ApiGroup: "config", Method: "POST"},
	{Path: "/v1/admin/config/getConfigByKey/:key", AuthorityType: multi.AdminAuthority, Description: "根据key配置项", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/getConfigByID/:id", AuthorityType: multi.AdminAuthority, Description: "根据id配置项", ApiGroup: "config", Method: "GET"},
	{Path: "/v1/admin/config/changeConfigStatus", AuthorityType: multi.AdminAuthority, Description: "修改配置状态", ApiGroup: "config", Method: "POST"},
	{Path: "/v1/admin/config/updateConfig/:id", AuthorityType: multi.AdminAuthority, Description: "更新配置项", ApiGroup: "config", Method: "PUT"},
	{Path: "/v1/admin/config/deleteConfig/:id", AuthorityType: multi.AdminAuthority, Description: "删除配置项", ApiGroup: "config", Method: "DELETE"},

	// 配置分类
	{Path: "/v1/admin/configCategory/getCreateConfigCategoryMap", AuthorityType: multi.AdminAuthority, Description: "配置分类创建表单", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/getUpdateConfigCategoryMap/:id", AuthorityType: multi.AdminAuthority, Description: "配置分类编辑表单", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/getConfigCategoryList", AuthorityType: multi.AdminAuthority, Description: "配置分类列表", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/createConfigCategory", AuthorityType: multi.AdminAuthority, Description: "添加配置分类", ApiGroup: "configCategory", Method: "POST"},
	{Path: "/v1/admin/configCategory/changeConfigCategoryStatus", AuthorityType: multi.AdminAuthority, Description: "修改配置分类状态", ApiGroup: "configCategory", Method: "POST"},
	{Path: "/v1/admin/configCategory/getConfigCategoryById/:id", AuthorityType: multi.AdminAuthority, Description: "配置分类", ApiGroup: "configCategory", Method: "GET"},
	{Path: "/v1/admin/configCategory/updateConfigCategory/:id", AuthorityType: multi.AdminAuthority, Description: "更新配置分类", ApiGroup: "configCategory", Method: "PUT"},
	{Path: "/v1/admin/configCategory/deleteConfigCategory/:id", AuthorityType: multi.AdminAuthority, Description: "删除配置分类", ApiGroup: "configCategory", Method: "DELETE"},

	// 商户
	{Path: "/v1/admin/tenancy/getTenancySelect", AuthorityType: multi.AdminAuthority, Description: "下拉列表", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/changePasswordMap/:id", AuthorityType: multi.AdminAuthority, Description: "商户管理员编辑表单", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/changeCopyMap/:id", AuthorityType: multi.AdminAuthority, Description: "修改商品复制次数map", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/getTenancies/:code", AuthorityType: multi.AdminAuthority, Description: "根据地区商户", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/getTenancyCount", AuthorityType: multi.AdminAuthority, Description: "Tenancy对应状态数量", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/getTenancyList", AuthorityType: multi.AdminAuthority, Description: "商户列表", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/loginTenancy/:id", AuthorityType: multi.AdminAuthority, Description: "登录商户", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/createTenancy", AuthorityType: multi.AdminAuthority, Description: "添加商户", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/setTenancyRegion", AuthorityType: multi.AdminAuthority, Description: "设置商户地区", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/setCopyProductNum/:id", AuthorityType: multi.AdminAuthority, Description: "设置商品复制次数", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/changeTenancyStatus", AuthorityType: multi.AdminAuthority, Description: "启用/禁用商户", ApiGroup: "tenancy", Method: "POST"},
	{Path: "/v1/admin/tenancy/getTenancyById/:id", AuthorityType: multi.AdminAuthority, Description: "商户详细信息", ApiGroup: "tenancy", Method: "GET"},
	{Path: "/v1/admin/tenancy/updateTenancy/:id", AuthorityType: multi.AdminAuthority, Description: "更新商户", ApiGroup: "tenancy", Method: "PUT"},
	{Path: "/v1/admin/tenancy/deleteTenancy/:id", AuthorityType: multi.AdminAuthority, Description: "删除商户", ApiGroup: "tenancy", Method: "DELETE"},
	// MQTT
	{Path: "/v1/admin/mqtt/getCreateMqttMap", AuthorityType: multi.AdminAuthority, Description: "添加mqtt客户端表单", ApiGroup: "mqtt", Method: "GET"},
	{Path: "/v1/admin/mqtt/getUpdateMqttMap/:id", AuthorityType: multi.AdminAuthority, Description: "编辑mqtt客户端表单", ApiGroup: "mqtt", Method: "GET"},
	{Path: "/v1/admin/mqtt/getMqttList", AuthorityType: multi.AdminAuthority, Description: "mqtt客户端列表", ApiGroup: "mqtt", Method: "POST"},
	{Path: "/v1/admin/mqtt/getMqttRecordList", AuthorityType: multi.AdminAuthority, Description: "mqtt消息记录", ApiGroup: "mqtt", Method: "POST"},
	{Path: "/v1/admin/mqtt/createMqtt", AuthorityType: multi.AdminAuthority, Description: "添加mqtt客户端", ApiGroup: "mqtt", Method: "POST"},
	{Path: "/v1/admin/mqtt/getMqttById/:id", AuthorityType: multi.AdminAuthority, Description: "mqtt客户端详情", ApiGroup: "mqtt", Method: "GET"},
	{Path: "/v1/admin/mqtt/changeMqttStatus", AuthorityType: multi.AdminAuthority, Description: "修改mqtt客户端状态", ApiGroup: "mqtt", Method: "POST"},
	{Path: "/v1/admin/mqtt/updateMqtt/:id", AuthorityType: multi.AdminAuthority, Description: "更新mqtt客户端", ApiGroup: "mqtt", Method: "PUT"},
	{Path: "/v1/admin/mqtt/deleteMqtt/:id", AuthorityType: multi.AdminAuthority, Description: "删除mqtt客户端", ApiGroup: "mqtt", Method: "DELETE"},
	// 定时任务
	{Path: "/v1/admin/job/getJobList", AuthorityType: multi.AdminAuthority, Description: "定时任务列表", ApiGroup: "job", Method: "GET"},
	{Path: "/v1/admin/job/startJob/:name", AuthorityType: multi.AdminAuthority, Description: "启动定时任务", ApiGroup: "job", Method: "GET"},
	{Path: "/v1/admin/job/stopJob/:name", AuthorityType: multi.AdminAuthority, Description: "停止定时任务", ApiGroup: "job", Method: "GET"},

	//商品分类
	{Path: "/v1/admin/productCategory/getCreateProductCategoryMap", AuthorityType: multi.AdminAuthority, Description: "商品分类添加表单", ApiGroup: "productCategory", Method: "GET"},
	{Path: "/v1/admin/productCategory/getUpdateProductCategoryMap/:id", AuthorityType: multi.AdminAuthority, Description: "商品分类编辑表单", ApiGroup: "productCategory", Method: "GET"},
	{Path: "/v1/admin/productCategory/getProductCategorySelect", AuthorityType: multi.AdminAuthority, Description: "商品分类选项", ApiGroup: "productCategory", Method: "GET"},
	{Path: "/v1/admin/productCategory/getProductCategoryList", AuthorityType: multi.AdminAuthority, Description: "商品分类列表", ApiGroup: "productCategory", Method: "GET"},
	{Path: "/v1/admin/productCategory/createProductCategory", AuthorityType: multi.AdminAuthority, Description: "添加商品分类", ApiGroup: "productCategory", Method: "POST"},
	{Path: "/v1/admin/productCategory/getProductCategoryById/:id", AuthorityType: multi.AdminAuthority, Description: "根据id商品分类", ApiGroup: "productCategory", Method: "GET"},
	{Path: "/v1/admin/productCategory/changeProductCategoryStatus", AuthorityType: multi.AdminAuthority, Description: "修改商品分类状态", ApiGroup: "productCategory", Method: "POST"},
	{Path: "/v1/admin/productCategory/updateProductCategory/:id", AuthorityType: multi.AdminAuthority, Description: "更新商品分类", ApiGroup: "productCategory", Method: "PUT"},
	{Path: "/v1/admin/productCategory/deleteProductCategory/:id", AuthorityType: multi.AdminAuthority, Description: "删除商品分类", ApiGroup: "productCategory", Method: "DELETE"},

	//商品
	{Path: "/v1/admin/product/getProductSelect/:tenancy_id", AuthorityType: multi.AdminAuthority, Description: "商品下拉选项", ApiGroup: "product", Method: "GET"},
	{Path: "/v1/admin/product/getEditProductFictiMap/:id", AuthorityType: multi.AdminAuthority, Description: "设置虚拟销量表单", ApiGroup: "product", Method: "GET"},
	{Path: "/v1/admin/product/setProductFicti/:id", AuthorityType: multi.AdminAuthority, Description: "设置虚拟销量", ApiGroup: "product", Method: "PUT"},
	{Path: "/v1/admin/product/getProductFilter", AuthorityType: multi.AdminAuthority, Description: "商品过滤参数", ApiGroup: "product", Method: "POST"},
	{Path: "/v1/admin/product/changeProductStatus", AuthorityType: multi.AdminAuthority, Description: "强制下架，重新审核", ApiGroup: "product", Method: "POST"},
	{Path: "/v1/admin/product/changeMutilProductStatus", AuthorityType: multi.AdminAuthority, Description: "批量强制下架，重新审核", ApiGroup: "product", Method: "POST"},
	{Path: "/v1/admin/product/getProductList", AuthorityType: multi.AdminAuthority, Description: "商品列表", ApiGroup: "product", Method: "POST"},
	{Path: "/v1/admin/product/getProductById/:id", AuthorityType: multi.AdminAuthority, Description: "商品详情", ApiGroup: "product", Method: "GET"},
	{Path: "/v1/admin/product/updateProduct/:id", AuthorityType: multi.AdminAuthority, Description: "编辑商品", ApiGroup: "product", Method: "PUT"},
	// 商品评论
	{Path: "/v1/admin/productReply/replyMap", AuthorityType: multi.AdminAuthority, Description: "虚拟评论表单", ApiGroup: "productReply", Method: "GET"},
	{Path: "/v1/admin/productReply/reply", AuthorityType: multi.AdminAuthority, Description: "添加虚拟评论", ApiGroup: "productReply", Method: "POST"},
	{Path: "/v1/admin/productReply/deleteProductReply/:id", AuthorityType: multi.AdminAuthority, Description: "删除评论", ApiGroup: "productReply", Method: "DELETE"},
	{Path: "/v1/admin/productReply/getProductReplyList", AuthorityType: multi.AdminAuthority, Description: "商品评论列表", ApiGroup: "productReply", Method: "POST"},

	//品牌分类
	{Path: "/v1/admin/category/getCreateBrandCategoryMap", AuthorityType: multi.AdminAuthority, Description: "品牌分类添加表单", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/category/getUpdateBrandCategoryMap/:id", AuthorityType: multi.AdminAuthority, Description: "品牌分类编辑表单", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/brandCategory/getBrandCategoryList", AuthorityType: multi.AdminAuthority, Description: "品牌分类列表", ApiGroup: "brandCategory", Method: "GET"},
	{Path: "/v1/admin/brandCategory/createBrandCategory", AuthorityType: multi.AdminAuthority, Description: "添加品牌分类", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/brandCategory/getBrandCategoryById/:id", AuthorityType: multi.AdminAuthority, Description: "根据id品牌分类", ApiGroup: "brandCategory", Method: "GET"},
	{Path: "/v1/admin/brandCategory/changeBrandCategoryStatus", AuthorityType: multi.AdminAuthority, Description: "修改品牌分类状态", ApiGroup: "brandCategory", Method: "POST"},
	{Path: "/v1/admin/brandCategory/updateBrandCategory/:id", AuthorityType: multi.AdminAuthority, Description: "更新品牌分类", ApiGroup: "brandCategory", Method: "PUT"},
	{Path: "/v1/admin/brandCategory/deleteBrandCategory/:id", AuthorityType: multi.AdminAuthority, Description: "删除品牌分类", ApiGroup: "brandCategory", Method: "DELETE"},

	//品牌
	{Path: "/v1/admin/brand/getBrandList", AuthorityType: multi.AdminAuthority, Description: "品牌列表", ApiGroup: "brand", Method: "POST"},
	{Path: "/v1/admin/brand/getCreateBrandMap", AuthorityType: multi.AdminAuthority, Description: "品牌添加表单", ApiGroup: "brand", Method: "GET"},
	{Path: "/v1/admin/brand/getUpdateBrandMap/:id", AuthorityType: multi.AdminAuthority, Description: "品牌编辑表单", ApiGroup: "brand", Method: "GET"},
	{Path: "/v1/admin/brand/createBrand", AuthorityType: multi.AdminAuthority, Description: "添加品牌", ApiGroup: "brand", Method: "POST"},
	{Path: "/v1/admin/brand/getBrandById/:id", AuthorityType: multi.AdminAuthority, Description: "根据id品牌", ApiGroup: "brand", Method: "GET"},
	{Path: "/v1/admin/brand/changeBrandStatus", AuthorityType: multi.AdminAuthority, Description: "修改品牌分类状态", ApiGroup: "brand", Method: "POST"},
	{Path: "/v1/admin/brand/updateBrand/:id", AuthorityType: multi.AdminAuthority, Description: "更新品牌", ApiGroup: "brand", Method: "PUT"},
	{Path: "/v1/admin/brand/deleteBrand/:id", AuthorityType: multi.AdminAuthority, Description: "删除品牌", ApiGroup: "brand", Method: "DELETE"},

	// 小程序
	{Path: "/v1/admin/mini/getMiniList", AuthorityType: multi.AdminAuthority, Description: "小程序列表", ApiGroup: "mini", Method: "POST"},
	{Path: "/v1/admin/mini/createMini", AuthorityType: multi.AdminAuthority, Description: "添加小程序", ApiGroup: "mini", Method: "POST"},
	{Path: "/v1/admin/mini/getMiniById/:id", AuthorityType: multi.AdminAuthority, Description: "小程序详细信息", ApiGroup: "mini", Method: "GET"},
	{Path: "/v1/admin/mini/updateMini/:id", AuthorityType: multi.AdminAuthority, Description: "更新小程序", ApiGroup: "mini", Method: "PUT"},
	{Path: "/v1/admin/mini/deleteMini/:id", AuthorityType: multi.AdminAuthority, Description: "删除小程序", ApiGroup: "mini", Method: "DELETE"},

	// 角色管理
	{Path: "/v1/admin/authority/getAdminAuthorityList", AuthorityType: multi.AdminAuthority, Description: "员工角色列表", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/getTenancyAuthorityList", AuthorityType: multi.AdminAuthority, Description: "商户角色列表", ApiGroup: "authority", Method: "POST"},
	{Path: "/v1/admin/authority/getGeneralAuthorityList", AuthorityType: multi.AdminAuthority, Description: "普通用户角色列表", ApiGroup: "authority", Method: "POST"},

	// 操作记录
	{Path: "/v1/admin/sysOperationRecord/getSysOperationRecordList", AuthorityType: multi.AdminAuthority, Description: "操作记录列表", ApiGroup: "sysOperationRecord", Method: "POST"},

	//订单
	{Path: "/v1/admin/order/getOrderList", AuthorityType: multi.AdminAuthority, Description: "订单列表", ApiGroup: "order", Method: "POST"},
	{Path: "/v1/admin/order/getOrderChart", AuthorityType: multi.AdminAuthority, Description: "订单表头数量", ApiGroup: "order", Method: "POST"},
	{Path: "/v1/admin/order/getOrderById/:id", AuthorityType: multi.AdminAuthority, Description: "订单详情", ApiGroup: "order", Method: "GET"},
	//退款订单
	{Path: "/v1/admin/refundOrder/getRefundOrderList", AuthorityType: multi.AdminAuthority, Description: "退款订单列表", ApiGroup: "refundOrder", Method: "POST"},

	//控制台
	{Path: "/v1/admin/statistics/main", AuthorityType: multi.AdminAuthority, Description: "运营数据", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/order", AuthorityType: multi.AdminAuthority, Description: "当日订单金额", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/orderNum", AuthorityType: multi.AdminAuthority, Description: "订单数量统计", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/orderUser", AuthorityType: multi.AdminAuthority, Description: "支付人数统计", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/merchantStock", AuthorityType: multi.AdminAuthority, Description: "商品销量排行", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/merchantVisit", AuthorityType: multi.AdminAuthority, Description: "商户访客量排行", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/merchantRate", AuthorityType: multi.AdminAuthority, Description: "商户销售额占比", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/userData", AuthorityType: multi.AdminAuthority, Description: "用户数据", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/user", AuthorityType: multi.AdminAuthority, Description: "成交用户统计", ApiGroup: "statistics", Method: "GET"},
	{Path: "/v1/admin/statistics/userRate", AuthorityType: multi.AdminAuthority, Description: "成交用户占比", ApiGroup: "statistics", Method: "GET"},

	//财务
	{Path: "/v1/admin/financialRecord/getFinancialRecordList", AuthorityType: multi.AdminAuthority, Description: "财务-资金流水", ApiGroup: "financialRecord", Method: "GET"},

	// other
	{Path: "/v1/admin/test/emailTest", AuthorityType: multi.AdminAuthority, Description: "发送测试邮件", ApiGroup: "email", Method: "POST"},
	{Path: "/v1/admin/test/pay", AuthorityType: multi.AdminAuthority, Description: "支付测试", ApiGroup: "email", Method: "POST"},
	{Path: "/v1/admin/test/mqtt", AuthorityType: multi.AdminAuthority, Description: "mqtt测试", ApiGroup: "email", Method: "GET"},

	// TODO:商户用户权限
	{Path: "/v1/merchant/config/getConfigMap/:category", AuthorityType: multi.TenancyAuthority, Description: "配置表单", ApiGroup: "configClient", Method: "GET"},
	// 配置值保存
	{Path: "/v1/merchant/configValue/saveConfigValue/:category", AuthorityType: multi.TenancyAuthority, Description: "保持配置表单", ApiGroup: "configValueClient", Method: "POST"},
	//菜单
	{Path: "/v1/merchant/menu/getMenu", AuthorityType: multi.TenancyAuthority, Description: "菜单树", ApiGroup: "menuClient", Method: "GET"},
	// 商户
	{Path: "/v1/merchant/tenancy/getTenancyInfo", AuthorityType: multi.TenancyAuthority, Description: "登录商户信息", ApiGroup: "tenancyClient", Method: "GET"},
	{Path: "/v1/merchant/tenancy/getUpdateTenancyMap", AuthorityType: multi.TenancyAuthority, Description: "登录商户信息表单", ApiGroup: "tenancyClient", Method: "GET"},
	{Path: "/v1/merchant/tenancy/getTenancyCopyCount", AuthorityType: multi.TenancyAuthority, Description: "商户商品复制次数", ApiGroup: "tenancyClient", Method: "GET"},
	{Path: "/v1/merchant/tenancy/updateTenancy/:id", AuthorityType: multi.TenancyAuthority, Description: "保存登录商户信息", ApiGroup: "tenancyClient", Method: "PUT"},
	// 媒体库
	{Path: "/v1/merchant/media/getUpdateMediaMap/:id", AuthorityType: multi.TenancyAuthority, Description: "媒体文件表单", ApiGroup: "mediaClient", Method: "GET"},
	{Path: "/v1/merchant/media/upload", AuthorityType: multi.TenancyAuthority, Description: "上传文件", ApiGroup: "mediaClient", Method: "POST"},
	{Path: "/v1/merchant/media/getFileList", AuthorityType: multi.TenancyAuthority, Description: "getFileList", ApiGroup: "mediaClient", Method: "POST"},
	{Path: "/v1/merchant/media/updateMediaName/:id", AuthorityType: multi.TenancyAuthority, Description: "修改媒体文件名称", ApiGroup: "mediaClient", Method: "POST"},
	{Path: "/v1/merchant/media/deleteFile", AuthorityType: multi.TenancyAuthority, Description: "删除媒体文件", ApiGroup: "mediaClient", Method: "DELETE"},
	//品牌
	{Path: "/v1/merchant/brand/getBrandList", AuthorityType: multi.TenancyAuthority, Description: "品牌列表", ApiGroup: "brandClient", Method: "GET"},
	//商品分类
	{Path: "/v1/merchant/productCategory/getCreateProductCategoryMap", AuthorityType: multi.TenancyAuthority, Description: "商品分类添加表单", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getUpdateProductCategoryMap/:id", AuthorityType: multi.TenancyAuthority, Description: "商品分类编辑表单", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getProductCategorySelect", AuthorityType: multi.TenancyAuthority, Description: "商品分类选项", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getAdminProductCategorySelect", AuthorityType: multi.TenancyAuthority, Description: "平台商品分类选项", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/getProductCategoryList", AuthorityType: multi.TenancyAuthority, Description: "商品分类列表", ApiGroup: "categoryClient", Method: "POST"},
	{Path: "/v1/merchant/productCategory/createProductCategory", AuthorityType: multi.TenancyAuthority, Description: "添加商品分类", ApiGroup: "categoryClient", Method: "POST"},
	{Path: "/v1/merchant/productCategory/getProductCategoryById/:id", AuthorityType: multi.TenancyAuthority, Description: "根据id商品分类", ApiGroup: "categoryClient", Method: "GET"},
	{Path: "/v1/merchant/productCategory/changeProductCategoryStatus", AuthorityType: multi.TenancyAuthority, Description: "修改商品分类状态", ApiGroup: "categoryClient", Method: "POST"},
	{Path: "/v1/merchant/productCategory/updateProductCategory/:id", AuthorityType: multi.TenancyAuthority, Description: "更新商品分类", ApiGroup: "categoryClient", Method: "PUT"},
	{Path: "/v1/merchant/productCategory/deleteProductCategory/:id", AuthorityType: multi.TenancyAuthority, Description: "删除商品分类", ApiGroup: "categoryClient", Method: "DELETE"},

	//规格参数
	{Path: "/v1/merchant/attrTemplate/getAttrTemplateList", AuthorityType: multi.TenancyAuthority, Description: "规格参数列表", ApiGroup: "attrTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/attrTemplate/createAttrTemplate", AuthorityType: multi.TenancyAuthority, Description: "添加规格参数", ApiGroup: "attrTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/attrTemplate/getAttrTemplateById/:id", AuthorityType: multi.TenancyAuthority, Description: "规格参数详情", ApiGroup: "attrTemplateClient", Method: "GET"},
	{Path: "/v1/merchant/attrTemplate/updateAttrTemplate/:id", AuthorityType: multi.TenancyAuthority, Description: "更新规格参数", ApiGroup: "attrTemplateClient", Method: "PUT"},
	{Path: "/v1/merchant/attrTemplate/deleteAttrTemplate/:id", AuthorityType: multi.TenancyAuthority, Description: "删除规格参数", ApiGroup: "attrTemplateClient", Method: "DELETE"},

	//运费模板
	{Path: "/v1/merchant/shippingTemplate/getShippingTemplateSelect", AuthorityType: multi.TenancyAuthority, Description: "运费模板下拉", ApiGroup: "shippingTemplateClient", Method: "GET"},
	{Path: "/v1/merchant/shippingTemplate/getShippingTemplateList", AuthorityType: multi.TenancyAuthority, Description: "运费模板列表", ApiGroup: "shippingTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/shippingTemplate/createShippingTemplate", AuthorityType: multi.TenancyAuthority, Description: "添加运费模板", ApiGroup: "shippingTemplateClient", Method: "POST"},
	{Path: "/v1/merchant/shippingTemplate/getShippingTemplateById/:id", AuthorityType: multi.TenancyAuthority, Description: "运费模板详情", ApiGroup: "shippingTemplateClient", Method: "GET"},
	{Path: "/v1/merchant/shippingTemplate/updateShippingTemplate/:id", AuthorityType: multi.TenancyAuthority, Description: "更新运费模板", ApiGroup: "shippingTemplateClient", Method: "PUT"},
	{Path: "/v1/merchant/shippingTemplate/deleteShippingTemplate/:id", AuthorityType: multi.TenancyAuthority, Description: "删除运费模板", ApiGroup: "attrTemplateClient", Method: "DELETE"},

	//商品
	{Path: "/v1/merchant/product/getEditProductFictiMap/:id", AuthorityType: multi.TenancyAuthority, Description: "设置虚拟销量表单", ApiGroup: "productClient", Method: "GET"},
	{Path: "/v1/merchant/product/setProductFicti/:id", AuthorityType: multi.TenancyAuthority, Description: "设置虚拟销量", ApiGroup: "productClient", Method: "PUT"},
	{Path: "/v1/merchant/product/getProductFilter", AuthorityType: multi.TenancyAuthority, Description: "商品过滤参数", ApiGroup: "productClient", Method: "POST"},
	{Path: "/v1/merchant/product/changeProductIsShow", AuthorityType: multi.TenancyAuthority, Description: "上下架商品", ApiGroup: "productClient", Method: "POST"},
	{Path: "/v1/merchant/product/getProductList", AuthorityType: multi.TenancyAuthority, Description: "商品列表", ApiGroup: "productClient", Method: "POST"},
	{Path: "/v1/merchant/product/createProduct", AuthorityType: multi.TenancyAuthority, Description: "添加商品", ApiGroup: "productClient", Method: "POST"},
	{Path: "/v1/merchant/product/getProductById/:id", AuthorityType: multi.TenancyAuthority, Description: "商品详情", ApiGroup: "productClient", Method: "GET"},
	{Path: "/v1/merchant/product/updateProduct/:id", AuthorityType: multi.TenancyAuthority, Description: "编辑商品", ApiGroup: "productClient", Method: "PUT"},
	{Path: "/v1/merchant/product/restoreProduct/:id", AuthorityType: multi.TenancyAuthority, Description: "还原商品", ApiGroup: "productClient", Method: "GET"},
	{Path: "/v1/merchant/product/deleteProduct/:id", AuthorityType: multi.TenancyAuthority, Description: "加入回收站", ApiGroup: "productClient", Method: "DELETE"},
	{Path: "/v1/merchant/product/destoryProduct/:id", AuthorityType: multi.TenancyAuthority, Description: "删除商品", ApiGroup: "productClient", Method: "DELETE"},

	// 商品评论
	{Path: "/v1/merchant/product/replyMap/:id", AuthorityType: multi.TenancyAuthority, Description: "回复评论表单", ApiGroup: "productClientReply", Method: "GET"},
	{Path: "/v1/merchant/product/reply/:id", AuthorityType: multi.TenancyAuthority, Description: "回复评论", ApiGroup: "productClientReply", Method: "POST"},
	{Path: "/v1/merchant/product/getProductReplyList", AuthorityType: multi.TenancyAuthority, Description: "商品评论列表", ApiGroup: "productClientReply", Method: "POST"},

	//订单
	{Path: "/v1/merchant/order/deliveryOrderMap/:id", AuthorityType: multi.TenancyAuthority, Description: "订单发货表单", ApiGroup: "orderClient", Method: "GET"},
	{Path: "/v1/merchant/order/getOrderRemarkMap/:id", AuthorityType: multi.TenancyAuthority, Description: "订单备注表单", ApiGroup: "orderClient", Method: "GET"},
	{Path: "/v1/merchant/order/getEditOrderMap/:id", AuthorityType: multi.TenancyAuthority, Description: "订单编辑表单", ApiGroup: "orderClient", Method: "GET"},
	{Path: "/v1/merchant/order/getOrderList", AuthorityType: multi.TenancyAuthority, Description: "订单列表", ApiGroup: "orderClient", Method: "POST"},
	{Path: "/v1/merchant/order/getOrderChart", AuthorityType: multi.TenancyAuthority, Description: "订单表头数量", ApiGroup: "orderClient", Method: "POST"},
	{Path: "/v1/merchant/order/getOrderFilter", AuthorityType: multi.TenancyAuthority, Description: "订单分类统计", ApiGroup: "orderClient", Method: "POST"},
	{Path: "/v1/merchant/order/getOrderById/:id", AuthorityType: multi.TenancyAuthority, Description: "订单详情", ApiGroup: "orderClient", Method: "GET"},
	{Path: "/v1/merchant/order/getOrderRecord/:id", AuthorityType: multi.TenancyAuthority, Description: "订单记录", ApiGroup: "orderClient", Method: "POST"},
	{Path: "/v1/merchant/order/deliveryOrder/:id", AuthorityType: multi.TenancyAuthority, Description: "订单发货", ApiGroup: "orderClient", Method: "POST"},
	{Path: "/v1/merchant/order/remarkOrder/:id", AuthorityType: multi.TenancyAuthority, Description: "订单备注", ApiGroup: "orderClient", Method: "POST"},
	{Path: "/v1/merchant/order/updateOrder/:id", AuthorityType: multi.TenancyAuthority, Description: "订单更新", ApiGroup: "orderClient", Method: "POST"},
	{Path: "/v1/merchant/order/deleteOrder/:id", AuthorityType: multi.TenancyAuthority, Description: "删除订单", ApiGroup: "orderClient", Method: "DELETE"},
	//退款订单
	{Path: "/v1/merchant/refundOrder/getRefundOrderMap/:id", AuthorityType: multi.TenancyAuthority, Description: "退款订单表单", ApiGroup: "refundOrderClient", Method: "GET"},
	{Path: "/v1/merchant/refundOrder/getRefundOrderRemarkMap/:id", AuthorityType: multi.TenancyAuthority, Description: "退款订单备注表单", ApiGroup: "refundOrderClient", Method: "GET"},
	{Path: "/v1/merchant/refundOrder/remarkRefundOrder/:id", AuthorityType: multi.TenancyAuthority, Description: "退款订单备注", ApiGroup: "refundOrderClient", Method: "POST"},
	{Path: "/v1/merchant/refundOrder/auditRefundOrder/:id", AuthorityType: multi.TenancyAuthority, Description: "退款订单审核", ApiGroup: "refundOrderClient", Method: "POST"},
	{Path: "/v1/merchant/refundOrder/getRefundOrderList", AuthorityType: multi.TenancyAuthority, Description: "退款订单列表", ApiGroup: "refundOrderClient", Method: "POST"},
	{Path: "/v1/merchant/refundOrder/getRefundOrderRecord/:id", AuthorityType: multi.TenancyAuthority, Description: "退款订单记录", ApiGroup: "refundOrderClient", Method: "POST"},
	{Path: "/v1/merchant/refundOrder/deleteRefundOrder/:id", AuthorityType: multi.TenancyAuthority, Description: "删除退款订单", ApiGroup: "refundOrderClient", Method: "DELETE"},
	// 物流公司
	{Path: "/v1/admin/express/getCreateExpressMap", AuthorityType: multi.TenancyAuthority, Description: "物流添加表单", ApiGroup: "express", Method: "GET"},
	{Path: "/v1/admin/express/getUpdateExpressMap/:id", AuthorityType: multi.TenancyAuthority, Description: "物流编辑表单", ApiGroup: "express", Method: "GET"},
	{Path: "/v1/admin/express/getExpressList", AuthorityType: multi.TenancyAuthority, Description: "物流列表", ApiGroup: "express", Method: "POST"},
	{Path: "/v1/admin/express/createExpress", AuthorityType: multi.TenancyAuthority, Description: "添加物流", ApiGroup: "express", Method: "POST"},
	{Path: "/v1/admin/express/getExpressById/:id", AuthorityType: multi.TenancyAuthority, Description: "物流详情", ApiGroup: "express", Method: "GET"},
	{Path: "/v1/admin/express/changeExpressStatus", AuthorityType: multi.TenancyAuthority, Description: "物流状态切换", ApiGroup: "express", Method: "POST"},
	{Path: "/v1/admin/express/updateExpress/:id", AuthorityType: multi.TenancyAuthority, Description: "更新物流", ApiGroup: "express", Method: "PUT"},
	{Path: "/v1/admin/express/deleteExpress/:id", AuthorityType: multi.TenancyAuthority, Description: "删除物流", ApiGroup: "express", Method: "DELETE"},
	// 用户分组
	{Path: "/v1/admin/userGroup/getCreateUserGroupMap", AuthorityType: multi.TenancyAuthority, Description: "用户分组添加表单", ApiGroup: "userGroup", Method: "GET"},
	{Path: "/v1/admin/userGroup/getUpdateUserGroupMap/:id", AuthorityType: multi.TenancyAuthority, Description: "用户分组编辑表单", ApiGroup: "userGroup", Method: "GET"},
	{Path: "/v1/admin/userGroup/getUserGroupList", AuthorityType: multi.TenancyAuthority, Description: "用户分组列表", ApiGroup: "userGroup", Method: "POST"},
	{Path: "/v1/admin/userGroup/createUserGroup", AuthorityType: multi.TenancyAuthority, Description: "添加用户分组", ApiGroup: "userGroup", Method: "POST"},
	{Path: "/v1/admin/userGroup/updateUserGroup/:id", AuthorityType: multi.TenancyAuthority, Description: "更新用户分组", ApiGroup: "userGroup", Method: "PUT"},
	{Path: "/v1/admin/userGroup/deleteUserGroup/:id", AuthorityType: multi.TenancyAuthority, Description: "删除用户分组", ApiGroup: "userGroup", Method: "DELETE"},
	// 用户标签
	{Path: "/v1/admin/userLabel/getCreateUserLabelMap", AuthorityType: multi.TenancyAuthority, Description: "用户标签添加表单", ApiGroup: "userLabel", Method: "GET"},
	{Path: "/v1/admin/userLabel/getUpdateUserLabelMap/:id", AuthorityType: multi.TenancyAuthority, Description: "用户标签编辑表单", ApiGroup: "userLabel", Method: "GET"},
	{Path: "/v1/admin/userLabel/getUserLabelList", AuthorityType: multi.TenancyAuthority, Description: "用户标签列表", ApiGroup: "userLabel", Method: "POST"},
	{Path: "/v1/admin/userLabel/createUserLabel", AuthorityType: multi.TenancyAuthority, Description: "添加用户标签", ApiGroup: "userLabel", Method: "POST"},
	{Path: "/v1/admin/userLabel/updateUserLabel/:id", AuthorityType: multi.TenancyAuthority, Description: "更新用户标签", ApiGroup: "userLabel", Method: "PUT"},
	{Path: "/v1/admin/userLabel/deleteUserLabel/:id", AuthorityType: multi.TenancyAuthority, Description: "删除用户标签", ApiGroup: "userLabel", Method: "DELETE"},
	// 用户手动标签
	{Path: "/v1/merchant/userLabel/getLabelList", AuthorityType: multi.TenancyAuthority, Description: "用户标签列表", ApiGroup: "userLabelClient", Method: "POST"},
	{Path: "/v1/merchant/userLabel/getCreateUserLabelMap", AuthorityType: multi.TenancyAuthority, Description: "用户标签添加表单", ApiGroup: "userLabelClient", Method: "GET"},
	{Path: "/v1/merchant/userLabel/getUpdateUserLabelMap/:id", AuthorityType: multi.TenancyAuthority, Description: "用户标签编辑表单", ApiGroup: "userLabelClient", Method: "GET"},
	{Path: "/v1/merchant/userLabel/createUserLabel", AuthorityType: multi.TenancyAuthority, Description: "添加用户标签", ApiGroup: "userLabelClient", Method: "POST"},
	{Path: "/v1/merchant/userLabel/updateUserLabel/:id", AuthorityType: multi.TenancyAuthority, Description: "更新用户标签", ApiGroup: "userLabelClient", Method: "PUT"},
	{Path: "/v1/merchant/userLabel/deleteUserLabel/:id", AuthorityType: multi.TenancyAuthority, Description: "删除用户标签", ApiGroup: "userLabelClient", Method: "DELETE"},
	// 用户自动标签
	{Path: "/v1/merchant/userLabel/auto/getLabelList", AuthorityType: multi.TenancyAuthority, Description: "用户标签列表", ApiGroup: "userAutoLabelClient", Method: "POST"},
	{Path: "/v1/merchant/userLabel/auto/createUserLabel", AuthorityType: multi.TenancyAuthority, Description: "添加用户标签", ApiGroup: "userAutoLabelClient", Method: "POST"},
	{Path: "/v1/merchant/userLabel/auto/updateUserLabel/:id", AuthorityType: multi.TenancyAuthority, Description: "更新用户标签", ApiGroup: "userAutoLabelClient", Method: "PUT"},
	{Path: "/v1/merchant/userLabel/auto/deleteUserLabel/:id", AuthorityType: multi.TenancyAuthority, Description: "删除用户标签", ApiGroup: "userAutoLabelClient", Method: "DELETE"},
	// 管理员管理
	{Path: "/v1/merchant/user/changeProfileMap", AuthorityType: multi.TenancyAuthority, Description: "修改个人信息表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/merchant/user/changeLoginPasswordMap", AuthorityType: multi.TenancyAuthority, Description: "修改登录用户密码表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/merchant/user/changePasswordMap/:id", AuthorityType: multi.TenancyAuthority, Description: "管理员修改密码表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/merchant/user/updateAdminMap/:id", AuthorityType: multi.TenancyAuthority, Description: "管理员编辑表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/merchant/user/registerAdminMap", AuthorityType: multi.TenancyAuthority, Description: "管理员注册表单", ApiGroup: "user", Method: "GET"},
	{Path: "/v1/merchant/user/registerAdmin", AuthorityType: multi.TenancyAuthority, Description: "管理员注册", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/merchant/user/changePassword", AuthorityType: multi.TenancyAuthority, Description: "修改密码", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/merchant/user/changeProfile", AuthorityType: multi.TenancyAuthority, Description: "修改个人信息", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/merchant/user/getAdminList", AuthorityType: multi.TenancyAuthority, Description: "管理员列表", ApiGroup: "user", Method: "POST"},
	{Path: "/v1/merchant/user/setUserInfo/:user_id", AuthorityType: multi.TenancyAuthority, Description: "设置用户信息", ApiGroup: "user", Method: "PUT"},
	{Path: "/v1/merchant/user/deleteUser", AuthorityType: multi.TenancyAuthority, Description: "删除用户", ApiGroup: "user", Method: "DELETE"},
	// 用户管理
	{Path: "/v1/merchant/cuser/getGeneralList", AuthorityType: multi.TenancyAuthority, Description: "c用户列表", ApiGroup: "cuserClient", Method: "POST"},
	{Path: "/v1/merchant/cuser/getGeneralDetail/:id", AuthorityType: multi.TenancyAuthority, Description: "c用户列表", ApiGroup: "cuserClient", Method: "GET"},
	{Path: "/v1/merchant/cuser/getOrderList/:id", AuthorityType: multi.TenancyAuthority, Description: "消费列表", ApiGroup: "cuserClient", Method: "POST"},
	{Path: "/v1/merchant/cuser/setUserLabel/:id", AuthorityType: multi.TenancyAuthority, Description: "设置用户标签", ApiGroup: "cuserClient", Method: "POST"},
	{Path: "/v1/merchant/cuser/setUserLabelMap/:id", AuthorityType: multi.TenancyAuthority, Description: "设置用户标签表单", ApiGroup: "cuserClient", Method: "GET"},

	//物流信息
	{Path: "/v1/merchant/express/getExpressByCode/:code", AuthorityType: multi.TenancyAuthority, Description: "物流信息", ApiGroup: "expressClient", Method: "GET"},
	// 操作记录
	{Path: "/v1/merchant/sysOperationRecord/getSysOperationRecordList", AuthorityType: multi.TenancyAuthority, Description: "操作记录列表", ApiGroup: "sysOperationRecordClient", Method: "POST"},

	//控制台
	{Path: "/v1/merchant/statistics/main", AuthorityType: multi.TenancyAuthority, Description: "主要数据", ApiGroup: "statisticsClient", Method: "GET"},
	{Path: "/v1/merchant/statistics/order", AuthorityType: multi.TenancyAuthority, Description: "支付订单", ApiGroup: "statisticsClient", Method: "GET"},
	{Path: "/v1/merchant/statistics/user", AuthorityType: multi.TenancyAuthority, Description: "成交用户", ApiGroup: "statisticsClient", Method: "GET"},
	{Path: "/v1/merchant/statistics/userRate", AuthorityType: multi.TenancyAuthority, Description: "成交用户占比", ApiGroup: "statisticsClient", Method: "GET"},
	{Path: "/v1/merchant/statistics/product", AuthorityType: multi.TenancyAuthority, Description: "商品支付排行", ApiGroup: "statisticsClient", Method: "GET"},
	{Path: "/v1/merchant/statistics/productVisit", AuthorityType: multi.TenancyAuthority, Description: "商品访客排行", ApiGroup: "statisticsClient", Method: "GET"},
	{Path: "/v1/merchant/statistics/productCart", AuthorityType: multi.TenancyAuthority, Description: "商品加购排行", ApiGroup: "statisticsClient", Method: "GET"},

	//TODO:: 用户权限
	// 我的地址
	{Path: "/v1/user/address/getAddressList", AuthorityType: multi.GeneralAuthority, Description: "地址列表", ApiGroup: "address", Method: "POST"},
	{Path: "/v1/user/address/createAddress", AuthorityType: multi.GeneralAuthority, Description: "添加地址", ApiGroup: "address", Method: "POST"},
	{Path: "/v1/user/address/getAddressById/:id", AuthorityType: multi.GeneralAuthority, Description: "地址详情", ApiGroup: "address", Method: "GET"},
	{Path: "/v1/user/address/updateAddress/:id", AuthorityType: multi.GeneralAuthority, Description: "更新地址", ApiGroup: "address", Method: "PUT"},
	{Path: "/v1/user/address/deleteAddress/:id", AuthorityType: multi.GeneralAuthority, Description: "删除地址", ApiGroup: "address", Method: "DELETE"},
	// 我的发票
	{Path: "/v1/user/receipt/getReceiptList", AuthorityType: multi.GeneralAuthority, Description: "发票列表", ApiGroup: "receipt", Method: "POST"},
	{Path: "/v1/user/receipt/createReceipt", AuthorityType: multi.GeneralAuthority, Description: "添加发票", ApiGroup: "receipt", Method: "POST"},
	{Path: "/v1/user/receipt/getReceiptById/:id", AuthorityType: multi.GeneralAuthority, Description: "发票详情", ApiGroup: "receipt", Method: "GET"},
	{Path: "/v1/user/receipt/updateReceipt/:id", AuthorityType: multi.GeneralAuthority, Description: "更新发票", ApiGroup: "receipt", Method: "PUT"},
	{Path: "/v1/user/receipt/deleteReceipt/:id", AuthorityType: multi.GeneralAuthority, Description: "删除发票", ApiGroup: "receipt", Method: "DELETE"},
}

//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, len(baseApis)}).Find(&[]model.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&baseApis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis 表初始数据成功!")
		return nil
	})
}
