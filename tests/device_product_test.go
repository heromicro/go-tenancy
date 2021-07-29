package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDeviceProductList(t *testing.T) {
	auth := deviceWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/device/product/getProductList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10, "type": "0", "cateId": 0, "isGiftBag": "", "keyword": "", "tenancyCategoryId": 0}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Ge(0)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("keyword",
		"isBest",
		"specType",
		"refusal",
		"rate",
		"oldId",
		"sysTenancyId",
		"productCates",
		"rank",
		"isNew",
		"updatedAt",
		"sales",
		"otPrice",
		"isHot",
		"isGood",
		"ficti",
		"id",
		"status",
		"cost",
		"image",
		"tempId",
		"createdAt",
		"barCode",
		"unitName",
		"stock",
		"replyCount",
		"productCategoryId",
		"sysTenancyName",
		"isShow",
		"isBenefit",
		"codePath",
		"videoLink",
		"sysBrandId",
		"cateName",
		"brandName",
		"storeInfo",
		"sort",
		"price",
		"careCount",
		"storeName",
		"productType",
		"browse",
		"isGiftBag")
	first.Value("id").Number().Ge(0)
}

func TestDeviceProductDetail(t *testing.T) {
	auth := deviceWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.GET(fmt.Sprintf("v1/device/product/getProductById/%d", 1)).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	product := obj.Value("data").Object()
	product.Value("id").Number().Equal(1)
	product.Value("storeName").String().Equal("领立裁腰带短袖连衣裙")
	product.Value("storeInfo").String().Equal("短袖连衣裙")
	product.Value("keyword").String().Equal("连衣裙")
	product.Value("unitName").String().Equal("件")
	product.Value("sort").Number().Equal(40)
	product.Value("sales").Number().Equal(1)
	product.Value("price").Number().Equal(80)
	product.Value("otPrice").Number().Equal(100)
	product.Value("stock").Number().Equal(399)
	product.Value("isHot").Number().Equal(2)
	product.Value("isBenefit").Number().Equal(2)
	product.Value("isBest").Number().Equal(2)
	product.Value("isNew").Number().Equal(2)
	product.Value("isGood").Number().Equal(1)
	product.Value("productType").Number().Equal(2)
	product.Value("ficti").Number().Equal(100)
	product.Value("specType").Number().Equal(1)
	product.Value("rate").Number().Equal(5)
	product.Value("isGiftBag").Number().Equal(2)
	product.Value("image").String().Equal("http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	product.Value("tempId").Number().Equal(99)
	product.Value("sysTenancyId").Number().Equal(1)
	product.Value("sysBrandId").Number().Equal(2)
	product.Value("productCategoryId").Number().Equal(162)
	product.Value("sysTenancyName").String().Equal("宝安中心人民医院")
	product.Value("cateName").String().Equal("男士上衣")
	product.Value("brandName").String().Equal("苹果")
	product.Value("tempName").String().Equal("")
	product.Value("content").String().Equal("<p>好手机</p>")
	product.Value("sliderImage").String().Equal("http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	product.Value("sliderImages").Array().First().String().Equal("http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	product.Value("sliderImages").Array().Last().String().Equal("http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	attr := product.Value("attr").Array()
	attr.Element(0).Object().Value("detail").Array().First().String().Equal("35")
	attr.Element(0).Object().Value("value").String().Equal("S")
	attr.Element(1).Object().Value("detail").Array().First().String().Equal("36")
	attr.Element(1).Object().Value("value").String().Equal("L")
	attr.Element(2).Object().Value("detail").Array().First().String().Equal("37")
	attr.Element(2).Object().Value("value").String().Equal("XL")
	attr.Element(3).Object().Value("detail").Array().First().String().Equal("38")
	attr.Element(3).Object().Value("value").String().Equal("XXL")
	attrValue := product.Value("attrValue").Array()
	attrValue.Element(0).Object().Value("sku").String().Equal("S")
	attrValue.Element(0).Object().Value("stock").Number().Equal(99)
	attrValue.Element(0).Object().Value("sales").Number().Equal(1)
	attrValue.Element(0).Object().Value("image").String().Equal("\thttp://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	attrValue.Element(0).Object().Value("barCode").String().Equal("123456")
	attrValue.Element(0).Object().Value("cost").Number().Equal(50)
	attrValue.Element(0).Object().Value("otPrice").Number().Equal(180)
	attrValue.Element(0).Object().Value("price").Number().Equal(160)
	attrValue.Element(0).Object().Value("volume").Number().Equal(1)
	attrValue.Element(0).Object().Value("weight").Number().Equal(1)
	attrValue.Element(0).Object().Value("extensionOne").Number().Equal(0)
	attrValue.Element(0).Object().Value("extensionTwo").Number().Equal(0)
	attrValue.Element(0).Object().Value("unique").String().NotEmpty()
	attrValue.Element(0).Object().Value("detail").Object().Value("尺寸").Equal("S")
	attrValue.Element(0).Object().Value("value0").Equal("S")

	attrValue.Element(1).Object().Value("sku").String().Equal("L")
	attrValue.Element(1).Object().Value("stock").Number().Equal(100)
	attrValue.Element(1).Object().Value("sales").Number().Equal(0)
	attrValue.Element(1).Object().Value("image").String().Equal("\thttp://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	attrValue.Element(1).Object().Value("barCode").String().Equal("123456")
	attrValue.Element(1).Object().Value("cost").Number().Equal(50)
	attrValue.Element(1).Object().Value("otPrice").Number().Equal(180)
	attrValue.Element(1).Object().Value("price").Number().Equal(160)
	attrValue.Element(1).Object().Value("volume").Number().Equal(1)
	attrValue.Element(1).Object().Value("weight").Number().Equal(1)
	attrValue.Element(1).Object().Value("extensionOne").Number().Equal(0)
	attrValue.Element(1).Object().Value("extensionTwo").Number().Equal(0)
	attrValue.Element(1).Object().Value("unique").String().NotEmpty()
	attrValue.Element(1).Object().Value("detail").Object().Value("尺寸").Equal("L")
	attrValue.Element(1).Object().Value("value0").Equal("L")

	attrValue.Element(2).Object().Value("sku").String().Equal("XL")
	attrValue.Element(2).Object().Value("stock").Number().Equal(100)
	attrValue.Element(2).Object().Value("sales").Number().Equal(0)
	attrValue.Element(2).Object().Value("image").String().Equal("\thttp://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	attrValue.Element(2).Object().Value("barCode").String().Equal("123456")
	attrValue.Element(2).Object().Value("cost").Number().Equal(50)
	attrValue.Element(2).Object().Value("otPrice").Number().Equal(180)
	attrValue.Element(2).Object().Value("price").Number().Equal(160)
	attrValue.Element(2).Object().Value("volume").Number().Equal(1)
	attrValue.Element(2).Object().Value("weight").Number().Equal(1)
	attrValue.Element(2).Object().Value("extensionOne").Number().Equal(0)
	attrValue.Element(2).Object().Value("extensionTwo").Number().Equal(0)
	attrValue.Element(2).Object().Value("unique").String().NotEmpty()
	attrValue.Element(2).Object().Value("detail").Object().Value("尺寸").Equal("XL")
	attrValue.Element(2).Object().Value("value0").Equal("XL")

	attrValue.Element(3).Object().Value("sku").String().Equal("XXL")
	attrValue.Element(3).Object().Value("stock").Number().Equal(100)
	attrValue.Element(3).Object().Value("sales").Number().Equal(0)
	attrValue.Element(3).Object().Value("image").String().Equal("\thttp://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg")
	attrValue.Element(3).Object().Value("barCode").String().Equal("123456")
	attrValue.Element(3).Object().Value("cost").Number().Equal(50)
	attrValue.Element(3).Object().Value("otPrice").Number().Equal(180)
	attrValue.Element(3).Object().Value("price").Number().Equal(160)
	attrValue.Element(3).Object().Value("volume").Number().Equal(1)
	attrValue.Element(3).Object().Value("weight").Number().Equal(1)
	attrValue.Element(3).Object().Value("extensionOne").Number().Equal(0)
	attrValue.Element(3).Object().Value("extensionTwo").Number().Equal(0)
	attrValue.Element(3).Object().Value("unique").String().NotEmpty()
	attrValue.Element(3).Object().Value("detail").Object().Value("尺寸").Equal("XXL")
	attrValue.Element(3).Object().Value("value0").Equal("XXL")
	product.Value("cateId").Number().Equal(162)
	product.Value("tenancyCategoryId").Array().Element(0).Number().Equal(174)
	product.Value("tenancyCategoryId").Array().Element(1).Number().Equal(173)
	product.Value("productCates").Array().Element(0).Object().Value("id").Number().Equal(174)
	product.Value("productCates").Array().Element(0).Object().Value("cateName").String().Equal("时尚女装")
	product.Value("productCates").Array().Element(1).Object().Value("id").Number().Equal(173)
	product.Value("productCates").Array().Element(1).Object().Value("cateName").String().Equal("品牌服饰")
}
