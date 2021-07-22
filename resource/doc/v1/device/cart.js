/**
 * @api {post} /v1/device/cart/createCart 添加购物车
 * @apiVersion 0.0.1
 * @apiName 添加购物车
 * @apiGroup 购物车管理
 * @apiPermission device
 *
 * @apiDescription 添加商品到购物车
 *     
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
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/cart/product/createCart
 *
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
