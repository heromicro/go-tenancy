/**
 * @api {get} /v1/device/productCategory/getProductCategoryList 商品分类列表
 * @apiVersion 0.0.1
 * @apiName 商品分类列表
 * @apiGroup 商品分类
 * @apiPermission device
 *
 * @apiDescription 获取商品分类数据
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://10.0.0.26:8085/v1/device/productCategory/getProductCategoryList
 *
 * @apiSuccess {Number}   id            分类id
 * @apiSuccess {Number}   pid           上级id
 * @apiSuccess {String}   cateName      分类名称
 * @apiSuccess {Number}   sort      排序
 * @apiSuccess {String}   pic      图片
 * @apiSuccess {Number}   level      等级
 * @apiSuccess {Object[]}   children      子分类
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
    "status": 200,
    "data": [
        {
            "id": 173,
            "createdAt": "2021-07-22T09:27:57+08:00",
            "updatedAt": "2021-07-22T09:27:57+08:00",
            "pid": 0,
            "cateName": "品牌服饰",
            "path": "/",
            "sort": 2,
            "pic": "",
            "status": 1,
            "level": 0,
            "children": [
                {
                    "id": 174,
                    "createdAt": "2021-07-22T09:27:57+08:00",
                    "updatedAt": "2021-07-22T09:27:57+08:00",
                    "pid": 173,
                    "cateName": "时尚女装",
                    "path": "/173/",
                    "sort": 0,
                    "pic": "",
                    "status": 1,
                    "level": 1,
                    "children": null
                }
            ]
        }
    ],
    "message": "获取成功"
 *     }
 */