package validator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"
)

type ComesWithType struct {
	Keys []string
}

func (j ComesWithType) Description(ctx context.Context) string {
	return ""
}

func (j ComesWithType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j ComesWithType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {

	result, _, _ := tftypes.WalkAttributePath(request.Config.Raw, tftypes.NewAttributePathWithSteps(request.AttributePath.Steps()))

	if result.(tftypes.Value).IsNull() {
		return
	} else if result.(tftypes.Value).Type().Is(tftypes.List{}) {
		if request.AttributeConfig.(types.List).Null {
			return
		}
	} else if result.(tftypes.Value).Type().Is(tftypes.Map{}) {
		if request.AttributeConfig.(types.Map).Null {
			return
		}
	} else if result.(tftypes.Value).Type().Is(tftypes.Set{}) {
		if request.AttributeConfig.(types.Set).Null {
			return
		}
	}

	for _, v := range j.Keys {
		paths := strings.Split(v, ".")
		var steps []tftypes.AttributePathStep
		if len(paths) > 0 {
			if !strings.HasPrefix(paths[0], "!") {
				if parentPath := request.AttributePath.WithoutLastStep(); parentPath != nil {
					steps = append(steps, parentPath.Steps()...)
				}
			} else {
				paths[0] = strings.TrimPrefix(paths[0], "!")
			}

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

			if isNull {
				response.Diagnostics.AddError(
					fmt.Sprintf("Come with missing"), // value in field: %v", ),
					fmt.Sprintf("Path '%v' comes with '%v", request.AttributePath.String(), tftypes.NewAttributePathWithSteps(steps).String()),
				)
				return
			}
		}

	}
}

func ComesWith(items ...string) ComesWithType {
	var values []string

	for _, v := range items {
		values = append(values, v)
	}

	return ComesWithType{Keys: values}
}
