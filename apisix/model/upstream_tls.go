package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
