package plan_modifier

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"math/big"
)

type DefaultValue struct {
	Value attr.Value
}

func (j DefaultValue) Description(ctx context.Context) string {

	return ""
}

func (j DefaultValue) MarkdownDescription(ctx context.Context) string {
	return ""
}

func (j DefaultValue) Modify(ctx context.Context, request tfsdk.ModifyAttributePlanRequest, response *tfsdk.ModifyAttributePlanResponse) {

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
		response.AttributePlan = j.Value
	}
}

func DefaultNumber(v float64) DefaultValue {
	return DefaultValue{Value: types.Number{Value: big.NewFloat(v)}}
}

func DefaultString(v string) DefaultValue {
	return DefaultValue{Value: types.String{Value: v}}
}

func DefaultBool(v bool) DefaultValue {
	return DefaultValue{Value: types.Bool{Value: v}}
}

func DefaultObject(t map[string]attr.Type, v map[string]attr.Value) DefaultValue {
	return DefaultValue{
		Value: types.Object{
			AttrTypes: t,
			Attrs:     v,
		},
	}
}

func DefaultListOfNumbers(items ...float64) DefaultValue {
	values := make([]attr.Value, 0)

	for _, v := range items {
		values = append(values, types.Number{Value: big.NewFloat(v)})
	}
	return DefaultValue{
		Value: types.List{
			ElemType: types.NumberType,
			Elems:    values,
		},
	}
}

func DefaultListOfStrings(items ...string) DefaultValue {
	values := make([]attr.Value, 0)

	for _, v := range items {
		values = append(values, types.String{Value: v})
	}
	return DefaultValue{
		Value: types.List{
			ElemType: types.StringType,
			Elems:    values,
		},
	}
}
