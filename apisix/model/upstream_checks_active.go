package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type UpstreamChecksActiveType struct {
	Type                   types.String                       `tfsdk:"type"`
	Timeout                types.Number                       `tfsdk:"timeout"`
	Concurrency            types.Number                       `tfsdk:"concurrency"`
	HTTPPath               types.String                       `tfsdk:"http_path"`
	Host                   types.String                       `tfsdk:"host"`
	Port                   types.Number                       `tfsdk:"port"`
	HTTPSVerifyCertificate types.Bool                         `tfsdk:"https_verify_certificate"`
	ReqHeaders             types.List                         `tfsdk:"req_headers"`
	Healthy                *UpstreamChecksActiveHealthyType   `tfsdk:"healthy"`
	Unhealthy              *UpstreamChecksActiveUnhealthyType `tfsdk:"unhealthy"`
}

var UpstreamChecksActiveSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"type": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("http", "https", "tcp"),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("http"),
			},
			Description: "The type of active check",
		},

		"timeout": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(1),
			},
			Description: "The timeout period of the active check (unit: second)",
		},
		"concurrency": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(10),
			},
			Description: "The number of targets to be checked at the same time during the active check",
		},
		"http_path": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("/"),
			},
			Description: "The HTTP request path that is actively checked",
		},
		"host": {
			Type:        types.StringType,
			Optional:    true,
			Description: "The hostname of the HTTP request actively checked",
		},
		"port": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.NumberGreatOrEqualThan(1),
				validator.NumberLessOrEqualThan(65535),
			},
			Description: "The host port of the HTTP request that is actively checked",
		},
		"https_verify_certificate": {
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(true),
			},
			Description: "Active check whether to check the SSL certificate of the remote host when HTTPS type checking is used",
		},
		"req_headers": {
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Description: "Active check When using HTTP or HTTPS type checking, set additional request header information",
		},
		"healthy": UpstreamChecksActiveHealthySchemaAttribute,

		"unhealthy": UpstreamChecksActiveUnhealthySchemaAttribute,
	}),
}

func UpstreamChecksActiveMapToState(data map[string]interface{}) *UpstreamChecksActiveType {
	v := data["active"]
	if v == nil {
		return nil
	}

	output := UpstreamChecksActiveType{}
	value := v.(map[string]interface{})

	utils.MapValueToStringTypeValue(value, "type", &output.Type)
	utils.MapValueToNumberTypeValue(value, "timeout", &output.Timeout)
	utils.MapValueToNumberTypeValue(value, "concurrency", &output.Concurrency)
	utils.MapValueToStringTypeValue(value, "http_path", &output.HTTPPath)
	utils.MapValueToStringTypeValue(value, "host", &output.Host)
	utils.MapValueToNumberTypeValue(value, "port", &output.Port)
	utils.MapValueToBoolTypeValue(value, "https_verify_certificate", &output.HTTPSVerifyCertificate)
	utils.MapValueToListTypeValue(value, "req_headers", &output.ReqHeaders)

	output.Healthy = UpstreamChecksActiveHealthyMapToState(value)
	output.Unhealthy = UpstreamChecksActiveUnhealthyMapToState(value)

	return &output
}

func UpstreamChecksActiveStateToMap(state *UpstreamChecksActiveType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["active"] = nil
		}
		return
	}

	output := make(map[string]interface{})

	utils.StringTypeValueToMap(state.Type, output, "type", isUpdate)
	utils.NumberTypeValueToMap(state.Timeout, output, "timeout", isUpdate)
	utils.NumberTypeValueToMap(state.Concurrency, output, "concurrency", isUpdate)
	utils.StringTypeValueToMap(state.HTTPPath, output, "http_path", isUpdate)
	utils.StringTypeValueToMap(state.Host, output, "host", isUpdate)
	utils.NumberTypeValueToMap(state.Port, output, "port", isUpdate)
	utils.BoolTypeValueToMap(state.HTTPSVerifyCertificate, output, "https_verify_certificate", isUpdate)
	utils.ListTypeValueToMap(state.ReqHeaders, output, "req_headers", isUpdate)

	UpstreamChecksActiveHealthyStateToMap(state.Healthy, output, isUpdate)
	UpstreamChecksActiveUnhealthyStateToMap(state.Unhealthy, output, isUpdate)

	dMap["active"] = output
}
