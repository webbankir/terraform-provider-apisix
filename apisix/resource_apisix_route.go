package apisix

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"log"
)

type ResourceRouteType struct {
	p      provider
	client ApiClient
}

func (r ResourceRouteType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return ResourceRouteType{
		p:      *(p.(*provider)),
		client: getCl(),
	}, nil
}

func (r ResourceRouteType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"desc": {
				Type:     types.StringType,
				Optional: true,
			},
			"uri": {
				Type:     types.StringType,
				Optional: true,
			},
			"uris": {
				Type:     types.ListType{ElemType: types.StringType},
				Optional: true,
			},
			"host": {
				Type:     types.StringType,
				Optional: true,
			},
			"hosts": {
				Type:     types.ListType{ElemType: types.StringType},
				Optional: true,
			},
			"remote_addr": {
				Type:     types.StringType,
				Optional: true,
			},
			"remote_addrs": {
				Type:     types.ListType{ElemType: types.StringType},
				Optional: true,
			},

			"methods": {
				Type:     types.ListType{ElemType: types.StringType},
				Optional: true,
			},
			"priority": {
				Type:     types.NumberType,
				Optional: true,
			},

			"is_enabled": {
				Type:     types.BoolType,
				Optional: true,
				//Default:  true,
			},
			"enable_websocket": {
				Type:     types.BoolType,
				Optional: true,
				//Default:  false,
			},

			"service_id": {
				Type:     types.StringType,
				Optional: true,
			},
			"upstream_id": {
				Type:     types.StringType,
				Optional: true,
				//ConflictsWith: []string{"upstream"},
			},

			"labels": {
				Optional: true,
				Type:     types.MapType{ElemType: types.StringType},
			},

			"timeout": {
				Required: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"connect": {
						Required: true,
						Type:     types.NumberType,
					},
					"send": {
						Required: true,
						Type:     types.NumberType,
					},
					"read": {
						Required: true,
						Type:     types.NumberType,
					},
				}),
			},
			"script": {
				Optional: true,
				Type:     types.StringType,
				//ConflictsWith: []string{"plugin_config_id"},
			},
			"plugin_config_id": {
				Optional: true,
				Type:     types.StringType,
				//ConflictsWith: []string{"script"},
			},

			"filter_func": {
				Optional: true,
				Type:     types.StringType,
			},
			"plugins": {
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"custom_jsons": {
						Optional: true,
						Type:     types.MapType{ElemType: types.StringType},
					},
					"proxy_rewrite": {
						Optional: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"disable": {
								Required: true,
								Type:     types.BoolType,
							},
							"scheme": {
								Optional: true,
								Type:     types.StringType,
								//ValidateFunc: validation.StringInSlice([]string{"http", "https"}, false),
							},
							"method": {
								Optional: true,
								Type:     types.StringType,
								//ValidateFunc: validation.StringInSlice([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS", "MKCOL", "COPY", "MOVE", "PROPFIND", "PROPFIND", "LOCK", "UNLOCK", "PATCH", "TRACE"}, false),
							},
							"uri": {
								Optional: true,
								Type:     types.StringType,
							},
							"host": {
								Optional: true,
								Type:     types.StringType,
							},
							"headers": {
								Optional: true,
								Type:     types.MapType{ElemType: types.StringType},
							},
							"regex_uri": {
								Optional: true,
								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
									"regex": {
										Required: true,
										Type:     types.StringType,
									},
									"replacement": {
										Required: true,
										Type:     types.StringType,
									},
								}),
							},
						}),
					},
					"ip_restriction": {
						Optional: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"disable": {
								Required: true,
								Type:     types.BoolType,
							},
							"whitelist": {
								Optional: true,
								Type:     types.ListType{ElemType: types.StringType},
							},
							"blacklist": {
								Optional: true,
								Type:     types.ListType{ElemType: types.StringType},
							},
							"message": {
								Optional: true,
								Type:     types.StringType,
							},
						}),
					},
				}),
			},
			"upstream": {
				Optional: true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"type": {
						Type:     types.StringType,
						Optional: true,
						//					ValidateFunc: validation.StringInSlice([]string{"roundrobin", "chash", "ewma", "least_conn"}, false),
						//					Default:      "roundrobin",
					},
					"service_name": {
						Type:     types.StringType,
						Optional: true,
					},
					"discovery_type": {
						Type:     types.StringType,
						Optional: true,
					},
					"timeout": {
						Required: true,
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"connect": {
								Required: true,
								Type:     types.NumberType,
							},
							"send": {
								Required: true,
								Type:     types.NumberType,
							},
							"read": {
								Required: true,
								Type:     types.NumberType,
							},
						}),
					},
					"name": {
						Type:     types.StringType,
						Optional: true,
					},
					"desc": {
						Type:     types.StringType,
						Optional: true,
					},

					"pass_host": {
						Type:     types.StringType,
						Optional: true,
						//					ValidateFunc: validation.StringInSlice([]string{"pass", "node", "rewrite"}, false),
						//					Default:      "pass",
					},
					"scheme": {
						Type:     types.StringType,
						Optional: true,
						//					ValidateFunc: validation.StringInSlice([]string{"http", "https"}, false),
						//					Default:      "http",

					},
					"retries": {
						Type:     types.NumberType,
						Optional: true,
					},
					"retry_timeout": {
						Type:     types.NumberType,
						Optional: true,
					},
					"upstream_host": {
						Type:     types.StringType,
						Optional: true,
					},
					"hash_on": {
						Type:     types.StringType,
						Optional: true,
					},
					"labels": {
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
				}),
			},

			//				//	optional	This option is only valid if the type is chash. Supported types vars(Nginx variables), header(custom header), cookie, consumer, the default value is vars.
			//				//key	optional	This option is only valid if the type is chash. Find the corresponding node id according to hash_on and key. When hash_on is set as vars, key is the required parameter, for now, it support nginx built-in variables like uri, server_name, server_addr, request_uri, remote_port, remote_addr, query_string, host, hostname, arg_***, arg_*** is arguments in the request line, Nginx variables list. When hash_on is set as header, key is the required parameter, and header name is customized. When hash_on is set to cookie, key is the required parameter, and cookie name is customized. When hash_on is set to consumer, key does not need to be set. In this case, the key adopted by the hash algorithm is the consumer_name authenticated. If the specified hash_on and key can not fetch values, it will be fetch remote_addr by default.
			//				//checks	optional	Configure the parameters of the health check. For details, refer to health-check.
			//				//create_time	optional	epoch timestamp in second, like 1602883670, will be created automatically if missing	1602883670
			//				//update_time	optional	epoch timestamp in second, like 1602883670, will be created automatically if missing	1602883670
			//				//tls.client_cert	optional	Set the client certificate when connecting to TLS upstream, see below for more details
			//				//tls.client_key	optional	Set the client private key when connecting to TLS upstream, see below for more details
			//				//keepalive_pool.size	optional	Set keepalive directive dynamically, see below for more details
			//				//keepalive_pool.idle_timeout	optional	Set keepalive_timeout directive dynamically, see below for more details
			//				//keepalive_pool.requests	optional	Set keepalive_requests directive dynamically, see below for more details
			//
			//			},
			//		},
			//	},
		},
	}, nil
}

