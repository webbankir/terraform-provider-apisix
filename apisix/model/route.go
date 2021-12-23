package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/common"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
	"math/big"
	"reflect"
)

type RouteType struct {
	ID              types.String  `tfsdk:"id"`
	Description     types.String  `tfsdk:"desc"`
	EnableWebsocket types.Bool    `tfsdk:"enable_websocket"`
	FilterFunc      types.String  `tfsdk:"filter_func"`
	Host            types.String  `tfsdk:"host"`
	Hosts           types.List    `tfsdk:"hosts"`
	IsEnabled       types.Bool    `tfsdk:"is_enabled"`
	Labels          types.Map     `tfsdk:"labels"`
	Methods         types.List    `tfsdk:"methods"`
	Name            types.String  `tfsdk:"name"`
	Plugins         *PluginsType  `tfsdk:"plugins"`
	PluginConfigId  types.String  `tfsdk:"plugin_config_id"`
	Priority        types.Number  `tfsdk:"priority"`
	RemoteAddr      types.String  `tfsdk:"remote_addr"`
	RemoteAddrs     types.List    `tfsdk:"remote_addrs"`
	Script          types.String  `tfsdk:"script"`
	ServiceId       types.String  `tfsdk:"service_id"`
	Timeout         *TimeoutType  `tfsdk:"timeout"`
	Upstream        *UpstreamType `tfsdk:"upstream"`
	UpstreamId      types.String  `tfsdk:"upstream_id"`
	Uri             types.String  `tfsdk:"uri"`
	Uris            types.List    `tfsdk:"uris"`
}

var RouteSchema = tfsdk.Schema{
	Version: 8,
	Attributes: map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"desc": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("Managed by Terraform"),
			},
		},
		"uri": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("uris"),
			},
		},
		"uris": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("uri"),
			},
		},
		"host": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("hosts"),
			},
		},
		"hosts": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("host"),
			},
		},
		"remote_addr": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("remote_addrs"),
			},
		},
		"remote_addrs": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("remote_addr"),
			},
		},
		"methods": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.StringOfStringInSlice(common.HttpMethods...),
			},
		},
		"priority": {
			Type:     types.NumberType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(0),
			},
		},

		"is_enabled": {
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(true),
			},
		},
		"enable_websocket": {
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"service_id": {
			Type:     types.StringType,
			Optional: true,
		},
		"upstream_id": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("upstream"),
			},
		},
		//
		"labels": {
			Optional: true,
			Type:     types.MapType{ElemType: types.StringType},
		},

		"timeout": TimeoutSchemaAttribute,
		"script": {
			Optional: true,
			Type:     types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("plugin_config_id"),
			},
		},
		"plugin_config_id": {
			Optional: true,
			Type:     types.StringType,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("script"),
			},
		},
		"filter_func": {
			Optional: true,
			Type:     types.StringType,
		},
		"plugins": {
			Optional:   true,
			Attributes: PluginsSchemaAttribute,
		},
		"upstream": UpstreamSchemaAttribute,
	},
}

