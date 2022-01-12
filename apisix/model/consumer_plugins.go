package model

import "github.com/hashicorp/terraform-plugin-framework/tfsdk"

type ConsumerPluginsType struct {
	BasicAuth *ConsumerPluginBasicAuthType `tfsdk:"basic_auth"`
}

var ConsumerPluginsSchemaAttribute = tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
	"basic_auth": ConsumerPluginBasicAuthSchemaAttribute,
})

type ConsumerPluginCommonInterface interface {
	Name() string
	StateToMap(m map[string]interface{})
	MapToState(v map[string]interface{}, pluginsType *ConsumerPluginsType)
}
