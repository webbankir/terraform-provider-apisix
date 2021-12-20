package apisix

//var schemeRouterResource = map[string]*schema.Schema{
//	"name": {
//		Type:     schema.TypeString,
//		Required: true,
//	},
//	"desc": {
//		Type:     schema.TypeString,
//		Optional: true,
//		Default:  "Managed by Terraform",
//	},
//	"uri": {
//		Type:          schema.TypeString,
//		Optional:      true,
//		ConflictsWith: []string{"uris"},
//	},
//	"uris": {
//		Type:     schema.TypeList,
//		Optional: true,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//		ConflictsWith: []string{"uri"},
//	},
//	"host": {
//		Type:          schema.TypeString,
//		Optional:      true,
//		ConflictsWith: []string{"hosts"},
//	},
//	"hosts": {
//		Type:     schema.TypeList,
//		Optional: true,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//		ConflictsWith: []string{"host"},
//	},
//
//	"remote_addr": {
//		Type:          schema.TypeString,
//		Optional:      true,
//		ConflictsWith: []string{"remote_addrs"},
//	},
//	"remote_addrs": {
//		Type:     schema.TypeList,
//		Optional: true,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//		ConflictsWith: []string{"remote_addr"},
//	},
//
//	"methods": {
//		Type:     schema.TypeList,
//		Optional: true,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//	},
//
//	"priority": {
//		Type:     schema.TypeInt,
//		Optional: true,
//	},
//
//	"is_enabled": {
//		Type:     schema.TypeBool,
//		Optional: true,
//		Default:  true,
//	},
//	"enable_websocket": {
//		Type:     schema.TypeBool,
//		Optional: true,
//		Default:  false,
//	},
//
//	"service_id": {
//		Type:     schema.TypeString,
//		Optional: true,
//	},
//
//	"upstream_id": {
//		Type:          schema.TypeString,
//		Optional:      true,
//		ConflictsWith: []string{"upstream"},
//	},
//
//	"labels": {
//		Optional: true,
//		Type:     schema.TypeMap,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//	},
//
//	"timeout": {
//		Required: true,
//		Type:     schema.TypeList,
//		Elem: &schema.Resource{
//			Schema: timeOutSchema,
//		},
//	},
//	"script": {
//		Optional:      true,
//		Type:          schema.TypeString,
//		ConflictsWith: []string{"plugin_config_id"},
//	},
//	"plugin_config_id": {
//		Optional:      true,
//		Type:          schema.TypeString,
//		ConflictsWith: []string{"script"},
//	},
//
//	"filter_func": {
//		Optional: true,
//		Type:     schema.TypeString,
//	},
//	"plugins": {
//		Optional: true,
//		Type:     schema.TypeList,
//		MaxItems: 1,
//		Elem: &schema.Resource{
//			Schema: map[string]*schema.Schema{
//				"custom_jsons": {
//					Optional: true,
//					Type:     schema.TypeList,
//					Elem: &schema.Schema{
//						Type: schema.TypeString,
//					},
//				},
//
//				"proxy_rewrite": {
//					Optional: true,
//					MaxItems: 1,
//					Type:     schema.TypeList,
//					Elem: &schema.Resource{
//						Schema: schemePluginProxyRewrite,
//					},
//				},
//				"ip_restriction": {
//					Optional: true,
//					MaxItems: 1,
//					Type:     schema.TypeList,
//					Elem: &schema.Resource{
//						Schema: schemePluginIpRestriction,
//					},
//				},
//			},
//		},
//	},
//
//	"upstream": {
//		Optional:      true,
//		Type:          schema.TypeList,
//		MaxItems:      1,
//		ConflictsWith: []string{"upstream_id"},
//		Elem: &schema.Resource{
//			Schema: map[string]*schema.Schema{
//				"type": {
//					Optional:     true,
//					Type:         schema.TypeString,
//					ValidateFunc: validation.StringInSlice([]string{"roundrobin", "chash", "ewma", "least_conn"}, false),
//					Default:      "roundrobin",
//				},
//				"service_name": {
//					Optional: true,
//					Type:     schema.TypeString,
//				},
//				"nodes": {
//					// TODO: Make!
//					//nodes	required, can't be used with service_name	Hash table or array. If it is a hash table, the key of the internal element is the upstream machine address list, the format is Address + (optional) Port, where the address part can be IP or domain name, such as 192.168.1.100:80, foo.com:80, etc. The value is the weight of node. If it is an array, each item is a hash table with key host/weight and optional port/priority. The nodes can be empty, which means it is a placeholder and will be filled later. Clients use such an upstream will get 502 response.	192.168.1.100:80
//					Optional: true,
//					Type:     schema.TypeString,
//					//ConflictsWith: []string{"service_name"},
//				},
//				"discovery_type": {
//					Optional:     true,
//					Type:         schema.TypeString,
//					ValidateFunc: validation.StringInSlice([]string{"dns", "consul_kv"}, false),
//				},
//				"timeout": {
//					Required: true,
//					Type:     schema.TypeList,
//					MaxItems: 1,
//					Elem: &schema.Resource{
//						Schema: timeOutSchema,
//					},
//				},
//				"name": {
//					Type:     schema.TypeString,
//					Optional: true,
//				},
//				"desc": {
//					Type:     schema.TypeString,
//					Optional: true,
//				},
//				"pass_host": {
//					Optional:     true,
//					Type:         schema.TypeString,
//					ValidateFunc: validation.StringInSlice([]string{"pass", "node", "rewrite"}, false),
//					Default:      "pass",
//				},
//				"scheme": {
//					Optional:     true,
//					Type:         schema.TypeString,
//					ValidateFunc: validation.StringInSlice([]string{"http", "https"}, false),
//					Default:      "http",
//				},
//				"retries": {
//					Optional: true,
//					Type:     schema.TypeInt,
//				},
//				"retry_timeout": {
//					Optional: true,
//					Type:     schema.TypeInt,
//				},
//				"labels": {
//					Optional: true,
//					Type:     schema.TypeMap,
//					Elem: &schema.Schema{
//						Type: schema.TypeString,
//					},
//				},
//				"upstream_host": {
//					Optional: true,
//					Type:     schema.TypeString,
//				},
//				"hash_on": {
//					Optional: true,
//					Type:     schema.TypeString,
//					Default:  "vars",
//				},
//
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
//
//	"last_updated": {
//		Type:     schema.TypeString,
//		Optional: true,
//		Computed: true,
//	},
//}
