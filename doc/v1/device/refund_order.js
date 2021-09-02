/**
 * @api {GET} /v1/device/refundOrder/getRfundOrderById/1 根据id获取退款订单详情
 * @apiVersion 0.0.1
 * @apiName 根据id获取退款订单详情
 * @apiGroup 订单管理-[床旁端]
 * @apiPermission device
 *
 * @apiDescription 根据id获取退款订单详情
 *   
 * 
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://127.0.0.1:8089/v1/device/refundOrder/getRfundOrderById/1
 *
 * @apiUse TokenError
 *         
 * @apiSuccess {Number} id 退款订单id
 * @apiSuccess {String} createdAt 创建时间
 * @apiSuccess {String} updatedAt 更新时间
 * @apiSuccess {String} refundOrderSn 退款订单号
 * @apiSuccess {String} deliveryType 快递公司
 * @apiSuccess {String} deliveryId 快递单号
 * @apiSuccess {String} deliveryMark 快递备注
 * @apiSuccess {String} deliveryPics 快递凭证
 * @apiSuccess {String} deliveryPhone 快递联系电话
 * @apiSuccess {String} merDeliveryUser 收货人
 * @apiSuccess {String} merDeliveryAddress 收货地址
 * @apiSuccess {String} Phone 收货人联系电话
 * @apiSuccess {String} Mark 备注
 * @apiSuccess {String} MerMark 商户备注
 * @apiSuccess {String} adminMark 平台备注
 * @apiSuccess {String} Pics 图片
 * @apiSuccess {Number} RefundType 退款类型 1:退款 2:退款退货
 * @apiSuccess {String} RefundMessage 退款原因
 * @apiSuccess {Number} RefundPrice 退款金额
 * @apiSuccess {Number} RefundNum 退款数
 * @apiSuccess {String} FailMessage 退款未通过原因
 * @apiSuccess {Number} Status 状态 1:待审核 2:待退货 3:待收货 4:已退款 5:审核未通过
 * @apiSuccess {String} StatusTime 状态改变时间
 * @apiSuccess {Number} patientId 床旁用户
 * @apiSuccess {Number} sysUserId 小程序用户
 * @apiSuccess {Number} sysTenancyId 商户

 * 
 *
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
        {
          "status": 200,
          "data": {
            "id": 1,
            "createdAt": "2021-09-02T15:00:27+08:00",
            "updatedAt": "2021-09-02T15:00:27+08:00",
            "refundOrderSn": "R2021090215002733323921448374272",
            "deliveryType": "",
            "deliveryId": "",
            "deliveryMark": "",
            "deliveryPics": "",
            "deliveryPhone": "",
            "merDeliveryUser": "",
            "merDeliveryAddress": "",
            "phone": "",
            "mark": "",
            "merMark": "",
            "adminMark": "",
            "pics": "",
            "refundType": 1,
            "refundMessage": "地址错了",
            "refundPrice": 1,
            "refundNum": 1,
            "failMessage": "",
            "status": 1,
            "statusTime": "2021-09-02T15:00:27+08:00",
            "isDel": 2,
            "isSystemDel": 2,
            "reconciliationId": 0,
            "patientId": 0,
            "orderId": 1,
            "sysUserId": 0,
            "sysTenancyId": 0
          },
          "message": "操作成功"
        }
 */


 /**
 * @api {post} /v1/device/refundOrder/getRfundOrderList 我的退款订单
 * @apiVersion 0.0.1
 * @apiName 我的退款订单
 * @apiGroup 订单管理-[床旁端]
 * @apiPermission device
 *
 * @apiDescription 床旁用户的退款订单列表
 *   
 * @apiBody {Number} pageSize 页数
 * @apiBody {Number} page 页码
 * @apiBody {String} status 状态 1:待审核 2:待退货 3:待收货 4:已退款 5:审核未通过
 * @apiBody {String} date 日期：today，yesterday，lately7，lately30，month，year或者日期范围:2021/08/01-2021/08/05 
 * 
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://127.0.0.1:8089/v1/device/refundOrder/getRfundOrderList
 *
 * @apiUse TokenError
 *         
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
                    "createdAt": "2021-09-02T15:00:27+08:00",
                    "updatedAt": "2021-09-02T15:00:27+08:00",
                    "refundOrderSn": "R2021090215002733323921448374272",
                    "deliveryType": "",
                    "deliveryId": "",
                    "deliveryMark": "",
                    "deliveryPics": "",
                    "deliveryPhone": "",
                    "merDeliveryUser": "",
                    "merDeliveryAddress": "",
                    "phone": "",
                    "mark": "",
                    "merMark": "",
                    "adminMark": "",
                    "pics": "",
                    "refundType": 1,
                    "refundMessage": "地址错了",
                    "refundPrice": 1,
                    "refundNum": 1,
                    "failMessage": "",
                    "status": 1,
                    "statusTime": "2021-09-02T15:00:27+08:00",
                    "isDel": 2,
                    "isSystemDel": 2,
                    "reconciliationId": 0,
                    "patientId": 0,
                    "orderId": 1,
                    "sysUserId": 0,
                    "sysTenancyId": 0
              ],
              "page": 1,
              "pageSize": 20,
              "total": 12
          },
          "message": "获取成功"
 *     }
 */
