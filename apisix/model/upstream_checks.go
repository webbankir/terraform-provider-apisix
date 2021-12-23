package model

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/plan_modifier"
	"github.com/webbankir/terraform-provider-apisix/apisix/utils"
	"github.com/webbankir/terraform-provider-apisix/apisix/validator"
	"math/big"
)

type UpstreamChecksType struct {
	Active  *UpstreamChecksActiveType  `tfsdk:"active"`
	Passive *UpstreamChecksPassiveType `tfsdk:"passive"`
}

type UpstreamChecksActiveType struct {
	Type                   types.String                       `tfsdk:"type"`
	Timeout                types.Number                       `tfsdk:"timeout"`
	Concurrency            types.Number                       `tfsdk:"concurrency"`
	HTTPPath               types.String                       `tfsdk:"http_path"`
	Host                   types.String                       `tfsdk:"host"`
	Port                   types.Number                       `tfsdk:"port"`
	HTTPSVerifyCertificate types.Bool                         `tfsdk:"https_verify_certificate"`
	ReqHeaders             types.List                         `tfsdk:"req_headers"`
	Healthy                *UpstreamChecksActiveHealthyType   `tfsdk:"healthy"`
	Unhealthy              *UpstreamChecksActiveUnhealthyType `tfsdk:"unhealthy"`
}

type UpstreamChecksPassiveType struct {
	Healthy   *UpstreamChecksPassiveHealthyType   `tfsdk:"healthy"`
	Unhealthy *UpstreamChecksPassiveUnhealthyType `tfsdk:"unhealthy"`
}

type UpstreamChecksActiveHealthyType struct {
	Interval     types.Number `tfsdk:"interval"`
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	Successes    types.Number `tfsdk:"successes"`
}

type UpstreamChecksActiveUnhealthyType struct {
	Interval     types.Number `tfsdk:"interval"`
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	TCPFailures  types.Number `tfsdk:"tcp_failures"`
	Timeouts     types.Number `tfsdk:"timeouts"`
	HTTPFailures types.Number `tfsdk:"http_failures"`
}

type UpstreamChecksPassiveHealthyType struct {
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	Successes    types.Number `tfsdk:"successes"`
}

type UpstreamChecksPassiveUnhealthyType struct {
	HTTPStatuses types.List   `tfsdk:"http_statuses"`
	TCPFailures  types.Number `tfsdk:"tcp_failures"`
	Timeouts     types.Number `tfsdk:"timeouts"`
	HTTPFailures types.Number `tfsdk:"http_failures"`
}

