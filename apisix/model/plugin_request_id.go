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

func (s PluginRequestIdType) DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType) {
	if v := v[s.Name()]; v != nil {
		jsonData := v.(map[string]interface{})
		item := PluginRequestIdType{}

		if v := jsonData["disable"]; v != nil {
			item.Disable = types.Bool{Value: v.(bool)}
		} else {
			item.Disable = types.Bool{Value: true}
		}

		if v := jsonData["header_name"]; v != nil {
			item.HeaderName = types.String{Value: v.(string)}
		} else {
			item.HeaderName = types.String{Null: true}
		}

		if v := jsonData["include_in_response"]; v != nil {
			item.IncludeInResponse = types.Bool{Value: v.(bool)}
		} else {
			item.IncludeInResponse = types.Bool{Null: true}
		}

		if v := jsonData["algorithm"]; v != nil {
			item.Algorithm = types.String{Value: v.(string)}
		} else {
			item.Algorithm = types.String{Null: true}
		}

		pluginsType.RequestId = &item
	}
}

func (s PluginRequestIdType) validate() error { return nil }

func (s PluginRequestIdType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.HeaderName, pluginValue, "header_name", true)
	utils.ValueToMap(s.IncludeInResponse, pluginValue, "include_in_response", true)
	utils.ValueToMap(s.Algorithm, pluginValue, "algorithm", true)

	m[s.Name()] = pluginValue
}
