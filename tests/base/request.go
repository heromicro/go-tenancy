package base

import (
	"net/http"

	"github.com/gavv/httpexpect"
)

func GetList(auth *httpexpect.Expect, url string, id uint, data map[string]interface{}, keys ResponseKeys, status int, message string) {
	obj := auth.GET(url).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
	if id > 0 && data != nil {
		obj.Value("data").Array().Length().Equal(1)
		first := obj.Value("data").Array().First().Object()
		keys.Test(first, id, data)
	}
}

func Create(auth *httpexpect.Expect, url string, create map[string]interface{}, status int, message string) uint {
	obj := auth.POST(url).
		WithJSON(create).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
	if status == http.StatusOK {
		return uint(obj.Value("data").Object().Value("id").Number().Raw())
	}
	return 0
}

func Update(auth *httpexpect.Expect, url string, update map[string]interface{}, id uint, status int, message string) {
	obj := auth.PUT(url).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}

func Get(auth *httpexpect.Expect, url string, update map[string]interface{}, id uint, keys ResponseKeys, status int, message string) {
	obj := auth.GET(url).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
	if id > 0 {
		brandCategory := obj.Value("data").Object()
		keys.Test(brandCategory, id, update)
	}
}

func Post(auth *httpexpect.Expect, url string, update map[string]interface{}, keys ResponseKeys, status int, message string) {
	obj := auth.POST(url).
		WithJSON(update).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}

func Delete(auth *httpexpect.Expect, url string, update map[string]interface{}, keys ResponseKeys, status int, message string) {
	obj := auth.DELETE(url).
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(status)
	obj.Value("message").String().Equal(message)
}
