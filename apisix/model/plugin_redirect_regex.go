package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginRedirectRegexType struct {
	Disable    types.Bool   `tfsdk:"disable"`
	Pattern    types.String `tfsdk:"pattern"`
	Replace    types.String `tfsdk:"replace"`
	StatusCode types.Number `tfsdk:"status_code"`
}

var PluginRedirectRegexSchemaAttribute = tfsdk.Attribute{
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
		"pattern": {
			Required: true,
			Type:     types.StringType,
		},
		"replace": {
			Required: true,
			Type:     types.StringType,
		},
		"status_code": {
			Optional: true,
			Computed: true,
			Type:     types.NumberType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(301),
			},
		},
	}),
}

func (s PluginRedirectRegexType) Name() string { return "redirect-regex" }

func (s PluginRedirectRegexType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginRedirectRegexType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "pattern", &item.Pattern)
	utils.MapValueToStringTypeValue(jsonData, "replace", &item.Replace)
	utils.MapValueToNumberTypeValue(jsonData, "status_code", &item.StatusCode)

	pluginsType.RedirectRegex = &item
}

func (s PluginRedirectRegexType) StateToMap(m map[string]interface{}, _ bool) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.StringTypeValueToMap(s.Pattern, pluginValue, "pattern", false)
	utils.StringTypeValueToMap(s.Replace, pluginValue, "replace", false)
	utils.NumberTypeValueToMap(s.StatusCode, pluginValue, "status_code", false)

	m[s.Name()] = pluginValue
}
