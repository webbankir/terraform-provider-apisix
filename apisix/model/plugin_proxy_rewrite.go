package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/common"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginProxyRewriteType struct {
	Disable  types.Bool   `tfsdk:"disable"`
	Scheme   types.String `tfsdk:"scheme"`
	Method   types.String `tfsdk:"method"`
	Uri      types.String `tfsdk:"uri"`
	Host     types.String `tfsdk:"host"`
	Headers  types.Map    `tfsdk:"headers"`
	RegexUri types.List   `tfsdk:"regex_uri"`
}

var PluginProxyRewriteSchemaAttribute = tfsdk.Attribute{
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
		"scheme": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("http", "https"),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("http"),
			},
		},
		"method": {
			Optional: true,
			Type:     types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice(common.HttpMethods...),
			},
		},
		"uri": {
			Optional: true,
			Type:     types.StringType,
		},
		"host": {
			Optional: true,
			Type:     types.StringType,
		},
		"headers": {
			Optional: true,
			Type:     types.MapType{ElemType: types.StringType},
		},
		"regex_uri": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
		},
	}),
}

func (s PluginProxyRewriteType) Name() string { return "proxy-rewrite" }

func (s PluginProxyRewriteType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}

	jsonData := v.(map[string]interface{})
	item := PluginProxyRewriteType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "scheme", &item.Scheme)
	utils.MapValueToStringTypeValue(jsonData, "method", &item.Method)
	utils.MapValueToStringTypeValue(jsonData, "uri", &item.Uri)
	utils.MapValueToStringTypeValue(jsonData, "host", &item.Host)
	utils.MapValueToMapTypeValue(jsonData, "headers", &item.Headers)
	utils.MapValueToListTypeValue(jsonData, "regex_uri", &item.RegexUri)

	pluginsType.ProxyRewrite = &item
}

func (s PluginProxyRewriteType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.StringTypeValueToMap(s.Scheme, pluginValue, "scheme")
	utils.StringTypeValueToMap(s.Uri, pluginValue, "uri")
	utils.MapTypeValueToMap(s.Headers, pluginValue, "headers")
	utils.StringTypeValueToMap(s.Host, pluginValue, "host")
	utils.StringTypeValueToMap(s.Method, pluginValue, "method")
	utils.ListTypeValueToMap(s.RegexUri, pluginValue, "regex_uri")

	m[s.Name()] = pluginValue
}
