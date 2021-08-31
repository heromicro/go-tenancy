/**
 * @api {get} /v1/device/patient/getPatientList 患者列表
 * @apiVersion 0.0.1
 * @apiName 患者列表
 * @apiGroup 患者管理-[床旁端]
 * @apiPermission device
 *
 * @apiDescription 获取医院患者数据
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://127.0.0.1:8089/v1/device/patient/getPatientList
 *
 * @apiSuccess {Number}   id            
 * @apiSuccess {String}   name           患者名称
 * @apiSuccess {String}   phone      手机
 * @apiSuccess {Number}   sex      性别
 * @apiSuccess {Number}   age      年龄
 * @apiSuccess {String}   locName      科室
 * @apiSuccess {String}   bedNum      床号
 * @apiSuccess {String}   hospitalNo      住院号
 * @apiSuccess {String}   disease      病种
 * @apiSuccess {String}   hospitalName      医院
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
                    "createdAt": "2021-07-26T12:24:42+08:00",
                    "updatedAt": "2021-07-26T12:24:42+08:00",
                    "name": "八两金",
                    "phone": "13845687419",
                    "sex": 2,
                    "age": 32,
                    "locName": "泌尿科一区",
                    "bedNum": "15",
                    "hospitalNo": "88956655",
                    "disease": "不孕不育",
                    "sysTenancyId": 1,
                    "hospitalName": "宝安中心人民医院"
                }
            ],
            "total": 1,
            "page": 1,
            "pageSize": -1
        },
        "message": "获取成功"
 *     }
 */


/**
 * @api {get} /v1/device/patient/getPatientDetail 获取当前患者
 * @apiVersion 0.0.1
 * @apiName 获取当前患者
 * @apiGroup 患者管理-[床旁端]
 * @apiPermission device
 *
 * @apiDescription 获取当前床旁设备患者
 *
 * @apiHeader {String} Authorization 接口需要带上此头信息
 * @apiHeaderExample {Header} Header-Example
 *     "Authorization: Bearer 5f048fe"
 *
 * @apiExample {bash} Curl example
 * curl -H "Authorization: Bearer 5f048fe" -i http://127.0.0.1:8089/v1/device/patient/getPatientDetail
 *
 * @apiSuccess {Number}   id            
 * @apiSuccess {String}   name           患者名称
 * @apiSuccess {String}   phone      手机
 * @apiSuccess {Number}   sex      性别
 * @apiSuccess {Number}   age      年龄
 * @apiSuccess {String}   locName      科室
 * @apiSuccess {String}   bedNum      床号
 * @apiSuccess {String}   hospitalNo      住院号
 * @apiSuccess {String}   disease      病种
 * @apiSuccess {String}   hospitalName      医院
 *
 * @apiUse TokenError
 * 
 * @apiSuccessExample Response:
 *     HTTP/1.1 200 OK
 *     {
        "status": 200,
        "data": {
            "id": 1,
            "createdAt": "2021-07-26T12:24:42+08:00",
            "updatedAt": "2021-07-26T17:28:06+08:00",
            "name": "八两金",
            "phone": "13845687419",
            "sex": 2,
            "age": 32,
            "locName": "泌尿科一区",
            "bedNum": "15",
            "hospitalNo": "88956655",
            "disease": "不孕不育",
            "sysTenancyId": 1
        },
        "message": "获取成功"
 *     }
 */