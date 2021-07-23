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
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/cart/getCartList
 *
 * @apiUse TokenError
 * 
 * @apiSuccess {Number}   sysTenancyId       商户id
 * @apiSuccess {String}   name     商户名称
 * @apiSuccess {String}   Avatar    商户图片
 * @apiSuccess {Object []}   products   商品集合
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
                        "productId": 1,
                        "products": [
                            {
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
                                "sliderImages": null,
                                "attr": null,
                                "attrValue": null,
                                "cateId": 0,
                                "tenancyCategoryId": null,
                                "productCates": null
                            }
                        ]
                    }
                ],
                "total": 1,
                "page": 0,
                "pageSize": 0
            },
            "message": "获取成功"
 *     }
 */
