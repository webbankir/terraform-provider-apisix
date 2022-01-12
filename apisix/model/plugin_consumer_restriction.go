package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/common"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginConsumerRestrictionType struct {
	Disable          types.Bool                                      `tfsdk:"disable"`
	Type             types.String                                    `tfsdk:"type"`
	WhiteList        types.List                                      `tfsdk:"whitelist"`
	BlackList        types.List                                      `tfsdk:"blacklist"`
	RejectedCode     types.Number                                    `tfsdk:"rejected_code"`
	AllowedByMethods *[]PluginConsumerRestrictionAllowedByMethodType `tfsdk:"allowed_by_methods"`
}

type PluginConsumerRestrictionAllowedByMethodType struct {
	User    types.String `tfsdk:"user"`
	Methods types.List   `tfsdk:"methods"`
}

var PluginConsumerRestrictionSchemaAttribute = tfsdk.Attribute{
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
		"type": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("consumer_name"),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("consumer_name", "service_id", "route_id"),
			},
		},
		"whitelist": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validator.ElementsGreatThan(0),
			},
		},
		"blacklist": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validator.ElementsGreatThan(0),
			},
		},
		"rejected_code": {
			Optional: true,
			Computed: true,
			Type:     types.NumberType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(403),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(200),
				validator.NumberLessOrEqualThan(599),
			},
		},
		"allowed_by_methods": {
			Optional: true,
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
				"users": {
					Required: true,
					Type:     types.StringType,
				},
				"methods": {
					Required: true,
					Type:     types.ListType{ElemType: types.StringType},
					Validators: []tfsdk.AttributeValidator{
						validator.StringOfStringInSlice(common.HttpMethods...),
					},
				},
			}, tfsdk.ListNestedAttributesOptions{MinItems: 1}),
		},
	}),
	Validators: []tfsdk.AttributeValidator{
		validator.HasOneOf("whitelist", "blacklist", "allowed_by_methods"),
	},
}

func (s PluginConsumerRestrictionType) Name() string { return "consumer-restriction" }

func (s PluginConsumerRestrictionType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginConsumerRestrictionType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "type", &item.Type)
	utils.MapValueToListTypeValue(jsonData, "whitelist", &item.WhiteList)
	utils.MapValueToListTypeValue(jsonData, "blacklist", &item.BlackList)
	utils.MapValueToNumberTypeValue(jsonData, "rejected_code", &item.RejectedCode)

	if v := jsonData["allowed_by_methods"]; v != nil {
		var subItems []PluginConsumerRestrictionAllowedByMethodType
		for _, vv := range v.([]interface{}) {
			subItem := PluginConsumerRestrictionAllowedByMethodType{}
			subV := vv.(map[string]interface{})
			utils.MapValueToStringTypeValue(subV, "user", &subItem.User)
			utils.MapValueToListTypeValue(subV, "methods", &subItem.Methods)
			subItems = append(subItems, subItem)
		}

		if len(subItems) > 0 {
			item.AllowedByMethods = &subItems
		} else {
			item.AllowedByMethods = nil
		}
	} else {
		item.AllowedByMethods = nil
	}

	pluginsType.ConsumerRestriction = &item
}

func (s PluginConsumerRestrictionType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.StringTypeValueToMap(s.Type, pluginValue, "type")
	utils.NumberTypeValueToMap(s.RejectedCode, pluginValue, "rejected_code")
	utils.ListTypeValueToMap(s.BlackList, pluginValue, "blacklist")
	utils.ListTypeValueToMap(s.WhiteList, pluginValue, "whitelist")

	if v := s.AllowedByMethods; v != nil {
		var subItems []map[string]interface{}
		for _, vv := range *v {
			subItem := make(map[string]interface{})
			utils.StringTypeValueToMap(vv.User, subItem, "user")
			utils.ListTypeValueToMap(vv.Methods, subItem, "methods")
		}

		pluginValue["allowed_by_methods"] = subItems

	}

	m[s.Name()] = pluginValue
}
