package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginBasicAuthType struct {
	Disable types.Bool `tfsdk:"disable"`
}

var PluginBasicAuthSchemaAttribute = tfsdk.Attribute{
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
	}),
}

func (s PluginBasicAuthType) Name() string { return "basic-auth" }

func (s PluginBasicAuthType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginBasicAuthType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)

	pluginsType.BasicAuth = &item
}

func (s PluginBasicAuthType) StateToMap(m map[string]interface{}) {
	var pluginValue = make(map[string]interface{})
	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	m[s.Name()] = pluginValue
}
