package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringTypeValueToMap(value types.String, dMap map[string]interface{}, mapKey string) {
	if !value.Null {
		dMap[mapKey] = value.Value
	}
}

func BoolTypeValueToMap(value types.Bool, dMap map[string]interface{}, mapKey string) {
	if !value.Null {
		dMap[mapKey] = value.Value
	}
}

func NumberTypeValueToMap(value types.Number, dMap map[string]interface{}, mapKey string) {
	if !value.Null {
		dMap[mapKey] = TypeNumberToInt(value)
	}
}

func ListTypeValueToMap(value types.List, dMap map[string]interface{}, mapKey string) {

	switch value.ElemType {
	case types.StringType:
		if !value.Null {
			var values []string
			for _, v := range value.Elems {
				values = append(values, v.(types.String).Value)
			}
			dMap[mapKey] = values
		}
	case types.NumberType:
		if !value.Null {
			var values []int
			for _, v := range value.Elems {
				values = append(values, TypeNumberToInt(v.(types.Number)))
			}
			dMap[mapKey] = values
		}
	default:
		panic(value.ElemType)
	}
}

func MapTypeValueToMap(value types.Map, dMap map[string]interface{}, mapKey string) {
	switch value.ElemType {
	case types.StringType:
		if !value.Null {
			values := make(map[string]interface{})
			for k, v := range value.Elems {
				values[k] = v.(types.String).Value
			}
			dMap[mapKey] = values
		}
	default:
		panic("WTF")
	}
}