var UpstreamChecksSchemaAttribute = tfsdk.Attribute{
	Optional: true,
	Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"active": {
			Optional: true,
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
				"type": {
					Type:     types.StringType,
					Optional: true,
					Computed: true,
					Validators: []tfsdk.AttributeValidator{
						validator.StringInSlice("http", "https", "tcp"),
					},
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultString("http"),
					},
					Description: "The type of active check",
				},

				"timeout": {
					Type:     types.NumberType,
					Optional: true,
					Computed: true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultNumber(1),
					},
					Description: "The timeout period of the active check (unit: second)",
				},
				"concurrency": {
					Type:     types.NumberType,
					Optional: true,
					Computed: true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultNumber(10),
					},
					Description: "The number of targets to be checked at the same time during the active check",
				},
				"http_path": {
					Type:     types.StringType,
					Optional: true,
					Computed: true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultString("/"),
					},
					Description: "The HTTP request path that is actively checked",
				},
				"host": {
					Type:        types.StringType,
					Optional:    true,
					Description: "The hostname of the HTTP request actively checked",
				},
				"port": {
					Type:     types.NumberType,
					Optional: true,
					Computed: true,
					Validators: []tfsdk.AttributeValidator{
						validator.NumberGreatOrEqualThan(1),
						validator.NumberLessOrEqualThan(65535),
					},
					Description: "The host port of the HTTP request that is actively checked",
				},
				"https_verify_certificate": {
					Type:     types.BoolType,
					Optional: true,
					Computed: true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						plan_modifier.DefaultBool(true),
					},
					Description: "Active check whether to check the SSL certificate of the remote host when HTTPS type checking is used",
				},
				"req_headers": {
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Description: "Active check When using HTTP or HTTPS type checking, set additional request header information",
				},
				"healthy": {
					Optional: true,
					Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
						"interval": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(1),
							},
							Description: "Active check (healthy node) check interval (unit: second)",
						},
						"http_statuses": {
							Type:     types.ListType{ElemType: types.NumberType},
							Optional: true,
							Computed: true,
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultListOfNumbers(200, 302),
							},
							Description: "Active check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node",
						},

						"successes": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(2),
							},
							Description: "Active check (healthy node) check interval (unit: second)",
						},
					}),
				},

				"unhealthy": {
					Optional: true,
					Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
						"interval": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(1),
							},
							Description: "Active check (unhealthy node) check interval (unit: second)",
						},
						"http_statuses": {
							Type:     types.ListType{ElemType: types.NumberType},
							Optional: true,
							Computed: true,
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultListOfNumbers(404, 429, 500, 501, 503, 504, 505),
							},
							Description: "Active check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node",
						},

						"http_failures": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(5),
							},
							Description: "Active check (unhealthy node) HTTP or HTTPS type check, determine the number of times that the node is not healthy",
						},

						"tcp_failures": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(2),
							},
							Description: "Active check (unhealthy node) TCP type check, determine the number of times that the node is not healthy",
						},

						"timeouts": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(3),
							},
							Description: "Active check (unhealthy node) to determine the number of timeouts for unhealthy nodes",
						},
					}),
				},
			}),
		},

		"passive": {
			Optional: true,
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
				"healthy": {
					Optional: true,
					Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
						"http_statuses": {
							Type:     types.ListType{ElemType: types.NumberType},
							Optional: true,
							Computed: true,
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultListOfNumbers(200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308),
							},
							Description: "Passive check (healthy node) HTTP or HTTPS type check, the HTTP status code of the healthy node",
						},

						"successes": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(5),
							},
							Description: "Passive checks (healthy node) determine the number of times a node is healthy",
						},
					}),
				},
				"unhealthy": {
					Optional: true,
					Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
						"http_statuses": {
							Type:     types.ListType{ElemType: types.NumberType},
							Optional: true,
							Computed: true,
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultListOfNumbers(429, 500, 503),
							},
							Description: "Passive check (unhealthy node) HTTP or HTTPS type check, the HTTP status code of the non-healthy node",
						},

						"http_failures": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(5),
							},
							Description: "Passive check (unhealthy node) The number of times that the node is not healthy during HTTP or HTTPS type checking",
						},

						"tcp_failures": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(2),
							},
							Description: "Passive check (unhealthy node) When TCP type is checked, determine the number of times that the node is not healthy",
						},

						"timeouts": {
							Type:     types.NumberType,
							Optional: true,
							Computed: true,
							Validators: []tfsdk.AttributeValidator{
								validator.NumberGreatOrEqualThan(1),
								validator.NumberLessOrEqualThan(254),
							},
							PlanModifiers: []tfsdk.AttributePlanModifier{
								plan_modifier.DefaultNumber(7),
							},
							Description: "Passive checks (unhealthy node) determine the number of timeouts for unhealthy nodes",
						},
					}),
				},
			}),
		},
	}),
}

