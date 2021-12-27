package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type UpstreamChecksActiveUnhealthyType struct {
	Interval     types.Number `tfsdk:"interval"`
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	TCPFailures  types.Number `tfsdk:"tcp_failures"`
	Timeouts     types.Number `tfsdk:"timeouts"`
	HTTPFailures types.Number `tfsdk:"http_failures"`
}

var UpstreamChecksActiveUnhealthySchemaAttribute = tfsdk.Attribute{
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
			Description: "Active check (unhealthy node) check interval (unit: second)",
		},
		"http_statuses": {
			Type:     types.ListType{ElemType: types.NumberType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfNumbers(404, 429, 500, 501, 503, 504, 505),
			},
			Description: "Active check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node",
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
			Description: "Active check (unhealthy node) HTTP or HTTPS type check, determine the number of times that the node is not healthy",
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
			Description: "Active check (unhealthy node) TCP type check, determine the number of times that the node is not healthy",
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
				plan_modifier.DefaultNumber(3),
			},
			Description: "Active check (unhealthy node) to determine the number of timeouts for unhealthy nodes",
		},
	}),
}

func UpstreamChecksActiveUnhealthyMapToState(data map[string]interface{}) *UpstreamChecksActiveUnhealthyType {
	v := data["unhealthy"]
	if v == nil {
		return nil
	}

	value := v.(map[string]interface{})
	output := UpstreamChecksActiveUnhealthyType{}

	utils.MapValueToNumberTypeValue(value, "interval", &output.Interval)
	utils.MapValueToNumberTypeValue(value, "tcp_failures", &output.TCPFailures)
	utils.MapValueToNumberTypeValue(value, "timeouts", &output.Timeouts)
	utils.MapValueToNumberTypeValue(value, "http_failures", &output.HTTPFailures)
	utils.MapValueToListTypeValue(value, "http_statuses", &output.HTTPStatuses)

	return &output
}

func UpstreamChecksActiveUnhealthyStateToMap(state *UpstreamChecksActiveUnhealthyType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["unhealthy"] = nil
		}
		return
	}

	output := make(map[string]interface{})
	utils.NumberTypeValueToMap(state.Interval, output, "interval", isUpdate)
	utils.NumberTypeValueToMap(state.TCPFailures, output, "tcp_failures", isUpdate)
	utils.NumberTypeValueToMap(state.Timeouts, output, "timeouts", isUpdate)
	utils.NumberTypeValueToMap(state.HTTPFailures, output, "http_failures", isUpdate)
	utils.ListTypeValueToMap(state.HTTPStatuses, output, "http_statuses", isUpdate)

	dMap["unhealthy"] = output
}
