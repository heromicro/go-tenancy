package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/chindeo/pkg/file"
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetEditProductFictiMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	ficti, err := GetProductFictiByID(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"number","value":"%s","title":"现有虚拟销量","props":{"type":"text","placeholder":"请输入现有虚拟销量","readonly":true}},{"type":"radio","field":"type","value":1,"title":"修改类型","props":{},"options":[{"value":1,"label":"增加"},{"value":2,"label":"减少"}]},{"type":"inputNumber","field":"ficti","value":0,"title":"修改虚拟销量数","props":{"placeholder":"请输入修改虚拟销量数"}}],"action":"","method":"PUT","title":"修改虚拟销量数","config":{}}`, strconv.FormatInt(int64(ficti), 10))

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/product/setProductFicti", id), ctx)
	return form, err
}

// 出售中 1: is_show' => 1, 'status' => 1
// 仓库中 2:'is_show' => 2, 'status' => 1
// 3,4,5 商户才有
// 已售罄 3:'is_show' => 1, 'stock' => 0, 'status' => 1
// 警戒库存 4:'stock' => $stock ? $stock : 0, 'status' => 1
// 回收站 5:'deleted_at' => not null
// 待审核 6:'status' => 2
// 审核未通过 7:'status' => 3

// GetProductFilter
func GetProductFilter(info request.ProductPageInfo, tenancyId uint, isTenancy bool) ([]response.ProductFilter, error) {
	wheres := getProductConditions(tenancyId, isTenancy)
	filters := []response.ProductFilter{}
	for _, where := range wheres {
		filter := response.ProductFilter{Name: where.Name, Type: where.Type}
		db := g.TENANCY_DB.Model(&model.Product{})
		// 显示软删除数据
		if where.IsDeleted {
			db = db.Unscoped()
		}

		if where.Conditions != nil && len(where.Conditions) > 0 {
			for key, cn := range where.Conditions {
				if cn == nil {
					db = db.Where(key)
				} else {
					db = db.Where(fmt.Sprintf("%s = ?", key), cn)
				}
			}
		}

		if isTenancy {
			db = CheckTenancyId(db, tenancyId, "")
		}

		db, err := getProductSearch(db, info, tenancyId)
		if err != nil {
			return filters, err
		}

		err = db.Count(&filter.Count).Error
		if err != nil {
			return filters, err
		}
		filters = append(filters, filter)
	}

	return filters, nil
}

// getProductConditions
func getProductConditions(tenancyId uint, isTenancy bool) []response.ProductCondition {
	stock := 0
	if config, err := GetTenancyConfigValue("mer_store_stock", tenancyId); err == nil {
		if value, err := strconv.Atoi(config.Value); err == nil {
			stock = value
		}
	}

	conditions := []response.ProductCondition{
		{Name: "出售中", Type: 1, Conditions: map[string]interface{}{"is_show": g.StatusTrue, "status": model.SuccessProductStatus}},
		{Name: "仓库中", Type: 2, Conditions: map[string]interface{}{"is_show": g.StatusFalse, "status": model.SuccessProductStatus}},

		{Name: "待审核", Type: 6, Conditions: map[string]interface{}{"status": model.AuditProductStatus}},
		{Name: "审核未通过", Type: 7, Conditions: map[string]interface{}{"status": model.FailProductStatus}},
	}

	if isTenancy {
		other := []response.ProductCondition{{Name: "已售罄", Type: 3, Conditions: map[string]interface{}{"is_show": g.StatusTrue, "stock": stock, "status": model.SuccessProductStatus}},
			{Name: "警戒库存", Type: 4, Conditions: map[string]interface{}{"stock": stock, "status": model.SuccessProductStatus}},
			{Name: "回收站", Type: 5, Conditions: map[string]interface{}{"deleted_at is not null": nil}, IsDeleted: true},
		}
		conditions = append(conditions, other...)
	}
	return conditions
}

// getProductConditionByType
func getProductConditionByType(tenancyId uint, isTenancy bool, t int) response.ProductCondition {
	conditions := getProductConditions(tenancyId, isTenancy)
	for _, condition := range conditions {
		if condition.Type == t {
			return condition
		}
	}
	return conditions[0]
}

// CreateProduct
func CreateProduct(req request.CreateProduct, tenancyId uint) (uint, error) {
	product := model.Product{
		BaseProduct: req.BaseProduct,
		SliderImage: strings.Join(req.SliderImages, ","),
	}
	product.SysTenancyID = tenancyId
	product.ProductCategoryID = req.CateId
	product.IsHot = g.StatusFalse
	product.IsBenefit = g.StatusFalse
	product.IsBest = g.StatusFalse
	product.IsNew = g.StatusFalse
	product.ProductType = model.GeneralSale

	tenancy, err := GetTenancyByID(tenancyId)
	if err != nil {
		return product.ID, err
	}

	// 开启商品审核的商家，审核商品
	if tenancy.IsAudit == g.StatusTrue {
		product.Status = model.AuditProductStatus
	} else {
		product.Status = model.SuccessProductStatus
	}

	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Product{}).Create(&product).Error
		if err != nil {
			return fmt.Errorf("create product %w", err)
		}

		if len(req.CategoryIds) > 0 {
			err = SetProductCategory(tx, product.ID, tenancyId, req.CategoryIds)
			if err != nil {
				return fmt.Errorf("create product product cate %w", err)
			}
		}

		err = SetProductContent(tx, product.ID, tenancyId, product.ProductType, req.Content)
		if err != nil {
			return fmt.Errorf("set product content  %w", err)
		}

		err = SetProductAttrValue(tx, product.ID, req.ProductType, req.AttrValue, req.Attr)
		if err != nil {
			return fmt.Errorf("set product attr %w", err)
		}

		return nil
	})
	if err != nil {
		return product.ID, err
	}

	return product.ID, nil
}

// ChangeProduct
func ChangeProduct(req request.UpdateProduct, id uint, ctx *gin.Context) error {
	// 更新商品重新审核，并下架
	tenancyId := multi.GetTenancyId(ctx)
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if multi.IsAdmin(ctx) {
			data := map[string]interface{}{"store_name": req.StoreName, "is_benefit": req.IsBenefit, "is_best": req.IsBest, "is_hot": req.IsHot, "is_new": req.IsNew, "rank": req.Rank}
			if err := UpdateProduct(tx, id, data); err != nil {
				return err
			}
		} else if multi.IsTenancy(ctx) {
			tenancy, err := GetTenancyByID(tenancyId)
			if err != nil {
				return err
			}
			product := model.Product{
				BaseProduct: req.BaseProduct,
			}

			// 开启审核商品需要审核
			if tenancy.IsAudit == g.StatusTrue {
				product.Status = model.AuditProductStatus
			}
			product.IsShow = g.StatusFalse
			product.ProductCategoryID = req.CateId
			product.SliderImage = strings.Join(req.SliderImages, ",")
			if err := tx.Model(&model.Product{}).Where("id = ?", id).Updates(&product).Error; err != nil {
				return err
			}
			err = SetProductAttrValue(tx, id, req.ProductType, req.AttrValue, req.Attr)
			if err != nil {
				return fmt.Errorf("set product attr %w", err)
			}
			err = tx.Model(&model.Cart{}).Where("product_id = ?", id).Update("is_fail", g.StatusTrue).Error
			if err != nil {
				return err
			}
			if len(req.CategoryIds) > 0 {
				err := SetProductCategory(tx, product.ID, tenancyId, req.CategoryIds)
				if err != nil {
					return fmt.Errorf("set product product_cate %w", err)
				}
			}
		}

		err := SetProductContent(tx, id, tenancyId, req.ProductType, req.Content)
		if err != nil {
			return fmt.Errorf("set product content  %w", err)
		}

		return nil
	})
	return err
}

// UpdateProduct 更新产品信息
func UpdateProduct(db *gorm.DB, id uint, data map[string]interface{}) error {
	if err := db.Model(&model.Product{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		g.TENANCY_LOG.Error("更新产品信息错误", zap.String("UpdateProduct()", err.Error()))
		return fmt.Errorf("更新产品信息错误 %w", err)
	}
	return nil
}

// GetProductAttrValues 产品规格
func GetProductAttrValues(productIds []uint, uniques []string) ([]model.ProductAttrValue, error) {
	attrValues := []model.ProductAttrValue{}
	db := g.TENANCY_DB.Model(&model.ProductAttrValue{})
	if len(productIds) > 0 {
		db = db.Where("product_id in ?", productIds)
	}
	if len(uniques) > 0 {
		db = db.Where("`unique` in ?", uniques)
	}
	err := db.Find(&attrValues).Error
	if err != nil {
		return attrValues, fmt.Errorf("get product attr value %w", err)
	}
	return attrValues, nil
}

// GetProductAttrs 产品规格
func GetProductAttrs(productIds []uint) ([]model.ProductAttr, error) {
	attrs := []model.ProductAttr{}
	db := g.TENANCY_DB.Model(&model.ProductAttr{})
	if len(productIds) > 0 {
		db = db.Where("product_id in ?", productIds)
	}
	err := db.Find(&attrs).Error
	if err != nil {
		return attrs, fmt.Errorf("get product attr  %w", err)
	}
	return attrs, nil
}

// SetProductAttrValue 设置产品规格
func SetProductAttrValue(tx *gorm.DB, productId uint, productType int32, reqAttrValue []request.ProductAttrValue, reqAttr []request.ProductAttr) error {
	err := tx.Unscoped().Where("product_id = ?", productId).Delete(&model.ProductAttrValue{}).Error
	if err != nil {
		return fmt.Errorf("delete product attr value %w", err)
	}

	productAttrValues := []model.ProductAttrValue{}
	for _, attrValue := range reqAttrValue {
		unique, err := file.Md5Byte([]byte(fmt.Sprintf("%s:%d", attrValue.Detail.String(), productId)))
		if err != nil {
			return fmt.Errorf("get product attr value unique %w", err)
		}
		unique = fmt.Sprintf("%s%d", unique[12:23], productType)
		attrValue.BaseProductAttrValue.Sku = attrValue.Value0
		attrValue.BaseProductAttrValue.Unique = unique
		productAttrValue := model.ProductAttrValue{
			ProductID:            productId,
			BaseProductAttrValue: attrValue.BaseProductAttrValue,
			Detail:               attrValue.Detail.String(),
			Type:                 productType,
		}
		productAttrValues = append(productAttrValues, productAttrValue)
	}

	err = tx.Model(&model.ProductAttrValue{}).Create(&productAttrValues).Error
	if err != nil {
		return fmt.Errorf("create product attr value %w", err)
	}

	err = tx.Unscoped().Where("product_id = ?", productId).Delete(&model.ProductAttr{}).Error
	if err != nil {
		return fmt.Errorf("delete product attr %w", err)
	}

	productAttrs := []model.ProductAttr{}
	for _, attr := range reqAttr {
		if attr.Detail == nil || attr.Value == "" {
			continue
		}
		productAttr := model.ProductAttr{
			ProductID:  productId,
			AttrName:   attr.Value,
			AttrValues: attr.Detail.String(),
			Type:       productType,
		}
		productAttrs = append(productAttrs, productAttr)
	}

	if len(productAttrs) > 0 {
		err = tx.Model(&model.ProductAttr{}).Create(&productAttrs).Error
		if err != nil {
			return fmt.Errorf("create product attr value %w", err)
		}
	}

	return nil
}

func GetProductCategoryIdsById(id uint) ([]uint, error) {
	ids := []uint{}
	err := g.TENANCY_DB.Model(&model.ProductProductCate{}).Select("product_category_id").Where("product_id = ?", id).Find(&ids).Error
	if err != nil {
		return ids, err
	}
	return ids, nil
}

func SetProductContent(tx *gorm.DB, productId, tenancyId uint, productType int32, content string) error {
	var conModel model.ProductContent
	err := g.TENANCY_DB.Model(&model.ProductContent{}).Where("product_id = ?", productId).First(&conModel).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		con := model.ProductContent{Content: content, ProductID: productId, Type: productType}
		if err := tx.Model(&model.ProductContent{}).Create(&con).Error; err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		if err := tx.Model(&model.ProductContent{}).Where("product_id = ?", productId).Updates(map[string]interface{}{"content": content}).Error; err != nil {
			return err
		}
	}
	return nil
}

// SetProductCategory
func SetProductCategory(tx *gorm.DB, id, tenancyId uint, reqIds []uint) error {
	cateIds, err := GetProductCategoryIdsById(id)
	if err != nil {
		return err
	}

	// 删除
	delIds := []uint{}
	for _, cateId := range cateIds {
		isDel := true
		for _, reqlId := range reqIds {
			if cateId == reqlId {
				isDel = false
				break
			}
		}
		if isDel {
			delIds = append(delIds, cateId)
		}
	}

	if len(delIds) > 0 {
		if err = tx.Where("product_id = ?", id).Where("sys_tenancy_id = ?", tenancyId).Where("product_category_id in ?", delIds).Delete(&model.ProductProductCate{}).Error; err != nil {
			return fmt.Errorf("delete product_product_categorys %w", err)
		}
	}

	// 增加
	addIds := []uint{}
	for _, reqId := range reqIds {
		isAdd := true
		for _, cateId := range cateIds {
			if reqId == cateId {
				isAdd = false
				break
			}
		}
		if isAdd {
			addIds = append(addIds, reqId)
		}
	}

	if len(addIds) > 0 {
		var cates []model.ProductProductCate
		for _, addId := range addIds {
			cates = append(cates, model.ProductProductCate{ProductID: id, ProductCategoryID: addId, SysTenancyID: tenancyId})
		}
		if err = tx.Model(&model.ProductProductCate{}).Create(&cates).Error; err != nil {
			return fmt.Errorf("create product_product_categorys %w", err)
		}
	}

	return nil
}

func GetProductForReplysByIds(ids []uint, tenancyId uint) ([]response.ProductForReply, error) {
	productForReplies := []response.ProductForReply{}
	db := g.TENANCY_DB.Model(&model.Product{}).Select("id,store_name,image")
	db = CheckTenancyId(db, tenancyId, "")
	err := db.Where("id in ?", ids).
		Find(&productForReplies).Error
	if err != nil {
		return productForReplies, err
	}
	return productForReplies, nil
}

func GetProductIdsByKeyword(keyword string, tenancyId uint) ([]uint, error) {
	ids := []uint{}
	db := g.TENANCY_DB.Model(&model.Product{}).Select("id")
	db = CheckTenancyId(db, tenancyId, "")
	err := db.Where(g.TENANCY_DB.Where("id like ?", keyword+"%").Or("store_name like ?", keyword+"%").Or("store_info like ?", keyword+"%").Or("keyword like ?", keyword+"%")).
		Find(&ids).Error
	if err != nil {
		return ids, err
	}
	return ids, nil
}

// GetCartProducts
func GetCartProducts(sysTenancyID, sysUserID, patientId uint, cartIds []uint) ([]response.CartProduct, error) {
	cartProducts := []response.CartProduct{}
	db := g.TENANCY_DB.Model(&model.Product{}).Where("products.is_show = ?", g.StatusTrue).Where("products.status = ?", model.SuccessProductStatus).Select("products.id as product_id,products.store_name,products.image,products.spec_type,products.price,carts.id,carts.cart_num,carts.sys_tenancy_id as sys_tenancy_id,carts.product_attr_unique,carts.is_fail").
		Joins("left join carts on products.id = carts.product_id")
	db = CheckTenancyId(db, sysTenancyID, "carts.")
	db.Where("carts.sys_user_id = ?", sysUserID).
		Where("carts.patient_id = ?", patientId).
		Where("carts.is_pay = ?", g.StatusFalse).
		Where("carts.deleted_at is null")

	if len(cartIds) > 0 {
		db = db.Where("carts.id in ?", cartIds)
	}

	err := db.Find(&cartProducts).Error
	if err != nil {
		return cartProducts, err
	}

	productIds := []uint{}
	uniques := []string{}
	for _, cartProduct := range cartProducts {
		if cartProduct.SpecType == model.SingleSpec {
			productIds = append(productIds, cartProduct.ProductID)
		} else if cartProduct.SpecType == model.DoubleSpec {
			productIds = append(productIds, cartProduct.ProductID)
		}

		uniques = append(uniques, cartProduct.ProductAttrUnique)
	}

	attrValues, err := GetProductAttrValues(productIds, uniques)
	if err != nil {
		return cartProducts, fmt.Errorf("get product attr value %w", err)
	}

	if len(attrValues) > 0 {
		for _, attrValue := range attrValues {
			for i := 0; i < len(cartProducts); i++ {
				if cartProducts[i].ProductID == attrValue.ProductID && cartProducts[i].ProductAttrUnique == attrValue.Unique {
					productAttrValue := request.ProductAttrValue{BaseProductAttrValue: attrValue.BaseProductAttrValue, Value0: attrValue.BaseProductAttrValue.Sku}
					if attrValue.Detail != "" {
						err := json.Unmarshal([]byte(attrValue.Detail), &productAttrValue.Detail)
						if err != nil {
							return cartProducts, fmt.Errorf("json product attr value detail marshal %w", err)
						}
					}
					productAttrValue.Value0 = attrValue.BaseProductAttrValue.Sku
					cartProducts[i].AttrValue = productAttrValue
				}
			}
		}
	}

	return cartProducts, err
}

// GetProductByID
func GetProductByID(id, tenancyId uint, isCuser bool) (response.ProductDetail, error) {
	var product response.ProductDetail
	db := g.TENANCY_DB.Model(&model.Product{})
	if isCuser {
		// 用户端成本价
		db = db.Where("products.is_show = ?", g.StatusTrue).Where("products.status = ?", model.SuccessProductStatus).Select("products.id,products.store_name,products.store_info,products.keyword,products.ficti,products.unit_name,products.sort,products.sales,products.price,products.ot_price,products.stock,products.is_hot,products.is_benefit,products.is_best,products.is_new,products.is_good,products.product_type,products.spec_type,products.rate,products.image,products.temp_id,products.sys_tenancy_id,products.sys_brand_id,products.product_category_id,products.slider_image,sys_tenancies.name as sys_tenancy_name,sys_brands.brand_name as brand_name,product_categories.cate_name as cate_name,product_contents.content as content,shipping_templates.name as temp_name")
	} else {
		db = db.Select("products.*,sys_tenancies.name as sys_tenancy_name,sys_brands.brand_name as brand_name,product_categories.cate_name as cate_name,product_contents.content as content,shipping_templates.name as temp_name")
	}

	db = db.Joins("left join sys_tenancies on products.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_brands on products.sys_brand_id = sys_brands.id").
		Joins("left join product_categories on products.product_category_id = product_categories.id").
		Joins("left join product_contents on product_contents.product_id = products.id").
		Joins("left join shipping_templates on products.temp_id = shipping_templates.id").
		Where("products.id = ?", id)
	db = CheckTenancyId(db, tenancyId, "products.")
	err := db.First(&product).Error
	if err != nil {
		return product, err
	}
	product.SliderImages = strings.Split(product.SliderImage, ",")

	attrValues, err := GetProductAttrValues([]uint{id}, nil)
	if err != nil {
		return product, err
	}
	productAttrValues := []request.ProductAttrValue{}
	if len(attrValues) > 0 {
		for _, attrValue := range attrValues {
			productAttrValue := request.ProductAttrValue{BaseProductAttrValue: attrValue.BaseProductAttrValue, Value0: attrValue.BaseProductAttrValue.Sku}
			if attrValue.Detail != "" {
				err := json.Unmarshal([]byte(attrValue.Detail), &productAttrValue.Detail)
				if err != nil {
					return product, fmt.Errorf("json product attr value detail marshal %w", err)
				}
			}
			productAttrValue.Value0 = attrValue.BaseProductAttrValue.Sku
			productAttrValues = append(productAttrValues, productAttrValue)
		}
	}
	product.AttrValue = productAttrValues

	attrs, err := GetProductAttrs([]uint{id})
	if err != nil {
		return product, err
	}
	productAttrs := []request.ProductAttr{}
	if len(attrs) > 0 {
		for _, attr := range attrs {
			productAttr := request.ProductAttr{Value: attr.AttrName}
			if attr.AttrValues != "" {
				err := json.Unmarshal([]byte(attr.AttrValues), &productAttr.Detail)
				if err != nil {
					return product, fmt.Errorf("json product attr detail marshal %w", err)
				}
			}
			productAttrs = append(productAttrs, productAttr)
		}
	}
	product.Attr = productAttrs

	product.CateId = product.ProductCategoryID
	product.SliderImages = strings.Split(product.SliderImage, ",")

	productCates, err := getProductCatesByProductId(product.ID, tenancyId)
	if err != nil {
		return product, err
	}
	product.ProductCates = productCates

	var categoryIds []uint
	for _, productCate := range productCates {
		categoryIds = append(categoryIds, productCate.ID)
	}
	product.CategoryIds = categoryIds

	return product, err
}

// GetProductFictiByID
func GetProductFictiByID(id uint) (int32, error) {
	var product response.ProductFicti
	err := g.TENANCY_DB.Model(&model.Product{}).
		Select("ficti").
		Where("products.id = ?", id).
		First(&product).Error
	return product.Ficti, err
}

// ChangeProductStatus
func ChangeProductStatus(changeStatus request.ChangeProductStatus) error {
	return g.TENANCY_DB.Model(&model.Product{}).Where("id in ?", changeStatus.Id).Updates(map[string]interface{}{"status": changeStatus.Status, "refusal": changeStatus.Refusal}).Error
}

// ChangeProductIsShow
func ChangeProductIsShow(changeStatus request.ChangeProductIsShow) error {
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Product{}).Where("id = ?", changeStatus.Id).Updates(map[string]interface{}{"is_show": changeStatus.IsShow}).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.Cart{}).Where("product_id = ?", changeStatus.Id).Update("is_fail", g.StatusTrue).Error
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// SetProductFicti
func SetProductFicti(req request.SetProductFicti, id uint) error {
	ficti, err := GetProductFictiByID(id)
	if err != nil {
		return err
	}
	// 增加
	if req.Type == 1 {
		ficti = ficti + req.Ficti
	} else if req.Type == 2 {
		if ficti <= req.Ficti {
			ficti = 0
		} else {
			ficti = ficti - req.Ficti
		}
	}
	if err := g.TENANCY_DB.Model(&model.Product{}).Where("id = ?", id).Updates(map[string]interface{}{"ficti": ficti}).Error; err != nil {
		return err
	}
	return err
}

// DeleteProduct
func DeleteProduct(id uint) error {
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&model.Product{}, id).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.Cart{}).Where("product_id = ?", id).Update("is_fail", g.StatusTrue).Error
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// RestoreProduct
func RestoreProduct(id uint) error {
	return g.TENANCY_DB.Model(&model.Product{}).Unscoped().Where("id = ?", id).Updates(map[string]interface{}{"deleted_at": nil}).Error
}

// ForceDeleteProduct
func ForceDeleteProduct(id uint) error {
	return g.TENANCY_DB.Unscoped().Delete(&model.Product{}, 1).Error
}

// GetProductInfoList
func GetProductInfoList(info request.ProductPageInfo, tenancyId uint, isTenancy, isCuser bool) ([]response.ProductList, int64, error) {
	productList := []response.ProductList{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Product{}).
		Joins("left join sys_tenancies on products.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_brands on products.sys_brand_id = sys_brands.id").
		Joins("left join product_categories on products.product_category_id = product_categories.id")

	if isCuser {
		db = db.Select("products.id,products.store_name,products.price,products.image,products.sales").
			Where("products.is_show = ?", g.StatusTrue).Where("products.status = ?", model.SuccessProductStatus)
	} else {
		db = db.Select("products.*,sys_tenancies.name as sys_tenancy_name,sys_brands.brand_name as brand_name,product_categories.cate_name as cate_name")
		if info.Type != "" {
			t, err := strconv.Atoi(info.Type)
			if err != nil {
				return productList, total, err
			}
			cond := getProductConditionByType(tenancyId, isTenancy, t)
			if cond.IsDeleted {
				db = db.Unscoped()
			}
			for key, cn := range cond.Conditions {
				if cn == nil {
					db = db.Where(fmt.Sprintf("%s%s", "products.", key))
				} else {
					db = db.Where(fmt.Sprintf("%s%s = ?", "products.", key), cn)
				}
			}
		}
	}

	db, err := getProductSearch(db, info, tenancyId)
	if err != nil {
		return productList, total, err
	}
	db = CheckTenancyId(db, tenancyId, "products.")

	err = db.Count(&total).Error
	if err != nil {
		return productList, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy, "products.")
	err = db.Limit(limit).Offset(offset).Find(&productList).Error

	for i := 0; i < len(productList); i++ {
		productCates, err := getProductCatesByProductId(productList[i].ID, productList[i].SysTenancyID)
		if err != nil {
			continue
		}
		productList[i].ProductCates = productCates
	}

	return productList, total, err
}

func getProductSearch(db *gorm.DB, info request.ProductPageInfo, tenancyId uint) (*gorm.DB, error) {
	if info.Keyword != "" {
		db = db.Where(g.TENANCY_DB.Where("products.store_name like ?", info.Keyword+"%").Or("products.store_info like ?", info.Keyword+"%").Or("products.keyword like ?", info.Keyword+"%").Or("products.bar_code like ?", info.Keyword+"%"))
	}
	// 平台分类id
	if info.ProductCategoryId > 0 {
		productIds, err := getProductIdsByProductCategoryId(info.ProductCategoryId, tenancyId)
		if err != nil {
			return db, err
		}
		db = db.Where("products.id in ?", productIds)
	}

	// 平台分类id
	if info.CateId > 0 {
		db = db.Where("products.product_category_id = ?", info.CateId)
	}

	return db, nil
}

func GetProductSelect(tenancyId uint) ([]response.SelectOption, error) {
	selects := []response.SelectOption{
		{ID: 0, Name: "请选择"},
	}
	var userSelects []response.SelectOption
	err := g.TENANCY_DB.Model(&model.Product{}).
		Select("id,store_name as name").
		Where("status = ?", g.StatusTrue).
		Where("is_show = ?", g.StatusTrue).
		Find(&userSelects).Error
	selects = append(selects, userSelects...)
	return selects, err
}

// UpdateProductAttrValue 更新商品规格
func UpdateProductAttrValue(db *gorm.DB, productId uint, unique string, data map[string]interface{}) error {
	err := db.Model(&model.ProductAttrValue{}).
		Where("id = ?", productId).
		Where("`unique` = ?", unique).
		Updates(&data).Error
	if err != nil {
		g.TENANCY_LOG.Error("更新商品规格错误", zap.String("UpdateProductAttrValue()", err.Error()))
		return fmt.Errorf("更新商品规格错误 %w", err)
	}
	return nil
}

// IncStock 回退商品库存
func IncStock(db *gorm.DB, id uint, inc int64) error {
	data := map[string]interface{}{"stock": gorm.Expr("stock+?", inc), "sales": gorm.Expr("sales-(IF(sales >?,?,0))", inc, inc)}
	return UpdateProduct(db, id, data)
}

// DecStock 减去商品库存
func DecStock(db *gorm.DB, id uint, dec int64) error {
	data := map[string]interface{}{"stock": gorm.Expr("stock-(IF(stock >?,?,0))", dec, dec), "sales": gorm.Expr("sales+?", dec)}
	return UpdateProduct(db, id, data)
}

// IncSkuStock 回退商品规格库存
func IncSkuStock(db *gorm.DB, id uint, unique string, inc int64) error {
	data := map[string]interface{}{"stock": gorm.Expr("stock+?", inc), "sales": gorm.Expr("sales-(IF(sales >?,?,0))", inc, inc)}
	return UpdateProductAttrValue(db, id, unique, data)
}

// DecSkuStock 减去商品规格库存
func DecSkuStock(db *gorm.DB, id uint, unique string, dec int64) error {
	data := map[string]interface{}{"stock": gorm.Expr("stock-(IF(stock >?,?,0))", dec, dec), "sales": gorm.Expr("sales+?", dec)}
	return UpdateProductAttrValue(db, id, unique, data)
}

