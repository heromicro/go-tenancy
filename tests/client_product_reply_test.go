package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/tests/base"
)

func TestClientProductReplyList(t *testing.T) {
	t.SkipNow()
	params := []base.Param{
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "", "nickname": ""}, length: 4},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "1", "nickname": ""}, length: 4},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "2", "nickname": ""}, length: 0},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "1", "nickname": "C"}, length: 4},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "", "isReply": 0, "keyword": "1", "nickname": "B"}, length: 0},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "year", "isReply": 0, "keyword": "1", "nickname": "C"}, length: 4},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "yesterday", "isReply": 0, "keyword": "1", "nickname": "C"}, length: 0},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "year", "isReply": 1, "keyword": "1", "nickname": "C"}, length: 4},
		// {args: map[string]interface{}{"page": 1, "pageSize": 10, "date": "year", "isReply": 2, "keyword": "1", "nickname": "C"}, length: 0},
	}
	for _, param := range params {
		fmt.Print(param)
		// productReplyClientlist(t, param.args, param.length)
	}
}

func productReplyClientlist(t *testing.T, params map[string]interface{}, length int) {
	auth, _ := base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(auth)
	obj := auth.POST("v1/merchant/productReply/getProductReplyList").
		WithJSON(params).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(length)

	if length > 0 {
		list := data.Value("list").Array()
		list.Length().Ge(0)
		first := list.First().Object()
		first.Keys().ContainsOnly(
			"id",
			"createdAt",
			"updatedAt",
			"productScore",
			"serviceScore",
			"postageScore",
			"rate",
			"comment",
			"pics",
			"merchantReplyContent",
			"merchantReplyTime",
			"isReply",
			"isVirtual",
			"avatar",
			"productId",
			"sysUserId",
			"nickname",
			"storeName",
			"image",
			"images",
		)
		first.Value("id").Number().Ge(0)
	}
}

func TestClientProductReply(t *testing.T) {
	var brandId, shipTempId, cateId, tenancyCategoryId, productId, cartId uint
	var unique string
	var productType int32
	var adminAuth, tenancyAuth, deviceAuth *httpexpect.Expect

	adminAuth = base.BaseWithLoginTester(t)
	defer base.BaseLogOut(adminAuth)

	tenancyAuth, _ = base.TenancyWithLoginTester(t)
	defer base.BaseLogOut(tenancyAuth)

	deviceAuth = base.DeviceWithLoginTester(t)
	defer base.BaseLogOut(deviceAuth)

	brandCategoryPid, _ := CreateBrandCategory(t, adminAuth, "箱包服饰_product_detail", 0, http.StatusOK, "创建成功")
	defer DeleteBrandCategory(adminAuth, brandCategoryPid)

	brandCategoryId, _ := CreateBrandCategory(t, adminAuth, "精品服饰_product_detail", brandCategoryPid, http.StatusOK, "创建成功")

	defer DeleteBrandCategory(adminAuth, brandCategoryId)

	brandId, _ = CreateBrand(t, adminAuth, "冈本_product_detail", brandCategoryId, http.StatusOK, "创建成功")
	defer DeleteBrand(adminAuth, brandId)

	cateId, _ = CreateCategory(adminAuth, "数码产品_product_detail", 0, http.StatusOK, "创建成功")
	if cateId == 0 {
		t.Error("添加分类失败")
		return
	}
	defer DeleteCategory(adminAuth, cateId, http.StatusOK, "删除成功")

	shipTempId, _ = CreateShippingTemplate(tenancyAuth, "ship_temp_name_物流邮费模板", http.StatusOK, "创建成功")
	if shipTempId == 0 {
		t.Error("添加物流模板失败")
		return
	}
	defer DeleteShippingTemplate(tenancyAuth, shipTempId, http.StatusOK, "删除成功")

	tenancyCategoryId, _ = ClientCreateCategory(tenancyAuth, "客户端数码产品_device_order", 0, http.StatusOK, "创建成功")
	if tenancyCategoryId == 0 {
		t.Error("添加商户分类失败")
		return
	}
	defer DeleteClientCategory(tenancyAuth, tenancyCategoryId, http.StatusOK, "删除成功")

	productId, productData := CreateProduct(tenancyAuth, cartId, brandId, shipTempId, tenancyCategoryId, http.StatusOK, "创建成功")
	if productId == 0 {
		t.Errorf("添加商品失败 商品id:%d ", productId)
		return
	}
	defer DeleteProduct(tenancyAuth, productId, http.StatusOK, "删除成功")

	unique, productType = GetProduct(tenancyAuth, productId, productData)
	if len(unique) == 0 || productType == 0 {
		t.Errorf("添加商品失败规格:%+v,商品类型:%d", unique, productType)
	}

	ChangeProductIsShow(tenancyAuth, productId, g.StatusTrue, http.StatusOK, "设置成功")

	obj := tenancyAuth.GET(fmt.Sprintf("v1/merchant/productReply/replyMap/%d", productId)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

	obj = tenancyAuth.POST(fmt.Sprintf("v1/merchant/productReply/reply/%d", productId)).
		WithJSON(map[string]interface{}{"content": "pageSize"}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")

}
