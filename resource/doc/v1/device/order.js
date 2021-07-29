/**
 * @api {post} /v1/device/order/checkOrder 用户下单数据
 * @apiVersion 0.0.1
 * @apiName 用户下单数据
 * @apiGroup 订单管理管理
 * @apiPermission device
 *
 * @apiDescription 获取用户下单数据详情
 *   
 *
 * @apiBody {Number[]} cartIds 购物车ids
 * 
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://127.0.0.1:8089/v1/device/order/checkOrder
 *
 * @apiUse TokenError
 *         
 * @apiSuccess {Number}   finalPrice            订单最终价格
 * @apiSuccess {Number}   finalOtPrice            订单最终原价
 * @apiSuccess {Number}   totalPrice            订单总价
 * @apiSuccess {Number}   totalOtPrice            订单原价
 * @apiSuccess {Number}   postagePrice            订单邮费
 * @apiSuccess {Number}   downPrice            订单优惠价格
 * @apiSuccess {Number}   totalNum            商品总数
 * @apiSuccess {Number}   orderType            订单类型 1：普通，2：自提
 * @apiSuccess {Object[]}   productPrices            商品价格
 * @apiSuccess {Object[]}   products            商品信息
 * 
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
        "status": 200,
        "data": {
            "sysTenancyId": 1,
            "name": "宝安中心人民医院",
            "Avatar": "",
            "products": [
                {
                    "id": 6,
                    "sysTenancyId": 1,
                    "specType": 1,
                    "productId": 1,
                    "storeName": "领立裁腰带短袖连衣裙",
                    "image": "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg",
                    "cartNum": 6,
                    "isFail": 2,
                    "productAttrUnique": "e2fe28308fd2",
                    "attrValue": {
                        "productId": 0,
                        "sku": "S",
                        "stock": 99,
                        "sales": 1,
                        "image": "\thttp://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg",
                        "barCode": "123456",
                        "cost": 50,
                        "otPrice": 180,
                        "price": 160,
                        "volume": 1,
                        "weight": 1,
                        "extensionOne": 0,
                        "extensionTwo": 0,
                        "unique": "e2fe28308fd2",
                        "detail": {
                            "尺寸": "S"
                        },
                        "value0": "S"
                    }
                }
            ],
            "finalPrice": "960",
            "finalOtPrice": "1080",
            "totalPrice": "960",
            "totalOtPrice": "1080",
            "postagePrice": "0",
            "downPrice": "0",
            "productPrices": {
                "1": {
                    "otPrice": "1080",
                    "price": "960"
                }
            },
            "totalNum": 6,
            "orderType": 2
        },
        "message": "获取成功"
 *     }
 */