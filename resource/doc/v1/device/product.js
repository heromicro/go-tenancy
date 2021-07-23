/**
 * @api {post} /v1/device/product/getProductList 商品列表
 * @apiVersion 0.0.1
 * @apiName 商品列表
 * @apiGroup 商品管理
 * @apiPermission device
 *
 * @apiDescription 获取商品列表数据
 *     
 * @apiParam {Number} page 页码
 * @apiParam {Number} pageSize 每页数量
 * @apiParam {Number} cateId 商户分类 id
 * @apiParam {Number} tenancyCategoryId 商城分类 id 
 * @apiParam {String} keyword 关键字搜索 
 * @apiParam {String} type 商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
 * @apiParam {String} [isGiftBag] 是否礼包 1 是 ，2 否
 * 
 * @apiBody {Number} page 页码
 * @apiBody {Number} pageSize 每页数量
 * @apiBody {Number} cateId 商户分类 id
 * @apiBody {Number} tenancyCategoryId 商城分类 id 
 * @apiBody {String} keyword 关键字搜索 
 * @apiBody {String} type 商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
 * @apiBody {String} [isGiftBag] 是否礼包 1 是 ，2 否
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/product/getProductList
 *
 * @apiUse TokenError
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
    "status": 200,
    "data": {
        "list": [
            {
                "id": 1,
                "storeName": "领立裁腰带短袖连衣裙",
                "sales": 1,
                "price": 80,
                "image": "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"
            }
        ],
        "total": 1,
        "page": 1,
        "pageSize": 10
    },
    "message": "获取成功"
 *     }
 */




