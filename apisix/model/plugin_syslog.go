package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginSyslogType struct {
	Disable      types.Bool   `tfsdk:"disable"`
	Host         types.String `tfsdk:"host"`
	Port         types.Number `tfsdk:"port"`
	BatchMaxSize types.Number `tfsdk:"batch_max_size"`
}

var PluginSyslogSchemaAttribute = tfsdk.Attribute{
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
			Description: "IP address or the Hostname.",
		},
		"port": {
			Required:    true,
			Type:        types.NumberType,
			Description: "Target upstream port.",
		},
		"batch_max_size": {
			Optional: true,
			Computed: true,
			Type:     types.NumberType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(1),
			},
		},
	}),
}

func (s PluginSyslogType) Name() string { return "syslog" }

func (s PluginSyslogType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginSyslogType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "host", &item.Host)
	utils.MapValueToNumberTypeValue(jsonData, "port", &item.Port)
	utils.MapValueToNumberTypeValue(jsonData, "batch_max_size", &item.BatchMaxSize)

	pluginsType.Syslog = &item
}

func (s PluginSyslogType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.StringTypeValueToMap(s.Host, pluginValue, "host")
	utils.NumberTypeValueToMap(s.Port, pluginValue, "port")
	utils.NumberTypeValueToMap(s.BatchMaxSize, pluginValue, "batch_max_size")

	m[s.Name()] = pluginValue
}
