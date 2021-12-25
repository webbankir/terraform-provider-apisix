package model

import (
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

	utils.MapValueToValue(jsonMap, "type", &newState.Type)
	utils.MapValueToValue(jsonMap, "service_name", &newState.ServiceName)
	utils.MapValueToValue(jsonMap, "discovery_type", &newState.DiscoveryType)
	utils.MapValueToValue(jsonMap, "name", &newState.Name)
	utils.MapValueToValue(jsonMap, "desc", &newState.Desc)
	utils.MapValueToValue(jsonMap, "pass_host", &newState.PassHost)
	utils.MapValueToValue(jsonMap, "scheme", &newState.Scheme)
	utils.MapValueToValue(jsonMap, "upstream_host", &newState.UpstreamHost)
	utils.MapValueToValue(jsonMap, "hash_on", &newState.HashOn)
	utils.MapValueToValue(jsonMap, "retries", &newState.Retries)
	utils.MapValueToValue(jsonMap, "retry_timeout", &newState.RetryTimeout)
	utils.MapValueToValue(jsonMap, "labels", &newState.Labels)

	//if v := jsonMap["labels"]; v != nil {
	//	values := make(map[string]attr.Value)
	//	for key, value := range v.(map[string]interface{}) {
	//		values[key] = types.String{Value: value.(string)}
	//	}
	//	newState.Labels = types.Map{ElemType: types.StringType, Elems: values}
	//} else {
	//	newState.Labels = types.Map{Null: true}
	//}

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

func UpstreamTypeStateToMap(state *UpstreamType, isUpdate bool) (map[string]interface{}, error) {
	if state == nil {
		return nil, nil
	}
	upstreamRequestObject := make(map[string]interface{})

	utils.ValueToMap(state.Type, upstreamRequestObject, "type", isUpdate)
	utils.ValueToMap(state.Name, upstreamRequestObject, "name", isUpdate)
	utils.ValueToMap(state.ServiceName, upstreamRequestObject, "service_name", isUpdate)
	utils.ValueToMap(state.DiscoveryType, upstreamRequestObject, "discovery_type", isUpdate)
	utils.ValueToMap(state.Desc, upstreamRequestObject, "desc", isUpdate)
	utils.ValueToMap(state.PassHost, upstreamRequestObject, "pass_host", isUpdate)
	utils.ValueToMap(state.Scheme, upstreamRequestObject, "scheme", isUpdate)
	utils.ValueToMap(state.Retries, upstreamRequestObject, "retries", isUpdate)
	utils.ValueToMap(state.RetryTimeout, upstreamRequestObject, "retry_timeout", isUpdate)
	utils.ValueToMap(state.Labels, upstreamRequestObject, "labels", isUpdate)
	utils.ValueToMap(state.UpstreamHost, upstreamRequestObject, "upstream_host", isUpdate)
	utils.ValueToMap(state.HashOn, upstreamRequestObject, "hash_on", isUpdate)

	if v := state.Timeout; v != nil {
		upstreamRequestObject["timeout"] = map[string]interface{}{
			"connect": utils.TypeNumberToInt(v.Connect),
			"send":    utils.TypeNumberToInt(v.Send),
			"read":    utils.TypeNumberToInt(v.Read),
		}
	} else if isUpdate {
		upstreamRequestObject["timeout"] = nil
	}

	if v := state.KeepalivePool; v != nil {
		upstreamRequestObject["keepalive_pool"] = map[string]interface{}{
			"size":         utils.TypeNumberToInt(v.Size),
			"idle_timeout": utils.TypeNumberToInt(v.IdleTimeout),
			"requests":     utils.TypeNumberToInt(v.Requests),
		}
	} else if isUpdate {
		upstreamRequestObject["keepalive_pool"] = nil
	}

	if v := state.Tls; v != nil {
		upstreamRequestObject["tls"] = map[string]interface{}{
			"client_cert": v.ClientCert.Value,
			"client_key":  v.ClientKey.Value,
		}
	} else if isUpdate {
		upstreamRequestObject["tls"] = nil
	}

	if v := UpstreamChecksStateToMap(state.Checks); v != nil {
		upstreamRequestObject["checks"] = v
	} else if isUpdate {
		upstreamRequestObject["checks"] = nil
	}

	if v := UpstreamNodesStateToMap(state.Nodes); v != nil {
		upstreamRequestObject["nodes"] = v
	} else if isUpdate {
		upstreamRequestObject["nodes"] = nil
	}

	return upstreamRequestObject, nil
}