/**
 * @api {get} /v1/device/product/getProductById/1 商品详情
 * @apiVersion 0.0.1
 * @apiName 商品详情
 * @apiGroup 商品管理
 * @apiPermission device
 *
 * @apiDescription 获取商品详情数据
 *     
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/product/getProductById/1
 *
 * @apiSuccess {Number}   id            商品id
 * @apiSuccess {String}   storeName     商品名称
 * @apiSuccess {String}   storeInfo     商品简介
 * @apiSuccess {String}   keyword       商品关键词
 * @apiSuccess {String}   barCode       商品条码
 * @apiSuccess {Number}   isShow        是否显示 1 是，2否
 * @apiSuccess {Number}   status        商品状态 1：审核通过,2：审核中 3: 未通过
 * @apiSuccess {String}   unitName      商品单位
 * @apiSuccess {Number}   sort      排序
 * @apiSuccess {Number}   rank      总后台排序
 * @apiSuccess {Number}   sales      商品销量
 * @apiSuccess {Number}   price      最低价格
 * @apiSuccess {Number}   cost      成本价
 * @apiSuccess {Number}   otPrice      原价
 * @apiSuccess {Number}   stock      库存
 * @apiSuccess {Number}   isHot      是否热卖 1 是，2否
 * @apiSuccess {Number}   isBenefit      促销单品 1 是，2否
 * @apiSuccess {Number}   isBest      是否精品 1 是，2否
 * @apiSuccess {Number}   isNew      是否新品 1 是，2否
 * @apiSuccess {Number}   isGood      是否优品推荐 1 是，2否
 * @apiSuccess {Number}   productType      商品分类 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
 * @apiSuccess {Number}   ficti      虚拟销量
 * @apiSuccess {Number}   browse      浏览量
 * @apiSuccess {String}   codePath      产品二维码地址
 * @apiSuccess {String}   videoLink      主图视频链接
 * @apiSuccess {Number}   specType       规格 1单 2多
 * @apiSuccess {String}   refusal       审核拒绝理由
 * @apiSuccess {Number}   Rate       评价分数
 * @apiSuccess {Number}   ReplyCount       评论数
 * @apiSuccess {Number}   isGiftBag       是否为礼包
 * @apiSuccess {Number}   careCount       收藏数
 * @apiSuccess {String}   image       商品图片
 * @apiSuccess {String}   sysTenancyName       医院名称
 * @apiSuccess {String}   cateName       后台分类名称
 * @apiSuccess {String}   brandName       品牌名称
 * @apiSuccess {String}   tempName       模板名称
 * @apiSuccess {String}   content       详情内容
 * @apiSuccess {String[]}   sliderImages      展示图片
 * @apiSuccess {String[]}   attr      规格
 * @apiSuccess {Object[]}   attrValue      规格详情
 * @apiSuccess {Object[]}   productCates  商品分类
 *
 * @apiUse TokenError
 * 
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
  "status": 200,
    "data": {
        "id": 1,
        "storeName": "领立裁腰带短袖连衣裙",
        "storeInfo": "短袖连衣裙",
        "keyword": "连衣裙",
        "unitName": "件",
        "sort": 40,
        "sales": 1,
        "price": 80,
        "otPrice": 100,
        "stock": 399,
        "isHot": 2,
        "isBenefit": 2,
        "isBest": 2,
        "isNew": 2,
        "isGood": 1,
        "productType": 2,
        "ficti": 100,
        "specType": 1,
        "rate": 5,
        "isGiftBag": 2,
        "image": "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
        "tempId": 99,
        "sysTenancyId": 1,
        "sysBrandId": 2,
        "productCategoryId": 162,
        "sysTenancyName": "宝安中心人民医院",
        "cateName": "男士上衣",
        "brandName": "苹果",
        "tempName": "",
        "content": "<p>好手机</p>",
        "sliderImage": "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg,http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
        "sliderImages": [
            "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
            "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg"
        ],
        "attr": [
            {
                "detail": [
                    "35"
                ],
                "value": "S"
            },
            {
                "detail": [
                    "36"
                ],
                "value": "L"
            },
            {
                "detail": [
                    "37"
                ],
                "value": "XL"
            },
            {
                "detail": [
                    "38"
                ],
                "value": "XXL"
            }
        ],
        "attrValue": [
            {
                "sku": "S",
                "stock": 99,
                "sales": 1,
                "image": "\thttp://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
                "barCode": "123456",
                "cost": 50,
                "otPrice": 180,
                "price": 160,
                "volume": 1,
                "weight": 1,
                "extensionOne": 0,
                "extensionTwo": 0,
                "unique": "e2fe28308fd0",
                "detail": {
                    "尺寸": "S"
                },
                "value0": "S"
            },
            {
                "sku": "L",
                "stock": 100,
                "sales": 0,
                "image": "\thttp://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
                "barCode": "123456",
                "cost": 50,
                "otPrice": 180,
                "price": 160,
                "volume": 1,
                "weight": 1,
                "extensionOne": 0,
                "extensionTwo": 0,
                "unique": "e2fe28308fd0",
                "detail": {
                    "尺寸": "L"
                },
                "value0": "L"
            },
            {
                "sku": "XL",
                "stock": 100,
                "sales": 0,
                "image": "\thttp://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
                "barCode": "123456",
                "cost": 50,
                "otPrice": 180,
                "price": 160,
                "volume": 1,
                "weight": 1,
                "extensionOne": 0,
                "extensionTwo": 0,
                "unique": "e2fe28308fd0",
                "detail": {
                    "尺寸": "XL"
                },
                "value0": "XL"
            },
            {
                "sku": "XXL",
                "stock": 100,
                "sales": 0,
                "image": "\thttp://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
                "barCode": "123456",
                "cost": 50,
                "otPrice": 180,
                "price": 160,
                "volume": 1,
                "weight": 1,
                "extensionOne": 0,
                "extensionTwo": 0,
                "unique": "e2fe28308fd0",
                "detail": {
                    "尺寸": "XXL"
                },
                "value0": "XXL"
            }
        ],
        "cateId": 162,
        "tenancyCategoryId": [
            174,
            173
        ],
        "productCates": [
            {
                "id": 174,
                "cateName": "时尚女装"
            },
            {
                "id": 173,
                "cateName": "品牌服饰"
            }
        ]
    },
    "message": "操作成功"
 *     }
 */