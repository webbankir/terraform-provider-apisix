package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type UpstreamChecksActiveHealthyType struct {
	Interval     types.Number `tfsdk:"interval"`
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	Successes    types.Number `tfsdk:"successes"`
}

var UpstreamChecksActiveHealthySchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"interval": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(1),
			},
			Description: "Active check (healthy node) check interval (unit: second)",
		},
		"http_statuses": {
			Type:     types.ListType{ElemType: types.NumberType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfNumbers(200, 302),
			},
			Description: "Active check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node",
		},

		"successes": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
				validator.NumberLessOrEqualThan(254),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(2),
			},
			Description: "Active check (healthy node) check interval (unit: second)",
		},
	}),
}

func UpstreamChecksActiveHealthyMapToState(data map[string]interface{}) *UpstreamChecksActiveHealthyType {
	v := data["healthy"]
	if v == nil {
		return nil
	}

	value := v.(map[string]interface{})
	output := UpstreamChecksActiveHealthyType{}

	utils.MapValueToNumberTypeValue(value, "interval", &output.Interval)
	utils.MapValueToNumberTypeValue(value, "successes", &output.Successes)
	utils.MapValueToListTypeValue(value, "http_statuses", &output.HTTPStatuses)
	utils.MapValueToNumberTypeValue(value, "interval", &output.Interval)

	return &output
}

func UpstreamChecksActiveHealthyStateToMap(state *UpstreamChecksActiveHealthyType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["healthy"] = nil
		}
		return
	}

	output := make(map[string]interface{})

	utils.NumberTypeValueToMap(state.Interval, output, "interval", isUpdate)
	utils.ListTypeValueToMap(state.HTTPStatuses, output, "http_statuses", isUpdate)
	utils.NumberTypeValueToMap(state.Successes, output, "successes", isUpdate)

	dMap["healthy"] = output

}
