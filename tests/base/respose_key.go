package base

import (
	"strings"

	"github.com/gavv/httpexpect"
)

var IdKeys = ResponseKeys{
	{Type: "uint", Key: "id", Value: uint(0)},
}

type Param struct {
	Name         string
	Args         map[string]interface{}
	ResponseKeys ResponseKeys
}

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
		if rk.Value == nil {
			continue
		}
		switch strings.ToLower(rk.Type) {
		case "string":
			object.Value(rk.Key).String().Equal(rk.Value.(string))
		case "float64":
			object.Value(rk.Key).Number().Equal(rk.Value.(float64))
		case "uint":
			object.Value(rk.Key).Number().Equal(rk.Value.(uint))
		case "int":
			object.Value(rk.Key).Number().Equal(rk.Value.(int))
		case "object":
			continue
		case "array":
			if subs, ok := rk.Value.([]ResponseKeys); ok {
				object.Value(rk.Key).Array().Length().Equal(len(subs))
				length := int(object.Value(rk.Key).Array().Length().Raw())
				if length > 0 && len(subs) == length {
					for i := 0; i < length; i++ {
						subs[i].Test(object.Value(rk.Key).Array().Element(i).Object())
					}
				}
			} else if subs, ok := rk.Value.([]uint); ok {
				object.Value(rk.Key).Array().Length().Equal(len(subs))
				length := int(object.Value(rk.Key).Array().Length().Raw())
				if length > 0 && len(subs) == length {
					for i := 0; i < length; i++ {
						object.Value(rk.Key).Array().Element(i).Number().Equal(subs[i])
					}
				}
			} else if subs, ok := rk.Value.([]string); ok {
				object.Value(rk.Key).Array().Length().Equal(len(subs))
				length := int(object.Value(rk.Key).Array().Length().Raw())
				if length > 0 && len(subs) == length {
					for i := 0; i < length; i++ {
						object.Value(rk.Key).Array().Element(i).String().Equal(subs[i])
					}
				}
			}
		case "notempty":
			object.Value(rk.Key).String().NotEmpty()
		default:
			object.Value(rk.Key).String().Equal(rk.Value.(string))
		}
	}
}

func (rks ResponseKeys) Scan(object *httpexpect.Object) {
	for k, rk := range rks {
		switch strings.ToLower(rk.Type) {
		case "string":
			rks[k].Value = object.Value(rk.Key).String().Raw()
		case "uint":
			rks[k].Value = uint(object.Value(rk.Key).Number().Raw())
		case "int":
			rks[k].Value = int(object.Value(rk.Key).Number().Raw())
		case "float64":
			rks[k].Value = object.Value(rk.Key).Number().Raw()
		case "array":
			length := int(object.Value(rk.Key).Array().Length().Raw())
			if length == 0 {
				continue
			}
			if reskey, ok := rks[k].Value.([]string); ok {
				for i := 0; i < length; i++ {
					reskey = append(reskey, object.Value(rk.Key).Array().Element(i).String().Raw())
				}
				rks[k].Value = reskey
			}

		case "object":
			continue
		default:
			rks[k].Value = object.Value(rk.Key).String().Raw()
		}
	}
}

func (rks ResponseKeys) GetStringValue(key string) string {
	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return ""
			}
			switch strings.ToLower(rk.Type) {
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
			switch strings.ToLower(rk.Type) {
			case "string":
				return rk.Value.([]string)
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
			switch strings.ToLower(rk.Type) {
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

func (rks ResponseKeys) GetIntValue(key string) int {
	for _, rk := range rks {
		if key == rk.Key {
			if rk.Value == nil {
				return 0
			}
			switch strings.ToLower(rk.Type) {
			case "float64":
				return int(rk.Value.(float64))
			case "int":
				return rk.Value.(int)
			case "uint":
				return int(rk.Value.(uint))
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
			switch strings.ToLower(rk.Type) {
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
