package validator

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/net/context"
	"math/big"
)

type NumberType struct {
	Min   float64
	Max   float64
	Slice []float64
	Type  string
}

func (j NumberType) Description(ctx context.Context) string {
	return "Default value for types.Number attribute !!! "
}

func (j NumberType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j NumberType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	if !request.AttributeConfig.(types.Number).Null {
		v := request.AttributeConfig.(types.Number).Value

		switch j.Type {
		case "gt":
			if v.Cmp(big.NewFloat(j.Min)) > 0 {
				return
			}
		case "gte":
			if v.Cmp(big.NewFloat(j.Min)) >= 0 {
				return
			}
		case "lt":
			if v.Cmp(big.NewFloat(j.Max)) < 0 {
				return
			}
		case "lte":
			if v.Cmp(big.NewFloat(j.Max)) <= 0 {
				return
			}

		}

		response.Diagnostics.AddError(
			fmt.Sprintf("Wrong value in field: %v", request.AttributePath.String()),
			fmt.Sprintf("Values must be more than: %v", j.Min),
		)
		return
	}

}

func NumberGreatThan(v float64) NumberType {
	return NumberType{Min: v, Type: "gt"}
}

func NumberGreatOrEqualThan(v float64) NumberType {
	return NumberType{Min: v, Type: "gte"}
}

func NumberLessThan(v float64) NumberType {
	return NumberType{Max: v, Type: "lt"}
}

func NumberLessOrEqualThan(v float64) NumberType {
	return NumberType{Max: v, Type: "lte"}
}
