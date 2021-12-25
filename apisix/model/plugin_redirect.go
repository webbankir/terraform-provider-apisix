package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
	"math/big"
)

type PluginRedirectType struct {
	Disable           types.Bool    `tfsdk:"disable"`
	HTTPToHTTPS       types.Bool    `tfsdk:"http_to_https"`
	URI               types.String  `tfsdk:"uri"`
	RegexUri          *RegexUriType `tfsdk:"regex_uri"`
	RetCode           types.Number  `tfsdk:"ret_code"`
	EncodeURI         types.Bool    `tfsdk:"encode_uri"`
	AppendQueryString types.Bool    `tfsdk:"append_query_string"`
}

var PluginRedirectSchemaAttribute = tfsdk.Attribute{
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
		"http_to_https": {
			Type:     types.BoolType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("uri", "regex_uri"),
			},
			Description: "When it is set to true and the request is HTTP, will be automatically redirected to HTTPS with 301 response code, and the URI will keep the same as client request",
		},
		"uri": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("http_to_https", "regex_uri"),
			},
			Description: "New URL which can contain Nginx variable, eg: /test/index.html, $uri/index.html. You can refer to variables in a way similar to ${xxx} to avoid ambiguity, eg: ${uri}foo/index.html. If you just need the original $ character, add \\ in front of it, like this one: /\\$foo/index.html. If you refer to a variable name that does not exist, this will not produce an error, and it will be used as an empty string",
		},
		"regex_uri": RegexUriSchemaAttribute,

		"ret_code": {
			Optional:    true,
			Computed:    true,
			Type:        types.NumberType,
			Description: "Response code",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(302),
			},
		},
		"encode_uri": {
			Optional:    true,
			Computed:    true,
			Type:        types.BoolType,
			Description: "When set to true the uri in Location header will be encoded as per RFC3986",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},

		"append_query_string": {
			Optional:    true,
			Computed:    true,
			Type:        types.BoolType,
			Description: "When set to true, add the query string from the original request to the location header. If the configured uri / regex_uri already contains a query string, the query string from request will be appended to that after an &. Caution: don't use this if you've already handled the query string, e.g. with nginx variable $request_uri, to avoid duplicates.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
	}),
}

func (s PluginRedirectType) Name() string { return "redirect" }

func (s PluginRedirectType) DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType) {
	if v := v[s.Name()]; v != nil {
		jsonData := v.(map[string]interface{})
		item := PluginRedirectType{}

		if v := jsonData["disable"]; v != nil {
			item.Disable = types.Bool{Value: v.(bool)}
		} else {
			item.Disable = types.Bool{Value: true}
		}

		if v := jsonData["http_to_https"]; v != nil {
			item.HTTPToHTTPS = types.Bool{Value: v.(bool)}
		} else {
			item.HTTPToHTTPS = types.Bool{Null: true}
		}

		if v := jsonData["uri"]; v != nil {
			item.URI = types.String{Value: v.(string)}
		} else {
			item.URI = types.String{Null: true}
		}

		if v := jsonData["ret_code"]; v != nil {
			item.RetCode = types.Number{Value: big.NewFloat(v.(float64))}
		} else {
			item.RetCode = types.Number{Null: true}
		}

		if v := jsonData["encode_uri"]; v != nil {
			item.EncodeURI = types.Bool{Value: v.(bool)}
		} else {
			item.EncodeURI = types.Bool{Null: true}
		}

		if v := jsonData["append_query_string"]; v != nil {
			item.AppendQueryString = types.Bool{Value: v.(bool)}
		} else {
			item.AppendQueryString = types.Bool{Null: true}
		}

		if v := jsonData["regex_uri"]; v != nil {
			item.RegexUri = &RegexUriType{
				Regex:       types.String{Value: v.([]interface{})[0].(string)},
				Replacement: types.String{Value: v.([]interface{})[1].(string)},
			}
		} else {
			item.RegexUri = nil
		}

		pluginsType.Redirect = &item
	}
}

func (s PluginRedirectType) validate() error { return nil }

func (s PluginRedirectType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.ValueToMap(s.HTTPToHTTPS, pluginValue, "http_to_https", true)
	utils.ValueToMap(s.URI, pluginValue, "uri", true)
	utils.ValueToMap(s.EncodeURI, pluginValue, "encode_uri", true)
	utils.ValueToMap(s.AppendQueryString, pluginValue, "append_query_string", true)
	utils.ValueToMap(s.RetCode, pluginValue, "ret_code", true)

	if s.RegexUri != nil {
		pluginValue["regex_uri"] = []string{s.RegexUri.Regex.Value, s.RegexUri.Replacement.Value}
	} else {
		pluginValue["regex_uri"] = nil
	}

	m[s.Name()] = pluginValue
}
