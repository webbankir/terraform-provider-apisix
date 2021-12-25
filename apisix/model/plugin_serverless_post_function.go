package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

func (s PluginServerlessPostFunctionType) DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType) {
	if v := v[s.Name()]; v != nil {
		jsonData := v.(map[string]interface{})
		item := PluginServerlessPostFunctionType{}

		if v := jsonData["disable"]; v != nil {
			item.Disable = types.Bool{Value: v.(bool)}
		} else {
			item.Disable = types.Bool{Value: true}
		}

		if v := jsonData["phase"]; v != nil {
			item.Phase = types.String{Value: v.(string)}
		} else {
			item.Phase = types.String{Null: true}
		}

		if v := jsonData["functions"]; v != nil {
			var values []attr.Value
			for _, value := range v.([]interface{}) {
				values = append(values, types.String{Value: value.(string)})
			}

			item.Functions = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.Functions = types.List{Null: true}
		}

		pluginsType.ServerlessPostFunction = &item
	}
}

func (s PluginServerlessPostFunctionType) validate() error { return nil }

func (s PluginServerlessPostFunctionType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.Phase, pluginValue, "phase", true)
	utils.ValueToMap(s.Functions, pluginValue, "functions", true)

	m[s.Name()] = pluginValue
}
