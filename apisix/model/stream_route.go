package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"math/big"
)

type StreamRouteType struct {
	ID         types.String  `tfsdk:"id"`
	RemoteAddr types.String  `tfsdk:"remote_addr"`
	ServerAddr types.String  `tfsdk:"server_addr"`
	ServerPort types.Number  `tfsdk:"server_port"`
	SNI        types.String  `tfsdk:"sni"`
	Upstream   *UpstreamType `tfsdk:"upstream"`
	UpstreamId types.String  `tfsdk:"upstream_id"`
}

var StreamRouteTypeSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"remote_addr": {
			Type:        types.StringType,
			Optional:    true,
			Description: "IP/CIDR client IP \"127.0.0.1/32\" or \"127.0.0.1\"",
		},
		"server_addr": {
			Type:        types.StringType,
			Optional:    true,
			Description: "IP/CIDR server IP \"127.0.0.1/32\" or \"127.0.0.1\"",
		},
		"server_port": {
			Type:        types.NumberType,
			Optional:    true,
			Description: "server port 9090",
		},
		"sni": {
			Type:        types.StringType,
			Optional:    true,
			Description: "server name indication \"test.com\"",
		},
		"upstream": UpstreamSchemaAttribute,
		"upstream_id": {
			Type:        types.StringType,
			Optional:    true,
			Description: "Specify the upstream id, see Upstream for more details",
		},
	},
}

func StreamRouteTypeMapToState(data map[string]interface{}) (*StreamRouteType, error) {
	var result StreamRouteType

	if v := data["id"]; v != nil {
		result.ID = types.String{Value: v.(string)}
	}

	if v := data["remote_addr"]; v != nil {
		result.RemoteAddr = types.String{Value: v.(string)}
	} else {
		result.RemoteAddr = types.String{Null: true}
	}

	if v := data["server_addr"]; v != nil {
		result.ServerAddr = types.String{Value: v.(string)}
	} else {
		result.ServerAddr = types.String{Null: true}
	}

	if v := data["sni"]; v != nil {
		result.SNI = types.String{Value: v.(string)}
	} else {
		result.SNI = types.String{Null: true}
	}

	if v := data["upstream_id"]; v != nil {
		result.UpstreamId = types.String{Value: v.(string)}
	} else {
		result.UpstreamId = types.String{Null: true}
	}

	if v := data["server_port"]; v != nil {
		result.ServerPort = types.Number{Value: big.NewFloat(v.(float64))}
	} else {
		result.ServerPort = types.Number{Null: true}
	}

	upstream, err := UpstreamTypeMapToState(data)
	if err != nil {
		return nil, err
	}

	result.Upstream = upstream

	return &result, nil
}

func StreamRouteTypeStateToMap(state *StreamRouteType, isUpdate bool) (map[string]interface{}, error) {
	if state == nil {
		return nil, nil
	}

	var result = make(map[string]interface{})

	utils.ValueToMap(state.RemoteAddr, result, "remote_addr", isUpdate)
	utils.ValueToMap(state.ServerAddr, result, "server_addr", isUpdate)
	utils.ValueToMap(state.ServerPort, result, "server_port", isUpdate)
	utils.ValueToMap(state.UpstreamId, result, "upstream_id", isUpdate)
	utils.ValueToMap(state.SNI, result, "sni", isUpdate)

	upstream, err := UpstreamTypeStateToMap(state.Upstream, isUpdate)
	if err != nil {
		return nil, err
	}

	if upstream != nil {
		result["upstream"] = upstream
	} else if isUpdate {
		result["upstream"] = nil
	}

	return result, nil
}
