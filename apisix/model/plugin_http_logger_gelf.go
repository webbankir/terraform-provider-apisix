package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginGELFHTTPLoggerType struct {
	Disable         types.Bool   `tfsdk:"disable"`
	Host            types.String `tfsdk:"host"`
	Port            types.Number `tfsdk:"port"`
	Timeout         types.Number `tfsdk:"timeout"`
	LoggerName      types.String `tfsdk:"name"`
	InactiveTimeout types.Number `tfsdk:"inactive_timeout"`
	BufferDuration  types.Number `tfsdk:"buffer_duration"`
	MaxRetryCount   types.Number `tfsdk:"max_retry_count"`
	RetryDelay      types.Number `tfsdk:"retry_delay"`
	IncludeReqBody  types.Bool   `tfsdk:"include_req_body"`
}

// host = { type = "string" },
//        port = { type = "integer" },
//        timeout = { type = "integer", minimum = 1, default = 3 },
//        name = { type = "string", default = "gelf http logger" },
//        max_retry_count = { type = "integer", minimum = 0, default = 0 },
//        retry_delay = { type = "integer", minimum = 0, default = 1 },
//        buffer_duration = { type = "integer", minimum = 1, default = 60 },
//        inactive_timeout = { type = "integer", minimum = 1, default = 5 },
//        include_req_body = { type = "boolean", default = false },

var PluginGELFHTTPLoggerSchemaAttribute = tfsdk.Attribute{
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
		"host": {
			Required:    true,
			Type:        types.StringType,
			Description: "host of graylog",
		},
		"port": {
			Required:    true,
			Type:        types.NumberType,
			Description: "port of graylog",
		},
		"name": {
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			Description: "A unique identifier to identity the logger.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("gelf http logger"),
			},
		},
		"timeout": {
			Optional:    true,
			Computed:    true,
			Type:        types.NumberType,
			Description: "Time to keep the connection alive after sending a request",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(3),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
			},
		},

		"inactive_timeout": {
			Optional:    true,
			Computed:    true,
			Type:        types.NumberType,
			Description: "The maximum time to refresh the buffer (in seconds). When the maximum refresh time is reached, all logs will be automatically pushed to the HTTP/HTTPS service regardless of whether the number of logs in the buffer reaches the maximum number set.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(5),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
			},
		},
		"buffer_duration": {
			Optional:    true,
			Computed:    true,
			Type:        types.NumberType,
			Description: "Maximum age in seconds of the oldest entry in a batch before the batch must be processed.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(60),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
			},
		},
		"max_retry_count": {
			Optional:    true,
			Computed:    true,
			Type:        types.NumberType,
			Description: "Maximum number of retries before removing from the processing pipe line.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(0),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(0),
			},
		},
		"retry_delay": {
			Optional:    true,
			Computed:    true,
			Type:        types.NumberType,
			Description: "Number of seconds the process execution should be delayed if the execution fails.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(1),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(0),
			},
		},
		"include_req_body": {
			Optional:    true,
			Computed:    true,
			Type:        types.BoolType,
			Description: "Whether to include the request body. false: indicates that the requested body is not included; true: indicates that the requested body is included.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
	}),
}

func (s PluginGELFHTTPLoggerType) Name() string { return "gelf-http-logger" }

func (s PluginGELFHTTPLoggerType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginGELFHTTPLoggerType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "host", &item.Host)
	utils.MapValueToNumberTypeValue(jsonData, "port", &item.Port)
	utils.MapValueToStringTypeValue(jsonData, "name", &item.LoggerName)
	utils.MapValueToBoolTypeValue(jsonData, "include_req_body", &item.IncludeReqBody)
	utils.MapValueToNumberTypeValue(jsonData, "timeout", &item.Timeout)
	utils.MapValueToNumberTypeValue(jsonData, "inactive_timeout", &item.InactiveTimeout)
	utils.MapValueToNumberTypeValue(jsonData, "buffer_duration", &item.BufferDuration)
	utils.MapValueToNumberTypeValue(jsonData, "max_retry_count", &item.MaxRetryCount)
	utils.MapValueToNumberTypeValue(jsonData, "retry_delay", &item.RetryDelay)

	pluginsType.GELFHTTPLogger = &item
}

func (s PluginGELFHTTPLoggerType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.BoolTypeValueToMap(s.IncludeReqBody, pluginValue, "include_req_body")
	utils.StringTypeValueToMap(s.Host, pluginValue, "host")
	utils.NumberTypeValueToMap(s.Port, pluginValue, "port")
	utils.StringTypeValueToMap(s.LoggerName, pluginValue, "name")
	utils.NumberTypeValueToMap(s.Timeout, pluginValue, "timeout")
	utils.NumberTypeValueToMap(s.InactiveTimeout, pluginValue, "inactive_timeout")
	utils.NumberTypeValueToMap(s.BufferDuration, pluginValue, "buffer_duration")
	utils.NumberTypeValueToMap(s.MaxRetryCount, pluginValue, "max_retry_count")
	utils.NumberTypeValueToMap(s.RetryDelay, pluginValue, "retry_delay")

	m[s.Name()] = pluginValue
}
