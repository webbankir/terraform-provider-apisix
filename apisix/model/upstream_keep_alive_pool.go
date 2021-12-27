package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"math/big"
)

type UpstreamKeepAlivePoolType struct {
	Size        types.Number `tfsdk:"size"`
	IdleTimeout types.Number `tfsdk:"idle_timeout"`
	Requests    types.Number `tfsdk:"requests"`
}

var UpstreamKeepAlivePoolSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Computed: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"size": {
			Required: true,
			Type:     types.NumberType,
		},
		"idle_timeout": {
			Required: true,
			Type:     types.NumberType,
		},
		"requests": {
			Required: true,
			Type:     types.NumberType,
		},
	}),
	PlanModifiers: []tfsdk.AttributePlanModifier{
		plan_modifier.DefaultObject(
			map[string]attr.Type{
				"size":         types.NumberType,
				"idle_timeout": types.NumberType,
				"requests":     types.NumberType,
			},
			map[string]attr.Value{
				"size":         types.Number{Value: big.NewFloat(320)},
				"idle_timeout": types.Number{Value: big.NewFloat(60)},
				"requests":     types.Number{Value: big.NewFloat(1000)},
			},
		),
	},
}

func UpstreamKeepAlivePoolMapToState(data map[string]interface{}) *UpstreamKeepAlivePoolType {
	v := data["keepalive_pool"]

	if v == nil {
		return nil
	}
	value := v.(map[string]interface{})
	output := UpstreamKeepAlivePoolType{}

	utils.MapValueToValue(value, "size", &output.Size)
	utils.MapValueToValue(value, "idle_timeout", &output.IdleTimeout)
	utils.MapValueToValue(value, "requests", &output.Requests)

	return &output
}

func UpstreamKeepAlivePoolStateToMap(state *UpstreamKeepAlivePoolType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["keepalive_pool"] = nil
		}
		return
	}

	output := make(map[string]interface{})
	utils.ValueToMap(state.Size, output, "size", isUpdate)
	utils.ValueToMap(state.IdleTimeout, output, "idle_timeout", isUpdate)
	utils.ValueToMap(state.Requests, output, "requests", isUpdate)

	dMap["keepalive_pool"] = output
}
