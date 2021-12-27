package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginServerlessPreFunctionType struct {
	Disable   types.Bool   `tfsdk:"disable"`
	Phase     types.String `tfsdk:"phase"`
	Functions types.List   `tfsdk:"functions"`
}

var PluginServerlessPreFunctionSchemaAttribute = tfsdk.Attribute{
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

func (s PluginServerlessPreFunctionType) Name() string { return "serverless-pre-function" }

func (s PluginServerlessPreFunctionType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}

	jsonData := v.(map[string]interface{})
	item := PluginServerlessPreFunctionType{}

	utils.MapValueToValue(jsonData, "disable", &item.Disable)
	utils.MapValueToValue(jsonData, "phase", &item.Phase)
	utils.MapValueToValue(jsonData, "functions", &item.Functions)

	pluginsType.ServerlessPreFunction = &item
}

func (s PluginServerlessPreFunctionType) StateToMap(m map[string]interface{}, isUpdate bool) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.Phase, pluginValue, "phase", isUpdate)
	utils.ValueToMap(s.Functions, pluginValue, "functions", isUpdate)

	m[s.Name()] = pluginValue
}
