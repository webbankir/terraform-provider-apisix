package apisix

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plugins"
)

type TimeoutType struct {
	Connect types.Number `tfsdk:"connect"`
	Send    types.Number `tfsdk:"send"`
	Read    types.Number `tfsdk:"read"`
}

type RouteType struct {
	ID              types.String  `tfsdk:"id"`
	Description     types.String  `tfsdk:"desc"`
	EnableWebsocket types.Bool    `tfsdk:"enable_websocket"`
	FilterFunc      types.String  `tfsdk:"filter_func"`
	Host            types.String  `tfsdk:"host"`
	Hosts           types.List    `tfsdk:"hosts"`
	IsEnabled       types.Bool    `tfsdk:"is_enabled"`
	Labels          types.Map     `tfsdk:"labels"`
	Methods         types.List    `tfsdk:"methods"`
	Name            types.String  `tfsdk:"name"`
	Plugins         *PluginsType  `tfsdk:"plugins"`
	PluginConfigId  types.String  `tfsdk:"plugin_config_id"`
	Priority        types.Number  `tfsdk:"priority"`
	RemoteAddr      types.String  `tfsdk:"remote_addr"`
	RemoteAddrs     types.List    `tfsdk:"remote_addrs"`
	Script          types.String  `tfsdk:"script"`
	ServiceId       types.String  `tfsdk:"service_id"`
	Timeout         *TimeoutType  `tfsdk:"timeout"`
	Upstream        *UpstreamType `tfsdk:"upstream"`
	UpstreamId      types.String  `tfsdk:"upstream_id"`
	Uri             types.String  `tfsdk:"uri"`
	Uris            types.List    `tfsdk:"uris"`
}

type PluginsType struct {
	CustomJsons   types.Map                  `tfsdk:"custom_jsons"`
	ProxyRewrite  *plugins.ProxyRewriteType  `tfsdk:"proxy_rewrite"`
	IpRestriction *plugins.IpRestrictionType `tfsdk:"ip_restriction"`
}

type UpstreamType struct {
	Type          types.String `tfsdk:"type"`
	ServiceName   types.String `tfsdk:"service_name"`
	DiscoveryType types.String `tfsdk:"discovery_type"`
	Timeout       *TimeoutType `tfsdk:"timeout"`
	Name          types.String `tfsdk:"name"`
	Desc          types.String `tfsdk:"desc"`
	PassHost      types.String `tfsdk:"pass_host"`
	Scheme        types.String `tfsdk:"scheme"`
	Retries       types.Number `tfsdk:"retries"`
	RetryTimeout  types.Number `tfsdk:"retry_timeout"`
	Labels        types.Map    `tfsdk:"labels"`
	UpstreamHost  types.String `tfsdk:"upstream_host"`
	HashOn        types.String `tfsdk:"hash_on"`
}