func (r ResourceRouteType) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	var plan RouteType
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
	// Check and build object

	//Description: plan.Description.Value,
	requestObjectJson, err := r.stateToJson(plan)
	if err != nil {
		response.Diagnostics.AddError(
			"Error in transformation from state to map",
			"Unexpected error: "+err.Error(),
		)
		return
	}
	log.Printf("[DEBUG] Request Object is %v", string(requestObjectJson))

	// New State
	//

	//	diags = resp.State.Set(ctx, &newState)
	//	resp.Diagnostics.Append(diags...)
	//	if resp.Diagnostics.HasError() {
	//		return
	//	}
	//
	//	snis, err := ParseCert(plan.Certificate.Value, plan.PrivateKey.Value)
	//

	//
	//	kk := SSL{
	//		Certificate: plan.Certificate.Value,
	//		PrivateKey:  plan.PrivateKey.Value,
	//		SNIS:        snis,
	//	}
	//
	//	result, err := r.client.CreateSsl(kk)
	//
	//	if err != nil {
	//		resp.Diagnostics.AddError(
	//			"Can't create new ssl resource",
	//			"Unexpected error: "+err.Error(),
	//		)
	//		return
	//	}
	//
	//	var newState = SslCertificate{
	//		ID:          types.String{Value: result.ID},
	//		Certificate: types.String{Value: result.Certificate},
	//		PrivateKey:  types.String{Value: kk.PrivateKey},
	//	}
	//

}
func (r ResourceRouteType) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	//TODO implement me
	panic("implement me")
}

