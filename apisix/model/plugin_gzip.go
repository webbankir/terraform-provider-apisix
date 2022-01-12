package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginGZIPType struct {
	Disable     types.Bool             `tfsdk:"disable"`
	Types       types.List             `tfsdk:"types"`
	MinLength   types.Number           `tfsdk:"min_length"`
	CompLevel   types.Number           `tfsdk:"comp_level"`
	HttpVersion types.Number           `tfsdk:"http_version"`
	Vary        types.Bool             `tfsdk:"vary"`
	Buffers     *PluginGZIPBuffersType `tfsdk:"buffers"`
}

type PluginGZIPBuffersType struct {
	Number types.Number `tfsdk:"number"`
	Size   types.Number `tfsdk:"size"`
}

var PluginGZIPSchemaAttribute = tfsdk.Attribute{
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

		"types": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfStrings("text/html"),
			},
			Description: "dynamically set the gzip_types directive, special value \"*\" matches any MIME type",
		},
		"min_length": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(20),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
			},
			Description: "dynamically set the gzip_min_length directive",
		},

		"comp_level": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(1),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
				validator.NumberLessOrEqualThan(9),
			},
			Description: "dynamically set the gzip_comp_level directive",
		},

		"http_version": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(1.1),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberInSlice(1.0, 1.1),
			},
			Description: "dynamically set the gzip_http_version directive",
		},

		"vary": {
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},

			Description: "dynamically set the gzip_vary directive",
		},

		"buffers": {
			Optional: true,
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
				"number": {
					Optional: true,
					Computed: true,
					Type:     types.NumberType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultNumber(32),
					},
					Validators: []tfsdk.AttributeValidator{
						validator.NumberGreatOrEqualThan(1),
					},
					Description: "dynamically set the gzip_buffers directive",
				},
				"size": {
					Optional: true,
					Computed: true,
					Type:     types.NumberType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultNumber(4096),
					},
					Validators: []tfsdk.AttributeValidator{
						validator.NumberGreatOrEqualThan(1),
					},
					Description: "dynamically set the gzip_buffers directive",
				},
			}),
		},
	}),
}

func (s PluginGZIPType) Name() string { return "gzip" }

func (s PluginGZIPType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}

	jsonData := v.(map[string]interface{})
	item := PluginGZIPType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToListTypeValue(jsonData, "types", &item.Types)
	utils.MapValueToNumberTypeValue(jsonData, "min_length", &item.MinLength)
	utils.MapValueToNumberTypeValue(jsonData, "comp_level", &item.CompLevel)
	utils.MapValueToNumberTypeValue(jsonData, "http_version", &item.HttpVersion)
	utils.MapValueToBoolTypeValue(jsonData, "vary", &item.Vary)
	if v := jsonData["buffers"]; v != nil {
		subJson := v.(map[string]interface{})
		subItem := PluginGZIPBuffersType{}

		utils.MapValueToNumberTypeValue(subJson, "number", &subItem.Number)
		utils.MapValueToNumberTypeValue(subJson, "size", &subItem.Size)

		item.Buffers = &subItem

	} else {
		item.Buffers = nil
	}

	pluginsType.GZIP = &item
}

func (s PluginGZIPType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.ListTypeValueToMap(s.Types, pluginValue, "types")
	utils.NumberTypeValueToMap(s.MinLength, pluginValue, "min_length")
	utils.NumberTypeValueToMap(s.CompLevel, pluginValue, "comp_level")
	utils.NumberTypeValueToMap(s.HttpVersion, pluginValue, "http_version")
	utils.BoolTypeValueToMap(s.Vary, pluginValue, "vary")

	if v := s.Buffers; v != nil {
		subPluginValue := map[string]interface{}{}

		utils.NumberTypeValueToMap(v.Number, subPluginValue, "number")
		utils.NumberTypeValueToMap(v.Size, subPluginValue, "size")

		pluginValue["buffers"] = subPluginValue
	}

	m[s.Name()] = pluginValue
}
