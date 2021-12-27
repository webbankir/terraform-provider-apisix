package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type UpstreamChecksPassiveUnhealthyType struct {
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	TCPFailures  types.Number `tfsdk:"tcp_failures"`
	Timeouts     types.Number `tfsdk:"timeouts"`
	HTTPFailures types.Number `tfsdk:"http_failures"`
}

var UpstreamChecksPassiveUnhealthySchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"http_statuses": {
			Type:     types.ListType{ElemType: types.NumberType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfNumbers(429, 500, 503),
			},
			Description: "Passive check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node",
		},

		"http_failures": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
				validator.NumberLessOrEqualThan(254),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(5),
			},
			Description: "Passive check (unhealthy node) The number of times that the node is not healthy during HTTP or HTTPS type checking",
		},

		"tcp_failures": {
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
			Description: "Passive check (unhealthy node) When TCP type is checked, determine the number of times that the node is not healthy",
		},

		"timeouts": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
				validator.NumberLessOrEqualThan(254),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(7),
			},
			Description: "Passive checks (unhealthy node) determine the number of timeouts for unhealthy nodes",
		},
	}),
}

func UpstreamChecksPassiveUnhealthyMapToState(data map[string]interface{}) *UpstreamChecksPassiveUnhealthyType {
	v := data["unhealthy"]
	if v == nil {
		return nil
	}

	value := v.(map[string]interface{})
	output := UpstreamChecksPassiveUnhealthyType{}

	utils.MapValueToValue(value, "tcp_failures", &output.TCPFailures)
	utils.MapValueToValue(value, "timeouts", &output.Timeouts)
	utils.MapValueToValue(value, "http_failures", &output.HTTPFailures)
	utils.MapValueToValue(value, "http_statuses", &output.HTTPStatuses)

	return &output
}

func UpstreamChecksPassiveUnhealthyStateToMap(state *UpstreamChecksPassiveUnhealthyType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["unhealthy"] = nil
		}
		return
	}

	output := make(map[string]interface{})
	utils.ValueToMap(state.TCPFailures, output, "tcp_failures", isUpdate)
	utils.ValueToMap(state.Timeouts, output, "timeouts", isUpdate)
	utils.ValueToMap(state.HTTPFailures, output, "http_failures", isUpdate)
	utils.ValueToMap(state.HTTPStatuses, output, "http_statuses", isUpdate)

	dMap["unhealthy"] = output

}
