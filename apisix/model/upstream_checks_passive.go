package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type UpstreamChecksPassiveType struct {
	Healthy   *UpstreamChecksPassiveHealthyType   `tfsdk:"healthy"`
	Unhealthy *UpstreamChecksPassiveUnhealthyType `tfsdk:"unhealthy"`
}

var UpstreamChecksPassiveSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"healthy":   UpstreamChecksPassiveHealthySchemaAttribute,
		"unhealthy": UpstreamChecksPassiveUnhealthySchemaAttribute,
	}),
}

func UpstreamChecksPassiveMapToState(data map[string]interface{}) *UpstreamChecksPassiveType {
	v := data["passive"]
	if v == nil {
		return nil
	}

	output := UpstreamChecksPassiveType{}
	value := v.(map[string]interface{})

	output.Healthy = UpstreamChecksPassiveHealthyMapToState(value)
	output.Unhealthy = UpstreamChecksPassiveUnhealthyMapToState(value)

	return &output
}

func UpstreamChecksPassiveStateToMap(state *UpstreamChecksPassiveType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["passive"] = nil
		}
		return
	}

	output := make(map[string]interface{})

	UpstreamChecksPassiveHealthyStateToMap(state.Healthy, output, isUpdate)
	UpstreamChecksPassiveUnhealthyStateToMap(state.Unhealthy, output, isUpdate)

	dMap["passive"] = output
}
