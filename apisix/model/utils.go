package model

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func varsMapToState(data map[string]interface{}) types.String {
	if v := data["vars"]; v != nil {
		//FIXME:
		d, _ := json.Marshal(v)
		return types.String{Value: string(d)}
	}
	return types.String{Null: true}

}

func varsStateToMap(state types.String, jsonMap map[string]interface{}, isUpdate bool) {

	if !state.Null {
		jj := make([]interface{}, 0)
		//FIXME:
		_ = json.Unmarshal([]byte(state.Value), &jj)

		jsonMap["vars"] = jj
	} else if isUpdate {
		jsonMap["vars"] = nil
	}
}
