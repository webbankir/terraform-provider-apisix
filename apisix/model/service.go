package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
	"reflect"
)

type ServiceType struct {
	ID              types.String  `tfsdk:"id"`
	Description     types.String  `tfsdk:"desc"`
	EnableWebsocket types.Bool    `tfsdk:"enable_websocket"`
	Hosts           types.List    `tfsdk:"hosts"`
	Labels          types.Map     `tfsdk:"labels"`
	Name            types.String  `tfsdk:"name"`
	Plugins         *PluginsType  `tfsdk:"plugins"`
	Upstream        *UpstreamType `tfsdk:"upstream"`
	UpstreamId      types.String  `tfsdk:"upstream_id"`
}

var ServiceSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"desc": {
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultString("Managed by Terraform"),
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
		"hosts": {
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"labels": {
			Optional: true,
			Type:     types.MapType{ElemType: types.StringType},
		},
		"name": {
			Type:     types.StringType,
			Required: true,
		},
		"plugins": {
			Optional:   true,
			Attributes: PluginsSchemaAttribute,
		},
		"upstream": UpstreamSchemaAttribute,
		"upstream_id": {
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validator.ConflictsWith("upstream"),
			},
		},
	},
}

func ServiceTypeMapToState(jsonMap map[string]interface{}) (*ServiceType, error) {
	newState := ServiceType{}

	utils.MapValueToStringTypeValue(jsonMap, "id", &newState.ID)
	utils.MapValueToStringTypeValue(jsonMap, "desc", &newState.Description)
	utils.MapValueToBoolTypeValue(jsonMap, "enable_websocket", &newState.EnableWebsocket)
	utils.MapValueToListTypeValue(jsonMap, "hosts", &newState.Hosts)
	utils.MapValueToMapTypeValue(jsonMap, "labels", &newState.Labels)
	utils.MapValueToStringTypeValue(jsonMap, "name", &newState.Name)
	utils.MapValueToStringTypeValue(jsonMap, "upstream_id", &newState.UpstreamId)

	upstream, err := UpstreamTypeMapToState(jsonMap)
	if err != nil {
		return nil, err
	}
	newState.Upstream = upstream

	if v := jsonMap["plugins"]; v != nil {
		value := v.(map[string]interface{})
		pluginsType := PluginsType{}

		e := reflect.ValueOf(&pluginsType).Elem()
		for i := 0; i < e.NumField(); i++ {
			switch e.Field(i).Interface().(type) {
			case PluginCommonInterface:
				reflect.New(e.Type().Field(i).Type.Elem()).Interface().(PluginCommonInterface).MapToState(value, &pluginsType)
			default:

			}
		}

		//PluginCustomTypeMapToState(value, &pluginsType, plan, state)
		newState.Plugins = &pluginsType
	} else {
		newState.Plugins = nil
	}
	return &newState, nil
}

func ServiceTypeStateToMap(state ServiceType) (map[string]interface{}, error) {

	output := make(map[string]interface{})
	utils.StringTypeValueToMap(state.Name, output, "name")
	utils.StringTypeValueToMap(state.Description, output, "desc")
	utils.ListTypeValueToMap(state.Hosts, output, "hosts")
	utils.BoolTypeValueToMap(state.EnableWebsocket, output, "enable_websocket")
	utils.StringTypeValueToMap(state.UpstreamId, output, "upstream_id")
	plugins := make(map[string]interface{})
	if state.Plugins != nil {
		statePlugins := state.Plugins

		e := reflect.ValueOf(statePlugins).Elem()
		for i := 0; i < e.NumField(); i++ {

			if !e.Field(i).IsNil() {
				switch e.Field(i).Interface().(type) {
				case PluginCommonInterface:
					e.Field(i).Interface().(PluginCommonInterface).StateToMap(plugins)
				default:

				}

			}
			//else if isUpdate {
			//	switch e.Field(i).Interface().(type) {
			//	case PluginCommonInterface:
			//		plugins[reflect.New(e.Type().Field(i).Type.Elem()).Interface().(PluginCommonInterface).Name()] = nil
			//	default:
			//	}
			//}
		}

		//PluginCustomTypeStateToMap(plugins, plan, state, isUpdate)

		output["plugins"] = plugins
	}

	upstream, err := UpstreamTypeStateToMap(state.Upstream)

	if err != nil {
		return nil, err
	}

	if upstream != nil {
		output["upstream"] = upstream
	}

	return output, nil
}
