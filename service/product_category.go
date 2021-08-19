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
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// GetProductCategoryMap
func GetProductCategoryMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	uploadUrl := SetUrl("/setting/uploadPicture?field=pic&type=1", ctx)
	if id > 0 {
		cate, err := GetProductCategoryByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"checkStrictly":true,"emitPath":false},"filterable":true},"validate":[{"required":true,"type":"integer","trigger":"change"}]},{"type":"input","field":"cateName","value":"%s","title":"分类名称","props":{"type":"text","placeholder":"请输入分类名称"},"validate":[{"message":"请输入分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"pic","value":"%s","title":"分类图片(110*110px)","props":{"type":"image","maxLength":1,"title":"请选择分类图片(110*110px)","src":"%s","width":"896px","height":"480px","footer":false,"modal":{"modal":false,"custom-class":"suibian-modal"}}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"","method":"PUT","title":"编辑分类","config":{}}`, cate.Pid, cate.CateName, cate.Pic, uploadUrl, cate.Status, cate.Sort)
	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"checkStrictly":true,"emitPath":false},"filterable":true},"validate":[{"required":true,"type":"integer","trigger":"change"}]},{"type":"input","field":"cateName","value":"%s","title":"分类名称","props":{"type":"text","placeholder":"请输入分类名称"},"validate":[{"message":"请输入分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"pic","value":"%s","title":"分类图片(110*110px)","props":{"type":"image","maxLength":1,"title":"请选择分类图片(110*110px)","src":"%s","width":"896px","height":"480px","footer":false,"modal":{"modal":false,"custom-class":"suibian-modal"}}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"","method":"POST","title":"添加分类","config":{}}`, 0, "", "", uploadUrl, 2, 0)
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}

	if id > 0 {
		form.SetAction(fmt.Sprintf("/productCategory/updateProductCategory/%d", id), ctx)
	} else {
		form.SetAction("/productCategory/createProductCategory", ctx)
	}
	opts, err := GetProductCategoriesOptions(multi.GetTenancyId(ctx), IsCuser(ctx))
	if err != nil {
		return form, err
	}
	form.Rule[0].Props["options"] = opts
	return form, err
}

// CreateProductCategory
func CreateProductCategory(productCategory model.ProductCategory, tenancyId uint) (model.ProductCategory, error) {
	err := g.TENANCY_DB.Where("cate_name = ?", productCategory.CateName).First(&productCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return productCategory, errors.New("名称已被注冊")
	}

	productCategory.SysTenancyID = tenancyId
	err = g.TENANCY_DB.Create(&productCategory).Error
	return productCategory, err
}

// GetProductCategoryByID
func GetProductCategoryByID(id uint) (model.ProductCategory, error) {
	var productCategory model.ProductCategory
	err := g.TENANCY_DB.Where("id = ?", id).First(&productCategory).Error
	return productCategory, err
}

// ChangeProductCategoryStatus
func ChangeProductCategoryStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.ProductCategory{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateProductCategory
func UpdateProductCategory(productCategory model.ProductCategory, id uint) (model.ProductCategory, error) {
	err := g.TENANCY_DB.Where("cate_name = ?", productCategory.CateName).Not("id = ?", id).First(&productCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return productCategory, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Updates(productCategory).Error
	return productCategory, err
}

// DeleteProductCategory
func DeleteProductCategory(id uint) error {
	var productCategory model.ProductCategory
	return g.TENANCY_DB.Where("id = ?", id).Delete(&productCategory).Error
}

// GetCategoryInfoList
func GetProductCategoryInfoList(tenancyId uint, isCuser bool) ([]response.ProductCategory, error) {
	productCategoryList := []response.ProductCategory{}
	treeMap, err := getProductCategoryMap(tenancyId, isCuser)
	productCategoryList = treeMap[0]
	for i := 0; i < len(productCategoryList); i++ {
		err = getProductCategoryBaseChildrenList(&productCategoryList[i], treeMap)
	}
	return productCategoryList, err
}

// getProductCategoryMap
func getProductCategoryMap(tenancyId uint, isCuser bool) (map[int32][]response.ProductCategory, error) {
	productCategoryList := []response.ProductCategory{}
	treeMap := make(map[int32][]response.ProductCategory)
	db := g.TENANCY_DB.Model(&model.ProductCategory{})
	db = CheckTenancyId(db, tenancyId, "")
	if isCuser {
		db = db.Where("status = ?", g.StatusTrue)
		db = db.Select("id,pid,cate_name,pic")
	}
	err := db.Order("sort").Find(&productCategoryList).Error
	for _, v := range productCategoryList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// getProductCategoryBaseChildrenList
func getProductCategoryBaseChildrenList(cate *response.ProductCategory, treeMap map[int32][]response.ProductCategory) (err error) {
	cate.Children = treeMap[int32(cate.ID)]
	for i := 0; i < len(cate.Children); i++ {
		err = getProductCategoryBaseChildrenList(&cate.Children[i], treeMap)
	}
	return err
}

// GetProductCategoriesOptions
func GetProductCategoriesOptions(tenancyId uint, isCuser bool) ([]Option, error) {
	options := []Option{}
	options = append(options, Option{Label: "请选择", Value: 0})
	treeMap, err := getProductCategoryMap(tenancyId, isCuser)

	for _, opt := range treeMap[0] {
		options = append(options, Option{Label: opt.CateName, Value: opt.ID})
	}
	for i := 0; i < len(options); i++ {
		getProductCategoriesOption(&options[i], treeMap)
	}

	return options, err
}

// getProductCategoriesOption
func getProductCategoriesOption(op *Option, treeMap map[int32][]response.ProductCategory) {
	id, ok := op.Value.(uint)
	if ok {
		for _, opt := range treeMap[int32(id)] {
			op.Children = append(op.Children, Option{Label: opt.CateName, Value: opt.ID})
		}
		for i := 0; i < len(op.Children); i++ {
			getProductCategoriesOption(&op.Children[i], treeMap)
		}
	}
}

func getProductCatesByProductId(productId, tenancyId uint) ([]response.ProductCate, error) {
	productCates := []response.ProductCate{}
	err := g.TENANCY_DB.Model(&model.ProductCategory{}).Select("product_categories.*").
		Joins("left join product_product_cates on product_product_cates.product_category_id = product_categories.id and product_product_cates.sys_tenancy_id = product_categories.sys_tenancy_id").
		Where("product_categories.sys_tenancy_id = ?", tenancyId).
		Where("product_product_cates.product_id = ?", productId).
		Find(&productCates).Error

	return productCates, err
}

func getProductIdsByProductCategoryId(productCategoryId, tenancyId uint) ([]uint, error) {
	productIds := []uint{}
	db := g.TENANCY_DB.Model(&model.ProductProductCate{}).Select("product_id").Where("product_category_id = ?", productCategoryId)
	db = CheckTenancyId(db, tenancyId, "")
	err := db.
		Find(&productIds).Error

	return productIds, err
}
