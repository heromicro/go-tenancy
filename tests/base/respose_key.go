package base

import (
	"reflect"
	"strings"

	"github.com/gavv/httpexpect"
)

type Param struct {
	Name         string
	Args         map[string]interface{}
	ResponseKeys ResponseKeys
}

type ResponseKeys []ResponseKey
type ResponseKey struct {
	Key   string
	Value interface{}
	Type  string
}

func (rks ResponseKeys) Keys() []string {
	keys := []string{}
	for _, rk := range rks {
		keys = append(keys, rk.Key)
	}
	return keys
}

func IdKeys() ResponseKeys {
	return ResponseKeys{
		{Key: "id", Value: uint(0)},
	}
}

func (rks ResponseKeys) Test(object *httpexpect.Object) {
	for _, rk := range rks {
		if rk.Value == nil {
			continue
		}
		switch reflect.TypeOf(rk.Value).String() {
		case "string":
			if strings.ToLower(rk.Type) == "notempty" {
				object.Value(rk.Key).String().NotEmpty()
			} else {
				object.Value(rk.Key).String().Equal(rk.Value.(string))
			}
		case "float64":
			if strings.ToLower(rk.Type) == "ge" {
				object.Value(rk.Key).Number().Ge(rk.Value.(float64))
			} else {
				object.Value(rk.Key).Number().Equal(rk.Value.(float64))
			}
		case "uint":
			if strings.ToLower(rk.Type) == "ge" {
				object.Value(rk.Key).Number().Ge(rk.Value.(uint))
			} else {
				object.Value(rk.Key).Number().Equal(rk.Value.(uint))
			}
		case "int":
			if strings.ToLower(rk.Type) == "ge" {
				object.Value(rk.Key).Number().Ge(rk.Value.(int))
			} else {
				object.Value(rk.Key).Number().Equal(rk.Value.(int))
			}
		case "[]base.ResponseKeys":
			length := int(object.Value(rk.Key).Array().Length().Raw())
			if length > 0 && len(rk.Value.([]ResponseKeys)) == length {
				for i := 0; i < length; i++ {
					rk.Value.([]ResponseKeys)[i].Test(object.Value(rk.Key).Array().Element(i).Object())
				}
			}
		case "[]uint":
			object.Value(rk.Key).Array().Length().Equal(len(rk.Value.([]uint)))
			length := int(object.Value(rk.Key).Array().Length().Raw())
			if length > 0 && len(rk.Value.([]uint)) == length {
				for i := 0; i < length; i++ {
					object.Value(rk.Key).Array().Element(i).Number().Equal(rk.Value.([]uint)[i])
				}
			}
		case "[]string":
			object.Value(rk.Key).Array().Length().Equal(len(rk.Value.([]string)))
			length := int(object.Value(rk.Key).Array().Length().Raw())
			if length > 0 && len(rk.Value.([]string)) == length {
				for i := 0; i < length; i++ {
					object.Value(rk.Key).Array().Element(i).String().Equal(rk.Value.([]string)[i])
				}
			}

		default:
			continue
		}
	}
}

func (rks ResponseKeys) Scan(object *httpexpect.Object) {
	for k, rk := range rks {
		if !Exist(object, rk.Key) {
			continue
		}
		if rk.Value == nil {
			continue
		}
		switch reflect.TypeOf(rk.Value).String() {
		case "string":
			rks[k].Value = object.Value(rk.Key).String().Raw()
		case "uint":
			rks[k].Value = uint(object.Value(rk.Key).Number().Raw())
		case "int":
			rks[k].Value = int(object.Value(rk.Key).Number().Raw())
		case "int32":
			rks[k].Value = int32(object.Value(rk.Key).Number().Raw())
		case "float64":
			rks[k].Value = object.Value(rk.Key).Number().Raw()
		case "[]base.ResponseKeys":
			length := len(rk.Value.([]ResponseKeys))
			for i := 0; i < length; i++ {
				rk.Value.([]ResponseKeys)[i].Scan(object.Value(rk.Key).Array().Element(i).Object())
			}
		case "[]string":
			length := int(object.Value(rk.Key).Array().Length().Raw())

			if length == 0 {
				continue
			}
			reskey, ok := rks[k].Value.([]string)
			if ok {
				var strings []string
				for i := 0; i < length; i++ {
					strings = append(reskey, object.Value(rk.Key).Array().Element(i).String().Raw())
				}
				rks[k].Value = strings
			}
		default:
			continue
		}
	}
}

func Exist(object *httpexpect.Object, key string) bool {
	objectKyes := object.Keys().Raw()
	for _, objectKey := range objectKyes {
		if key == objectKey.(string) {
			return true
		}
	}
	return false
}

func (rks ResponseKeys) GetStringValue(key string) string {

	if strings.Contains(key, ".") {
		keys := strings.Split(key, ".")
		if len(keys) != 2 {
			return ""
		}
		res := rks.GetResponseKeysValue(keys[0])
		if len(res) == 0 {
			return ""
		}
		return res[0].GetStringValue(keys[1])
	}

	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return ""
			}
			switch reflect.TypeOf(rk.Value).String() {
			case "string":
				return rk.Value.(string)
			}
		}
	}
	return ""
}

func (rks ResponseKeys) GetStringArrayValue(key string) []string {
	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return nil
			}
			switch reflect.TypeOf(rk.Value).String() {
			case "[]string":
				return rk.Value.([]string)
			}
		}
	}
	return nil
}

func (rks ResponseKeys) GetResponseKeysValue(key string) []ResponseKeys {
	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return nil
			}
			switch reflect.TypeOf(rk.Value).String() {
			case "[]base.ResponseKeys":
				return rk.Value.([]ResponseKeys)
			}
		}
	}
	return nil
}

func (rks ResponseKeys) GetUintValue(key string) uint {
	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return 0
			}
			switch reflect.TypeOf(rk.Value).String() {
			case "float64":
				return uint(rk.Value.(float64))
			case "int32":
				return uint(rk.Value.(int32))
			case "uint":
				return rk.Value.(uint)
			case "int":
				return uint(rk.Value.(int))
			}
		}
	}
	return 0
}

func (rks ResponseKeys) GetIntValue(key string) int {
	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return 0
			}
			switch reflect.TypeOf(rk.Value).String() {
			case "float64":
				return int(rk.Value.(float64))
			case "int":
				return rk.Value.(int)
			case "int32":
				return int(rk.Value.(int32))
			case "uint":
				return int(rk.Value.(uint))
			}
		}
	}
	return 0
}
func (rks ResponseKeys) GetInt32Value(key string) int32 {
	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return 0
			}
			switch reflect.TypeOf(rk.Value).String() {
			case "float64":
				return int32(rk.Value.(float64))
			case "int32":
				return rk.Value.(int32)
			case "int":
				return int32(rk.Value.(int))
			case "uint":
				return int32(rk.Value.(uint))
			}
		}
	}
	return 0
}

func (rks ResponseKeys) GetId() uint {
	for _, rk := range rks {
		if rk.Key == "id" {
			if rk.Value == nil {
				return 0
			}
			switch reflect.TypeOf(rk.Value).String() {
			case "float64":
				return uint(rk.Value.(float64))
			case "uint":
				return rk.Value.(uint)
			case "int":
				return uint(rk.Value.(int))
			}
		}
	}
	return 0
}
