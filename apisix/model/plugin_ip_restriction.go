package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginIpRestrictionType struct {
	Disable   types.Bool   `tfsdk:"disable"`
	Message   types.String `tfsdk:"message"`
	WhiteList types.List   `tfsdk:"whitelist"`
	BlackList types.List   `tfsdk:"blacklist"`
}

var PluginIpRestrictionSchemaAttribute = tfsdk.Attribute{
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
		"whitelist": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("blacklist"),
			},
		},
		"blacklist": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("whitelist"),
			},
		},
		"message": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("Your IP address is not allowed"),
			},
		},
	}),
}

func (s PluginIpRestrictionType) DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType) {
	if v := v["ip-restriction"]; v != nil {
		jsonData := v.(map[string]interface{})
		item := PluginIpRestrictionType{}

		if v := jsonData["disable"]; v != nil {
			item.Disable = types.Bool{Value: v.(bool)}
		} else {
			item.Disable = types.Bool{Value: true}
		}

		if v := jsonData["message"]; v != nil {
			item.Message = types.String{Value: v.(string)}
		} else {
			item.Message = types.String{Null: true}
		}

		if v := jsonData["whitelist"]; v != nil {
			var values []attr.Value
			for _, value := range v.([]interface{}) {
				values = append(values, types.String{Value: value.(string)})
			}

			item.WhiteList = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.WhiteList = types.List{Null: true}
		}

		if v := jsonData["blacklist"]; v != nil {
			var values []attr.Value
			for _, value := range v.([]interface{}) {
				values = append(values, types.String{Value: value.(string)})
			}

			item.BlackList = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.BlackList = types.List{Null: true}
		}
		pluginsType.IpRestriction = &item
	}
}
func (s PluginIpRestrictionType) validate() error { return nil }

func (s PluginIpRestrictionType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	if !s.BlackList.Null {
		var values []string
		for _, v := range s.BlackList.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["blacklist"] = values
	}

	if !s.WhiteList.Null {
		var values []string
		for _, v := range s.WhiteList.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["whitelist"] = values
	}

	if !s.Message.Null {
		pluginValue["message"] = s.Message.Value
	}

	m["ip-restriction"] = pluginValue
}
