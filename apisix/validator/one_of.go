package validator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"
)

type OneOfType struct {
	Keys []string
}

func (j OneOfType) Description(ctx context.Context) string {
	return ""
}

func (j OneOfType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j OneOfType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {

	//result, _, _ := tftypes.WalkAttributePath(request.Config.Raw, tftypes.NewAttributePathWithSteps(request.AttributePath.Steps()))
	//
	//if result.(tftypes.Value).IsNull() {
	//	return
	//} else if result.(tftypes.Value).Type().Is(tftypes.List{}) {
	//	if request.AttributeConfig.(types.List).Null {
	//		return
	//	}
	//} else if result.(tftypes.Value).Type().Is(tftypes.Map{}) {
	//	if request.AttributeConfig.(types.Map).Null {
	//		return
	//	}
	//} else if result.(tftypes.Value).Type().Is(tftypes.Set{}) {
	//	if request.AttributeConfig.(types.Set).Null {
	//		return
	//	}
	//}

	exists := false
	for _, v := range j.Keys {
		paths := strings.Split(v, ".")
		var steps []tftypes.AttributePathStep
		steps = append(steps, request.AttributePath.Steps()...)
		if len(paths) > 0 {
			for _, path := range paths {
				steps = append(steps, tftypes.AttributeName(path))
			}
			result, _, err := tftypes.WalkAttributePath(request.Config.Raw, tftypes.NewAttributePathWithSteps(steps))

			if err != nil {
				response.Diagnostics.AddError(
					fmt.Sprintf("Unknown path way: %v", v),
					fmt.Sprintf("Unexpected error: %v", err),
				)
				return

			}

			isNull := true

			if result.(tftypes.Value).Type().Is(tftypes.List{}) {
				vv := types.List{}
				request.Config.GetAttribute(ctx, tftypes.NewAttributePathWithSteps(steps), &vv)
				isNull = vv.Null
			} else if result.(tftypes.Value).Type().Is(tftypes.String) {
				vv := types.String{}
				request.Config.GetAttribute(ctx, tftypes.NewAttributePathWithSteps(steps), &vv)
				isNull = vv.Null
			} else if result.(tftypes.Value).Type().Is(tftypes.Number) {
				vv := types.Number{}
				request.Config.GetAttribute(ctx, tftypes.NewAttributePathWithSteps(steps), &vv)
				isNull = vv.Null
			} else if result.(tftypes.Value).Type().Is(tftypes.Bool) {
				vv := types.Bool{}
				request.Config.GetAttribute(ctx, tftypes.NewAttributePathWithSteps(steps), &vv)
				isNull = vv.Null
			} else if result.(tftypes.Value).Type().Is(tftypes.Map{}) {
				vv := types.Map{}
				request.Config.GetAttribute(ctx, tftypes.NewAttributePathWithSteps(steps), &vv)
				isNull = vv.Null
			} else if result.(tftypes.Value).Type().Is(tftypes.Set{}) {
				vv := types.Set{}
				request.Config.GetAttribute(ctx, tftypes.NewAttributePathWithSteps(steps), &vv)
				isNull = vv.Null
			}

			if !isNull {
				if !exists {
					exists = true
				} else {
					response.Diagnostics.AddError(
						fmt.Sprintf("Validation OneOf is failed"), // value in field: %v", ),
						fmt.Sprintf("TBA:%v, %v, %v", request.AttributePath.String(), tftypes.NewAttributePathWithSteps(steps).String(), j.Keys),
					)
					return
				}
			}
		}

	}

	if !exists {
		response.Diagnostics.AddError(
			fmt.Sprintf("Validation OneOf is failed"), // value in field: %v", ),
			fmt.Sprintf("None of the above options are specified TBA:%v, %v", request.AttributePath.String(), j.Keys),
		)
		return
	}
}

func OneOf(items ...string) OneOfType {
	var values []string

	for _, v := range items {
		values = append(values, v)
	}

	return OneOfType{Keys: values}
}
