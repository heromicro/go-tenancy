/**
 * @api {post} /v1/public/device/login 床旁授权登录
 * @apiVersion 0.0.1
 * @apiName 床旁授权登录
 * @apiGroup 公共方法
 * @apiPermission none
 *
 * @apiDescription 用于医院床旁设备登录使用.
 * 医院床旁设备必须授权登录后才可以调用平台其他接口，用于确定用户在哪个医院，患者信息用于用户注册患者和更新患者信息.
 *
 * @apiQuery {String} uuid 医院唯一标识，由后台定义.
 * @apiQuery {String} name 患者名称.
 * @apiQuery {String} phone 手机号.
 * @apiQuery {String} sex 性别 0 女，1男，2未知.
 * @apiQuery {String} age 年龄.
 * @apiQuery {String} locName 科室名.
 * @apiQuery {String} bedNum 床号.
 * @apiQuery {String} hospitalNo 住院号.
 * @apiQuery {String} disease 病种.
 *
 * @apiBody {String} age Age of the User
 *
 * @apiSuccessExample {json} Success-Response:
 *     HTTP/1.1 200 OK
 *     {
    "status": 200,
    "data": {
        "user": {
            "patient": {
                "id": 1,
                "createdAt": "2021-07-21T17:16:22+08:00",
                "updatedAt": "2021-07-21T17:39:16.715+08:00",
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
            "tenancy": {
                "id": 1,
                "createdAt": "2021-07-21T17:16:20+08:00",
                "updatedAt": "2021-07-21T17:16:20+08:00",
                "uuid": "c976999e-b004-403c-96b7-e2390f64fbb7",
                "name": "宝安中心人民医院",
                "tele": "0755-23568911",
                "address": "xxx街道888号",
                "businessTime": "08:30-17:30",
                "status": 1,
                "Keyword": "",
                "Avatar": "",
                "Banner": "",
                "sales": 0,
                "productScore": 5,
                "serviceScore": 5,
                "postageScore": 5,
                "mark": "",
                "regAdminId": 0,
                "sort": 0,
                "isAudit": 2,
                "isBest": 2,
                "isTrader": 2,
                "State": 1,
                "Info": "",
                "servicePhone": "",
                "careCount": 0,
                "copyProductNum": 0,
                "sysRegionCode": 1,
                "region": {
                    "code": 0,
                    "pCode": 0,
                    "name": "",
                    "subRegions": null
                }
            }
        },
        "AccessToken": "TVRReE56YzRNVEl4TWpReE9UY3lNekkyTkEuTWpBeU1TMHdOeTB5TVZReE56b3pPVG94Tmlzd09Eb3dNQQ.MTQxNzc4MTIxMjQxOTcyMzI2NA"
    },
    "message": "登录成功"
 *     }
 *
 */