func (r ResourceRouteType) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	//TODO implement me
	panic("implement me")
}

func (r ResourceRouteType) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	//TODO implement me
	panic("implement me")
}

func (r ResourceRouteType) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	//TODO implement me
	panic("implement me")
}

func (r ResourceRouteType) JsonToState(jsonR []byte) (RouteType, error) {
	jsonMap := make(map[string]interface{})
	newState := RouteType{}
	if err := json.Unmarshal(jsonR, &jsonMap); err != nil {
		return RouteType{}, err
	}
	if v := jsonMap["name"]; v != nil {
		newState.Name = types.String{Value: v.(string)}
	}

	// sudo yum -y install https://packages.endpoint.com/rhel/7/os/x86_6
	panic("JOPKA")
}

func (r ResourceRouteType) stateToJson(state RouteType) ([]byte, error) {
	requestObject := make(map[string]interface{})

	if !state.Name.Null {
		requestObject["name"] = state.Name.Value
	}
	if !state.Description.Null {
		requestObject["desc"] = state.Description.Value
	} else {
		requestObject["desc"] = "Managed by terraform"
	}

	if !state.Description.Null {
		requestObject["desc"] = state.Description.Value
	}

	if !state.Uri.Null {
		requestObject["uri"] = state.Uri.Value
	}

	if !state.Uris.Null {
		var values []string
		for _, v := range state.Uris.Elems {
			values = append(values, v.(types.String).Value)
		}
		requestObject["uris"] = values
	}

	if !state.Host.Null {
		requestObject["host"] = state.Host.Value
	}

	if !state.Hosts.Null {
		var values []string
		for _, v := range state.Hosts.Elems {
			values = append(values, v.(types.String).Value)
		}
		requestObject["host"] = values
	}

	if !state.RemoteAddr.Null {
		requestObject["remote_addr"] = state.RemoteAddr.Value
	}

	if !state.RemoteAddrs.Null {
		var values []string
		for _, v := range state.RemoteAddrs.Elems {
			values = append(values, v.(types.String).Value)
		}
		requestObject["remote_addrs"] = values
	}

	if !state.Methods.Null {
		var values []string
		for _, v := range state.Methods.Elems {
			values = append(values, v.(types.String).Value)
		}
		requestObject["methods"] = values
	}

	if !state.Priority.Null {
		requestObject["priority"] = state.Priority.Value
	}

	if !state.IsEnabled.Null {
		if state.IsEnabled.Value {
			requestObject["status"] = 1
		} else {
			requestObject["status"] = 0
		}
	}

	if !state.EnableWebsocket.Null {
		requestObject["enable_websocket"] = state.EnableWebsocket.Value
	}

	if !state.ServiceId.Null {
		requestObject["service_id"] = state.ServiceId.Value
	}

	if !state.UpstreamId.Null {
		requestObject["upstream_id"] = state.UpstreamId.Value
	}

	if !state.Labels.Null {
		values := make(map[string]interface{})
		for k, v := range state.Labels.Elems {
			values[k] = v.(types.String).Value
		}
		requestObject["labels"] = values
	}

	if state.Timeout != nil {
		requestObject["timeout"] = map[string]interface{}{
			"connect": state.Timeout.Connect.Value,
			"send":    state.Timeout.Send.Value,
			"read":    state.Timeout.Read.Value,
		}
	}

	if !state.Script.Null {
		requestObject["script"] = state.Script.Value
	}

	if !state.PluginConfigId.Null {
		requestObject["plugin_config_id"] = state.PluginConfigId.Value
	}

	if !state.FilterFunc.Null {
		requestObject["filter_func"] = state.FilterFunc.Value
	}

	plugins := make(map[string]interface{})
	if state.Plugins != nil {
		planPlugins := state.Plugins
		if !planPlugins.CustomJsons.Null {
			for k, v := range planPlugins.CustomJsons.Elems {
				jsonItem := make(map[string]interface{})

				if err := json.Unmarshal([]byte(v.(types.String).Value), &jsonItem); err != nil {
					return []byte{}, err
				}
				plugins[k] = jsonItem
			}
		}

		if planPlugins.ProxyRewrite != nil {
			planPlugins.ProxyRewrite.EncodeToMap(plugins)
		}

		if planPlugins.IpRestriction != nil {
			planPlugins.IpRestriction.EncodeToMap(plugins)
		}
	}

	requestObject["plugins"] = plugins

	return json.Marshal(requestObject)
}
