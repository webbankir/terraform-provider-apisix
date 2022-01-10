package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func varsMapToState(data map[string]interface{}) types.List {

	if v := data["vars"]; v != nil {

		data, ok := v.([]interface{})
		if !ok {
			return types.List{Null: true, ElemType: types.ListType{ElemType: types.StringType}}
		} else {
			var values []attr.Value

			for _, v := range data {
				var subValues []attr.Value

				subData, subOk := v.([]interface{})

				if subOk {
					for _, sv := range subData {
						subValues = append(subValues, types.String{Value: sv.(string)})
					}
				}

				values = append(values, types.List{ElemType: types.StringType, Elems: subValues})
			}
			return types.List{Elems: values, ElemType: types.ListType{ElemType: types.StringType}}
		}
	}

	//log.Printf("[DEBUG] KUKU - %v", data["vars"])
	return types.List{Null: true, ElemType: types.ListType{ElemType: types.StringType}}
}

func varsStateToMap(state types.List, jsonMap map[string]interface{}, isUpdate bool) {
	if !state.Null {
		var values = make([][]string, 0)
		for _, v := range state.Elems {
			var subValues []string
			for _, sv := range v.(types.List).Elems {
				subValues = append(subValues, sv.(types.String).Value)
			}
			values = append(values, subValues)
		}

		jsonMap["vars"] = values
	} else if isUpdate {
		jsonMap["vars"] = nil
	}
}
