package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginRedirectRegexType struct {
	Disable  types.Bool                       `tfsdk:"disable"`
	Variants []PluginRedirectRegexVariantType `tfsdk:"variants"`
}

type PluginRedirectRegexVariantType struct {
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
		"variants": {
			Required: true,
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
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
			},
				tfsdk.ListNestedAttributesOptions{MinItems: 1},
			),
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

	var subItems []PluginRedirectRegexVariantType

	for _, variant := range jsonData["variants"].([]interface{}) {
		subItem := PluginRedirectRegexVariantType{}
		utils.MapValueToStringTypeValue(variant.(map[string]interface{}), "pattern", &subItem.Pattern)
		utils.MapValueToStringTypeValue(variant.(map[string]interface{}), "replace", &subItem.Replace)
		utils.MapValueToNumberTypeValue(variant.(map[string]interface{}), "status_code", &subItem.StatusCode)
		subItems = append(subItems, subItem)
	}
	item.Variants = subItems

	pluginsType.RedirectRegex = &item
}

func (s PluginRedirectRegexType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}
	var variants []map[string]interface{}
	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	for _, v := range s.Variants {
		variant := map[string]interface{}{}
		utils.StringTypeValueToMap(v.Pattern, variant, "pattern")
		utils.StringTypeValueToMap(v.Replace, variant, "replace")
		utils.NumberTypeValueToMap(v.StatusCode, variant, "status_code")
		variants = append(variants, variant)
	}
	pluginValue["variants"] = variants

	m[s.Name()] = pluginValue
}
