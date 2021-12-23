package plan_modifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type DefaultValueFunction struct {
	Function func(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) (attr.Value, error)
}

func (j DefaultValueFunction) Description(ctx context.Context) string {

	return ""
}

func (j DefaultValueFunction) MarkdownDescription(ctx context.Context) string {
	return ""
}

func (j DefaultValueFunction) Modify(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) {

	result, _, _ := tftypes.WalkAttributePath(request.Config.Raw, tftypes.NewAttributePathWithSteps(request.AttributePath.Steps()))

	isNull := false
	if result.(tftypes.Value).IsNull() {
		isNull = true
	} else if result.(tftypes.Value).Type().Is(tftypes.List{}) {
		if request.AttributeConfig.(types.List).Null {
			isNull = true
		}
	} else if result.(tftypes.Value).Type().Is(tftypes.Map{}) {
		if request.AttributeConfig.(types.Map).Null {
			isNull = true
		}
	} else if result.(tftypes.Value).Type().Is(tftypes.Set{}) {
		if request.AttributeConfig.(types.Set).Null {
			isNull = true
		}
	}

	if isNull {
		v, err := j.Function(ctx, request, response)
		if err != nil {
			response.Diagnostics.AddError(
				fmt.Sprintf("Can't call default function on attr:%v", request.AttributePath.String()),
				"Unexpected error:"+err.Error(),
			)
			return
		}
		response.AttributePlan = v
	}

}

func DefaultFunction(f func(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) (attr.Value, error)) DefaultValueFunction {
	return DefaultValueFunction{Function: f}
}
