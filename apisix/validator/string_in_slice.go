package validator

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"golang.org/x/net/context"
)

type StringInSliceType struct {
	Slice []string
}

func (j StringInSliceType) Description(ctx context.Context) string {
	return ""
}

func (j StringInSliceType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j StringInSliceType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {

	if !request.AttributeConfig.(types.String).Null {
		if !utils.StringContainsInSlice(j.Slice, request.AttributeConfig.(types.String).Value) {
			response.Diagnostics.AddError(
				fmt.Sprintf("Wrong value '%v' in field: %v", request.AttributeConfig.(types.String).Value, request.AttributePath.String()),
				fmt.Sprintf("Values must be one of: %v", j.Slice),
			)
			return
		}
	}
}

func StringInSlice(items ...string) StringInSliceType {
	var values []string

	for _, v := range items {
		values = append(values, v)
	}

	return StringInSliceType{Slice: values}
}
