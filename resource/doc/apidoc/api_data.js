define({ "api": [
  {
    "type": "post",
    "url": "/v1/public/device/login",
    "title": "床旁授权登录",
    "version": "0.0.1",
    "name": "床旁授权登录",
    "group": "公共方法",
    "permission": [
      {
        "name": "none"
      }
    ],
    "description": "<p>用于医院床旁设备登录使用。 医院床旁设备必须授权登录后才可以调用平台其他接口，用于确定用户在哪个医院，患者信息用于用户注册患者和更新患者信息。</p>",
    "body": [
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "uuid",
        "description": "<p>c976999e-b004-403c-96b7-e2390f64fbb7</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "name",
        "description": "<p>八两金</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "phone",
        "description": "<p>13845687419</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "sex",
        "description": "<p>性别 0 女，1男，2未知</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "age",
        "description": "<p>年龄.</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "locName",
        "description": "<p>泌尿科一区</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "bedNum",
        "description": "<p>15</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "hospitalNo",
        "description": "<p>88956655</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "disease",
        "description": "<p>不孕不育</p>"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\"status\": 200,\n\"data\": {\n    \"user\": {\n        \"patient\": {\n            \"id\": 1,\n            \"createdAt\": \"2021-07-21T17:16:22+08:00\",\n            \"updatedAt\": \"2021-07-21T17:39:16.715+08:00\",\n            \"name\": \"八两金\",\n            \"phone\": \"13845687419\",\n            \"sex\": 2,\n            \"age\": 32,\n            \"locName\": \"泌尿科一区\",\n            \"bedNum\": \"15\",\n            \"hospitalNo\": \"88956655\",\n            \"disease\": \"不孕不育\",\n            \"sysTenancyId\": 1\n        },\n        \"tenancy\": {\n            \"id\": 1,\n            \"createdAt\": \"2021-07-21T17:16:20+08:00\",\n            \"updatedAt\": \"2021-07-21T17:16:20+08:00\",\n            \"uuid\": \"c976999e-b004-403c-96b7-e2390f64fbb7\",\n            \"name\": \"宝安中心人民医院\",\n            \"tele\": \"0755-23568911\",\n            \"address\": \"xxx街道888号\",\n            \"businessTime\": \"08:30-17:30\",\n            \"status\": 1,\n            \"Keyword\": \"\",\n            \"Avatar\": \"\",\n            \"Banner\": \"\",\n            \"sales\": 0,\n            \"productScore\": 5,\n            \"serviceScore\": 5,\n            \"postageScore\": 5,\n            \"mark\": \"\",\n            \"regAdminId\": 0,\n            \"sort\": 0,\n            \"isAudit\": 2,\n            \"isBest\": 2,\n            \"isTrader\": 2,\n            \"State\": 1,\n            \"Info\": \"\",\n            \"servicePhone\": \"\",\n            \"careCount\": 0,\n            \"copyProductNum\": 0,\n            \"sysRegionCode\": 1,\n            \"region\": {\n                \"code\": 0,\n                \"pCode\": 0,\n                \"name\": \"\",\n                \"subRegions\": null\n            }\n        }\n    },\n    \"AccessToken\": \"TVRReE56YzRNVEl4TWpReE9UY3lNekkyTkEuTWpBeU1TMHdOeTB5TVZReE56b3pPVG94Tmlzd09Eb3dNQQ.MTQxNzc4MTIxMjQxOTcyMzI2NA\"\n},\n\"message\": \"登录成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/public.js",
    "groupTitle": "公共方法",
    "sampleRequest": [
      {
        "url": "http://10.0.0.26:8085/v1/public/device/login"
      }
    ]
  },
  {
    "type": "get",
    "url": "/v1/device/productCategory/getProductCategoryList",
    "title": "商品分类列表",
    "version": "0.0.1",
    "name": "商品分类列表",
    "group": "商品分类",
    "permission": [
      {
        "name": "device"
      }
    ],
    "description": "<p>获取商品分类数据</p>",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>接口需要带上此头信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example",
          "content": "\"Authorization: Bearer 5f048fe\"",
          "type": "Header"
        }
      ]
    },
    "examples": [
      {
        "title": "Curl example",
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://10.0.0.26:8085/v1/device/productCategory/getProductCategoryList",
        "type": "bash"
      }
    ],
    "error": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n\"status\": 200,\n\"data\": [\n    {\n        \"id\": 173,\n        \"createdAt\": \"2021-07-22T09:27:57+08:00\",\n        \"updatedAt\": \"2021-07-22T09:27:57+08:00\",\n        \"pid\": 0,\n        \"cateName\": \"品牌服饰\",\n        \"path\": \"/\",\n        \"sort\": 2,\n        \"pic\": \"\",\n        \"status\": 1,\n        \"level\": 0,\n        \"children\": [\n            {\n                \"id\": 174,\n                \"createdAt\": \"2021-07-22T09:27:57+08:00\",\n                \"updatedAt\": \"2021-07-22T09:27:57+08:00\",\n                \"pid\": 173,\n                \"cateName\": \"时尚女装\",\n                \"path\": \"/173/\",\n                \"sort\": 0,\n                \"pic\": \"\",\n                \"status\": 1,\n                \"level\": 1,\n                \"children\": null\n            }\n        ]\n    }\n],\n\"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/product_category.js",
    "groupTitle": "商品分类",
    "sampleRequest": [
      {
        "url": "http://10.0.0.26:8085/v1/device/productCategory/getProductCategoryList"
      }
    ]
  },
  {
    "type": "post",
    "url": "/v1/device/product/getProductList",
    "title": "商品列表",
    "version": "0.0.1",
    "name": "商品列表",
    "group": "商品",
    "permission": [
      {
        "name": "device"
      }
    ],
    "description": "<p>获取商品数据</p>",
    "body": [
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "page",
        "description": "<p>页码</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "pageSize",
        "description": "<p>每页数量</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "cateId",
        "description": "<p>商户分类 id</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "tenancyCategoryId",
        "description": "<p>商城分类 id</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "keyword",
        "description": "<p>关键字搜索</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "type",
        "description": "<p>商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": true,
        "field": "isGiftBag",
        "description": "<p>是否礼包 1 是 ，2 否</p>"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": "<p>接口需要带上此头信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example",
          "content": "\"Authorization: Bearer 5f048fe\"",
          "type": "Header"
        }
      ]
    },
    "examples": [
      {
        "title": "Curl example",
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://10.0.0.26:8085/v1/device/product/getProductList",
        "type": "bash"
      }
    ],
    "error": {
      "examples": [
        {
          "title": "Response:",
          "content": "  HTTP/1.1 200 OK\n  {\n\"status\": 200,\n  \"data\": {\n      \"list\": [\n          {\n              \"id\": 1,\n              \"createdAt\": \"2021-07-22T09:27:57+08:00\",\n              \"updatedAt\": \"2021-07-22T09:27:57+08:00\",\n              \"storeName\": \"领立裁腰带短袖连衣裙\",\n              \"storeInfo\": \"短袖连衣裙\",\n              \"keyword\": \"连衣裙\",\n              \"barCode\": \"\",\n              \"isShow\": 1,\n              \"status\": 1,\n              \"unitName\": \"件\",\n              \"sort\": 40,\n              \"rank\": 0,\n              \"sales\": 1,\n              \"price\": 80,\n              \"cost\": 50,\n              \"otPrice\": 100,\n              \"stock\": 399,\n              \"isHot\": 2,\n              \"isBenefit\": 2,\n              \"isBest\": 2,\n              \"isNew\": 2,\n              \"isGood\": 1,\n              \"productType\": 2,\n              \"ficti\": 100,\n              \"browse\": 0,\n              \"codePath\": \"\",\n              \"videoLink\": \"\",\n              \"specType\": 1,\n              \"extensionType\": 2,\n              \"refusal\": \"\",\n              \"rate\": 5,\n              \"replyCount\": 0,\n              \"isGiftBag\": 2,\n              \"careCount\": 0,\n              \"image\": \"http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n              \"oldId\": 0,\n              \"tempId\": 99,\n              \"sysTenancyId\": 1,\n              \"sysBrandId\": 2,\n              \"productCategoryId\": 162,\n              \"sysTenancyName\": \"宝安中心人民医院\",\n              \"cateName\": \"男士上衣\",\n              \"brandName\": \"苹果\",\n              \"productCates\": [\n                  {\n                      \"id\": 174,\n                      \"cateName\": \"时尚女装\"\n                  },\n                  {\n                      \"id\": 173,\n                      \"cateName\": \"品牌服饰\"\n                  }\n              ]\n          }\n      ],\n      \"total\": 1,\n      \"page\": 1,\n      \"pageSize\": 10\n  },\n  \"message\": \"获取成功\"\n  }",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/product.js",
    "groupTitle": "商品",
    "sampleRequest": [
      {
        "url": "http://10.0.0.26:8085/v1/device/product/getProductList"
      }
    ]
  }
] });