func RouteTypeMapToState(jsonMap map[string]interface{}) (*RouteType, error) {
	newState := RouteType{}

	if v := jsonMap["id"]; v != nil {
		newState.ID = types.String{Value: v.(string)}
	}

	if v := jsonMap["status"]; v != nil {
		if v.(float64) == 1 {
			newState.IsEnabled = types.Bool{Value: true}
		} else {
			newState.IsEnabled = types.Bool{Value: false}
		}
	}

	if v := jsonMap["name"]; v != nil {
		newState.Name = types.String{Value: v.(string)}
	} else {
		newState.Name = types.String{Null: true}
	}

	if v := jsonMap["desc"]; v != nil {
		newState.Description = types.String{Value: v.(string)}
	} else {
		newState.Description = types.String{Null: true}
	}

	if v := jsonMap["uri"]; v != nil {
		newState.Uri = types.String{Value: v.(string)}
	} else {
		newState.Uri = types.String{Null: true}
	}

	if v := jsonMap["uris"]; v != nil {
		var values []attr.Value
		for _, value := range v.([]interface{}) {
			values = append(values, types.String{Value: value.(string)})
		}

		newState.Uris = types.List{ElemType: types.StringType, Elems: values}
	} else {
		newState.Uris = types.List{Null: true}
	}

	if v := jsonMap["host"]; v != nil {
		newState.Host = types.String{Value: v.(string)}
	} else {
		newState.Host = types.String{Null: true}
	}

	if v := jsonMap["hosts"]; v != nil {
		var values []attr.Value
		for _, value := range v.([]interface{}) {
			values = append(values, types.String{Value: value.(string)})
		}

		newState.Hosts = types.List{ElemType: types.StringType, Elems: values}
	} else {
		newState.Hosts = types.List{Null: true}
	}

	if v := jsonMap["remote_addr"]; v != nil {
		newState.RemoteAddr = types.String{Value: v.(string)}
	} else {
		newState.RemoteAddr = types.String{Null: true}
	}

	if v := jsonMap["remote_addrs"]; v != nil {
		var values []attr.Value
		for _, value := range v.([]interface{}) {
			values = append(values, types.String{Value: value.(string)})
		}

		newState.RemoteAddrs = types.List{ElemType: types.StringType, Elems: values}
	} else {
		newState.RemoteAddrs = types.List{Null: true}
	}

	if v := jsonMap["methods"]; v != nil {
		var values []attr.Value
		for _, value := range v.([]interface{}) {
			values = append(values, types.String{Value: value.(string)})
		}

		newState.Methods = types.List{ElemType: types.StringType, Elems: values}
	} else {
		newState.Methods = types.List{Null: true}
	}

	if v := jsonMap["priority"]; v != nil {
		newState.Priority = types.Number{Value: big.NewFloat(v.(float64))}
	} else {
		newState.Priority = types.Number{Null: true}
	}

	if v := jsonMap["filter_func"]; v != nil {
		newState.FilterFunc = types.String{Value: v.(string)}
	} else {
		newState.FilterFunc = types.String{Null: true}
	}

	if v := jsonMap["script"]; v != nil {
		newState.Script = types.String{Value: v.(string)}
	} else {
		newState.Script = types.String{Null: true}
	}

	if v := jsonMap["upstream_id"]; v != nil {
		newState.UpstreamId = types.String{Value: v.(string)}
	} else {
		newState.UpstreamId = types.String{Null: true}
	}

	if v := jsonMap["service_id"]; v != nil {
		newState.ServiceId = types.String{Value: v.(string)}
	} else {
		newState.ServiceId = types.String{Null: true}
	}

	if v := jsonMap["plugin_config_id"]; v != nil {
		newState.PluginConfigId = types.String{Value: v.(string)}
	} else {
		newState.PluginConfigId = types.String{Null: true}
	}

	if v := jsonMap["enable_websocket"]; v != nil {
		newState.EnableWebsocket = types.Bool{Value: v.(bool)}
	} else {
		newState.EnableWebsocket = types.Bool{Null: true}
	}

	if v := jsonMap["status"]; v != nil {
		if v.(float64) == 1 {
			newState.IsEnabled = types.Bool{Value: true}
		} else {
			newState.IsEnabled = types.Bool{Value: false}
		}
	} else {
		newState.IsEnabled = types.Bool{Null: true}
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

	upstream, err := UpstreamTypeMapToState(jsonMap)
	if err != nil {
		return nil, err
	}

	newState.Upstream = upstream

	if v := jsonMap["timeout"]; v != nil {
		timeout := v.(map[string]interface{})

		newState.Timeout = &TimeoutType{
			Connect: types.Number{Value: big.NewFloat(timeout["connect"].(float64))},
			Send:    types.Number{Value: big.NewFloat(timeout["send"].(float64))},
			Read:    types.Number{Value: big.NewFloat(timeout["read"].(float64))},
		}
	} else {
		newState.Timeout = nil
	}

	if v := jsonMap["plugins"]; v != nil {
		value := v.(map[string]interface{})
		pluginsType := PluginsType{}

		e := reflect.ValueOf(&pluginsType).Elem()
		for i := 0; i < e.NumField(); i++ {
			reflect.New(e.Type().Field(i).Type.Elem()).Interface().(PluginCommonInterface).DecodeFomMap(value, &pluginsType)
		}
		newState.Plugins = &pluginsType
	} else {
		newState.Plugins = nil
	}

	return &newState, nil
}

func RouteTypeStateToMap(state RouteType) (map[string]interface{}, error) {

	routeRequestObject := make(map[string]interface{})

	if !state.Name.Null {
		routeRequestObject["name"] = state.Name.Value
	}

	if !state.Description.Null {
		routeRequestObject["desc"] = state.Description.Value
	} else {
		routeRequestObject["desc"] = "Managed by terraform"
	}

	if state.IsEnabled.Value {
		routeRequestObject["status"] = 1
	} else {
		routeRequestObject["status"] = 0
	}

	if !state.Uri.Null {
		routeRequestObject["uri"] = state.Uri.Value
	}

	if !state.Uris.Null {
		var values []string
		for _, v := range state.Uris.Elems {
			values = append(values, v.(types.String).Value)
		}
		routeRequestObject["uris"] = values
	}

	if !state.Host.Null {
		routeRequestObject["host"] = state.Host.Value
	}

	if !state.Hosts.Null {
		var values []string
		for _, v := range state.Hosts.Elems {
			values = append(values, v.(types.String).Value)
		}
		routeRequestObject["host"] = values
	}

	if !state.RemoteAddr.Null {
		routeRequestObject["remote_addr"] = state.RemoteAddr.Value
	}

	if !state.RemoteAddrs.Null {
		var values []string
		for _, v := range state.RemoteAddrs.Elems {
			values = append(values, v.(types.String).Value)
		}
		routeRequestObject["remote_addrs"] = values
	}

	if !state.Methods.Null {
		var values []string
		for _, v := range state.Methods.Elems {
			values = append(values, v.(types.String).Value)
		}
		routeRequestObject["methods"] = values
	}

	if !state.Priority.Null {
		routeRequestObject["priority"] = utils.TypeNumberToInt(state.Priority)
	}

	if !state.IsEnabled.Null {
		if state.IsEnabled.Value {
			routeRequestObject["status"] = 1
		} else {
			routeRequestObject["status"] = 0
		}
	}

	if !state.EnableWebsocket.Null {
		routeRequestObject["enable_websocket"] = state.EnableWebsocket.Value
	}

	if !state.ServiceId.Null {
		routeRequestObject["service_id"] = state.ServiceId.Value
	}

	if !state.UpstreamId.Null {
		routeRequestObject["upstream_id"] = state.UpstreamId.Value
	}

	if !state.Labels.Null {
		values := make(map[string]interface{})
		for k, v := range state.Labels.Elems {
			values[k] = v.(types.String).Value
		}
		routeRequestObject["labels"] = values
	}

	if state.Timeout != nil {
		routeRequestObject["timeout"] = map[string]interface{}{
			"connect": utils.TypeNumberToInt(state.Timeout.Connect),
			"send":    utils.TypeNumberToInt(state.Timeout.Send),
			"read":    utils.TypeNumberToInt(state.Timeout.Read),
		}

	}

	if !state.Script.Null {
		routeRequestObject["script"] = state.Script.Value
	}

	if !state.PluginConfigId.Null {
		routeRequestObject["plugin_config_id"] = state.PluginConfigId.Value
	}

	if !state.FilterFunc.Null {
		routeRequestObject["filter_func"] = state.FilterFunc.Value
	}

	plugins := make(map[string]interface{})
	if state.Plugins != nil {
		planPlugins := state.Plugins

		e := reflect.ValueOf(planPlugins).Elem()
		for i := 0; i < e.NumField(); i++ {
			if !e.Field(i).IsNil() {
				e.Field(i).Interface().(PluginCommonInterface).EncodeToMap(plugins)
			}
		}
		routeRequestObject["plugins"] = plugins
	}

	upstream, err := UpstreamTypeStateToMap(state.Upstream)

	if err != nil {
		return nil, err
	}

	if upstream != nil {
		routeRequestObject["upstream"] = upstream
	}

	return routeRequestObject, nil
}
