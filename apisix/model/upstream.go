package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
	"math/big"
)

type UpstreamType struct {
	ID            types.String        `tfsdk:"id"`
	Type          types.String        `tfsdk:"type"`
	ServiceName   types.String        `tfsdk:"service_name"`
	DiscoveryType types.String        `tfsdk:"discovery_type"`
	Timeout       *TimeoutType        `tfsdk:"timeout"`
	Name          types.String        `tfsdk:"name"`
	Desc          types.String        `tfsdk:"desc"`
	PassHost      types.String        `tfsdk:"pass_host"`
	Scheme        types.String        `tfsdk:"scheme"`
	Retries       types.Number        `tfsdk:"retries"`
	RetryTimeout  types.Number        `tfsdk:"retry_timeout"`
	Labels        types.Map           `tfsdk:"labels"`
	UpstreamHost  types.String        `tfsdk:"upstream_host"`
	HashOn        types.String        `tfsdk:"hash_on"`
	KeepalivePool *KeepAlivePoolType  `tfsdk:"keepalive_pool"`
	Tls           *UpstreamTLSType    `tfsdk:"tls"`
	Checks        *UpstreamChecksType `tfsdk:"checks"`
	Nodes         *[]UpstreamNodeType `tfsdk:"nodes"`
}

var UpstreamSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"type": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("roundrobin"),
			},
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("roundrobin", "chash", "ewma", "least_conn"),
			},
		},
		"service_name": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ComesWith("discovery_type"),
			},
		},
		"discovery_type": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ComesWith("service_name"),
			},
		},
		"timeout": TimeoutSchemaAttribute,
		"name": {
			Type:     types.StringType,
			Optional: true,
		},
		"desc": {
			Type:     types.StringType,
			Optional: true,
		},

		"pass_host": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("pass", "node", "rewrite"),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("pass"),
			},
		},
		"scheme": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("http", "https"),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("http"),
			},
		},
		"retries": {
			Type:     types.NumberType,
			Optional: true,
		},
		"retry_timeout": {
			Type:     types.NumberType,
			Optional: true,
		},
		"upstream_host": {
			Type:     types.StringType,
			Optional: true,
		},
		"hash_on": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("vars"),
			},
		},
		"labels": {
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},

		"keepalive_pool": KeepAlivePoolSchemaAttribute,
		"tls":            UpstreamTLSSchemaAttribute,
		"checks":         UpstreamChecksSchemaAttribute,
		"nodes":          UpstreamNodesSchemaAttribute,
	}),

	Validators: []tfsdk.AttributeValidator{
		validator.OneOf("nodes", "discovery_type"),
	},
}

func UpstreamTypeMapToState(data map[string]interface{}) (*UpstreamType, error) {
	v := data["upstream"]
	if v == nil {
		return nil, nil
	}

	jsonMap := v.(map[string]interface{})

	newState := UpstreamType{}

	if v := jsonMap["id"]; v != nil {
		newState.ID = types.String{Value: v.(string)}
	}

	if v := jsonMap["type"]; v != nil {
		newState.Type = types.String{Value: v.(string)}
	} else {
		newState.Type = types.String{Null: true}
	}

	if v := jsonMap["service_name"]; v != nil {
		newState.ServiceName = types.String{Value: v.(string)}
	} else {
		newState.ServiceName = types.String{Null: true}
	}

	if v := jsonMap["discovery_type"]; v != nil {
		newState.DiscoveryType = types.String{Value: v.(string)}
	} else {
		newState.DiscoveryType = types.String{Null: true}
	}

	if v := jsonMap["name"]; v != nil {
		newState.Name = types.String{Value: v.(string)}
	} else {
		newState.Name = types.String{Null: true}
	}

	if v := jsonMap["desc"]; v != nil {
		newState.Desc = types.String{Value: v.(string)}
	} else {
		newState.Desc = types.String{Null: true}
	}

	if v := jsonMap["pass_host"]; v != nil {
		newState.PassHost = types.String{Value: v.(string)}
	} else {
		newState.PassHost = types.String{Null: true}
	}

	if v := jsonMap["scheme"]; v != nil {
		newState.Scheme = types.String{Value: v.(string)}
	} else {
		newState.Scheme = types.String{Null: true}
	}

	if v := jsonMap["upstream_host"]; v != nil {
		newState.UpstreamHost = types.String{Value: v.(string)}
	} else {
		newState.UpstreamHost = types.String{Null: true}
	}

	if v := jsonMap["hash_on"]; v != nil {
		newState.HashOn = types.String{Value: v.(string)}
	} else {
		newState.HashOn = types.String{Null: true}
	}

	if v := jsonMap["retries"]; v != nil {
		newState.Retries = types.Number{Value: big.NewFloat(v.(float64))}
	} else {
		newState.Retries = types.Number{Null: true}
	}

	if v := jsonMap["retry_timeout"]; v != nil {
		newState.RetryTimeout = types.Number{Value: big.NewFloat(v.(float64))}
	} else {
		newState.RetryTimeout = types.Number{Null: true}
	}

	if v := jsonMap["labels"]; v != nil {
		values := make(map[string]attr.Value)
		for key, value := range v.(map[string]interface{}) {
			values[key] = types.String{Value: value.(string)}
		}
		newState.Labels = types.Map{ElemType: types.StringType, Elems: values}
	} else {
		newState.Labels = types.Map{Null: true}
	}

	if v := jsonMap["timeout"]; v != nil {
		timeout := v.(map[string]interface{})

		//FIXME:
		newState.Timeout = &TimeoutType{
			Connect: types.Number{Value: big.NewFloat(timeout["connect"].(float64))},
			Send:    types.Number{Value: big.NewFloat(timeout["send"].(float64))},
			Read:    types.Number{Value: big.NewFloat(timeout["read"].(float64))},
		}
	} else {
		newState.Timeout = nil
	}

	if v := jsonMap["keepalive_pool"]; v != nil {
		keepAlivePool := v.(map[string]interface{})
		//FIXME:
		newState.KeepalivePool = &KeepAlivePoolType{
			Size:        types.Number{Value: big.NewFloat(keepAlivePool["size"].(float64))},
			IdleTimeout: types.Number{Value: big.NewFloat(keepAlivePool["idle_timeout"].(float64))},
			Requests:    types.Number{Value: big.NewFloat(keepAlivePool["requests"].(float64))},
		}
	}

	if v := jsonMap["tls"]; v != nil {
		tls := v.(map[string]interface{})

		newState.Tls = &UpstreamTLSType{
			ClientCert: types.String{Value: tls["client_cert"].(string)},
			ClientKey:  types.String{Value: tls["client_key"].(string)},
		}
	}

	if v := UpstreamChecksMapToState(jsonMap); v != nil {
		newState.Checks = v
	}

	if v := UpstreamNodesMapToState(jsonMap); v != nil {
		newState.Nodes = v
	}

	return &newState, nil

}

