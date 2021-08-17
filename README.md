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

POSTMAN 
- 测试导入地址 https://www.getpostman.com/collections/07881b7e98c809fa20cf
- 环境导入文件 [多商户运营平台.postman_environment.json](./多商户运营平台.postman_environment.json)

GO TEST 

在 tests 目录下增加 `main_test.go` 文件
```go
package tests

import (
	"fmt"
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/core"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/go-tenancy/initialize/cache"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
)

func TestMain(m *testing.M) {
	g.TENANCY_VP = core.Viper()     // 初始化Viper
	g.TENANCY_LOG = core.Zap()      // 初始化zap日志库
	g.TENANCY_CACHE = cache.Cache() // redis缓存
	// 初始化认证服务
	initialize.Auth()

	uuid := uuid.NewV3(uuid.NewV4(), uuid.NamespaceOID.String()).String()
	mysqlConfig := request.InitDB{
		SqlType: "mysql",
		Sql: request.Sql{
			Host:     "127.0.0.1",
			Port:     "3306",
			UserName: "root",
			Password: "",
			DBName:   uuid,
		},
		CacheType: "redis",
		Cache: request.Cache{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "",
		},
		Addr:  8089,
		Level: "test",
		Env:   "dev",
	}
	service.InitDB(mysqlConfig)

	req := request.CreateTenancy{
		SysTenancy: model.SysTenancy{
			BaseTenancy: model.BaseTenancy{
				Name:          "多商户平台直营医院",
				Tele:          "0755-23568911",
				Address:       "xxx街道666号",
				BusinessTime:  "08:30-17:30",
				Status:        g.StatusTrue,
				SysRegionCode: 1,
				IsAudit:       g.StatusFalse, // 商品无需审核
			},
		},
		Username: "tenancy_hospital",
	}
	tennancyId, tenancyUUID, username, err := service.CreateTenancy(req)
	if err != nil {
		fmt.Printf("初始化商户错误： %v\n", err)
		return
	}
	cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":username", username, 0)
	cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":id", tennancyId, 0)
	cache.SetCache(g.TENANCY_CONFIG.Mysql.Dbname+":uuid", tenancyUUID, 0)

	// call flag.Parse() here if TestMain uses flags
	// 如果 TestMain 使用了 flags，这里应该加上 flag.Parse()
	code := m.Run()

	err = dorpDB(uuid)
	if err != nil {
		fmt.Printf("初始化商户错误： %v\n", err)
		return
	}

	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":username")
	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":id")
	cache.DeleteCache(g.TENANCY_CONFIG.Mysql.Dbname + ":uuid")

	db, _ := g.TENANCY_DB.DB()
	db.Close()
	multi.AuthDriver.Close()

	os.Exit(code)
}

func dorpDB(uuid string) error {
	// 删除表和视图
	var sqls []string
	if err := g.TENANCY_DB.Raw("select CASE table_type WHEN 'VIEW' THEN concat('drop view ', table_name, ';') ELSE concat('drop table ', table_name, ';') END  from information_schema.tables where table_schema='%s';", uuid).Scan(&sqls).Error; err != nil {
		return err
	}

	for _, sql := range sqls {
		if err := g.TENANCY_DB.Exec(sql).Error; err != nil {
			continue
		}
	}

	if err := g.TENANCY_DB.Exec(fmt.Sprintf("drop database if exists `%s`;", uuid)).Error; err != nil {
		return err
	}

	return nil
}

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

