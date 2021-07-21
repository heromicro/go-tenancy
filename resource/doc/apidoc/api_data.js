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
    "description": "<p>用于医院床旁设备登录使用. 医院床旁设备必须授权登录后才可以调用平台其他接口，用于确定用户在哪个医院，患者信息用于用户注册患者和更新患者信息.</p>",
    "query": [
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "uuid",
        "description": "<p>医院唯一标识，由后台定义.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "name",
        "description": "<p>患者名称.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "phone",
        "description": "<p>手机号.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "sex",
        "description": "<p>性别 0 女，1男，2未知.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "age",
        "description": "<p>年龄.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "locName",
        "description": "<p>科室名.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "bedNum",
        "description": "<p>床号.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "hospitalNo",
        "description": "<p>住院号.</p>"
      },
      {
        "group": "Query",
        "type": "String",
        "optional": false,
        "field": "disease",
        "description": "<p>病种.</p>"
      }
    ],
    "body": [
      {
        "group": "Body",
        "type": "String",
        "optional": false,
        "field": "age",
        "description": "<p>Age of the User</p>"
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
    "filename": "v1/public.js",
    "groupTitle": "公共方法",
    "sampleRequest": [
      {
        "url": "https://apidoc.free.beeceptor.com/v1/public/device/login"
      }
    ]
  }
] });
