package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginProxyCacheType struct {
	Disable          types.Bool   `tfsdk:"disable"`
	CacheStrategy    types.String `tfsdk:"cache_strategy"`
	CacheZone        types.String `tfsdk:"cache_zone"`
	CacheKey         types.List   `tfsdk:"cache_key"`
	CacheBypass      types.List   `tfsdk:"cache_bypass"`
	CacheMethod      types.List   `tfsdk:"cache_method"`
	CacheHTTPStatus  types.List   `tfsdk:"cache_http_status"`
	HideCacheHeaders types.Bool   `tfsdk:"hide_cache_headers"`
	CacheControl     types.Bool   `tfsdk:"cache_control"`
	NoCache          types.List   `tfsdk:"no_cache"`
	CacheTTL         types.Number `tfsdk:"cache_ttl"`
}

var PluginProxyCacheSchemaAttribute = tfsdk.Attribute{
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
		"cache_strategy": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("disk"),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("disk", "memory"),
			},
		},
		"cache_zone": {
			Optional: true,
			Computed: true,
			Type:     types.StringType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("disk_cache_one"),
			},
		},
		"cache_key": {
			Optional: true,
			Computed: true,
			Type:     types.ListType{ElemType: types.StringType},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfStrings("$host", "$request_uri"),
			},
		},
		"cache_bypass": {
			Optional: true,
			Type:     types.ListType{ElemType: types.StringType},
		},
		"cache_method": {
			Optional: true,
			Computed: true,
			Type:     types.ListType{ElemType: types.StringType},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfStrings("GET", "HEAD"),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.StringOfStringInSlice("GET", "POST", "HEAD"),
			},
		},
		"cache_http_status": {
			Optional: true,
			Computed: true,
			Type:     types.ListType{ElemType: types.NumberType},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfNumbers(200, 301, 404),
			},
			//Validators: []tfsdk.AttributeValidator{
			//	validator.StringOfStringInSlice("GET", "POST", "HEAD"),
			//},
		},
		"hide_cache_headers": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"cache_control": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"no_cache": {
			Optional: true,
			Type:     types.ListType{ElemType: types.NumberType},
		},
		"cache_ttl": {
			Optional: true,
			Computed: true,
			Type:     types.NumberType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(300),
			},
		},
	}),
}

func (s PluginProxyCacheType) Name() string { return "proxy-cache" }

func (s PluginProxyCacheType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginProxyCacheType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "cache_strategy", &item.CacheStrategy)
	utils.MapValueToStringTypeValue(jsonData, "cache_zone", &item.CacheZone)
	utils.MapValueToListTypeValue(jsonData, "cache_key", &item.CacheKey)
	utils.MapValueToListTypeValue(jsonData, "cache_bypass", &item.CacheBypass)
	utils.MapValueToListTypeValue(jsonData, "cache_method", &item.CacheMethod)
	utils.MapValueToListTypeValue(jsonData, "cache_http_status", &item.CacheHTTPStatus)
	utils.MapValueToBoolTypeValue(jsonData, "hide_cache_headers", &item.HideCacheHeaders)
	utils.MapValueToListTypeValue(jsonData, "no_cache", &item.NoCache)
	utils.MapValueToNumberTypeValue(jsonData, "cache_ttl", &item.CacheTTL)

	pluginsType.ProxyCache = &item
}

func (s PluginProxyCacheType) StateToMap(m map[string]interface{}) {
	var pluginValue = make(map[string]interface{})

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.StringTypeValueToMap(s.CacheStrategy, pluginValue, "cache_strategy")
	utils.StringTypeValueToMap(s.CacheZone, pluginValue, "cache_zone")
	utils.ListTypeValueToMap(s.CacheKey, pluginValue, "cache_key")
	utils.ListTypeValueToMap(s.CacheBypass, pluginValue, "cache_bypass")
	utils.ListTypeValueToMap(s.CacheMethod, pluginValue, "cache_method")
	utils.ListTypeValueToMap(s.CacheHTTPStatus, pluginValue, "cache_http_status")
	utils.BoolTypeValueToMap(s.HideCacheHeaders, pluginValue, "hide_cache_headers")
	utils.BoolTypeValueToMap(s.CacheControl, pluginValue, "cache_control")
	utils.ListTypeValueToMap(s.NoCache, pluginValue, "no_cache")
	utils.NumberTypeValueToMap(s.CacheTTL, pluginValue, "cache_ttl")

	m[s.Name()] = pluginValue
}
