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
	BasicAuth              *PluginBasicAuthType              `tfsdk:"basic-auth"`
	ConsumerRestriction    *PluginConsumerRestrictionType    `tfsdk:"consumer-restriction"`
	Cors                   *PluginCorsType                   `tfsdk:"cors"`
	ExtPluginPostReqType   *PluginExtPluginPostReqType       `tfsdk:"ext-plugin-post-req"`
	ExtPluginPreReqType    *PluginExtPluginPreReqType        `tfsdk:"ext-plugin-pre-req"`
	GELFUDPLogger          *PluginGELFUDPLoggerType          `tfsdk:"gelf-udp-logger"`
	GZIP                   *PluginGZIPType                   `tfsdk:"gzip"`
	HTTPLogger             *PluginHTTPLoggerType             `tfsdk:"http-logger"`
	Headers                *PluginHeadersType                `tfsdk:"headers"`
	IpRestriction          *PluginIpRestrictionType          `tfsdk:"ip-restriction"`
	MultiResponseRewrite   *PluginMultiResponseRewriteType   `tfsdk:"multi-response-rewrite"`
	Prometheus             *PluginPrometheusType             `tfsdk:"prometheus"`
	ProxyCache             *PluginProxyCacheType             `tfsdk:"proxy-cache"`
	ProxyRewrite           *PluginProxyRewriteType           `tfsdk:"proxy-rewrite"`
	RealIP                 *PluginRealIPType                 `tfsdk:"real-ip"`
	Redirect               *PluginRedirectType               `tfsdk:"redirect"`
	RedirectRegex          *PluginRedirectRegexType          `tfsdk:"redirect-regex"`
	RequestId              *PluginRequestIdType              `tfsdk:"request-id"`
	ResponseRewrite        *PluginResponseRewriteType        `tfsdk:"response-rewrite"`
	ServerlessPostFunction *PluginServerlessPostFunctionType `tfsdk:"serverless-post-function"`
	ServerlessPreFunction  *PluginServerlessPreFunctionType  `tfsdk:"serverless-pre-function"`
	Syslog                 *PluginSyslogType                 `tfsdk:"syslog"`
}

var PluginsSchemaAttribute = tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
	"basic-auth":               PluginBasicAuthSchemaAttribute,
	"consumer-restriction":     PluginConsumerRestrictionSchemaAttribute,
	"cors":                     PluginCorsSchemaAttribute,
	"ext-plugin-post-req":      PluginExtPluginPostReqSchemaAttribute,
	"ext-plugin-pre-req":       PluginExtPluginPreReqSchemaAttribute,
	"gelf-udp-logger":          PluginGELFUDPLoggerSchemaAttribute,
	"gzip":                     PluginGZIPSchemaAttribute,
	"headers":                  PluginHeadersSchemaAttribute,
	"http-logger":              PluginHTTPLoggerSchemaAttribute,
	"ip-restriction":           PluginIpRestrictionSchemaAttribute,
	"multi-response-rewrite":   PluginMultiResponseRewriteSchemaAttribute,
	"prometheus":               PluginPrometheusSchemaAttribute,
	"proxy-cache":              PluginProxyCacheSchemaAttribute,
	"proxy-rewrite":            PluginProxyRewriteSchemaAttribute,
	"real-ip":                  PluginRealIPSchemaAttribute,
	"redirect":                 PluginRedirectSchemaAttribute,
	"redirect-regex":           PluginRedirectRegexSchemaAttribute,
	"request-id":               PluginRequestIdSchemaAttribute,
	"response-rewrite":         PluginResponseRewriteSchemaAttribute,
	"serverless-post-function": PluginServerlessPostFunctionSchemaAttribute,
	"serverless-pre-function":  PluginServerlessPreFunctionSchemaAttribute,
	"syslog":                   PluginSyslogSchemaAttribute,
	//"custom":                   PluginCustomSchemaAttribute,
})
