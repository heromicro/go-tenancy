package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/g"
)

func TestMqttList(t *testing.T) {
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/mqtt/getMqttList").
		WithJSON(map[string]interface{}{"page": 1, "pageSize": 10}).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("获取成功")

	data := obj.Value("data").Object()
	data.Keys().ContainsOnly("list", "total", "page", "pageSize")
	data.Value("pageSize").Number().Equal(10)
	data.Value("page").Number().Equal(1)
	data.Value("total").Number().Equal(3)

	list := data.Value("list").Array()
	list.Length().Ge(0)
	first := list.First().Object()
	first.Keys().ContainsOnly("id", "host", "status", "port", "username", "password", "clientId", "createdAt", "updatedAt")
}

func TestMqttProcess(t *testing.T) {
	data := map[string]interface{}{
		"host":     "127.0.0.2",
		"port":     1883,
		"username": "1",
		"password": "1",
		"status":   1,
	}
	auth := baseWithLoginTester(t)
	defer baseLogOut(auth)
	obj := auth.POST("v1/admin/mqtt/createMqtt").
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("创建成功")

	mqtt := obj.Value("data").Object()
	mqtt.Value("id").Number().Ge(0)
	mqttId := mqtt.Value("id").Number().Raw()
	if mqttId > 0 {

		update := map[string]interface{}{
			"host":     "127.0.0.2",
			"port":     1883,
			"username": "2",
			"password": "2",
			"status":   1,
		}

		obj = auth.PUT(fmt.Sprintf("v1/admin/mqtt/updateMqtt/%d", int(mqttId))).
			WithJSON(update).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("更新成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/mqtt/getMqttById/%d", int(mqttId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("操作成功")
		mqtt = obj.Value("data").Object()

		mqtt.Value("id").Number().Ge(0)
		mqtt.Value("name").String().Equal(update["name"].(string))
		mqtt.Value("status").Number().Equal(update["status"].(int))
		mqtt.Value("code").String().Equal(update["code"].(string))
		mqtt.Value("sort").Number().Equal(update["sort"].(int))

		obj = auth.POST("v1/admin/mqtt/changeMqttStatus").
			WithJSON(map[string]interface{}{
				"id":     mqttId,
				"status": g.StatusTrue,
			}).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("设置成功")

		obj = auth.GET("v1/admin/mqtt/getCreateMqttMap").
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		obj = auth.GET(fmt.Sprintf("v1/admin/mqtt/getUpdateMqttMap/%d", int(mqttId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("获取成功")

		// deleteMqtt
		obj = auth.DELETE(fmt.Sprintf("v1/admin/mqtt/deleteMqtt/%d", int(mqttId))).
			Expect().Status(http.StatusOK).JSON().Object()
		obj.Keys().ContainsOnly("status", "data", "message")
		obj.Value("status").Number().Equal(200)
		obj.Value("message").String().Equal("删除成功")
	}
}
