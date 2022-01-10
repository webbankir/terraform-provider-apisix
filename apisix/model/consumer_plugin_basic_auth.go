package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type ConsumerPluginBasicAuthType struct {
	Disable  types.Bool   `tfsdk:"disable"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

var ConsumerPluginBasicAuthSchemaAttribute = tfsdk.Attribute{
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
		"username": {
			Required: true,
			Type:     types.StringType,
		},
		"password": {
			Required: true,
			Type:     types.StringType,
		},
	}),
}

func (s ConsumerPluginBasicAuthType) Name() string { return "basic-auth" }

func (s ConsumerPluginBasicAuthType) MapToState(data map[string]interface{}, pluginsType *ConsumerPluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := ConsumerPluginBasicAuthType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "username", &item.Username)
	utils.MapValueToStringTypeValue(jsonData, "password", &item.Password)

	pluginsType.BasicAuth = &item
}

func (s ConsumerPluginBasicAuthType) StateToMap(m map[string]interface{}, _ bool) {
	var pluginValue = make(map[string]interface{})

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable", false)
	utils.StringTypeValueToMap(s.Username, pluginValue, "username", false)
	utils.StringTypeValueToMap(s.Password, pluginValue, "password", false)

	m[s.Name()] = pluginValue
}
