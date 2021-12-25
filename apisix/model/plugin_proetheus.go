package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginPrometheusType struct {
	Disable    types.Bool `tfsdk:"disable"`
	PreferName types.Bool `tfsdk:"prefer_name"`
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

func (s PluginPrometheusType) DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType) {
	if v := v[s.Name()]; v != nil {
		jsonData := v.(map[string]interface{})
		item := PluginPrometheusType{}

		if v := jsonData["disable"]; v != nil {
			item.Disable = types.Bool{Value: v.(bool)}
		} else {
			item.Disable = types.Bool{Value: true}
		}

		if v := jsonData["prefer_name"]; v != nil {
			item.PreferName = types.Bool{Value: v.(bool)}
		} else {
			item.PreferName = types.Bool{Null: true}
		}

		pluginsType.Prometheus = &item
	}
}

func (s PluginPrometheusType) validate() error { return nil }

func (s PluginPrometheusType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.PreferName, pluginValue, "prefer_name", true)

	m[s.Name()] = pluginValue
}