func UpstreamChecksStateToMap(state *UpstreamChecksType) *map[string]interface{} {
	if state == nil {
		return nil
	}

	result := make(map[string]interface{})

	if v := state.Active; v != nil {
		valueActive := make(map[string]interface{})

		if !v.Type.Null {
			valueActive["type"] = v.Type.Value
		}

		if !v.Timeout.Null {
			valueActive["timeout"] = utils.TypeNumberToInt(v.Timeout)
		}

		if !v.Concurrency.Null {
			valueActive["concurrency"] = utils.TypeNumberToInt(v.Concurrency)
		}

		if !v.HTTPPath.Null {
			valueActive["http_path"] = v.HTTPPath.Value
		}

		if !v.Host.Null {
			valueActive["host"] = v.Host.Value
		}

		if !v.Port.Null {
			valueActive["port"] = utils.TypeNumberToInt(v.Port)
		}

		if !v.HTTPSVerifyCertificate.Null {
			valueActive["https_verify_certificate"] = v.HTTPSVerifyCertificate.Value
		}

		if !v.ReqHeaders.Null {
			var values []string

			for _, v := range v.ReqHeaders.Elems {
				if !v.(types.String).Null {
					values = append(values, v.(types.String).Value)
				}
			}

			valueActive["req_headers"] = values
		}

		if healthy := v.Healthy; healthy != nil {
			valueActiveHealthy := make(map[string]interface{})

			if !healthy.Interval.Null {
				valueActiveHealthy["interval"] = utils.TypeNumberToInt(healthy.Interval)
			}

			if !healthy.HTTPStatuses.Null {
				var values []int

				for _, v := range healthy.HTTPStatuses.Elems {
					if !v.(types.String).Null {
						values = append(values, utils.TypeNumberToInt(v.(types.Number)))
					}
				}

				valueActive["http_statuses"] = values
			}

			if !healthy.Successes.Null {
				valueActiveHealthy["successes"] = utils.TypeNumberToInt(healthy.Successes)
			}

			valueActive["healthy"] = valueActiveHealthy
		}

		if unhealthy := v.Unhealthy; unhealthy != nil {
			valueActiveUnhealthy := make(map[string]interface{})

			if !unhealthy.Interval.Null {
				valueActiveUnhealthy["interval"] = utils.TypeNumberToInt(unhealthy.Interval)
			}

			if !unhealthy.TCPFailures.Null {
				valueActiveUnhealthy["tcp_failures"] = utils.TypeNumberToInt(unhealthy.TCPFailures)
			}

			if !unhealthy.Timeouts.Null {
				valueActiveUnhealthy["timeouts"] = utils.TypeNumberToInt(unhealthy.Timeouts)
			}

			if !unhealthy.HTTPFailures.Null {
				valueActiveUnhealthy["http_failures"] = utils.TypeNumberToInt(unhealthy.HTTPFailures)
			}

			if !unhealthy.HTTPStatuses.Null {
				var values []int

				for _, v := range unhealthy.HTTPStatuses.Elems {
					if !v.(types.String).Null {
						values = append(values, utils.TypeNumberToInt(v.(types.Number)))
					}
				}

				valueActiveUnhealthy["http_statuses"] = values
			}

			valueActive["unhealthy"] = valueActiveUnhealthy
		}

		result["active"] = valueActive
	}

	if v := state.Passive; v != nil {
		valuePassive := make(map[string]interface{})

		if healthy := v.Healthy; healthy != nil {
			valuePassiveHealthy := make(map[string]interface{})

			if !healthy.HTTPStatuses.Null {
				var values []int

				for _, v := range healthy.HTTPStatuses.Elems {
					if !v.(types.String).Null {
						values = append(values, utils.TypeNumberToInt(v.(types.Number)))
					}
				}

				valuePassiveHealthy["http_statuses"] = values
			}

			if !healthy.Successes.Null {
				valuePassiveHealthy["successes"] = utils.TypeNumberToInt(healthy.Successes)
			}

			valuePassive["healthy"] = valuePassiveHealthy
		}

		if unhealthy := v.Unhealthy; unhealthy != nil {
			valuePassiveUnhealthy := make(map[string]interface{})

			if !unhealthy.TCPFailures.Null {
				valuePassiveUnhealthy["tcp_failures"] = utils.TypeNumberToInt(unhealthy.TCPFailures)
			}

			if !unhealthy.Timeouts.Null {
				valuePassiveUnhealthy["timeouts"] = utils.TypeNumberToInt(unhealthy.Timeouts)
			}

			if !unhealthy.HTTPFailures.Null {
				valuePassiveUnhealthy["http_failures"] = utils.TypeNumberToInt(unhealthy.HTTPFailures)
			}

			if !unhealthy.HTTPStatuses.Null {
				var values []int

				for _, v := range unhealthy.HTTPStatuses.Elems {
					if !v.(types.String).Null {
						values = append(values, utils.TypeNumberToInt(v.(types.Number)))
					}
				}

				valuePassiveUnhealthy["http_statuses"] = values
			}

			valuePassive["unhealthy"] = valuePassiveUnhealthy
		}

		result["passive"] = valuePassive
	}

	return &result

}

