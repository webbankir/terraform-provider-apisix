package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
)

func MapValueToStringTypeValue(sMap map[string]interface{}, mapKey string, value *types.String) {
	v := sMap[mapKey]

	if v == nil {
		value.Null = true
	} else {
		value.Value = v.(string)
	}
}

func MapValueToBoolTypeValue(sMap map[string]interface{}, mapKey string, value *types.Bool) {
	v := sMap[mapKey]

	if v == nil {
		value.Null = true
	} else {
		value.Value = v.(bool)
	}
}

func MapValueToNumberTypeValue(sMap map[string]interface{}, mapKey string, value *types.Number) {
	v := sMap[mapKey]

	if v == nil {
		value.Null = true
	} else {
		value.Value = big.NewFloat(v.(float64))
	}
}

func MapValueToListTypeValue(sMap map[string]interface{}, mapKey string, value *types.List) {
	v := sMap[mapKey]

	if v == nil {
		value.Null = true
	} else {
		var values []attr.Value
		for _, vv := range v.([]interface{}) {
			switch dd := vv.(type) {
			case string:
				value.ElemType = types.StringType
				values = append(values, types.String{Value: dd})
			default:
				panic("WTF")
			}
		}
		value.Elems = values
	}
}

func MapValueToMapTypeValue(sMap map[string]interface{}, mapKey string, value *types.Map) {
	v := sMap[mapKey]

	if v == nil {
		value.Null = true
	} else {
		values := make(map[string]attr.Value)
		for key, vv := range v.(map[string]interface{}) {
			switch dd := vv.(type) {
			case string:
				value.ElemType = types.StringType
				values[key] = types.String{Value: dd}
			default:
				panic("WTF")
			}

		}
		value.Elems = values
	}
}
