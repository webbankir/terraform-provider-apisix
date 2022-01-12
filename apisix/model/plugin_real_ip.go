package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type PluginRealIPType struct {
	Disable          types.Bool   `tfsdk:"disable"`
	Source           types.String `tfsdk:"source"`
	TrustedAddresses types.List   `tfsdk:"trusted_addresses"`
}

var PluginRealIPSchemaAttribute = tfsdk.Attribute{
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
		"source": {
			Required:    true,
			Type:        types.StringType,
			Description: "Any Nginx variable like arg_realip or http_x_forwarded_for dynamically set the client's IP and port in APISIX's view, according to the value of variable. If the value doesn't contain a port, the client's port won't be changed.",
		},
		"trusted_addresses": {
			Optional:    true,
			Type:        types.ListType{ElemType: types.StringType},
			Description: "List of IPs or CIDR ranges dynamically set the set_real_ip_from directive",
		},
	}),
}

func (s PluginRealIPType) Name() string { return "real-ip" }

func (s PluginRealIPType) MapToState(data map[string]interface{}, pluginsType *PluginsType) {
	v := data[s.Name()]
	if v == nil {
		return
	}
	jsonData := v.(map[string]interface{})
	item := PluginRealIPType{}

	utils.MapValueToBoolTypeValue(jsonData, "disable", &item.Disable)
	utils.MapValueToStringTypeValue(jsonData, "max_age", &item.Source)
	utils.MapValueToListTypeValue(jsonData, "trusted_addresses", &item.TrustedAddresses)

	pluginsType.RealIP = &item
}

func (s PluginRealIPType) StateToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{}

	utils.BoolTypeValueToMap(s.Disable, pluginValue, "disable")
	utils.StringTypeValueToMap(s.Source, pluginValue, "source")
	utils.ListTypeValueToMap(s.TrustedAddresses, pluginValue, "trusted_addresses")

	m[s.Name()] = pluginValue
}
