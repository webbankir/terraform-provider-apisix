package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"log"
	"math/big"
	"strings"
)

type PluginCorsType struct {
	Disable             types.Bool   `tfsdk:"disable"`
	AllowOrigins        types.List   `tfsdk:"allow_origins"`
	AllowMethods        types.List   `tfsdk:"allow_methods"`
	AllowHeaders        types.List   `tfsdk:"allow_headers"`
	ExposeHeaders       types.List   `tfsdk:"expose_headers"`
	MaxAge              types.Number `tfsdk:"max_age"`
	AllowCredential     types.Bool   `tfsdk:"allow_credential"`
	AllowOriginsByRegex types.List   `tfsdk:"allow_origins_by_regex"`
}

var PluginCorsSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"disable": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"allow_credential": {
			Optional: true,
			Computed: true,
			Type:     types.BoolType,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultBool(false),
			},
		},
		"allow_origins": {
			Optional:    true,
			Computed:    true,
			Type:        types.ListType{ElemType: types.StringType},
			Description: "Which Origins is allowed to enable CORS, format as: scheme://host:port, for example: https://somehost.com:8081. Multiple origin use , to split. When allow_credential is false, you can use * to indicate allow any origin. you also can allow all any origins forcefully using ** even already enable allow_credential, but it will bring some security risks.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfStrings("*"),
			},
		},
		"allow_methods": {
			Optional:    true,
			Computed:    true,
			Type:        types.ListType{ElemType: types.StringType},
			Description: "Which Method is allowed to enable CORS, such as: GET, POST etc. Multiple method use , to split. When allow_credential is false, you can use * to indicate allow all any method. You also can allow any method forcefully using ** even already enable allow_credential, but it will bring some security risks.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfStrings("*"),
			},
		},
		"allow_headers": {
			Optional:    true,
			Computed:    true,
			Type:        types.ListType{ElemType: types.StringType},
			Description: "Which headers are allowed to set in request when access cross-origin resource. Multiple value use , to split. When allow_credential is false, you can use * to indicate allow all request headers. You also can allow any header forcefully using ** even already enable allow_credential, but it will bring some security risks.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfStrings("*"),
			},
		},
		"expose_headers": {
			Optional:    true,
			Computed:    true,
			Type:        types.ListType{ElemType: types.StringType},
			Description: "Which headers are allowed to set in response when access cross-origin resource. Multiple value use , to split. When allow_credential is false, you can use * to indicate allow any header. You also can allow any header forcefully using ** even already enable allow_credential, but it will bring some security risks.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultListOfStrings("*"),
			},
		},
		"max_age": {
			Optional:    true,
			Computed:    true,
			Type:        types.NumberType,
			Description: "Maximum number of seconds the results can be cached. Within this time range, the browser will reuse the last check result. -1 means no cache. Please note that the maximum value is depended on browser, please refer to MDN for details.",
			PlanModifiers: []tfsdk.AttributePlanModifier{
				plan_modifier.DefaultNumber(5),
			},
		},
		"allow_origins_by_regex": {
			Optional:    true,
			Type:        types.ListType{ElemType: types.StringType},
			Description: "Use regex expressions to match which origin is allowed to enable CORS, for example, [\".*.test.com\"] can use to match all subdomain of test.com",
		},
	}),
}

func (s PluginCorsType) DecodeFomMap(v map[string]interface{}, pluginsType *PluginsType) {
	if v := v["cors"]; v != nil {
		jsonData := v.(map[string]interface{})
		item := PluginCorsType{}

		if v := jsonData["disable"]; v != nil {
			item.Disable = types.Bool{Value: v.(bool)}
		} else {
			item.Disable = types.Bool{Value: true}
		}

		if v := jsonData["max_age"]; v != nil {
			item.MaxAge = types.Number{Value: big.NewFloat(v.(float64))}
		} else {
			item.MaxAge = types.Number{Null: true}
		}

		if v := jsonData["allow_credential"]; v != nil {
			item.AllowCredential = types.Bool{Value: v.(bool)}
		} else {
			item.AllowCredential = types.Bool{Null: true}
		}

		if v := jsonData["allow_origins_by_regex"]; v != nil {
			var values []attr.Value
			for _, value := range v.([]interface{}) {
				values = append(values, types.String{Value: value.(string)})
			}
			item.AllowOriginsByRegex = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.AllowOriginsByRegex = types.List{Null: true}
		}

		if v := jsonData["expose_headers"]; v != nil {
			var values []attr.Value
			for _, value := range strings.Split(v.(string), ",") {
				values = append(values, types.String{Value: value})
			}
			item.ExposeHeaders = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.ExposeHeaders = types.List{Null: true}
		}

		if v := jsonData["allow_headers"]; v != nil {
			var values []attr.Value
			for _, value := range strings.Split(v.(string), ",") {
				values = append(values, types.String{Value: value})
			}
			item.AllowHeaders = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.AllowHeaders = types.List{Null: true}
		}

		if v := jsonData["allow_origins"]; v != nil {
			var values []attr.Value
			for _, value := range strings.Split(v.(string), ",") {
				values = append(values, types.String{Value: value})
			}
			item.AllowOrigins = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.AllowOrigins = types.List{Null: true}
		}

		if v := jsonData["allow_methods"]; v != nil {
			var values []attr.Value
			for _, value := range strings.Split(v.(string), ",") {
				values = append(values, types.String{Value: value})
			}
			item.AllowMethods = types.List{
				ElemType: types.StringType,
				Elems:    values,
			}
		} else {
			item.AllowMethods = types.List{Null: true}
		}

		pluginsType.Cors = &item
	}
}
func (s PluginCorsType) validate() error { return nil }

func (s PluginCorsType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	log.Printf("[DEBUG] got state: %v", s)
	if !s.AllowOrigins.Null {
		var values []string
		for _, v := range s.AllowOrigins.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["allow_origins"] = strings.Join(values, ",")
	}

	if !s.AllowMethods.Null {
		var values []string
		for _, v := range s.AllowMethods.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["allow_methods"] = strings.Join(values, ",")
	}

	if !s.AllowHeaders.Null {
		var values []string
		for _, v := range s.AllowHeaders.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["allow_headers"] = strings.Join(values, ",")
	}

	if !s.ExposeHeaders.Null {
		var values []string
		for _, v := range s.ExposeHeaders.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["expose_headers"] = strings.Join(values, ",")
	}

	if !s.AllowOriginsByRegex.Null {
		var values []string
		for _, v := range s.AllowOriginsByRegex.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["allow_origins_by_regex"] = values
	}

	if !s.MaxAge.Null {
		pluginValue["max_age"] = utils.TypeNumberToInt(s.MaxAge)
	}
	if !s.AllowCredential.Null {
		pluginValue["allow_credential"] = s.AllowCredential.Value
	}

	m["cors"] = pluginValue
}
