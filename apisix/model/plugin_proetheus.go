package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginPrometheusType struct {
	Disable    types.Bool   `tfsdk:"disable"`
	PreferName types.String `tfsdk:"prefer_name"`
}

var PluginPrometheusSchemaAttribute = tfsdk.Attribute{
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
		"prefer_name": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(true),
			},
		},
	}),
}

func (s PluginPrometheusType) Name() string { return "prometheus" }

func (s PluginPrometheusType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginPrometheusType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "prefer_name", &item.PreferName)

	pluginsType.Prometheus = &item
}

func (s PluginPrometheusType) StateToMap(m map[string]interface{}, _ bool) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.StringTypeValueToMap(s.PreferName, pluginValue, "prefer_name", false)

	m[s.Name()] = pluginValue
}
