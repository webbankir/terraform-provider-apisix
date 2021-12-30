package validator

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"golang.org/x/net/context"
)

type StringOfStringInSliceType struct {
	Slice []string
}

func (j StringOfStringInSliceType) Description(ctx context.Context) string {
	return ""
}

func (j StringOfStringInSliceType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j StringOfStringInSliceType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	if !request.AttributeConfig.(types.List).Null {
		var values []string

		for _, v := range request.AttributeConfig.(types.List).Elems {
			values = append(values, v.(types.String).Value)
		}

		for k, v := range values {
			if !utils.StringContainsInSlice(j.Slice, v) {
				response.Diagnostics.AddError(
					fmt.Sprintf("Wrong value in field: %v", request.AttributePath.String()),
					fmt.Sprintf("Wrong value %v in position %v. Values must be in: %v", v, k, j.Slice),
				)
				return
			}
		}

	}
}

func StringOfStringInSlice(items ...string) StringOfStringInSliceType {
	var values []string

	for _, v := range items {
		values = append(values, v)
	}

	return StringOfStringInSliceType{Slice: values}
}
