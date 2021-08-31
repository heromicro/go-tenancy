/**
 * @api {get} /v1/device/productCategory/getProductCategoryList 商品分类列表
 * @apiVersion 0.0.1
 * @apiName 商品分类列表
 * @apiGroup 商品分类管理-[床旁端]
 * @apiPermission device
 *
 * @apiDescription 获取商品分类数据
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://127.0.0.1:8089/v1/device/productCategory/getProductCategoryList
 *
 * @apiSuccess {Number}   id            分类id
 * @apiSuccess {Number}   pid           上级id
 * @apiSuccess {String}   cateName      分类名称
 * @apiSuccess {Number}   sort      排序
 * @apiSuccess {String}   pic      图片
 * @apiSuccess {Number}   level      等级
 * @apiSuccess {Object[]}   children      子分类
 *
 * @apiUse TokenError
 * 
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
    "status": 200,
    "data": [
        {
            "id": 173,
            "cateName": "品牌服饰",
            "pic": "",
            "children": [
                {
                    "id": 174,
                    "pid": 173,
                    "cateName": "时尚女装",
                    "pic": ""
                }
            ]
        }
    ],
    "message": "获取成功"
 *     }
 */