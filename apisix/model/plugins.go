package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

//batch-requests
//echo
//server-info
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
	StateToMap(m map[string]interface{})
	MapToState(v map[string]interface{}, pluginsType *PluginsType)
}

type PluginsType struct {
	//Custom                 *[]PluginCustomType               `tfsdk:"custom"`
	BasicAuth              *PluginBasicAuthType              `tfsdk:"basic_auth"`
	ConsumerRestriction    *PluginConsumerRestrictionType    `tfsdk:"consumer_restriction"`
	Cors                   *PluginCorsType                   `tfsdk:"cors"`
	ExtPluginPostReqType   *PluginExtPluginPostReqType       `tfsdk:"ext_plugin_post_req"`
	ExtPluginPreReqType    *PluginExtPluginPreReqType        `tfsdk:"ext_plugin_pre_req"`
	GELFUDPLogger          *PluginGELFUDPLoggerType          `tfsdk:"gelf_udp_logger"`
	GZIP                   *PluginGZIPType                   `tfsdk:"gzip"`
	HTTPLogger             *PluginHTTPLoggerType             `tfsdk:"http_logger"`
	Headers                *PluginHeadersType                `tfsdk:"headers"`
	IpRestriction          *PluginIpRestrictionType          `tfsdk:"ip_restriction"`
	MultiResponseRewrite   *PluginMultiResponseRewriteType   `tfsdk:"multi_response_rewrite"`
	Prometheus             *PluginPrometheusType             `tfsdk:"prometheus"`
	ProxyCache             *PluginProxyCacheType             `tfsdk:"proxy_cache"`
	ProxyRewrite           *PluginProxyRewriteType           `tfsdk:"proxy_rewrite"`
	RealIP                 *PluginRealIPType                 `tfsdk:"real_ip"`
	Redirect               *PluginRedirectType               `tfsdk:"redirect"`
	RedirectRegex          *PluginRedirectRegexType          `tfsdk:"redirect_regex"`
	RequestId              *PluginRequestIdType              `tfsdk:"request_id"`
	ResponseRewrite        *PluginResponseRewriteType        `tfsdk:"response_rewrite"`
	ServerlessPostFunction *PluginServerlessPostFunctionType `tfsdk:"serverless_post_function"`
	ServerlessPreFunction  *PluginServerlessPreFunctionType  `tfsdk:"serverless_pre_function"`
	Syslog                 *PluginSyslogType                 `tfsdk:"syslog"`
}

var PluginsSchemaAttribute = tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
	"basic_auth":               PluginBasicAuthSchemaAttribute,
	"consumer_restriction":     PluginConsumerRestrictionSchemaAttribute,
	"cors":                     PluginCorsSchemaAttribute,
	"ext_plugin_post_req":      PluginExtPluginPostReqSchemaAttribute,
	"ext_plugin_pre_req":       PluginExtPluginPreReqSchemaAttribute,
	"gelf_udp_logger":          PluginGELFUDPLoggerSchemaAttribute,
	"gzip":                     PluginGZIPSchemaAttribute,
	"headers":                  PluginHeadersSchemaAttribute,
	"http_logger":              PluginHTTPLoggerSchemaAttribute,
	"ip_restriction":           PluginIpRestrictionSchemaAttribute,
	"multi_response_rewrite":   PluginMultiResponseRewriteSchemaAttribute,
	"prometheus":               PluginPrometheusSchemaAttribute,
	"proxy_cache":              PluginProxyCacheSchemaAttribute,
	"proxy_rewrite":            PluginProxyRewriteSchemaAttribute,
	"real_ip":                  PluginRealIPSchemaAttribute,
	"redirect":                 PluginRedirectSchemaAttribute,
	"redirect_regex":           PluginRedirectRegexSchemaAttribute,
	"request_id":               PluginRequestIdSchemaAttribute,
	"response_rewrite":         PluginResponseRewriteSchemaAttribute,
	"serverless_post_function": PluginServerlessPostFunctionSchemaAttribute,
	"serverless_pre_function":  PluginServerlessPreFunctionSchemaAttribute,
	"syslog":                   PluginSyslogSchemaAttribute,
	//"custom":                   PluginCustomSchemaAttribute,
})
