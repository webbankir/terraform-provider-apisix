package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"reflect"
)

type ConsumerType struct {
	Username    types.String         `tfsdk:"username"`
	Description types.String         `tfsdk:"desc"`
	Labels      types.Map            `tfsdk:"labels"`
	Plugins     *ConsumerPluginsType `tfsdk:"plugins"`
}

var ConsumerSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
		"username": {
			Type:        types.StringType,
			Required:    true,
			Description: "Consumer name",
		},
		"desc": {
			Type:        types.StringType,
			Optional:    true,
			Description: "Identifies route names, usage scenarios, and more.",
		},
		"labels": {
			Optional:    true,
			Type:        types.MapType{ElemType: types.StringType},
			Description: "Key/value pairs to specify attributes",
		},

		"plugins": {
			Optional:   true,
			Attributes: ConsumerPluginsSchemaAttribute,
		},
	},
}

func ConsumerTypeMapToState(jsonMap map[string]interface{}) (*ConsumerType, error) {
	newState := ConsumerType{}

	utils.MapValueToStringTypeValue(jsonMap, "username", &newState.Username)
	utils.MapValueToStringTypeValue(jsonMap, "desc", &newState.Description)
	utils.MapValueToMapTypeValue(jsonMap, "labels", &newState.Labels)

	if v := jsonMap["plugins"]; v != nil {
		value := v.(map[string]interface{})
		pluginsType := ConsumerPluginsType{}

		e := reflect.ValueOf(&pluginsType).Elem()
		for i := 0; i < e.NumField(); i++ {
			switch e.Field(i).Interface().(type) {
			case ConsumerPluginCommonInterface:
				reflect.New(e.Type().Field(i).Type.Elem()).Interface().(ConsumerPluginCommonInterface).MapToState(value, &pluginsType)
			default:

			}
		}
		newState.Plugins = &pluginsType
	} else {
		newState.Plugins = nil
	}
	return &newState, nil
}

func ConsumerTypeStateToMap(state ConsumerType) (map[string]interface{}, error) {

	output := make(map[string]interface{})

	utils.StringTypeValueToMap(state.Username, output, "username")
	utils.StringTypeValueToMap(state.Description, output, "desc")
	utils.MapTypeValueToMap(state.Labels, output, "labels")

	plugins := make(map[string]interface{})
	if state.Plugins != nil {
		planPlugins := state.Plugins

		e := reflect.ValueOf(planPlugins).Elem()
		for i := 0; i < e.NumField(); i++ {

			if !e.Field(i).IsNil() {
				switch e.Field(i).Interface().(type) {
				case ConsumerPluginCommonInterface:
					e.Field(i).Interface().(ConsumerPluginCommonInterface).StateToMap(plugins)
				default:

				}

			}
		}

		//PluginCustomTypeStateToMap(plugins, plan, state, isUpdate)

		output["plugins"] = plugins
	}

	return output, nil
}