func UpstreamChecksMapToState(data map[string]interface{}) *UpstreamChecksType {
	checks := data["checks"]

	if checks == nil {
		return nil
	}

	result := UpstreamChecksType{}

	if v := checks.(map[string]interface{})["active"]; v != nil {
		activeItem := UpstreamChecksActiveType{}
		value := v.(map[string]interface{})

		if v := value["type"]; v != nil {
			activeItem.Type = types.String{Value: v.(string)}
		} else {
			activeItem.Type = types.String{Null: true}
		}

		if v := value["timeout"]; v != nil {
			activeItem.Timeout = types.Number{Value: big.NewFloat(v.(float64))}
		} else {
			activeItem.Timeout = types.Number{Null: true}
		}

		if v := value["concurrency"]; v != nil {
			activeItem.Concurrency = types.Number{Value: big.NewFloat(v.(float64))}
		} else {
			activeItem.Concurrency = types.Number{Null: true}
		}

		if v := value["http_path"]; v != nil {
			activeItem.HTTPPath = types.String{Value: v.(string)}
		} else {
			activeItem.HTTPPath = types.String{Null: true}
		}

		if v := value["host"]; v != nil {
			activeItem.Host = types.String{Value: v.(string)}
		} else {
			activeItem.Host = types.String{Null: true}
		}

		if v := value["port"]; v != nil {
			activeItem.Port = types.Number{Value: big.NewFloat(v.(float64))}
		} else {
			activeItem.Port = types.Number{Null: true}
		}

		if v := value["https_verify_certificate"]; v != nil {
			activeItem.HTTPSVerifyCertificate = types.Bool{Value: v.(bool)}
		} else {
			activeItem.HTTPSVerifyCertificate = types.Bool{Null: true}
		}

		if v := value["req_headers"]; v != nil {
			var values []attr.Value
			for _, v := range v.([]interface{}) {
				values = append(values, types.String{Value: v.(string)})
			}
			activeItem.ReqHeaders = types.List{ElemType: types.StringType, Elems: values}
		} else {
			activeItem.ReqHeaders = types.List{Null: true}
		}

		if v := value["healthy"]; v != nil {
			healthy := v.(map[string]interface{})
			activeHealthyItem := UpstreamChecksActiveHealthyType{}

			if v := healthy["interval"]; v != nil {
				activeHealthyItem.Interval = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				activeHealthyItem.Interval = types.Number{Null: true}
			}

			if v := healthy["successes"]; v != nil {
				activeHealthyItem.Successes = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				activeHealthyItem.Successes = types.Number{Null: true}
			}

			if v := healthy["http_statuses"]; v != nil {
				var values []attr.Value
				for _, v := range v.([]interface{}) {
					values = append(values, types.String{Value: v.(string)})
				}
				activeHealthyItem.HTTPStatuses = types.List{ElemType: types.StringType, Elems: values}
			} else {
				activeHealthyItem.HTTPStatuses = types.List{Null: true}
			}

			activeItem.Healthy = &activeHealthyItem
		}

		if v := value["unhealthy"]; v != nil {
			unhealthy := v.(map[string]interface{})
			activeUnhealthyItem := UpstreamChecksActiveUnhealthyType{}

			if v := unhealthy["interval"]; v != nil {
				activeUnhealthyItem.Interval = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				activeUnhealthyItem.Interval = types.Number{Null: true}
			}

			if v := unhealthy["tcp_failures"]; v != nil {
				activeUnhealthyItem.TCPFailures = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				activeUnhealthyItem.TCPFailures = types.Number{Null: true}
			}

			if v := unhealthy["timeouts"]; v != nil {
				activeUnhealthyItem.Timeouts = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				activeUnhealthyItem.Timeouts = types.Number{Null: true}
			}

			if v := unhealthy["http_failures"]; v != nil {
				activeUnhealthyItem.HTTPFailures = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				activeUnhealthyItem.HTTPFailures = types.Number{Null: true}
			}

			if v := unhealthy["http_statuses"]; v != nil {
				var values []attr.Value
				for _, v := range v.([]interface{}) {
					values = append(values, types.String{Value: v.(string)})
				}
				activeUnhealthyItem.HTTPStatuses = types.List{ElemType: types.StringType, Elems: values}
			} else {
				activeUnhealthyItem.HTTPStatuses = types.List{Null: true}
			}

			activeItem.Unhealthy = &activeUnhealthyItem
		}

		result.Active = &activeItem
	}

	if v := checks.(map[string]interface{})["passive"]; v != nil {
		passiveItem := UpstreamChecksPassiveType{}
		value := v.(map[string]interface{})

		if v := value["healthy"]; v != nil {
			healthy := v.(map[string]interface{})
			passiveHealthyItem := UpstreamChecksPassiveHealthyType{}

			if v := healthy["successes"]; v != nil {
				passiveHealthyItem.Successes = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				passiveHealthyItem.Successes = types.Number{Null: true}
			}

			if v := healthy["http_statuses"]; v != nil {
				var values []attr.Value
				for _, v := range v.([]interface{}) {
					values = append(values, types.String{Value: v.(string)})
				}
				passiveHealthyItem.HTTPStatuses = types.List{ElemType: types.StringType, Elems: values}
			} else {
				passiveHealthyItem.HTTPStatuses = types.List{Null: true}
			}

			passiveItem.Healthy = &passiveHealthyItem
		}

		if v := value["unhealthy"]; v != nil {
			unhealthy := v.(map[string]interface{})
			passiveUnhealthyItem := UpstreamChecksPassiveUnhealthyType{}

			if v := unhealthy["tcp_failures"]; v != nil {
				passiveUnhealthyItem.TCPFailures = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				passiveUnhealthyItem.TCPFailures = types.Number{Null: true}
			}

			if v := unhealthy["timeouts"]; v != nil {
				passiveUnhealthyItem.Timeouts = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				passiveUnhealthyItem.Timeouts = types.Number{Null: true}
			}

			if v := unhealthy["http_failures"]; v != nil {
				passiveUnhealthyItem.HTTPFailures = types.Number{Value: big.NewFloat(v.(float64))}
			} else {
				passiveUnhealthyItem.HTTPFailures = types.Number{Null: true}
			}

			if v := unhealthy["http_statuses"]; v != nil {
				var values []attr.Value
				for _, v := range v.([]interface{}) {
					values = append(values, types.String{Value: v.(string)})
				}
				passiveUnhealthyItem.HTTPStatuses = types.List{ElemType: types.StringType, Elems: values}
			} else {
				passiveUnhealthyItem.HTTPStatuses = types.List{Null: true}
			}

			passiveItem.Unhealthy = &passiveUnhealthyItem
		}

		result.Passive = &passiveItem
	}

	return &result
}
