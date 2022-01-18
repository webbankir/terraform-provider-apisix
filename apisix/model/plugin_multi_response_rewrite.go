package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginMultiResponseRewriteType struct {
	Disable  types.Bool                              `tfsdk:"disable"`
	Variants []PluginMultiResponseRewriteVariantType `tfsdk:"variants"`
}

type PluginMultiResponseRewriteVariantType struct {
	StatusCode types.Number `tfsdk:"status_code"`
	Body       types.String `tfsdk:"body"`
	BodyBase64 types.Bool   `tfsdk:"body_base64"`
	Headers    types.Map    `tfsdk:"headers"`
	Vars       types.String `tfsdk:"vars"`
}

var PluginMultiResponseRewriteSchemaAttribute = tfsdk.Attribute{
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
		"variants": {
			Required: true,
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
				"status_code": {
					Optional: true,
					Type:     types.NumberType,
					Validators: []tfsdk.AttributeValidator{
						validator.NumberGreatOrEqualThan(200),
						validator.NumberLessOrEqualThan(598),
					},
				},
				"body": {
					Optional: true,
					Type:     types.StringType,
				},
				"body_base64": {
					Optional: true,
					Computed: true,
					Type:     types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultBool(false),
					},
				},

				"headers": {
					Optional: true,
					Type:     types.MapType{ElemType: types.StringType},
				},

				"vars": {
					Required:    true,
					Type:        types.StringType,
					Description: "JSON string",
				},
			},
				tfsdk.ListNestedAttributesOptions{MinItems: 1},
			),

			Validators: []tfsdk.AttributeValidator{
				//validator.HasOneOf("status_code", "body", "headers"),
			},
		},
	}),
}

func (s PluginMultiResponseRewriteType) Name() string { return "multi-response-rewrite" }

func (s PluginMultiResponseRewriteType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginMultiResponseRewriteType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	var subItems []PluginMultiResponseRewriteVariantType

	for _, variant := range jsonData["variants"].([]interface{}) {
		subItem := PluginMultiResponseRewriteVariantType{}
		utils.MapValueToNumberTypeValue(variant.(map[string]interface{}), "status_code", &subItem.StatusCode)
		utils.MapValueToStringTypeValue(variant.(map[string]interface{}), "body", &subItem.Body)
		utils.MapValueToBoolTypeValue(variant.(map[string]interface{}), "body_base64", &subItem.BodyBase64)
		utils.MapValueToMapTypeValue(variant.(map[string]interface{}), "headers", &subItem.Headers)
		subItem.Vars = varsMapToState(variant.(map[string]interface{}))
		subItems = append(subItems, subItem)
	}

	item.Variants = subItems

	pluginsType.MultiResponseRewrite = &item
}

func (s PluginMultiResponseRewriteType) StateToMap(m map[string]interface{}) {
	var pluginValue = make(map[string]interface{})

	var variants []map[string]interface{}
	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")

	for _, v := range s.Variants {
		variant := map[string]interface{}{}
		utils.NumberTypeValueToMap(v.StatusCode, variant, "status_code")
		utils.StringTypeValueToMap(v.Body, variant, "body")
		utils.BoolTypeValueToMap(v.BodyBase64, variant, "body_base64")
		utils.MapTypeValueToMap(v.Headers, variant, "headers")

		varsStateToMap(v.Vars, variant)
		variants = append(variants, variant)
	}
	pluginValue["variants"] = variants

	m[s.Name()] = pluginValue
}
