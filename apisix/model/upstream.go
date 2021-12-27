package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type UpstreamType struct {
	ID            types.String               `tfsdk:"id"`
	Type          types.String               `tfsdk:"type"`
	ServiceName   types.String               `tfsdk:"service_name"`
	DiscoveryType types.String               `tfsdk:"discovery_type"`
	Timeout       *TimeoutType               `tfsdk:"timeout"`
	Name          types.String               `tfsdk:"name"`
	Desc          types.String               `tfsdk:"desc"`
	PassHost      types.String               `tfsdk:"pass_host"`
	Scheme        types.String               `tfsdk:"scheme"`
	Retries       types.Number               `tfsdk:"retries"`
	RetryTimeout  types.Number               `tfsdk:"retry_timeout"`
	Labels        types.Map                  `tfsdk:"labels"`
	UpstreamHost  types.String               `tfsdk:"upstream_host"`
	HashOn        types.String               `tfsdk:"hash_on"`
	KeepalivePool *UpstreamKeepAlivePoolType `tfsdk:"keepalive_pool"`
	TLS           *UpstreamTLSType           `tfsdk:"tls"`
	Checks        *UpstreamChecksType        `tfsdk:"checks"`
	Nodes         *[]UpstreamNodeType        `tfsdk:"nodes"`
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

		"keepalive_pool": UpstreamKeepAlivePoolSchemaAttribute,
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

	utils.MapValueToValue(jsonMap, "id", &newState.ID)
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

	newState.Timeout = TimeoutMapToState(jsonMap)
	newState.KeepalivePool = UpstreamKeepAlivePoolMapToState(jsonMap)
	newState.TLS = UpstreamTLSMapToState(jsonMap)
	newState.Checks = UpstreamChecksMapToState(jsonMap)
	newState.Nodes = UpstreamNodesMapToState(jsonMap)

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

	TimeoutStateToMap(state.Timeout, upstreamRequestObject, isUpdate)
	UpstreamKeepAlivePoolStateToMap(state.KeepalivePool, upstreamRequestObject, isUpdate)
	UpstreamTLSStateToMap(state.TLS, upstreamRequestObject, isUpdate)
	UpstreamChecksStateToMap(state.Checks, upstreamRequestObject, isUpdate)
	UpstreamNodesStateToMap(state.Nodes, upstreamRequestObject, isUpdate)

	return upstreamRequestObject, nil
}
