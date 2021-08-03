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
 * @apiSuccess {Number}   totalPrice            商品总价
 * @apiSuccess {Number}   totalOtPrice            商品原价
 * @apiSuccess {Number}   orderPrice            订单总价
 * @apiSuccess {Number}   orderOtPrice            订单原价
 * @apiSuccess {Number}   postagePrice            订单邮费
 * @apiSuccess {Number}   downPrice            订单优惠价格
 * @apiSuccess {Number}   totalNum            商品总数
 * @apiSuccess {Number}   orderType            订单类型 1：普通，2：自提
 * @apiSuccess {Object[]}   productPrices       商品价格
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
                    "id": 5,
                    "sysTenancyId": 1,
                    "specType": 2,
                    "productId": 1,
                    "storeName": "领立裁腰带短袖连衣裙",
                    "image": "http://127.0.0.1:8089/uploads/file/9a6a2e1231fb19517ed1de71206a0657.jpg",
                    "cartNum": 2,
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
            "totalPrice": "320",
            "totalOtPrice": "360",
            "orderPrice": "320",
            "orderOtPrice": "360",
            "postagePrice": "0",
            "downPrice": "0",
            "productPrices": {
                "1": {
                    "otPrice": "360",
                    "price": "320"
                }
            },
            "totalNum": 2,
            "orderType": 2
        },
        "message": "获取成功"
 *     }
 */


/**
 * @api {post} /v1/device/order/createOrder 用户结算订单
 * @apiVersion 0.0.1
 * @apiName 用户结算订单
 * @apiGroup 订单管理管理
 * @apiPermission device
 *
 * @apiDescription 用户结算订单，并生成待支付订单和支付二维码，用户通过支付宝或者微信扫码支付
 *   
 *
 * @apiBody {Number[]} cartIds 购物车ids
 * @apiBody {Number} orderType 订单类型：1：普通，2：自提
 * @apiBody {Number} [remark] 订单备注
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
 * @apiSuccess {String}   qrcode  扫码支付二维码 base64 数据，需要在前面加上 data:image/png;base64, 才能显示为图片；例如data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACWklEQVR42uyYMbLjIBBEW0VAyBF0FF3sl5BqL8ZROAIhAaXe6sH+9l9vbtg1gUuFnoMR0zM94LM+679cC0leno4nEzMAn4FNu2kioALYfXJlqXDa9nkpG4AwFbCy7J4s8MR6BzzbcIA+/uVTQywgD8wJKMwdW1tPAWeZD1BGYbNkapIJ8JeUezPQxZtaOOvm8h5SXc/yqu6xgdvaHC9sWC9sPsfyWk/HBpYKhUSEw7P1CE9JPD4K6QDA6ngogS6l/RHoeQFWkyYCmJvedQXnXWdhPw9dTAAAvXI2xKowgSqxW2OYB7iH6cpCCtiqalTFn9/hrUBdVX3geHg6cwarZdTxLd4JAPR2hQb4xG5y9sD6OKwZgIXUOaUWyNR1jkXtt4WBgG5yYD2rlyC54oqvyQDrWXYqN8nu2vl6pNz4gA0k3ar5BKnbdFHX415IBwB0FtKtnrpB0z9UO+cC0A2xnpp26nqp/bYnkzM+wGy2p4WThN7ZYTF/DQUQgQog3rx9RSzbU0bNAABwMmWOJ8mMQCDq6ckajw8s2tTpFEjddljUnOUeJuf9QF2tp7o+N9kwqOE7PxXSCQB13t0T4ZI90Nydo8KMz/ZgeEBWzVyxpVVvBzI5LUwFADJAt54AbMw2rI8E3NpvQk+rs1+aWT/GPMD9lpUlKlYWhVk2z+PnTc7gQL9llTjMr8lwRoXpvif3EYB+s2cKNldcrdrz14yATiVZJ+iN6zFGzQPALqVSg72TOH440vcDdssKK5a0CWq1p8iZgLt49YPvO5CXu+Kxgc/6rH9s/Q4AAP//PZzXiSo38ugAAAAASUVORK5CYII=
 * 
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
      "status": 200,
      "data": {
          "qrcode": "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACWklEQVR42uyYMbLjIBBEW0VAyBF0FF3sl5BqL8ZROAIhAaXe6sH+9l9vbtg1gUuFnoMR0zM94LM+679cC0leno4nEzMAn4FNu2kioALYfXJlqXDa9nkpG4AwFbCy7J4s8MR6BzzbcIA+/uVTQywgD8wJKMwdW1tPAWeZD1BGYbNkapIJ8JeUezPQxZtaOOvm8h5SXc/yqu6xgdvaHC9sWC9sPsfyWk/HBpYKhUSEw7P1CE9JPD4K6QDA6ngogS6l/RHoeQFWkyYCmJvedQXnXWdhPw9dTAAAvXI2xKowgSqxW2OYB7iH6cpCCtiqalTFn9/hrUBdVX3geHg6cwarZdTxLd4JAPR2hQb4xG5y9sD6OKwZgIXUOaUWyNR1jkXtt4WBgG5yYD2rlyC54oqvyQDrWXYqN8nu2vl6pNz4gA0k3ar5BKnbdFHX415IBwB0FtKtnrpB0z9UO+cC0A2xnpp26nqp/bYnkzM+wGy2p4WThN7ZYTF/DQUQgQog3rx9RSzbU0bNAABwMmWOJ8mMQCDq6ckajw8s2tTpFEjddljUnOUeJuf9QF2tp7o+N9kwqOE7PxXSCQB13t0T4ZI90Nydo8KMz/ZgeEBWzVyxpVVvBzI5LUwFADJAt54AbMw2rI8E3NpvQk+rs1+aWT/GPMD9lpUlKlYWhVk2z+PnTc7gQL9llTjMr8lwRoXpvif3EYB+s2cKNldcrdrz14yATiVZJ+iN6zFGzQPALqVSg72TOH440vcDdssKK5a0CWq1p8iZgLt49YPvO5CXu+Kxgc/6rH9s/Q4AAP//PZzXiSo38ugAAAAASUVORK5CYII="
      },
      "message": "获取成功"
 *     }
 */