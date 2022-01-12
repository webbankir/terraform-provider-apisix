package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginRedirectType struct {
	Disable           types.Bool   `tfsdk:"disable"`
	HTTPToHTTPS       types.Bool   `tfsdk:"http_to_https"`
	URI               types.String `tfsdk:"uri"`
	RegexUri          types.List   `tfsdk:"regex_uri"`
	RetCode           types.Number `tfsdk:"ret_code"`
	EncodeURI         types.Bool   `tfsdk:"encode_uri"`
	AppendQueryString types.Bool   `tfsdk:"append_query_string"`
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
		"regex_uri": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
		},

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

func (s PluginRedirectType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}

	jsonData := v.(map[string]interface{})
	item := PluginRedirectType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToBoolTypeValue(jsonData, "http_to_https", &item.HTTPToHTTPS)
	utils.MapValueToStringTypeValue(jsonData, "uri", &item.URI)
	utils.MapValueToNumberTypeValue(jsonData, "ret_code", &item.RetCode)
	utils.MapValueToBoolTypeValue(jsonData, "encode_uri", &item.EncodeURI)
	utils.MapValueToBoolTypeValue(jsonData, "append_query_string", &item.AppendQueryString)
	utils.MapValueToListTypeValue(jsonData, "regex_uri", &item.RegexUri)

	pluginsType.Redirect = &item
}

func (s PluginRedirectType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	utils.BoolTypeValueToMap(s.HTTPToHTTPS, pluginValue, "http_to_https")
	utils.StringTypeValueToMap(s.URI, pluginValue, "uri")
	utils.BoolTypeValueToMap(s.EncodeURI, pluginValue, "encode_uri")
	utils.BoolTypeValueToMap(s.AppendQueryString, pluginValue, "append_query_string")
	utils.NumberTypeValueToMap(s.RetCode, pluginValue, "ret_code")
	utils.ListTypeValueToMap(s.RegexUri, pluginValue, "regex_uri")

	m[s.Name()] = pluginValue
}
