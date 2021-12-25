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

func RouteTypeStateToMap(state RouteType, isUpdate bool) (map[string]interface{}, error) {

	routeRequestObject := make(map[string]interface{})

	utils.ValueToMap(state.Name, routeRequestObject, "name", isUpdate)
	utils.ValueToMap(state.Description, routeRequestObject, "desc", isUpdate)
	utils.ValueToMap(state.Uri, routeRequestObject, "uri", isUpdate)
	utils.ValueToMap(state.Uris, routeRequestObject, "uris", isUpdate)
	utils.ValueToMap(state.Host, routeRequestObject, "host", isUpdate)
	utils.ValueToMap(state.Hosts, routeRequestObject, "hosts", isUpdate)
	utils.ValueToMap(state.RemoteAddr, routeRequestObject, "remote_addr", isUpdate)
	utils.ValueToMap(state.RemoteAddrs, routeRequestObject, "remote_addrs", isUpdate)
	utils.ValueToMap(state.Methods, routeRequestObject, "methods", isUpdate)
	utils.ValueToMap(state.Priority, routeRequestObject, "priority", isUpdate)

	if !state.IsEnabled.Null {
		if state.IsEnabled.Value {
			routeRequestObject["status"] = 1
		} else {
			routeRequestObject["status"] = 0
		}
	} else if isUpdate {
		routeRequestObject["status"] = nil
	}

	utils.ValueToMap(state.EnableWebsocket, routeRequestObject, "enable_websocket", isUpdate)
	utils.ValueToMap(state.ServiceId, routeRequestObject, "service_id", isUpdate)
	utils.ValueToMap(state.UpstreamId, routeRequestObject, "upstream_id", isUpdate)
	utils.ValueToMap(state.Labels, routeRequestObject, "labels", isUpdate)

	if state.Timeout != nil {
		// TODO: FIXME
		routeRequestObject["timeout"] = map[string]interface{}{
			"connect": utils.TypeNumberToInt(state.Timeout.Connect),
			"send":    utils.TypeNumberToInt(state.Timeout.Send),
			"read":    utils.TypeNumberToInt(state.Timeout.Read),
		}
	} else if isUpdate {
		routeRequestObject["timeout"] = nil
	}

	utils.ValueToMap(state.Script, routeRequestObject, "script", isUpdate)
	utils.ValueToMap(state.PluginConfigId, routeRequestObject, "plugin_config_id", isUpdate)
	utils.ValueToMap(state.FilterFunc, routeRequestObject, "filter_func", isUpdate)

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

	upstream, err := UpstreamTypeStateToMap(state.Upstream, isUpdate)

	if err != nil {
		return nil, err
	}

	routeRequestObject["upstream"] = upstream

	return routeRequestObject, nil
}
