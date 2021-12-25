package model

import "github.com/hashicorp/terraform-plugin-framework/tfsdk"

type PluginCommonInterface interface {
	Name() string
	EncodeToMap(m map[string]interface{})
	DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType)
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
})
