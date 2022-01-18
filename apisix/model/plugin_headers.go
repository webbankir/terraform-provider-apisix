package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginHeadersType struct {
	Disable  types.Bool            `tfsdk:"disable"`
	Request  types.Map             `tfsdk:"request"`
	Response types.Map             `tfsdk:"response"`
	STS      *PluginHeadersSTSType `tfsdk:"sts"`
}

type PluginHeadersSTSType struct {
	MaxAge            types.Number `tfsdk:"max_age"`
	IncludeSubDomains types.Bool   `tfsdk:"include_sub_domains"`
	Preload           types.Bool   `tfsdk:"preload"`
}

var PluginHeadersSchemaAttribute = tfsdk.Attribute{
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
		"request": {
			Optional: true,
			Type:     types.MapType{ElemType: types.StringType},
		},
		"response": {
			Optional: true,
			Type:     types.MapType{ElemType: types.StringType},
		},
		"sts": {
			Optional: true,
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
				"max_age": {
					Optional: true,
					Computed: true,
					Type:     types.NumberType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultNumber(31536000),
					},
				},
				"include_sub_domains": {
					Optional: true,
					Computed: true,
					Type:     types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultBool(true),
					},
				},
				"preload": {
					Optional: true,
					Computed: true,
					Type:     types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultBool(true),
					},
				},
			}),
		},
	}),
}

func (s PluginHeadersType) Name() string { return "headers" }

func (s PluginHeadersType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}

	jsonData := v.(map[string]interface{})
	item := PluginHeadersType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToMapTypeValue(jsonData, "request", &item.Request)
	utils.MapValueToMapTypeValue(jsonData, "response", &item.Response)

	if v := jsonData["sts"]; v != nil {
		subJsonData := v.(map[string]interface{})
		subItem := PluginHeadersSTSType{}

		utils.MapValueToNumberTypeValue(subJsonData, "max_age", &subItem.MaxAge)
		utils.MapValueToBoolTypeValue(subJsonData, "include_sub_domains", &subItem.IncludeSubDomains)
		utils.MapValueToBoolTypeValue(subJsonData, "preload", &subItem.Preload)

		item.STS = &subItem
	}

	pluginsType.Headers = &item
}

func (s PluginHeadersType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.MapTypeValueToMap(s.Request, pluginValue, "request")
	utils.MapTypeValueToMap(s.Response, pluginValue, "response")

	if v := s.STS; v != nil {
		subPluginValue := map[string]interface{}{}
		utils.NumberTypeValueToMap(v.MaxAge, subPluginValue, "max_age")
		utils.BoolTypeValueToMap(v.IncludeSubDomains, subPluginValue, "include_sub_domains")
		utils.BoolTypeValueToMap(v.Preload, subPluginValue, "preload")
		pluginValue["sts"] = subPluginValue
	}

	m[s.Name()] = pluginValue
}
