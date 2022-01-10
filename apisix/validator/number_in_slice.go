package validator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
)

type NumberInSliceType struct {
	Slice []float64
}

func (j NumberInSliceType) Description(ctx context.Context) string {
	return ""
}

func (j NumberInSliceType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j NumberInSliceType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	if !request.AttributeConfig.(types.Number).Null {
		f64, _ := request.AttributeConfig.(types.Number).Value.Float64()
		if !utils.NumberContainsInSlice(j.Slice, f64) {
			response.Diagnostics.AddError(
				fmt.Sprintf("Wrong value in field: %v", request.AttributePath.String()),
				fmt.Sprintf("Values must be one of: %v", j.Slice),
			)
			return
		}
	}
}

func NumberInSlice(items ...float64) NumberInSliceType {
	var values []float64

	for _, v := range items {
		values = append(values, v)
	}

	return NumberInSliceType{Slice: values}
}
