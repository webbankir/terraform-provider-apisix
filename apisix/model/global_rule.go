package model

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"reflect"
)

type GlobalRuleType struct {
	ID      types.String `tfsdk:"id"`
	Plugins *PluginsType `tfsdk:"plugins"`
}

var GlobalRuleSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"plugins": {
			Optional:   true,
			Attributes: PluginsSchemaAttribute,
		},
	},
}

func GlobalRuleTypeMapToState(jsonMap map[string]interface{}) (*GlobalRuleType, error) {
	newState := GlobalRuleType{}

	utils.MapValueToStringTypeValue(jsonMap, "id", &newState.ID)

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

func GlobalRuleTypeStateToMap(state GlobalRuleType) (map[string]interface{}, error) {

	output := make(map[string]interface{})
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

	return output, nil
}
