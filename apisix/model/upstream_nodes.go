package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
	"math/big"
	"strconv"
	"strings"
)

type UpstreamNodeType struct {
	Host   types.String `tfsdk:"host"`
	Port   types.Number `tfsdk:"port"`
	Weight types.Number `tfsdk:"weight"`
}

var UpstreamNodesSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"host": {
			Type:     types.StringType,
			Required: true,
		},
		"port": {
			Type:     types.NumberType,
			Required: true,
		},
		"weight": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(1),
			},
		},
	}, tfsdk.ListNestedAttributesOptions{}),
	Validators: []tfsdk.AttributeValidator{
		validator.ConflictsWith("discovery_type", "service_name"),
	},
}

func UpstreamNodesMapToState(data map[string]interface{}) *[]UpstreamNodeType {
	v := data["nodes"]

	if v == nil {
		return nil
	}

	var result []UpstreamNodeType

	switch v.(type) {
	case []interface{}:
		for _, v := range v.([]interface{}) {
			v := v.(map[string]interface{})
			result = append(result, UpstreamNodeType{
				Host:   types.String{Value: v["host"].(string)},
				Port:   types.Number{Value: big.NewFloat(v["port"].(float64))},
				Weight: types.Number{Value: big.NewFloat(v["weight"].(float64))},
			})
		}

	case map[string]interface{}:

		for k, v := range v.(map[string]interface{}) {
			ss := strings.Split(k, ":")
			var host string
			var port int

			if len(ss) == 1 {
				host = ss[0]
				port = 80
			} else if len(ss) == 2 {
				host = ss[0]
				// TODO: Fix this - error
				port, _ = strconv.Atoi(ss[1])
			} else {
				// TODO: Fix this
				panic("Ohhh, upstream node item is bad")
			}

			result = append(result, UpstreamNodeType{
				Host:   types.String{Value: host},
				Port:   types.Number{Value: big.NewFloat(float64(port))},
				Weight: types.Number{Value: big.NewFloat(v.(float64))},
			})
		}

	default:
		return nil

	}

	return &result
}

func UpstreamNodesStateToMap(state *[]UpstreamNodeType, dMap map[string]interface{}, isUpdate bool) {
	if state == nil {
		if isUpdate {
			dMap["nodes"] = nil
		}
		return
	}

	var result []map[string]interface{}

	for _, v := range *state {
		item := map[string]interface{}{}
		utils.ValueToMap(v.Host, item, "host", isUpdate)
		utils.ValueToMap(v.Port, item, "port", isUpdate)
		utils.ValueToMap(v.Weight, item, "weight", isUpdate)
		result = append(result, item)
	}

	if len(result) == 0 {
		return
	}

	dMap["nodes"] = result
}
