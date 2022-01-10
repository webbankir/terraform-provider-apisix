package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

//batch-requests
//echo
//gzip
//real-ip
//server-info
//ext-plugin-pre-req
//ext-plugin-post-req
//grpc-transcode
//fault-injection
//key-auth
//jwt-auth
//basic-auth
//authz-keycloak
//wolf-rbac
//openid-connect
//hmac-auth
//authz-casbin
//ldap-auth
//uri-blocker
//ua-restriction
//referer-restriction
//consumer-restriction
//limit-req
//limit-conn
//limit-count
//request-validation
//proxy-mirror
//api-breaker
//traffic-split
//client-control
//Zipkin
//skywalking
//node-status
//datadog
//http-logger
//skywalking-logger
//tcp-logger
//kafka-logger
//udp-logger
//syslog
//log-rotate
//error-log-logger
//sls-logger
//azure-functions
//dubbo-proxy
//mqtt-proxy

type PluginCommonInterface interface {
	Name() string
	StateToMap(m map[string]interface{}, isUpdate bool)
	MapToState(v map[string]interface{}, pluginsType *PluginsType)
}

type PluginsType struct {
	Custom                 *[]PluginCustomType               `tfsdk:"custom"`
	Cors                   *PluginCorsType                   `tfsdk:"cors"`
	IpRestriction          *PluginIpRestrictionType          `tfsdk:"ip_restriction"`
	Prometheus             *PluginPrometheusType             `tfsdk:"prometheus"`
	ProxyCache             *PluginProxyCacheType             `tfsdk:"proxy_cache"`
	ProxyRewrite           *PluginProxyRewriteType           `tfsdk:"proxy_rewrite"`
	Redirect               *PluginRedirectType               `tfsdk:"redirect"`
	RedirectRegex          *PluginRedirectRegexType          `tfsdk:"redirect_regex"`
	RequestId              *PluginRequestIdType              `tfsdk:"request_id"`
	ResponseRewrite        *PluginResponseRewriteType        `tfsdk:"response_rewrite"`
	ServerlessPostFunction *PluginServerlessPostFunctionType `tfsdk:"serverless_post_function"`
	ServerlessPreFunction  *PluginServerlessPreFunctionType  `tfsdk:"serverless_pre_function"`
}

var PluginsSchemaAttribute = tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
	"cors":                     PluginCorsSchemaAttribute,
	"ip_restriction":           PluginIpRestrictionSchemaAttribute,
	"prometheus":               PluginPrometheusSchemaAttribute,
	"proxy_cache":              PluginProxyCacheSchemaAttribute,
	"proxy_rewrite":            PluginProxyRewriteSchemaAttribute,
	"redirect":                 PluginRedirectSchemaAttribute,
	"redirect_regex":           PluginRedirectRegexSchemaAttribute,
	"request_id":               PluginRequestIdSchemaAttribute,
	"response_rewrite":         PluginResponseRewriteSchemaAttribute,
	"serverless_post_function": PluginServerlessPostFunctionSchemaAttribute,
	"serverless_pre_function":  PluginServerlessPreFunctionSchemaAttribute,
	"custom":                   PluginCustomSchemaAttribute,
})
