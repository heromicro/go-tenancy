package base

import (
	"strings"

	"github.com/gavv/httpexpect"
)

type ResponseKeys []ResponseKey
type ResponseKey struct {
	Type  string
	Key   string
	Value interface{}
}

func (rks ResponseKeys) Keys() []string {
	keys := []string{}
	for _, rk := range rks {
		keys = append(keys, rk.Key)
	}
	return keys
}

func (rks ResponseKeys) Test(object *httpexpect.Object) {
	for _, rk := range rks {
		object.Keys().Contains(rk.Key)
		switch strings.ToLower(rk.Type) {
		case "string":
			object.Value(rk.Key).String().Equal(rk.Value.(string))
		case "number":
			object.Value(rk.Key).Number().Equal(rk.Value.(int))
		case "object":
			continue
		case "array":
			if rk.Value == nil {
				continue
			}
			subs := rk.Value.([]ResponseKeys)
			object.Value(rk.Key).Array().Length().Equal(len(subs))
			length := int(object.Value(rk.Key).Array().Length().Raw())
			if length > 0 && len(subs) == length {
				for i := 0; i < length; i++ {
					subs[i].Test(object.Value(rk.Key).Array().Element(i).Object())
				}
			}

		default:
			object.Value(rk.Key).String().Equal(rk.Value.(string))
		}
	}
}

func (rks ResponseKeys) Scan(object *httpexpect.Object) {
	for _, rk := range rks {
		switch strings.ToLower(rk.Type) {
		case "string":
			rk.Value = object.Value(rk.Key).String().Raw()
		case "number":
			rk.Value = object.Value(rk.Key).Number().Raw()
		case "object":
			continue
		default:
			rk.Value = object.Value(rk.Key).String().Raw()
		}
	}
}

func (rks ResponseKeys) GetStringValue(key string) string {
	for _, rk := range rks {
		if key == rk.Key {
			switch strings.ToLower(rk.Type) {
			case "string":
				return rk.Value.(string)
			}
		}
	}
	return ""
}
func (rks ResponseKeys) GetUintValue(key string) uint {
	for _, rk := range rks {
		if key == rk.Key {
			switch strings.ToLower(rk.Type) {
			case "number":
				return uint(rk.Value.(int))
			}
		}
	}
	return 0
}
