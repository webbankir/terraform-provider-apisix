package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
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

func (s PluginIpRestrictionType) Name() string { return "ip-restriction" }

func (s PluginIpRestrictionType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginIpRestrictionType{}

	utils.MapValueToValue(jsonData, "disable", &item.Disable)
	utils.MapValueToValue(jsonData, "message", &item.Message)
	utils.MapValueToValue(jsonData, "whitelist", &item.WhiteList)
	utils.MapValueToValue(jsonData, "blacklist", &item.BlackList)

	pluginsType.IpRestriction = &item
}

func (s PluginIpRestrictionType) StateToMap(m map[string]interface{}, isUpdate bool) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.BlackList, pluginValue, "blacklist", isUpdate)
	utils.ValueToMap(s.WhiteList, pluginValue, "whitelist", isUpdate)
	utils.ValueToMap(s.Message, pluginValue, "message", isUpdate)

	m[s.Name()] = pluginValue
}
