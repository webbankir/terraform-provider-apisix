package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type UpstreamChecksType struct {
	Active  *UpstreamChecksActiveType  `tfsdk:"active"`
	Passive *UpstreamChecksPassiveType `tfsdk:"passive"`
}

var UpstreamChecksSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"active":  UpstreamChecksActiveSchemaAttribute,
		"passive": UpstreamChecksPassiveSchemaAttribute,
	}),
}

func UpstreamChecksStateToMap(state *UpstreamChecksType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["checks"] = nil
		}
		return
	}

	output := make(map[string]interface{})

	UpstreamChecksActiveStateToMap(state.Active, output, isUpdate)
	UpstreamChecksPassiveStateToMap(state.Passive, output, isUpdate)

	dMap["checks"] = output
}

func UpstreamChecksMapToState(data map[string]interface{}) *UpstreamChecksType {
	checks := data["checks"]

	if checks == nil {
		return nil
	}

	value := checks.(map[string]interface{})
	result := UpstreamChecksType{}

	result.Active = UpstreamChecksActiveMapToState(value)
	result.Passive = UpstreamChecksPassiveMapToState(value)

	return &result
}
