package validator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"strings"
)

type OneOfType struct {
	Keys        []string
	CanMultiple bool
}

func (j OneOfType) Description(ctx context.Context) string {
	return ""
}

func (j OneOfType) MarkdownDescription(ctx context.Context) string {
	return ""

}

func (j OneOfType) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	exists := false

	if utils.IsAttributeIsNull(request.AttributeConfig) {
		return
	}

	for _, v := range j.Keys {

		steps := request.AttributePath
		for _, v := range strings.Split(v, ".") {
			steps = steps.WithAttributeName(v)
		}

		result, _, err := tftypes.WalkAttributePath(request.Config.Raw, steps)

		if err != nil {
			response.Diagnostics.AddError(
				fmt.Sprintf("Unknown path way: %v", steps),
				fmt.Sprintf("Unexpected error: %v", err),
			)
			return
		}

		if !result.(tftypes.Value).IsNull() {
			if !exists {
				exists = true
			} else if !j.CanMultiple {
				response.Diagnostics.AddError(
					fmt.Sprintf("Validation OneOf is failed"), // value in field: %v", ),
					fmt.Sprintf("TBA:%v, %v, %v", request.AttributePath.String(), steps, j.Keys),
				)
				return
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

func HasOneOf(items ...string) OneOfType {
	var values []string

	for _, v := range items {
		values = append(values, v)
	}

	return OneOfType{Keys: values, CanMultiple: true}
}