func UpstreamTypeStateToMap(state *UpstreamType) (map[string]interface{}, error) {
	if state == nil {
		return nil, nil
	}
	upstreamRequestObject := make(map[string]interface{})

	if !state.Type.Null {
		upstreamRequestObject["type"] = state.Type.Value
	}

	if !state.Name.Null {
		upstreamRequestObject["name"] = state.Name.Value
	}

	if !state.ServiceName.Null {
		upstreamRequestObject["service_name"] = state.ServiceName.Value
	}

	if !state.DiscoveryType.Null {
		upstreamRequestObject["discovery_type"] = state.DiscoveryType.Value
	}

	if !state.Desc.Null {
		upstreamRequestObject["desc"] = state.Desc.Value
	}

	if !state.PassHost.Null {
		upstreamRequestObject["pass_host"] = state.PassHost.Value
	}

	if !state.Scheme.Null {
		upstreamRequestObject["scheme"] = state.Scheme.Value
	}

	if !state.Retries.Null {
		upstreamRequestObject["retries"] = utils.TypeNumberToInt(state.Retries)
	}

	if !state.RetryTimeout.Null {
		upstreamRequestObject["retry_timeout"] = utils.TypeNumberToInt(state.RetryTimeout)
	}

	if !state.Labels.Null {
		values := make(map[string]interface{})
		for k, v := range state.Labels.Elems {
			values[k] = v.(types.String).Value
		}
		upstreamRequestObject["labels"] = values
	}

	if !state.UpstreamHost.Null {
		upstreamRequestObject["upstream_host"] = state.UpstreamHost.Value
	}

	upstreamRequestObject["hash_on"] = state.HashOn.Value

	if v := state.Timeout; v != nil {
		upstreamRequestObject["timeout"] = map[string]interface{}{
			"connect": utils.TypeNumberToInt(v.Connect),
			"send":    utils.TypeNumberToInt(v.Send),
			"read":    utils.TypeNumberToInt(v.Read),
		}
	}

	if v := state.KeepalivePool; v != nil {
		upstreamRequestObject["keepalive_pool"] = map[string]interface{}{
			"size":         utils.TypeNumberToInt(v.Size),
			"idle_timeout": utils.TypeNumberToInt(v.IdleTimeout),
			"requests":     utils.TypeNumberToInt(v.Requests),
		}
	}

	if v := state.Tls; v != nil {
		upstreamRequestObject["tls"] = map[string]interface{}{
			"client_cert": v.ClientCert.Value,
			"client_key":  v.ClientKey.Value,
		}
	}

	if v := UpstreamChecksStateToMap(state.Checks); v != nil {
		upstreamRequestObject["checks"] = v
	}

	if v := UpstreamNodesStateToMap(state.Nodes); v != nil {
		upstreamRequestObject["nodes"] = v
	}

	return upstreamRequestObject, nil
}
