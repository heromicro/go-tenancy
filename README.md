<h1 align="center">GoTenancy</h1>

<div align="center">
    <a href="https://app.wercker.com/project/byKey/38763d8e14b612f57ad87f50a2b70f10">
      <img alt="Wercker status" src="https://app.wercker.com/status/38763d8e14b612f57ad87f50a2b70f10/s/master">
    </a>
    <a href="https://codecov.io/gh/snowlyg/go-tenancy"><img src="https://codecov.io/gh/snowlyg/go-tenancy/branch/master/graph/badge.svg" alt="Code Coverage"></a>
    <a href="https://goreportcard.com/report/github.com/snowlyg/go-tenancy"><img src="https://goreportcard.com/badge/github.com/snowlyg/go-tenancy" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/snowlyg/go-tenancy"><img src="https://godoc.org/github.com/snowlyg/go-tenancy?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/snowlyg/go-tenancy/blob/master/LICENSE"><img src="https://img.shields.io/github/license/snowlyg/go-tenancy" alt="Licenses"></a>
    <h5 align="center">多商户管理平台</h5>
</div>

###### `Iris-go` 学习交流 QQ 群 ：`676717248`
<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=cc99ccf86be594e790eacc91193789746af7df4a88e84fe949e61e5c6d63537c"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="Iris-go" title="Iris-go"></a>

If you don't have a QQ account, you can into the [iris-go-tenancy/community](https://gitter.im/iris-go-tenancy/community?utm_source=share-link&utm_medium=link&utm_campaign=share-link) .

- 基于 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 项目二次开发
- 生成 apidoc 文档
  ```shell
  cd resource/doc/ | apidoc -i v1/ -o apidoc/ -t template/
  ```
- 访问 http://127.0.0.1/doc

###### 价格逻辑

商品价格 
- 商品售价 price
- 商品原价 ot_price
- 商品成本价 cost
  
订单价格 
- 订单商品总价 total_price 
- 订单邮费 total_postage
- 订单支付总价 pay_price = total_price+total_postage
- 订单支付邮费 pay_postage = total_postage
- 订单平台手续费 commission_rate
- 订单成本价 cost = 商品成本价

###### 支付宝沙箱调试
- 需要设置 is-prod 为 false
- 下载 https://sandbox.alipaydev.com/user/downloadApp.htm 对应客户端
- 登录沙箱提供的账号


###### 接口测试

GO TEST 

在 tests 目录下增加 `password.txt` 文件,内容填写数据库密码,如果redis和mysql 密码不一样，需要处理一下获取到的文本内容；
```go
 pwds := strings.Split(password, ";")
 // pwds[1] mysql密码
 // pwds[2] redis密码
```


全局测试
```go
 go test -timeout 60s -run [^TestInitDB$] github.com/snowlyg/go-tenancy/tests 
```

迁移数据库，填充数据
```go
 go test -v -run ^TestInitDB$ github.com/snowlyg/go-tenancy/tests
```

使用 vscode 执行测试
- 用 vscode 打开项目 =》 终端 =》 运行任务 =》 选择对应任务执行
- `init db` 初始化数据库，填充数据
- `test all` 执行接口单元测试
- `build linux` 编译 linux 版本
- `apidoc` 更新接口文档


###### 使用第三方库
- 聚合支付 [gopay](https://github.com/go-pay/gopay)
- 雪花算法 [snowflake](https://github.com/bwmarrin/snowflake)
- 缓存 [redis](https://github.com/go-redis/redis/v8)
- 定时任务 [cron](https://github.com/robfig/cron/v3)
- 浮点计算 [decimal](https://github.com/shopspring/decimal)
- 二维码 [go-qrcode](https://github.com/skip2/go-qrcode)
- 认证 [multi](https://github.com/snowlyg/multi)
- 授权 [casbin](https://github.com/casbin/casbin/v2)
- 辅助 [pkg](https://github.com/chindeo/pkg) 
- 迁移 [migrate](https://github.com/golang-migrate/migrate) 
- 异步队列 [asynq](https://github.com/hibiken/asynq) 

