package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"math/big"
)

type TimeoutType struct {
	Connect types.Number `tfsdk:"connect"`
	Send    types.Number `tfsdk:"send"`
	Read    types.Number `tfsdk:"read"`
}

var TimeoutSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Computed: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"connect": {
			Required: true,
			Type:     types.NumberType,
		},
		"send": {
			Required: true,
			Type:     types.NumberType,
		},
		"read": {
			Required: true,
			Type:     types.NumberType,
		},
	}),
	PlanModifiers: []tfsdk.AttributePlanModifier{
		plan_modifier.DefaultObject(
			map[string]attr.Type{
				"connect": types.NumberType,
				"send":    types.NumberType,
				"read":    types.NumberType,
			},
			map[string]attr.Value{
				"connect": types.Number{Value: big.NewFloat(60)},
				"send":    types.Number{Value: big.NewFloat(60)},
				"read":    types.Number{Value: big.NewFloat(60)},
			},
		),
	},
}
