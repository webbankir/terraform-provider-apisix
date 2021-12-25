package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/common"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginProxyRewriteType struct {
	Disable  types.Bool    `tfsdk:"disable"`
	Scheme   types.String  `tfsdk:"scheme"`
	Method   types.String  `tfsdk:"method"`
	Uri      types.String  `tfsdk:"uri"`
	Host     types.String  `tfsdk:"host"`
	Headers  types.Map     `tfsdk:"headers"`
	RegexUri *RegexUriType `tfsdk:"regex_uri"`
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
		"regex_uri": RegexUriSchemaAttribute,
	}),
}

func (s PluginProxyRewriteType) Name() string { return "proxy-rewrite" }

func (s PluginProxyRewriteType) DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType) {
	if v := v[s.Name()]; v != nil {
		jsonData := v.(map[string]interface{})
		item := PluginProxyRewriteType{}

		if v := jsonData["disable"]; v != nil {
			item.Disable = types.Bool{Value: v.(bool)}
		} else {
			item.Disable = types.Bool{Value: true}
		}

		if v := jsonData["scheme"]; v != nil {
			item.Scheme = types.String{Value: v.(string)}
		} else {
			item.Scheme = types.String{Null: true}
		}

		if v := jsonData["method"]; v != nil {
			item.Method = types.String{Value: v.(string)}
		} else {
			item.Method = types.String{Null: true}
		}

		if v := jsonData["uri"]; v != nil {
			item.Uri = types.String{Value: v.(string)}
		} else {
			item.Uri = types.String{Null: true}
		}

		if v := jsonData["host"]; v != nil {
			item.Host = types.String{Value: v.(string)}
		} else {
			item.Host = types.String{Null: true}
		}

		if v := jsonData["headers"]; v != nil {
			items := make(map[string]attr.Value)

			for k, v := range v.(map[string]interface{}) {
				items[k] = types.String{Value: v.(string)}
			}

			item.Headers = types.Map{
				ElemType: types.StringType,
				Elems:    items,
			}
		} else {
			item.Headers = types.Map{Null: true}
		}

		if v := jsonData["regex_uri"]; v != nil {
			item.RegexUri = &RegexUriType{
				Regex:       types.String{Value: v.([]interface{})[0].(string)},
				Replacement: types.String{Value: v.([]interface{})[1].(string)},
			}
		} else {
			item.RegexUri = nil
		}

		pluginsType.ProxyRewrite = &item
	}
}

func (s PluginProxyRewriteType) validate() error { return nil }

func (s PluginProxyRewriteType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.Scheme, pluginValue, "scheme", true)
	utils.ValueToMap(s.Uri, pluginValue, "uri", true)
	utils.ValueToMap(s.Headers, pluginValue, "headers", true)
	utils.ValueToMap(s.Host, pluginValue, "host", true)
	utils.ValueToMap(s.Method, pluginValue, "method", true)

	if s.RegexUri != nil {
		pluginValue["regex_uri"] = []string{s.RegexUri.Regex.Value, s.RegexUri.Replacement.Value}
	} else {
		pluginValue["regex_uri"] = nil
	}

	m[s.Name()] = pluginValue
}
