package model

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
)

type PluginMetadataLogFormatType struct {
	Name      types.String `tfsdk:"name"`
	LogFormat types.String `tfsdk:"log_format"`
}

var PluginMetadataHTTPLoggerSchema = tfsdk.Schema{
	Attributes: map[string]tfsdk.Attribute{
		"log_format": {
			Type:     types.StringType,
			Required: true,
		},
		"name": {
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validator.StringInSlice("http-logger", "skywalking-logger", "kafka-logger"),
			},
		},
	},
}

func PluginMetadataHTTPLoggerTypeMapToState(data map[string]interface{}) (*PluginMetadataLogFormatType, error) {
	newState := PluginMetadataLogFormatType{}

	// FIXME:
	bb, _ := json.Marshal(data["log_format"])
	newState.LogFormat = types.String{Value: string(bb)}
	return &newState, nil
}

func PluginMetadataHTTPLoggerTypeStateToMap(state PluginMetadataLogFormatType) (map[string]interface{}, error) {

	requestObject := make(map[string]interface{})
	object := make(map[string]interface{})
	// FIXME:
	_ = json.Unmarshal([]byte(state.LogFormat.Value), &object)
	requestObject["log_format"] = object
	return requestObject, nil
}
