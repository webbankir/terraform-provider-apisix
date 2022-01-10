package model

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginCustomType struct {
	Disable        types.Bool   `tfsdk:"disable"`
	Name           types.String `tfsdk:"name"`
	JSON           types.String `tfsdk:"json"`
	JSONFromServer types.String `tfsdk:"json_from_server"`
}

var PluginCustomSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"disable": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"name": {
			Required: true,
			Type:     types.StringType,
		},
		"json": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("{}"),
			},
		},
		"json_from_server": {
			Computed: true,
			Type:     types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{}),
}

func PluginCustomTypeMapToState(data map[string]interface{}, pluginsType *PluginsType, plan *RouteType, state *RouteType) {
	var (
		knownPlugins []string
		items        []PluginCustomType
	)

	for k, _ := range PluginsSchemaAttribute.GetAttributes() {
		if k != "custom" {
			knownPlugins = append(knownPlugins, k)
		}
	}

	for k, v := range data {
		if !utils.StringContainsInSlice(knownPlugins, k) {
			item := PluginCustomType{}
			value := v.(map[string]interface{})

			item.Name = types.String{Value: k}
			utils.MapValueToBoolTypeValue(value, "disable", &item.Disable)
			delete(value, "name")
			delete(value, "disable")
			// FIXME:
			bb, _ := json.Marshal(value)

			jsonFromState := ""

			if plan != nil && plan.Plugins.Custom != nil {
				for _, v := range *(plan.Plugins.Custom) {
					if v.Name.Value == k {
						jsonFromState = v.JSON.Value
					}
				}
			}

			if jsonFromState == "" && state != nil && state.Plugins.Custom != nil {
				for _, v := range *(state.Plugins.Custom) {
					if v.Name.Value == k {
						jsonFromState = v.JSON.Value
					}
				}
			}

			item.JSONFromServer = types.String{Value: string(bb)}
			item.JSON = types.String{Value: jsonFromState}
			items = append(items, item)
		}
	}

	if len(items) > 0 {
		pluginsType.Custom = &items
	} else {
		pluginsType.Custom = nil
	}
}

func PluginCustomTypeStateToMap(m map[string]interface{}, plan RouteType, state *RouteType, isUpdate bool) {

	if state != nil {
		if state.Plugins.Custom != nil {
			for _, v := range *(state.Plugins.Custom) {
				m[v.Name.Value] = nil
			}
		}
	}

	if plan.Plugins.Custom != nil {
		for _, v := range *(plan.Plugins.Custom) {
			jsonData := make(map[string]interface{})
			// FIXME:
			_ = json.Unmarshal([]byte(v.JSON.Value), &jsonData)
			jsonData["disable"] = v.Disable.Value

			m[v.Name.Value] = jsonData
		}
	}
}
