package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringTypeValueToMap(value types.String, dMap map[string]interface{}, mapKey string, bindAsNil bool) {
	if !value.Null {
		dMap[mapKey] = value.Value
	} else if bindAsNil {
		dMap[mapKey] = nil
	}
}

func BoolTypeValueToMap(value types.Bool, dMap map[string]interface{}, mapKey string, bindAsNil bool) {
	if !value.Null {
		dMap[mapKey] = value.Value
	} else if bindAsNil {
		dMap[mapKey] = nil
	}
}

func NumberTypeValueToMap(value types.Number, dMap map[string]interface{}, mapKey string, bindAsNil bool) {
	if !value.Null {
		dMap[mapKey] = TypeNumberToInt(value)
	} else if bindAsNil {
		dMap[mapKey] = nil
	}
}

func ListTypeValueToMap(value types.List, dMap map[string]interface{}, mapKey string, bindAsNil bool) {

	switch value.ElemType {
	case types.StringType:
		if !value.Null {
			var values []string
			for _, v := range value.Elems {
				values = append(values, v.(types.String).Value)
			}
			dMap[mapKey] = values
		} else if bindAsNil {
			dMap[mapKey] = nil
		}
	default:
		panic("WTF")
	}
}

func MapTypeValueToMap(value types.Map, dMap map[string]interface{}, mapKey string, bindAsNil bool) {
	switch value.ElemType {
	case types.StringType:
		if !value.Null {
			values := make(map[string]interface{})
			for k, v := range value.Elems {
				values[k] = v.(types.String).Value
			}
			dMap[mapKey] = values
		} else if bindAsNil {
			dMap[mapKey] = nil
		}
	default:
		panic("WTF")
	}
}
