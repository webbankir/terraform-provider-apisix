package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginExtPluginPostReqType struct {
	Disable types.Bool                       `tfsdk:"disable"`
	Config  []PluginExtPluginPostReqConfType `tfsdk:"conf"`
}

type PluginExtPluginPostReqConfType struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var PluginExtPluginPostReqSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"disable": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"conf": {
			Required: true,
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
				"name": {
					Required: true,
					Type:     types.StringType,
				},
				"value": {
					Required: true,
					Type:     types.StringType,
				},
			}, tfsdk.ListNestedAttributesOptions{MinItems: 1}),
		},
	}),
}

func (s PluginExtPluginPostReqType) Name() string { return "ext-plugin-pre-req" }

func (s PluginExtPluginPostReqType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginExtPluginPostReqType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)

	var subItems []PluginExtPluginPostReqConfType
	for _, vv := range jsonData["conf"].([]interface{}) {
		subItem := PluginExtPluginPostReqConfType{}
		subV := vv.(map[string]interface{})
		utils.MapValueToStringTypeValue(subV, "name", &subItem.Name)
		utils.MapValueToStringTypeValue(subV, "value", &subItem.Value)
		subItems = append(subItems, subItem)
	}

	item.Config = subItems
	pluginsType.ExtPluginPostReqType = &item
}

func (s PluginExtPluginPostReqType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")

	var subItems []map[string]interface{}
	for _, vv := range s.Config {
		subItem := make(map[string]interface{})
		utils.StringTypeValueToMap(vv.Name, subItem, "name")
		utils.StringTypeValueToMap(vv.Value, subItem, "value")
	}

	pluginValue["config"] = subItems

	m[s.Name()] = pluginValue
}
