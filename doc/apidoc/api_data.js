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
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "uuid",
            "description": "<p>c976999e-b004-403c-96b7-e2390f64fbb7</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>八两金</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "phone",
            "description": "<p>13845687419</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "sex",
            "description": "<p>性别 0 女，1男，2未知</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "age",
            "description": "<p>年龄.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "locName",
            "description": "<p>泌尿科一区</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "bedNum",
            "description": "<p>15</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "hospitalNo",
            "description": "<p>88956655</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "disease",
            "description": "<p>不孕不育</p>"
          }
        ]
      }
    },
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
        "url": "http://127.0.0.1:8089/v1/public/device/login"
      }
    ]
  },
  {
    "type": "get",
    "url": "/v1/auth/logout",
    "title": "退出登录",
    "version": "0.0.1",
    "name": "退出登录",
    "group": "公共方法",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>退出当前登录用户</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/auth/logout",
        "type": "bash"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n     \"status\": 200,\n     \"data\": {},\n     \"message\": \"退出登录\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/public.js",
    "groupTitle": "公共方法",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/auth/logout"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/device/productCategory/getProductCategoryList",
    "title": "商品分类列表",
    "version": "0.0.1",
    "name": "商品分类列表",
    "group": "商品分类管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/productCategory/getProductCategoryList",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>分类id</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "pid",
            "description": "<p>上级id</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "cateName",
            "description": "<p>分类名称</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sort",
            "description": "<p>排序</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pic",
            "description": "<p>图片</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "level",
            "description": "<p>等级</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "children",
            "description": "<p>子分类</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n\"status\": 200,\n\"data\": [\n    {\n        \"id\": 173,\n        \"cateName\": \"品牌服饰\",\n        \"pic\": \"\",\n        \"children\": [\n            {\n                \"id\": 174,\n                \"pid\": 173,\n                \"cateName\": \"时尚女装\",\n                \"pic\": \"\"\n            }\n        ]\n    }\n],\n\"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/product_category.js",
    "groupTitle": "商品分类管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/productCategory/getProductCategoryList"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/device/product/getProductList",
    "title": "商品列表",
    "version": "0.0.1",
    "name": "商品列表",
    "group": "商品管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>获取商品列表数据</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "page",
            "description": "<p>页码</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "pageSize",
            "description": "<p>每页数量</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "cateId",
            "description": "<p>商户分类 id</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "tenancyCategoryId",
            "description": "<p>商城分类 id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "keyword",
            "description": "<p>关键字搜索</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "type",
            "description": "<p>商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": true,
            "field": "isGiftBag",
            "description": "<p>是否礼包 1 是 ，2 否</p>"
          }
        ]
      }
    },
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/product/getProductList",
        "type": "bash"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n\"status\": 200,\n\"data\": {\n    \"list\": [\n        {\n            \"id\": 1,\n            \"storeName\": \"领立裁腰带短袖连衣裙\",\n            \"sales\": 1,\n            \"price\": 80,\n            \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\"\n        }\n    ],\n    \"total\": 1,\n    \"page\": 1,\n    \"pageSize\": 10\n},\n\"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/product.js",
    "groupTitle": "商品管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/product/getProductList"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/device/product/getProductById/1",
    "title": "商品详情",
    "version": "0.0.1",
    "name": "商品详情",
    "group": "商品管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>获取商品详情数据</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/product/getProductById/1",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>商品id</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "storeName",
            "description": "<p>商品名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "storeInfo",
            "description": "<p>商品简介</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "keyword",
            "description": "<p>商品关键词</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "barCode",
            "description": "<p>商品条码</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isShow",
            "description": "<p>是否显示 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>商品状态 1：审核通过,2：审核中 3: 未通过</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "unitName",
            "description": "<p>商品单位</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sort",
            "description": "<p>排序</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "rank",
            "description": "<p>总后台排序</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sales",
            "description": "<p>商品销量</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "price",
            "description": "<p>最低价格</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "cost",
            "description": "<p>成本价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "otPrice",
            "description": "<p>原价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "stock",
            "description": "<p>库存</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isHot",
            "description": "<p>是否热卖 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isBenefit",
            "description": "<p>促销单品 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isBest",
            "description": "<p>是否精品 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isNew",
            "description": "<p>是否新品 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isGood",
            "description": "<p>是否优品推荐 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "productType",
            "description": "<p>商品分类 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "ficti",
            "description": "<p>虚拟销量</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "browse",
            "description": "<p>浏览量</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "codePath",
            "description": "<p>产品二维码地址</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "videoLink",
            "description": "<p>主图视频链接</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "specType",
            "description": "<p>规格 1单 2多</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "refusal",
            "description": "<p>审核拒绝理由</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "Rate",
            "description": "<p>评价分数</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "ReplyCount",
            "description": "<p>评论数</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isGiftBag",
            "description": "<p>是否为礼包</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "careCount",
            "description": "<p>收藏数</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "image",
            "description": "<p>商品图片</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "sysTenancyName",
            "description": "<p>医院名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "cateName",
            "description": "<p>后台分类名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "brandName",
            "description": "<p>品牌名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "tempName",
            "description": "<p>模板名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "content",
            "description": "<p>详情内容</p>"
          },
          {
            "group": "Success 200",
            "type": "String[]",
            "optional": false,
            "field": "sliderImages",
            "description": "<p>展示图片</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "attrValue",
            "description": "<p>规格详情</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "productCates",
            "description": "<p>商品分类</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "  HTTP/1.1 200 OK\n  {\n\"status\": 200,\n  \"data\": {\n      \"id\": 1,\n      \"storeName\": \"领立裁腰带短袖连衣裙\",\n      \"storeInfo\": \"短袖连衣裙\",\n      \"keyword\": \"连衣裙\",\n      \"unitName\": \"件\",\n      \"sort\": 40,\n      \"sales\": 1,\n      \"price\": 80,\n      \"otPrice\": 100,\n      \"stock\": 399,\n      \"isHot\": 2,\n      \"isBenefit\": 2,\n      \"isBest\": 2,\n      \"isNew\": 2,\n      \"isGood\": 1,\n      \"productType\": 2,\n      \"ficti\": 100,\n      \"specType\": 1,\n      \"rate\": 5,\n      \"isGiftBag\": 2,\n      \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n      \"tempId\": 99,\n      \"sysTenancyId\": 1,\n      \"sysBrandId\": 2,\n      \"productCategoryId\": 162,\n      \"sysTenancyName\": \"宝安中心人民医院\",\n      \"cateName\": \"男士上衣\",\n      \"brandName\": \"苹果\",\n      \"tempName\": \"\",\n      \"content\": \"<p>好手机</p>\",\n      \"sliderImage\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n      \"sliderImages\": [\n          \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n          \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\"\n      ],\n      \"attrValue\": [\n          {\n              \"sku\": \"S\",\n              \"stock\": 99,\n              \"sales\": 1,\n              \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n              \"barCode\": \"123456\",\n              \"cost\": 50,\n              \"otPrice\": 180,\n              \"price\": 160,\n              \"volume\": 1,\n              \"weight\": 1,\n              \"extensionOne\": 0,\n              \"extensionTwo\": 0,\n              \"unique\": \"e2fe28308fd0\",\n              \"detail\": {\n                  \"尺寸\": \"S\"\n              },\n              \"value0\": \"S\"\n          },\n          {\n              \"sku\": \"L\",\n              \"stock\": 100,\n              \"sales\": 0,\n              \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n              \"barCode\": \"123456\",\n              \"cost\": 50,\n              \"otPrice\": 180,\n              \"price\": 160,\n              \"volume\": 1,\n              \"weight\": 1,\n              \"extensionOne\": 0,\n              \"extensionTwo\": 0,\n              \"unique\": \"e2fe28308fd0\",\n              \"detail\": {\n                  \"尺寸\": \"L\"\n              },\n              \"value0\": \"L\"\n          },\n          {\n              \"sku\": \"XL\",\n              \"stock\": 100,\n              \"sales\": 0,\n              \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n              \"barCode\": \"123456\",\n              \"cost\": 50,\n              \"otPrice\": 180,\n              \"price\": 160,\n              \"volume\": 1,\n              \"weight\": 1,\n              \"extensionOne\": 0,\n              \"extensionTwo\": 0,\n              \"unique\": \"e2fe28308fd0\",\n              \"detail\": {\n                  \"尺寸\": \"XL\"\n              },\n              \"value0\": \"XL\"\n          },\n          {\n              \"sku\": \"XXL\",\n              \"stock\": 100,\n              \"sales\": 0,\n              \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n              \"barCode\": \"123456\",\n              \"cost\": 50,\n              \"otPrice\": 180,\n              \"price\": 160,\n              \"volume\": 1,\n              \"weight\": 1,\n              \"extensionOne\": 0,\n              \"extensionTwo\": 0,\n              \"unique\": \"e2fe28308fd0\",\n              \"detail\": {\n                  \"尺寸\": \"XXL\"\n              },\n              \"value0\": \"XXL\"\n          }\n      ],\n      \"cateId\": 162,\n      \"tenancyCategoryId\": [\n          174,\n          173\n      ],\n      \"productCates\": [\n          {\n              \"id\": 174,\n              \"cateName\": \"时尚女装\"\n          },\n          {\n              \"id\": 173,\n              \"cateName\": \"品牌服饰\"\n          }\n      ]\n  },\n  \"message\": \"操作成功\"\n  }",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/product.js",
    "groupTitle": "商品管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/product/getProductById/1"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/device/patient/getPatientList",
    "title": "患者列表",
    "version": "0.0.1",
    "name": "患者列表",
    "group": "患者管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>获取医院患者数据</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/patient/getPatientList",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>患者名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "phone",
            "description": "<p>手机</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sex",
            "description": "<p>性别</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "age",
            "description": "<p>年龄</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "locName",
            "description": "<p>科室</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "bedNum",
            "description": "<p>床号</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "hospitalNo",
            "description": "<p>住院号</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "disease",
            "description": "<p>病种</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "hospitalName",
            "description": "<p>医院</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {\n        \"list\": [\n            {\n                \"id\": 1,\n                \"createdAt\": \"2021-07-26T12:24:42+08:00\",\n                \"updatedAt\": \"2021-07-26T12:24:42+08:00\",\n                \"name\": \"八两金\",\n                \"phone\": \"13845687419\",\n                \"sex\": 2,\n                \"age\": 32,\n                \"locName\": \"泌尿科一区\",\n                \"bedNum\": \"15\",\n                \"hospitalNo\": \"88956655\",\n                \"disease\": \"不孕不育\",\n                \"sysTenancyId\": 1,\n                \"hospitalName\": \"宝安中心人民医院\"\n            }\n        ],\n        \"total\": 1,\n        \"page\": 1,\n        \"pageSize\": -1\n    },\n    \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/patient.js",
    "groupTitle": "患者管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/patient/getPatientList"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/device/patient/getPatientDetail",
    "title": "获取当前患者",
    "version": "0.0.1",
    "name": "获取当前患者",
    "group": "患者管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>获取当前床旁设备患者</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/patient/getPatientDetail",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>患者名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "phone",
            "description": "<p>手机</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sex",
            "description": "<p>性别</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "age",
            "description": "<p>年龄</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "locName",
            "description": "<p>科室</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "bedNum",
            "description": "<p>床号</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "hospitalNo",
            "description": "<p>住院号</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "disease",
            "description": "<p>病种</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "hospitalName",
            "description": "<p>医院</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {\n        \"id\": 1,\n        \"createdAt\": \"2021-07-26T12:24:42+08:00\",\n        \"updatedAt\": \"2021-07-26T17:28:06+08:00\",\n        \"name\": \"八两金\",\n        \"phone\": \"13845687419\",\n        \"sex\": 2,\n        \"age\": 32,\n        \"locName\": \"泌尿科一区\",\n        \"bedNum\": \"15\",\n        \"hospitalNo\": \"88956655\",\n        \"disease\": \"不孕不育\",\n        \"sysTenancyId\": 1\n    },\n    \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/patient.js",
    "groupTitle": "患者管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/patient/getPatientDetail"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/v1/device/order/cancelOrder/1",
    "title": "取消订单",
    "version": "0.0.1",
    "name": "取消订单",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>用户取消未支付的订单，其他订单无法取消</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/cancelOrder/1",
        "type": "bash"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {},\n    \"message\": \"操作成功\" \n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/cancelOrder/1"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/device/order/getOrderList",
    "title": "我的订单",
    "version": "0.0.1",
    "name": "我的订单",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>床旁用户的订单列表</p>",
    "body": [
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "pageSize",
        "description": "<p>页数</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "page",
        "description": "<p>页码</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "status",
        "description": "<p>订单状态0：待付款 1:待发货 2：待收货 3：待评价 4：已完成 5：已退款 6：已取消</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "date",
        "description": "<p>日期：today，yesterday，lately7，lately30，month，year或者日期范围:2021/08/01-2021/08/05</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "orderType",
        "description": "<p>订单类型</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/getOrderList",
        "type": "bash"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 200,\n      \"data\": {\n          \"list\": [\n              {\n                  \"id\": 7,\n                  \"createdAt\": \"2021-08-09T16:53:49+08:00\",\n                  \"updatedAt\": \"2021-08-09T17:09:00+08:00\",\n                  \"orderSn\": \"12021080916534924655141100851200\",\n                  \"realName\": \"八两金\",\n                  \"userPhone\": \"13845687419\",\n                  \"userAddress\": \"宝安中心人民医院-泌尿科一区-15床\",\n                  \"totalNum\": 20,\n                  \"totalPrice\": 3200,\n                  \"totalPostage\": 0,\n                  \"payPrice\": 3200,\n                  \"payPostage\": 0,\n                  \"commissionRate\": 0,\n                  \"orderType\": 1,\n                  \"paid\": 2,\n                  \"payTime\": \"0001-01-01T00:00:00Z\",\n                  \"payType\": 0,\n                  \"status\": 6,\n                  \"deliveryType\": 0,\n                  \"deliveryName\": \"\",\n                  \"deliveryId\": \"\",\n                  \"mark\": \"\",\n                  \"remark\": \"remark\",\n                  \"adminMark\": \"\",\n                  \"verifyCode\": \"\",\n                  \"verifyTime\": \"0001-01-01T00:00:00Z\",\n                  \"activityType\": 1,\n                  \"cost\": 1000,\n                  \"isDel\": 2,\n                  \"isSystemDel\": 2,\n                  \"groupOrderSn\": \"G2021080916534924655141088268288\",\n                  \"tenancyName\": \"宝安中心人民医院\",\n                  \"isTrader\": 2,\n                  \"sysUserId\": 1,\n                  \"sysTenancyId\": 1,\n                  \"groupOrderId\": 6,\n                  \"reconciliationId\": 0,\n                  \"orderProduct\": [\n                      {\n                          \"id\": 7,\n                          \"cartInfo\": {\n                              \"product\": {\n                                  \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n                                  \"storeName\": \"领立裁腰带短袖连衣裙\",\n                                  \"temp\": {\n                                      \"name\": \"\",\n                                      \"type\": 0,\n                                      \"appoint\": 0,\n                                      \"undelivery\": 0,\n                                      \"isDefault\": 0,\n                                      \"sort\": 0\n                                  }\n                              },\n                              \"productAttr\": {\n                                  \"price\": 160,\n                                  \"sku\": \"S\"\n                              }\n                          },\n                          \"productSku\": \"S\",\n                          \"isRefund\": 0,\n                          \"productNum\": 20,\n                          \"productType\": 1,\n                          \"refundNum\": 0,\n                          \"isReply\": 2,\n                          \"productPrice\": 160,\n                          \"orderID\": 7,\n                          \"productId\": 1\n                      }\n                  ]\n              }\n          ],\n          \"page\": 1,\n          \"pageSize\": 20,\n          \"stat\": [\n              {\n                  \"className\": \"el-icon-s-goods\",\n                  \"count\": 5,\n                  \"field\": \"件\",\n                  \"name\": \"已支付订单数量\"\n              },\n              {\n                  \"className\": \"el-icon-s-order\",\n                  \"count\": 673,\n                  \"field\": \"元\",\n                  \"name\": \"实际支付金额\"\n              },\n              {\n                  \"className\": \"el-icon-s-cooperation\",\n                  \"count\": 0,\n                  \"field\": \"元\",\n                  \"name\": \"已退款金额\"\n              },\n              {\n                  \"className\": \"el-icon-s-cooperation\",\n                  \"count\": 673,\n                  \"field\": \"元\",\n                  \"name\": \"微信支付金额\"\n              },\n              {\n                  \"className\": \"el-icon-s-finance\",\n                  \"count\": 0,\n                  \"field\": \"元\",\n                  \"name\": \"余额支付金额\"\n              },\n              {\n                  \"className\": \"el-icon-s-cooperation\",\n                  \"count\": 0,\n                  \"field\": \"元\",\n                  \"name\": \"支付宝支付金额\"\n              }\n          ],\n          \"total\": 12\n      },\n      \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/getOrderList"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/v1/device/order/refundOrder/1",
    "title": "提交退款",
    "version": "0.0.1",
    "name": "提交退款",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>用户支付的订单，用户提交退款申请</p>",
    "body": [
      {
        "group": "Body",
        "type": "Array",
        "optional": false,
        "field": "ids",
        "description": "<p>订单商品id</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "refundMessage",
        "description": "<p>退款原因</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "RefundPrice",
        "description": "<p>退款金额</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "RefundType",
        "description": "<p>退款类型 1：退款，2：退款退货</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "Num",
        "description": "<p>退款数量</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": true,
        "field": "Mark",
        "description": "<p>备注</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/refundOrder/1",
        "type": "bash"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {\n      \"id\":1,\n    },\n    \"message\": \"操作成功\" \n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/refundOrder/1"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/v1/device/order/getOrderById/1",
    "title": "根据id获取订单详情",
    "version": "0.0.1",
    "name": "根据id获取订单详情",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>根据id获取订单详情</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/getOrderById/1",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>订单id</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "createdAt",
            "description": "<p>创建时间</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "updatedAt",
            "description": "<p>更新时间</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "orderSn",
            "description": "<p>订单号</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "realName",
            "description": "<p>用户姓名</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userPhone",
            "description": "<p>用户电话</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userAddress",
            "description": "<p>用户地址</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "totalNum",
            "description": "<p>订单商品数量</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "totalPrice",
            "description": "<p>订单商品总价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "totalPostage",
            "description": "<p>订单邮费</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "payPrice",
            "description": "<p>订单支付总价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "payPostage",
            "description": "<p>订单支付邮费</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "commissionRate",
            "description": "<p>平台手续费</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "orderType",
            "description": "<p>订单类型</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "paid",
            "description": "<p>支付状态 1支付，2未支付</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "payTime",
            "description": "<p>支付时间</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "payType",
            "description": "<p>支付方式  1=微信 2=小程序 3=h5 4=余额 5=支付宝</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "status",
            "description": "<p>订单状态（0：待付款 1:待发货 2：待收货 3：待评价 4：已完成 5：已退款 6：已取消）</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "deliveryType",
            "description": "<p>发货类型(0:未发货 1:发货 2: 送货 3: 虚拟)</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "deliveryName",
            "description": "<p>快递名称/送货人姓名</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "deliveryId",
            "description": "<p>快递单号/手机号</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "mark",
            "description": "<p>备注</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "remark",
            "description": "<p>商户备注备注</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "adminMark",
            "description": "<p>后台备注</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "verifyCode",
            "description": "<p>核销码</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "verifyTime",
            "description": "<p>核销时间</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "activityType",
            "description": "<p>活动类型 1：普通 2:秒杀 3:预售 4:助力</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "cost",
            "description": "<p>成本价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isDel",
            "description": "<p>是否删除</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isSystemDel",
            "description": "<p>后台是否删除</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sysUserId",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sysTenancyId",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "groupOrderId",
            "description": ""
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "reconciliationId",
            "description": "<p>对账id</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "userNickName",
            "description": "<p>用户昵称</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n\"status\": 200,\n  \"data\": {\n      \"id\": 1,\n      \"createdAt\": \"2021-08-04T17:26:23+08:00\",\n      \"updatedAt\": \"2021-08-04T17:26:23+08:00\",\n      \"orderSn\": \"1202108041726161422851368560889861\",\n      \"realName\": \"real_name\",\n      \"userPhone\": \"user_phone\",\n      \"userAddress\": \"user_address\",\n      \"totalNum\": 10,\n      \"totalPrice\": 20,\n      \"totalPostage\": 30,\n      \"payPrice\": 50,\n      \"payPostage\": 30,\n      \"commissionRate\": 15,\n      \"orderType\": 1,\n      \"paid\": 1,\n      \"payTime\": \"2021-08-04T17:26:16+08:00\",\n      \"payType\": 1,\n      \"status\": 5,\n      \"deliveryType\": 1,\n      \"deliveryName\": \"delivery_name\",\n      \"deliveryId\": \"delivery_id\",\n      \"mark\": \"mark\",\n      \"remark\": \"remark\",\n      \"adminMark\": \"admin_mark\",\n      \"verifyCode\": \"\",\n      \"verifyTime\": \"2021-08-04T17:26:16+08:00\",\n      \"activityType\": 1,\n      \"cost\": 5,\n      \"isDel\": 2,\n      \"isSystemDel\": 2,\n      \"sysUserId\": 7,\n      \"sysTenancyId\": 1,\n      \"groupOrderId\": 1,\n      \"reconciliationId\": 0,\n      \"userNickName\": \"C端用户\"\n  },\n  \"message\": \"操作成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/getOrderById/1"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/device/order/checkOrder",
    "title": "用户下单数据",
    "version": "0.0.1",
    "name": "用户下单数据",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>获取用户下单数据详情</p>",
    "body": [
      {
        "group": "Body",
        "type": "Number[]",
        "optional": false,
        "field": "ids",
        "description": "<p>购物车ids</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/checkOrder",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "totalPrice",
            "description": "<p>商品总价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "totalOtPrice",
            "description": "<p>商品原价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "orderPrice",
            "description": "<p>订单总价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "orderOtPrice",
            "description": "<p>订单原价</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "postagePrice",
            "description": "<p>订单邮费</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "downPrice",
            "description": "<p>订单优惠价格</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "totalNum",
            "description": "<p>商品总数</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "orderType",
            "description": "<p>订单类型 1：普通，2：自提</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "productPrices",
            "description": "<p>商品价格</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "products",
            "description": "<p>商品信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {\n        \"sysTenancyId\": 1,\n        \"name\": \"宝安中心人民医院\",\n        \"Avatar\": \"\",\n        \"products\": [\n            {\n                \"id\": 5,\n                \"sysTenancyId\": 1,\n                \"specType\": 2,\n                \"productId\": 1,\n                \"storeName\": \"领立裁腰带短袖连衣裙\",\n                \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n                \"cartNum\": 2,\n                \"isFail\": 2,\n                \"productAttrUnique\": \"e2fe28308fd2\",\n                \"attrValue\": {\n                    \"productId\": 0,\n                    \"sku\": \"S\",\n                    \"stock\": 99,\n                    \"sales\": 1,\n                    \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n                    \"barCode\": \"123456\",\n                    \"cost\": 50,\n                    \"otPrice\": 180,\n                    \"price\": 160,\n                    \"volume\": 1,\n                    \"weight\": 1,\n                    \"extensionOne\": 0,\n                    \"extensionTwo\": 0,\n                    \"unique\": \"e2fe28308fd2\",\n                    \"detail\": {\n                        \"尺寸\": \"S\"\n                    },\n                    \"value0\": \"S\"\n                }\n            }\n        ],\n        \"totalPrice\": \"320\",\n        \"totalOtPrice\": \"360\",\n        \"orderPrice\": \"320\",\n        \"orderOtPrice\": \"360\",\n        \"postagePrice\": \"0\",\n        \"downPrice\": \"0\",\n        \"productPrices\": {\n            \"1\": {\n                \"otPrice\": \"360\",\n                \"price\": \"320\"\n            }\n        },\n        \"totalNum\": 2,\n        \"orderType\": 2\n    },\n    \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/checkOrder"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/device/order/createOrder",
    "title": "用户结算订单",
    "version": "0.0.1",
    "name": "用户结算订单",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>用户结算订单，并生成待支付订单和支付二维码，用户通过支付宝或者微信扫码支付</p>",
    "body": [
      {
        "group": "Body",
        "type": "Number[]",
        "optional": false,
        "field": "cartIds",
        "description": "<p>购物车ids</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "orderType",
        "description": "<p>订单类型：1：普通，2：自提</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": true,
        "field": "remark",
        "description": "<p>订单备注</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/createOrder",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "qrcode",
            "description": "<p>扫码支付二维码 base64 数据，需要在前面加上 data:image/png;base64, 才能显示为图片；例如data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACWklEQVR42uyYMbLjIBBEW0VAyBF0FF3sl5BqL8ZROAIhAaXe6sH+9l9vbtg1gUuFnoMR0zM94LM+679cC0leno4nEzMAn4FNu2kioALYfXJlqXDa9nkpG4AwFbCy7J4s8MR6BzzbcIA+/uVTQywgD8wJKMwdW1tPAWeZD1BGYbNkapIJ8JeUezPQxZtaOOvm8h5SXc/yqu6xgdvaHC9sWC9sPsfyWk/HBpYKhUSEw7P1CE9JPD4K6QDA6ngogS6l/RHoeQFWkyYCmJvedQXnXWdhPw9dTAAAvXI2xKowgSqxW2OYB7iH6cpCCtiqalTFn9/hrUBdVX3geHg6cwarZdTxLd4JAPR2hQb4xG5y9sD6OKwZgIXUOaUWyNR1jkXtt4WBgG5yYD2rlyC54oqvyQDrWXYqN8nu2vl6pNz4gA0k3ar5BKnbdFHX415IBwB0FtKtnrpB0z9UO+cC0A2xnpp26nqp/bYnkzM+wGy2p4WThN7ZYTF/DQUQgQog3rx9RSzbU0bNAABwMmWOJ8mMQCDq6ckajw8s2tTpFEjddljUnOUeJuf9QF2tp7o+N9kwqOE7PxXSCQB13t0T4ZI90Nydo8KMz/ZgeEBWzVyxpVVvBzI5LUwFADJAt54AbMw2rI8E3NpvQk+rs1+aWT/GPMD9lpUlKlYWhVk2z+PnTc7gQL9llTjMr8lwRoXpvif3EYB+s2cKNldcrdrz14yATiVZJ+iN6zFGzQPALqVSg72TOH440vcDdssKK5a0CWq1p8iZgLt49YPvO5CXu+Kxgc/6rH9s/Q4AAP//PZzXiSo38ugAAAAASUVORK5CYII=</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"data\": {\n      \"qrcode\": \"iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACWklEQVR42uyYMbLjIBBEW0VAyBF0FF3sl5BqL8ZROAIhAaXe6sH+9l9vbtg1gUuFnoMR0zM94LM+679cC0leno4nEzMAn4FNu2kioALYfXJlqXDa9nkpG4AwFbCy7J4s8MR6BzzbcIA+/uVTQywgD8wJKMwdW1tPAWeZD1BGYbNkapIJ8JeUezPQxZtaOOvm8h5SXc/yqu6xgdvaHC9sWC9sPsfyWk/HBpYKhUSEw7P1CE9JPD4K6QDA6ngogS6l/RHoeQFWkyYCmJvedQXnXWdhPw9dTAAAvXI2xKowgSqxW2OYB7iH6cpCCtiqalTFn9/hrUBdVX3geHg6cwarZdTxLd4JAPR2hQb4xG5y9sD6OKwZgIXUOaUWyNR1jkXtt4WBgG5yYD2rlyC54oqvyQDrWXYqN8nu2vl6pNz4gA0k3ar5BKnbdFHX415IBwB0FtKtnrpB0z9UO+cC0A2xnpp26nqp/bYnkzM+wGy2p4WThN7ZYTF/DQUQgQog3rx9RSzbU0bNAABwMmWOJ8mMQCDq6ckajw8s2tTpFEjddljUnOUeJuf9QF2tp7o+N9kwqOE7PxXSCQB13t0T4ZI90Nydo8KMz/ZgeEBWzVyxpVVvBzI5LUwFADJAt54AbMw2rI8E3NpvQk+rs1+aWT/GPMD9lpUlKlYWhVk2z+PnTc7gQL9llTjMr8lwRoXpvif3EYB+s2cKNldcrdrz14yATiVZJ+iN6zFGzQPALqVSg72TOH440vcDdssKK5a0CWq1p8iZgLt49YPvO5CXu+Kxgc/6rH9s/Q4AAP//PZzXiSo38ugAAAAASUVORK5CYII=\"\n  },\n  \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/createOrder"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/v1/device/order/checkRefundOrder/1",
    "title": "申请退款",
    "version": "0.0.1",
    "name": "申请退款",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>用户支付的订单，可以申请退款</p>",
    "body": [
      {
        "group": "Body",
        "type": "Array",
        "optional": false,
        "field": "ids",
        "description": "<p>订单商品集合id</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/checkRefundOrder/1",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "totalRefundPrice",
            "description": "<p>退款金额</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "postagePrice",
            "description": "<p>退回邮费</p>"
          },
          {
            "group": "Success 200",
            "type": "Object",
            "optional": false,
            "field": "product",
            "description": "<p>退款商品</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "    HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"data\": {\n    \"totalRefundPrice\": 2,\n    \"postagePrice\": 0,\n    \"product\": [\n      {\n        \"id\": 1,\n        \"cartInfo\": {\n          \"product\": {\n            \"image\": \"http://127.0.0.1:8089/uploads/file/b39024efbc6de61976f585c8421c6bba_20210702150027.png\",\n            \"storeName\": \"是防守打法发\",\n            \"temp\": {\n              \"name\": \"\",\n              \"type\": 0,\n              \"appoint\": 0,\n              \"undelivery\": 0,\n              \"isDefault\": 0,\n              \"sort\": 0\n            }\n          },\n          \"productAttr\": {\n            \"price\": 1,\n            \"sku\": \"S\"\n          }\n        },\n        \"productSku\": \"S\",\n        \"isRefund\": 0,\n        \"productNum\": 2,\n        \"productType\": 1,\n        \"refundNum\": 2,\n        \"isReply\": 2,\n        \"productPrice\": 1,\n        \"orderID\": 1,\n        \"productId\": 1\n      }\n    ],\n    \"status\": 1\n  },\n  \"message\": \"操作成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/checkRefundOrder/1"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "GET",
    "url": "/v1/device/order/payOrder/1?orderType=1",
    "title": "重新获取支付二维码",
    "version": "0.0.1",
    "name": "重新获取支付二维码",
    "group": "订单管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>重新获取支付二维码，用户通过支付宝或者微信扫码支付 扫码支付后，后台会通过 mqtt 发送主题为 tenancy_notify_pay 的消息体: { &quot;orderId&quot;: 2, // 订单 &quot;tenancyId&quot;: 2, // 商户 &quot;userId&quot;: 2, //用户 &quot;orderType&quot;: 2, // 订单类型 &quot;payType&quot;: 2, // 支付类型 &quot;createdAt&quot;: 2, //回调时间 }</p>",
    "query": [
      {
        "group": "Query",
        "type": "Number",
        "optional": false,
        "field": "orderType",
        "description": "<p>订单类型</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/order/payOrder/1?orderType=1",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "qrcode",
            "description": "<p>扫码支付二维码 base64 数据，需要在前面加上 data:image/png;base64, 才能显示为图片；例如data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACWklEQVR42uyYMbLjIBBEW0VAyBF0FF3sl5BqL8ZROAIhAaXe6sH+9l9vbtg1gUuFnoMR0zM94LM+679cC0leno4nEzMAn4FNu2kioALYfXJlqXDa9nkpG4AwFbCy7J4s8MR6BzzbcIA+/uVTQywgD8wJKMwdW1tPAWeZD1BGYbNkapIJ8JeUezPQxZtaOOvm8h5SXc/yqu6xgdvaHC9sWC9sPsfyWk/HBpYKhUSEw7P1CE9JPD4K6QDA6ngogS6l/RHoeQFWkyYCmJvedQXnXWdhPw9dTAAAvXI2xKowgSqxW2OYB7iH6cpCCtiqalTFn9/hrUBdVX3geHg6cwarZdTxLd4JAPR2hQb4xG5y9sD6OKwZgIXUOaUWyNR1jkXtt4WBgG5yYD2rlyC54oqvyQDrWXYqN8nu2vl6pNz4gA0k3ar5BKnbdFHX415IBwB0FtKtnrpB0z9UO+cC0A2xnpp26nqp/bYnkzM+wGy2p4WThN7ZYTF/DQUQgQog3rx9RSzbU0bNAABwMmWOJ8mMQCDq6ckajw8s2tTpFEjddljUnOUeJuf9QF2tp7o+N9kwqOE7PxXSCQB13t0T4ZI90Nydo8KMz/ZgeEBWzVyxpVVvBzI5LUwFADJAt54AbMw2rI8E3NpvQk+rs1+aWT/GPMD9lpUlKlYWhVk2z+PnTc7gQL9llTjMr8lwRoXpvif3EYB+s2cKNldcrdrz14yATiVZJ+iN6zFGzQPALqVSg72TOH440vcDdssKK5a0CWq1p8iZgLt49YPvO5CXu+Kxgc/6rH9s/Q4AAP//PZzXiSo38ugAAAAASUVORK5CYII=</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"data\": {\n      \"qrcode\": \"iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACWklEQVR42uyYMbLjIBBEW0VAyBF0FF3sl5BqL8ZROAIhAaXe6sH+9l9vbtg1gUuFnoMR0zM94LM+679cC0leno4nEzMAn4FNu2kioALYfXJlqXDa9nkpG4AwFbCy7J4s8MR6BzzbcIA+/uVTQywgD8wJKMwdW1tPAWeZD1BGYbNkapIJ8JeUezPQxZtaOOvm8h5SXc/yqu6xgdvaHC9sWC9sPsfyWk/HBpYKhUSEw7P1CE9JPD4K6QDA6ngogS6l/RHoeQFWkyYCmJvedQXnXWdhPw9dTAAAvXI2xKowgSqxW2OYB7iH6cpCCtiqalTFn9/hrUBdVX3geHg6cwarZdTxLd4JAPR2hQb4xG5y9sD6OKwZgIXUOaUWyNR1jkXtt4WBgG5yYD2rlyC54oqvyQDrWXYqN8nu2vl6pNz4gA0k3ar5BKnbdFHX415IBwB0FtKtnrpB0z9UO+cC0A2xnpp26nqp/bYnkzM+wGy2p4WThN7ZYTF/DQUQgQog3rx9RSzbU0bNAABwMmWOJ8mMQCDq6ckajw8s2tTpFEjddljUnOUeJuf9QF2tp7o+N9kwqOE7PxXSCQB13t0T4ZI90Nydo8KMz/ZgeEBWzVyxpVVvBzI5LUwFADJAt54AbMw2rI8E3NpvQk+rs1+aWT/GPMD9lpUlKlYWhVk2z+PnTc7gQL9llTjMr8lwRoXpvif3EYB+s2cKNldcrdrz14yATiVZJ+iN6zFGzQPALqVSg72TOH440vcDdssKK5a0CWq1p8iZgLt49YPvO5CXu+Kxgc/6rH9s/Q4AAP//PZzXiSo38ugAAAAASUVORK5CYII=\"\n  },\n  \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/order.js",
    "groupTitle": "订单管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/order/payOrder/1?orderType=1"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/device/cart/changeCartNum/1",
    "title": "修改购物车商品数量",
    "version": "0.0.1",
    "name": "修改购物车商品数量",
    "group": "购物车管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>修改购物车内商品数量</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "optional": false,
            "field": "id",
            "description": "<p>路径上使用购物车id</p>"
          }
        ]
      }
    },
    "body": [
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "cartNum",
        "description": "<p>商品数量</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/cart/changeCartNum/1",
        "type": "bash"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {},\n    \"message\": \"操作成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/cart.js",
    "groupTitle": "购物车管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/cart/changeCartNum/1"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "delete",
    "url": "/v1/device/cart/deleteCart",
    "title": "删除购物车商品",
    "version": "0.0.1",
    "name": "删除购物车商品",
    "group": "购物车管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>批量删除购物车商品</p>",
    "body": [
      {
        "group": "Body",
        "type": "Number[]",
        "optional": false,
        "field": "ids",
        "description": "<p>购物车id数组</p>"
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/cart/getProductCount",
        "type": "bash"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {\n        \"total\": 1\n    },\n    \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/cart.js",
    "groupTitle": "购物车管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/cart/deleteCart"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/device/cart/createCart",
    "title": "添加购物车",
    "version": "0.0.1",
    "name": "添加购物车",
    "group": "购物车管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>添加商品到购物车</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "cartNum",
            "description": "<p>商品数量</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "isNew",
            "description": "<p>是否为立即购买 1 是，2否</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "productType",
            "description": "<p>商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "productId",
            "description": "<p>商品 id</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "productAttrUnique",
            "description": "<p>商品规格唯一值</p>"
          }
        ]
      }
    },
    "body": [
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "cartNum",
        "description": "<p>商品数量</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "isNew",
        "description": "<p>是否为立即购买 1 是，2否</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "productType",
        "description": "<p>商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品</p>"
      },
      {
        "group": "Body",
        "type": "Number",
        "optional": false,
        "field": "productId",
        "description": "<p>商品 id</p>"
      },
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "productAttrUnique",
        "description": "<p>商品规格唯一值</p>"
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
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>购物车id</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "cartNum",
            "description": "<p>购物车商品数量</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isPay",
            "description": "<p>是否支付 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isDel",
            "description": "<p>是否删除 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isNew",
            "description": "<p>是否为立即购买 1 是，2否</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "isFail",
            "description": "<p>是否过期 1 是，2否</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 200,\n          \"data\": {\n              \"id\": 5,\n              \"createdAt\": \"2021-07-22T16:23:44.537+08:00\",\n              \"updatedAt\": \"2021-07-22T16:23:44.537+08:00\",\n              \"productType\": 1,\n              \"productAttrUnique\": \"e2fe28308fd2\",\n              \"cartNum\": 2,\n              \"source\": 0,\n              \"sourceId\": 0,\n              \"isPay\": 2,\n              \"isDel\": 2,\n              \"isNew\": 2,\n              \"isFail\": 2,\n              \"productId\": 1,\n              \"sysUserId\": 1,\n              \"sysTenancyId\": 1\n          },\n          \"message\": \"创建成功\"\n}",
          "type": "json"
        }
      ]
    },
    "examples": [
      {
        "title": "Curl example",
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/cart/createCart",
        "type": "bash"
      }
    ],
    "filename": "v1/device/cart.js",
    "groupTitle": "购物车管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/cart/createCart"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/device/cart/getProductCount",
    "title": "获取购物车商品数",
    "version": "0.0.1",
    "name": "获取购物车商品数",
    "group": "购物车管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>获取购物车商品数</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/cart/getProductCount",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "total",
            "description": "<p>商品总数</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n    \"status\": 200,\n    \"data\": {\n        \"total\": 1\n    },\n    \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/cart.js",
    "groupTitle": "购物车管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/cart/getProductCount"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/device/cart/getCartList",
    "title": "购物车商品列表",
    "version": "0.0.1",
    "name": "购物车商品列表",
    "group": "购物车管理-[床旁端]",
    "permission": [
      {
        "name": "device",
        "title": "床旁设备授权",
        "description": "<p>床旁设备授权，区分设备所在医院</p> <p>床旁设备请求平台接口之前都需要获取授权，并将授权凭证放置在头部信息中。</p>"
      }
    ],
    "description": "<p>获取购物车内商品列表信息</p>",
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
        "content": "curl -H \"Authorization: Bearer 5f048fe\" -i http://127.0.0.1:8089/v1/device/cart/getCartList",
        "type": "bash"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "sysTenancyId",
            "description": "<p>商户id</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>商户名称</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Avatar",
            "description": "<p>商户图片</p>"
          },
          {
            "group": "Success 200",
            "type": "Object[]",
            "optional": false,
            "field": "products",
            "description": "<p>商品集合</p>"
          },
          {
            "group": "Success 200",
            "type": "Number",
            "optional": false,
            "field": "total",
            "description": "<p>商品总数</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response:",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 200,\n      \"data\": {\n          fails\": [], // 失效商品\n          \"list\": [\n              {\n                  \"sysTenancyId\": 1,\n                  \"name\": \"宝安中心人民医院\",\n                  \"Avatar\": \"\",\n                  \"products\": [\n                      {\n                          \"id\": 1, // 购物车id\n                          \"sysTenancyId\": 1,\n                          \"productId\": 3,\n                          \"storeName\": \"精梳棉修身短袖T恤（圆/V领）\",\n                          \"image\": \"http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg\",\n                          \"price\": \"40.00\",\n                          \"cartNum\": 10,\n                          \"productAttrUnique\": \"e2fe28308fd2\",\n                          \"attrValue\": {\n                              \"productId\": 0,\n                              \"sku\": \"\",\n                              \"stock\": 0,\n                              \"sales\": 0,\n                              \"image\": \"\",\n                              \"barCode\": \"\",\n                              \"cost\": 0,\n                              \"otPrice\": 0,\n                              \"price\": 0,\n                              \"volume\": 0,\n                              \"weight\": 0,\n                              \"extensionOne\": 0,\n                              \"extensionTwo\": 0,\n                              \"unique\": \"\",\n                              \"detail\": null,\n                              \"value0\": \"\"\n                          }\n                      }\n                  ]\n              }\n          ],\n          \"total\": 1\n      },\n      \"message\": \"获取成功\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "v1/device/cart.js",
    "groupTitle": "购物车管理-[床旁端]",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8089/v1/device/cart/getCartList"
      }
    ],
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "status",
            "description": "<p>4001 授权错误时返回的状态码，得到次状态码需要重新授权。</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "message",
            "description": "<p>授权失败的具体描述信息</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Response (example):",
          "content": "HTTP/1.1 200 OK\n{\n      \"status\": 4001,\n      \"data\": {},\n      \"message\": \"mutil: invalid token\"\n}",
          "type": "json"
        }
      ]
    }
  }
] });
