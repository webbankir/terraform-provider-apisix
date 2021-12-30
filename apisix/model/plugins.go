package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type PluginCommonInterface interface {
	Name() string
	StateToMap(m map[string]interface{}, isUpdate bool)
	MapToState(v map[string]interface{}, pluginsType *PluginsType)
}

type PluginsType struct {
	ProxyRewrite           *PluginProxyRewriteType           `tfsdk:"proxy_rewrite"`
	IpRestriction          *PluginIpRestrictionType          `tfsdk:"ip_restriction"`
	RequestId              *PluginRequestIdType              `tfsdk:"request_id"`
	ServerlessPreFunction  *PluginServerlessPreFunctionType  `tfsdk:"serverless_pre_function"`
	ServerlessPostFunction *PluginServerlessPostFunctionType `tfsdk:"serverless_post_function"`
	Prometheus             *PluginPrometheusType             `tfsdk:"prometheus"`
	Redirect               *PluginRedirectType               `tfsdk:"redirect"`
	Cors                   *PluginCorsType                   `tfsdk:"cors"`
	RedirectRegex          *PluginRedirectRegexType          `tfsdk:"redirect_regex"`
	ResponseRewrite        *PluginResponseRewriteType        `tfsdk:"response_rewrite"`
	//Custom                 *[]PluginCustomType               `tfsdk:"custom"`
}

var PluginsSchemaAttribute = tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
	"proxy_rewrite":            PluginProxyRewriteSchemaAttribute,
	"ip_restriction":           PluginIpRestrictionSchemaAttribute,
	"request_id":               PluginRequestIdSchemaAttribute,
	"serverless_pre_function":  PluginServerlessPreFunctionSchemaAttribute,
	"serverless_post_function": PluginServerlessPostFunctionSchemaAttribute,
	"prometheus":               PluginPrometheusSchemaAttribute,
	"redirect":                 PluginRedirectSchemaAttribute,
	"cors":                     PluginCorsSchemaAttribute,
	"redirect_regex":           PluginRedirectRegexSchemaAttribute,
	//"custom":                   PluginCustomSchemaAttribute,
	"response_rewrite": PluginResponseRewriteSchemaAttribute,
})
