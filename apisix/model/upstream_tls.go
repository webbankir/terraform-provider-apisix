package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type UpstreamTLSType struct {
	ClientCert types.String `tfsdk:"client_cert"`
	ClientKey  types.String `tfsdk:"client_key"`
}

var UpstreamTLSSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"client_cert": {
			Type:     types.StringType,
			Required: true,
		},
		"client_key": {
			Type:     types.StringType,
			Required: true,
		},
	}),
}

func UpstreamTLSMapToState(data map[string]interface{}) *UpstreamTLSType {
	v := data["tls"]

	if v == nil {
		return nil
	}
	value := v.(map[string]interface{})
	output := UpstreamTLSType{}

	utils.MapValueToStringTypeValue(value, "client_cert", &output.ClientCert)
	utils.MapValueToStringTypeValue(value, "client_key", &output.ClientKey)

	return &output
}

func UpstreamTLSStateToMap(state *UpstreamTLSType, dMap map[string]interface{}) {
	if state == nil {
		return
	}

	output := make(map[string]interface{})
	utils.StringTypeValueToMap(state.ClientCert, output, "client_cert")
	utils.StringTypeValueToMap(state.ClientKey, output, "client_key")

	dMap["tls"] = output
}
