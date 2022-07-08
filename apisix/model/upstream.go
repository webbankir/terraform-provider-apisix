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

var UpstreamSchema = tfsdk.Schema{

	Attributes: map[string]tfsdk.Attribute{
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
			// FIXME: Currently can't check when in for_each
			//Validators: []tfsdk.AttributeValidator{
			//	validator.StringInSlice("http", "https", "grpc", "grpcs"),
			//},
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
	},
}

var UpstreamSchemaSAttribute = tfsdk.Attribute{
	Optional:   true,
	Attributes: tfsdk.SingleNestedAttributes(UpstreamSchema.Attributes),
	Validators: []tfsdk.AttributeValidator{
		validator.OneOf("nodes", "discovery_type"),
	},
}

func UpstreamTypeMapToState(data map[string]interface{}, needId bool) (*UpstreamType, error) {
	v := data["upstream"]
	if v == nil {
		return nil, nil
	}

	jsonMap := v.(map[string]interface{})

	newState := UpstreamType{}

	if needId {
		utils.MapValueToStringTypeValue(jsonMap, "id", &newState.ID)
	}
	utils.MapValueToStringTypeValue(jsonMap, "type", &newState.Type)
	utils.MapValueToStringTypeValue(jsonMap, "service_name", &newState.ServiceName)
	utils.MapValueToStringTypeValue(jsonMap, "discovery_type", &newState.DiscoveryType)
	utils.MapValueToStringTypeValue(jsonMap, "name", &newState.Name)
	utils.MapValueToStringTypeValue(jsonMap, "desc", &newState.Desc)
	utils.MapValueToStringTypeValue(jsonMap, "pass_host", &newState.PassHost)
	utils.MapValueToStringTypeValue(jsonMap, "scheme", &newState.Scheme)
	utils.MapValueToStringTypeValue(jsonMap, "upstream_host", &newState.UpstreamHost)
	utils.MapValueToStringTypeValue(jsonMap, "hash_on", &newState.HashOn)
	utils.MapValueToNumberTypeValue(jsonMap, "retries", &newState.Retries)
	utils.MapValueToNumberTypeValue(jsonMap, "retry_timeout", &newState.RetryTimeout)
	utils.MapValueToMapTypeValue(jsonMap, "labels", &newState.Labels)

	newState.Timeout = TimeoutMapToState(jsonMap)
	newState.KeepalivePool = UpstreamKeepAlivePoolMapToState(jsonMap)
	newState.TLS = UpstreamTLSMapToState(jsonMap)
	newState.Checks = UpstreamChecksMapToState(jsonMap)
	newState.Nodes = UpstreamNodesMapToState(jsonMap)

	return &newState, nil

}

// UpstreamTypeStateToMap Fucking golang!
func UpstreamTypeStateToMap(state *UpstreamType) (map[string]interface{}, error) {

	if state == nil {
		return nil, nil
	}
	requestObject := make(map[string]interface{})

	//utils.StringTypeValueToMap(state.ID, requestObject, "id")
	utils.StringTypeValueToMap(state.Type, requestObject, "type")
	utils.StringTypeValueToMap(state.Name, requestObject, "name")
	utils.StringTypeValueToMap(state.ServiceName, requestObject, "service_name")
	utils.StringTypeValueToMap(state.DiscoveryType, requestObject, "discovery_type")
	utils.StringTypeValueToMap(state.Desc, requestObject, "desc")
	utils.StringTypeValueToMap(state.PassHost, requestObject, "pass_host")
	utils.StringTypeValueToMap(state.Scheme, requestObject, "scheme")
	utils.NumberTypeValueToMap(state.Retries, requestObject, "retries")
	utils.NumberTypeValueToMap(state.RetryTimeout, requestObject, "retry_timeout")
	utils.MapTypeValueToMap(state.Labels, requestObject, "labels")
	utils.StringTypeValueToMap(state.UpstreamHost, requestObject, "upstream_host")
	utils.StringTypeValueToMap(state.HashOn, requestObject, "hash_on")

	TimeoutStateToMap(state.Timeout, requestObject)
	UpstreamKeepAlivePoolStateToMap(state.KeepalivePool, requestObject)
	UpstreamTLSStateToMap(state.TLS, requestObject)
	UpstreamChecksStateToMap(state.Checks, requestObject)
	UpstreamNodesStateToMap(state.Nodes, requestObject)

	return requestObject, nil
}
