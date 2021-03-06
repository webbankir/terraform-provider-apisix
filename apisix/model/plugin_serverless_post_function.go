package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginServerlessPostFunctionType struct {
	Disable   types.Bool   `tfsdk:"disable"`
	Phase     types.String `tfsdk:"phase"`
	Functions types.List   `tfsdk:"functions"`
}

var PluginServerlessPostFunctionSchemaAttribute = tfsdk.Attribute{
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
		"phase": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("access"),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("rewrite", "access", "header_filter", "body_filter", "log", "balancer"),
			},
		},
		"functions": {
			Required: true,
			Type:     types.ListType{ElemType: types.StringType},
		},
	}),
}

func (s PluginServerlessPostFunctionType) Name() string { return "serverless-post-function" }

func (s PluginServerlessPostFunctionType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginServerlessPostFunctionType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "phase", &item.Phase)
	utils.MapValueToListTypeValue(jsonData, "functions", &item.Functions)

	pluginsType.ServerlessPostFunction = &item

}

func (s PluginServerlessPostFunctionType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.StringTypeValueToMap(s.Phase, pluginValue, "phase")
	utils.ListTypeValueToMap(s.Functions, pluginValue, "functions")

	m[s.Name()] = pluginValue
}
