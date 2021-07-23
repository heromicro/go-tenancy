/**
 * @api {post} /v1/device/cart/createCart 添加购物车
 * @apiVersion 0.0.1
 * @apiName 添加购物车
 * @apiGroup 购物车管理
 * @apiPermission device
 *
 * @apiDescription 添加商品到购物车
 *     
 * @apiParam {Number} cartNum 商品数量
 * @apiParam {Number} isNew 是否为立即购买 1 是，2否
 * @apiParam {Number} productType 商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
 * @apiParam {Number} productId 商品 id 
 * @apiParam {String} productAttrUnique 商品规格唯一值 
 * 
 * @apiBody {Number} cartNum 商品数量
 * @apiBody {Number} isNew 是否为立即购买 1 是，2否
 * @apiBody {Number} productType 商品类型 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
 * @apiBody {Number} productId 商品 id 
 * @apiBody {String} productAttrUnique 商品规格唯一值 
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiSuccess {Number}   id            购物车id
 * @apiSuccess {Number}   cartNum     购物车商品数量
 * @apiSuccess {Number}   isPay     是否支付 1 是，2否
 * @apiSuccess {Number}   isDel     是否删除 1 是，2否
 * @apiSuccess {Number}   isNew     是否为立即购买 1 是，2否
 * @apiSuccess {Number}   isFail     是否过期 1 是，2否
 * 
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/cart/createCart
 *
 * @apiUse TokenError
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
          "status": 200,
              "data": {
                  "id": 5,
                  "createdAt": "2021-07-22T16:23:44.537+08:00",
                  "updatedAt": "2021-07-22T16:23:44.537+08:00",
                  "productType": 1,
                  "productAttrUnique": "e2fe28308fd2",
                  "cartNum": 2,
                  "source": 0,
                  "sourceId": 0,
                  "isPay": 2,
                  "isDel": 2,
                  "isNew": 2,
                  "isFail": 2,
                  "productId": 1,
                  "sysUserId": 1,
                  "sysTenancyId": 1
              },
              "message": "创建成功"
 *     }
 */


/**
 * @api {get} /v1/device/cart/getCartList 购物车商品列表
 * @apiVersion 0.0.1
 * @apiName 购物车商品列表
 * @apiGroup 购物车管理
 * @apiPermission device
 *
 * @apiDescription 获取购物车内商品列表信息    
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/cart/getCartList
 *
 * @apiUse TokenError
 * 
 * @apiSuccess {Number}   sysTenancyId       商户id
 * @apiSuccess {String}   name     商户名称
 * @apiSuccess {String}   Avatar    商户图片
 * @apiSuccess {Object[]}   products   商品集合
 * @apiSuccess {Number}   total   商品总数
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
        "status": 200,
            "data": {
                "list": [
                    {
                        "sysTenancyId": 1,
                        "name": "宝安中心人民医院",
                        "Avatar": "",
                        "products": [
                            {
                                "sysTenancyId": 1,
                                "productId": 1,
                                "storeName": "领立裁腰带短袖连衣裙",
                                "image": "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
                                "price": "80.00",
                                "cartNum": 213
                            },
                            {
                                "sysTenancyId": 1,
                                "productId": 3,
                                "storeName": "精梳棉修身短袖T恤（圆/V领）",
                                "image": "http://127.0.0.1:8090/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
                                "price": "40.00",
                                "cartNum": 6
                            }
                        ]
                    }
                ],
                "total": 2
            },
            "message": "获取成功"
 *     }
 */

            
/**
 * @api {post} /v1/device/cart/changeCartNum/1 修改购物车商品数量
 * @apiVersion 0.0.1
 * @apiName 修改购物车商品数量
 * @apiGroup 购物车管理
 * @apiPermission device
 *
 * @apiDescription 修改购物车内商品数量
 *   
 * @apiParam id 路径上使用商品id   
 *
 * @apiBody {Number} cartNum 商品数量
 * 
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/cart/changeCartNum/1
 *
 * @apiUse TokenError
 * 
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
        "status": 200,
        "data": {},
        "message": "操作成功"
 *     }
 */


 /**
 * @api {get} /v1/device/cart/getProductCount 获取购物车商品数
 * @apiVersion 0.0.1
 * @apiName 获取购物车商品数
 * @apiGroup 购物车管理
 * @apiPermission device
 *
 * @apiDescription 获取购物车商品数    
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/cart/getProductCount
 *
 * @apiUse TokenError
 * 
 * @apiSuccess {Number}   total   商品总数
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
        "status": 200,
        "data": {
            "total": 1
        },
        "message": "获取成功"
 *     }
 */

 /**
 * @api {delete} /v1/device/cart/deleteCart 删除购物车商品
 * @apiVersion 0.0.1
 * @apiName 删除购物车商品
 * @apiGroup 购物车管理
 * @apiPermission device
 *
 * @apiDescription 批量删除购物车商品  
 * 
 * @apiBody {Number[]} ids 商品id数组
 * 
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/cart/getProductCount
 *
 * @apiUse TokenError
 * 
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
        "status": 200,
        "data": {
            "total": 1
        },
        "message": "获取成功"
 *     }
 */