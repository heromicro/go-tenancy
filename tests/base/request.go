package base

import (
	"net/http"

	"github.com/gavv/httpexpect"
)

var PageRes = map[string]interface{}{"page": 1, "pageSize": 10}
var PageKeys = ResponseKeys{
	{Type: "int", Key: "pageSize", Value: 10},
	{Type: "int", Key: "page", Value: 1},
}

func PostList(auth *httpexpect.Expect, url string, id uint, res map[string]interface{}, pageKeys ResponseKeys, status int, message string) {
	obj := auth.POST(url).
		WithJSON(res).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
	pageKeys.Test(obj.Value("data").Object())
}

func GetList(auth *httpexpect.Expect, url string, id uint, keys ResponseKeys, status int, message string) {
	obj := auth.GET(url).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
	if id > 0 {
		obj.Value("data").Array().Length().Equal(1)
		first := obj.Value("data").Array().First().Object()
		keys.Test(first)
	}
}

func Create(auth *httpexpect.Expect, url string, create map[string]interface{}, keys ResponseKeys, status int, message string) {
	obj := auth.POST(url).
		WithJSON(create).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
	if status == http.StatusOK {
		if keys != nil {
			keys.Scan(obj.Value("data").Object())
		}
	}
}

func Update(auth *httpexpect.Expect, url string, update map[string]interface{}, status int, message string) {
	obj := auth.PUT(url).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}

func Get(auth *httpexpect.Expect, url string, status int, message string) {
	obj := auth.GET(url).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}

func GetById(auth *httpexpect.Expect, url string, id uint, keys ResponseKeys, status int, message string) {
	obj := auth.GET(url).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
	if id > 0 {
		keys.Test(obj.Value("data").Object())
	}
}

func Post(auth *httpexpect.Expect, url string, data map[string]interface{}, status int, message string) {
	obj := auth.POST(url).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}

func Delete(auth *httpexpect.Expect, url string, status int, message string) {
	obj := auth.DELETE(url).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}

func DeleteMutil(auth *httpexpect.Expect, url string, data map[string]interface{}, status int, message string) {
	obj := auth.DELETE(url).
		WithJSON(data).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}
