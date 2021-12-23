package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegexUriType struct {
	Regex       types.String `tfsdk:"regex"`
	Replacement types.String `tfsdk:"replacement"`
}

var RegexUriSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"regex": {
			Required: true,
			Type:     types.StringType,
		},
		"replacement": {
			Required: true,
			Type:     types.StringType,
		},
	}),
}
