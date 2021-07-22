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
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/product/getProductList
 *
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
