package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"math/big"
)

type KeepAlivePoolType struct {
	Size        types.Number `tfsdk:"size"`
	IdleTimeout types.Number `tfsdk:"idle_timeout"`
	Requests    types.Number `tfsdk:"requests"`
}

var KeepAlivePoolSchemaAttribute = tfsdk.Attribute{
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
