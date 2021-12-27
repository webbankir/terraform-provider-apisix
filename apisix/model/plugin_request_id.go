package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginRequestIdType struct {
	Disable           types.Bool   `tfsdk:"disable"`
	HeaderName        types.String `tfsdk:"header_name"`
	IncludeInResponse types.Bool   `tfsdk:"include_in_response"`
	Algorithm         types.String `tfsdk:"algorithm"`
}

var PluginRequestIdSchemaAttribute = tfsdk.Attribute{
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
		"header_name": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("X-Request-Id"),
			},
		},
		"include_in_response": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(true),
			},
		},
		"algorithm": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("uuid", "snowflake"),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("uuid"),
			},
		},
	}),
}

func (s PluginRequestIdType) Name() string { return "request-id" }

func (s PluginRequestIdType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginRequestIdType{}

	utils.MapValueToValue(jsonData, "disable", &item.Disable)
	utils.MapValueToValue(jsonData, "header_name", &item.HeaderName)
	utils.MapValueToValue(jsonData, "include_in_response", &item.IncludeInResponse)
	utils.MapValueToValue(jsonData, "algorithm", &item.Algorithm)

	pluginsType.RequestId = &item
}

func (s PluginRequestIdType) StateToMap(m map[string]interface{}, isUpdate bool) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.HeaderName, pluginValue, "header_name", isUpdate)
	utils.ValueToMap(s.IncludeInResponse, pluginValue, "include_in_response", isUpdate)
	utils.ValueToMap(s.Algorithm, pluginValue, "algorithm", isUpdate)

	m[s.Name()] = pluginValue
}
