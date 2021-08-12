package base

import (
	"strings"

	"github.com/gavv/httpexpect"
)

type ResponseKeys map[string]string

func (rk ResponseKeys) Keys() []string {
	keys := []string{}
	for k := range rk {
		keys = append(keys, k)
	}
	return keys
}

func (rk ResponseKeys) Test(object *httpexpect.Object, id uint, data map[string]interface{}) {
	object.Keys().ContainsOnly(rk.Keys())
	for k, v := range rk {
		switch strings.ToLower(v) {
		case "string":
			object.Value(k).String().Equal(data[k].(string))
		case "number":
			object.Value(k).Number().Equal(data[k].(int))
		case "object":
			continue
		default:
			object.Value(k).String().Equal(data[k].(string))
		}
	}
}
