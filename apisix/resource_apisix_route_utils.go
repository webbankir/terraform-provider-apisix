package apisix

//func __updateRouteResource(d *schema.ResourceData, data Route) error {
//	if err := d.Set("name", data.Name); err != nil {
//		return err
//	}
//
//	if len(data.Description) > 0 {
//		if err := d.Set("desc", data.Description); err != nil {
//			return err
//		}
//	}
//
//	if len(data.Uri) > 0 {
//		if err := d.Set("uri", data.Uri); err != nil {
//			return err
//		}
//	}
//
//	if len(data.Uris) > 0 {
//		if err := d.Set("uris", data.Uris); err != nil {
//			return err
//		}
//	}
//
//	if len(data.Host) > 0 {
//		if err := d.Set("host", data.Host); err != nil {
//			return err
//		}
//	}
//
//	if len(data.Hosts) > 0 {
//		if err := d.Set("hosts", data.Hosts); err != nil {
//			return err
//		}
//	}
//
//	if len(data.RemoteAddr) > 0 {
//		if err := d.Set("remote_addr", data.RemoteAddr); err != nil {
//			return err
//		}
//	}
//
//	if len(data.RemoteAddrs) > 0 {
//		if err := d.Set("remote_addrs", data.RemoteAddrs); err != nil {
//			return err
//		}
//	}
//
//	if len(data.Methods) > 0 {
//		if err := d.Set("methods", data.Methods); err != nil {
//			return err
//		}
//	}
//
//	if err := d.Set("priority", data.Priority); err != nil {
//		return err
//	}
//
//	if data.Status == 1 {
//		if err := d.Set("is_enabled", true); err != nil {
//			return err
//		}
//	} else {
//		if err := d.Set("is_enabled", false); err != nil {
//			return err
//		}
//	}
//
//	if err := d.Set("enable_websocket", data.EnableWebsocket); err != nil {
//		return err
//	}
//
//	if len(data.ServiceId) > 0 {
//		if err := d.Set("service_id", data.ServiceId); err != nil {
//			return err
//		}
//	}
//
//	if len(data.UpstreamId) > 0 {
//		if err := d.Set("upstream_id", data.UpstreamId); err != nil {
//			return err
//		}
//	}
//
//	if len(data.Labels) > 0 {
//		if err := d.Set("labels", data.Labels); err != nil {
//			return err
//		}
//	}
//	if len(data.Script) > 0 {
//		if err := d.Set("script", data.Script); err != nil {
//			return err
//		}
//	}
//
//	if len(data.PluginConfigId) > 0 {
//		if err := d.Set("plugin_config_id", data.PluginConfigId); err != nil {
//			return err
//		}
//	}
//
//	if len(data.FilterFunc) > 0 {
//		if err := d.Set("filter_func", data.FilterFunc); err != nil {
//			return err
//		}
//	}
//
//	timeout := map[string]interface{}{
//		"connect": data.Timeout.Connect,
//		"send":    data.Timeout.Send,
//		"read":    data.Timeout.Read,
//	}
//
//	if err := d.Set("timeout", []interface{}{timeout}); err != nil {
//		return err
//	}
//
//	upstream := __updateRouteResourceUpstream(data)
//	if len(upstream) > 0 {
//		if err := d.Set("upstream", upstream); err != nil {
//			return fmt.Errorf("can't update upstream shit, with error, %v, data is %v", err, upstream)
//		}
//	}
//
//	if v := data.Plugins; len(v) > 0 {
//		plugins := make(map[string]interface{})
//		if plugin := v["proxy-rewrite"].(map[string]interface{}); plugin != nil {
//			item := map[string]interface{}{}
//
//			if v := plugin["disable"]; v != nil {
//				item["disable"] = v
//			}
//
//			if v := plugin["scheme"]; v != nil {
//				item["scheme"] = v
//			}
//
//			if v := plugin["uri"]; v != nil {
//				item["uri"] = v
//			}
//
//			if v := plugin["method"]; v != nil {
//				item["method"] = v
//			}
//
//			if v := plugin["regex_uri"]; v != nil {
//				item["regex_uri"] = map[string]interface{}{
//					"regex":       v.([]interface{})[0],
//					"replacement": v.([]interface{})[1],
//				}
//			}
//
//			if v := plugin["host"]; v != nil {
//				item["host"] = v
//			}
//
//			if v := plugin["headers"]; v != nil {
//				item["headers"] = v
//			}
//
//			plugins["proxy_rewrite"] = item
//		}
//
//		if err := d.Set("plugins", []map[string]interface{}{plugins}); err != nil {
//			return nil
//		}
//
//	}
//
//	//nodes	required, can't be used with service_name	Hash table or array. If it is a hash table, the key of the internal element is the upstream machine address list, the format is Address + (optional) Port, where the address part can be IP or domain name, such as 192.168.1.100:80, foo.com:80, etc. The value is the weight of node. If it is an array, each item is a hash table with key host/weight and optional port/priority. The nodes can be empty, which means it is a placeholder and will be filled later. Clients use such an upstream will get 502 response.	192.168.1.100:80
//
//	//hash_on	optional	This option is only valid if the type is chash. Supported types vars(Nginx variables), header(custom header), cookie, consumer, the default value is vars.
//	//key	optional	This option is only valid if the type is chash. Find the corresponding node id according to hash_on and key. When hash_on is set as vars, key is the required parameter, for now, it support nginx built-in variables like uri, server_name, server_addr, request_uri, remote_port, remote_addr, query_string, host, hostname, arg_***, arg_*** is arguments in the request line, Nginx variables list. When hash_on is set as header, key is the required parameter, and header name is customized. When hash_on is set to cookie, key is the required parameter, and cookie name is customized. When hash_on is set to consumer, key does not need to be set. In this case, the key adopted by the hash algorithm is the consumer_name authenticated. If the specified hash_on and key can not fetch values, it will be fetch remote_addr by default.
//	//checks	optional	Configure the parameters of the health check. For details, refer to health-check.
//
//	//tls.client_cert	optional	Set the client certificate when connecting to TLS upstream, see below for more details
//	//tls.client_key	optional	Set the client private key when connecting to TLS upstream, see below for more details
//	//keepalive_pool.size	optional	Set keepalive directive dynamically, see below for more details
//	//keepalive_pool.idle_timeout	optional	Set keepalive_timeout directive dynamically, see below for more details
//	//keepalive_pool.requests	optional	Set keepalive_requests directive dynamically, see below for more details
//
//	//			"plugin_proxy_rewrite": {
//	//			"plugins": {
//
//	return nil
//}
//
//func __updateRouteResourceUpstream(data Route) []map[string]interface{} {
//	result := make(map[string]interface{})
//
//	if len(data.Upstream.Type) == 0 {
//		return []map[string]interface{}{}
//	}
//
//	result["type"] = data.Upstream.Type
//
//	if len(data.Upstream.ServiceName) > 0 {
//		result["service_name"] = data.Upstream.ServiceName
//	}
//
//	if len(data.Upstream.DiscoveryType) > 0 {
//		result["discovery_type"] = data.Upstream.DiscoveryType
//	}
//
//	if len(data.Upstream.Name) > 0 {
//		result["name"] = data.Upstream.Name
//	}
//
//	if len(data.Upstream.Description) > 0 {
//		result["desc"] = data.Upstream.Description
//	}
//
//	if len(data.Upstream.PassHost) > 0 {
//		result["pass_host"] = data.Upstream.PassHost
//	}
//
//	if len(data.Upstream.UpstreamHost) > 0 {
//		result["upstream_host"] = data.Upstream.UpstreamHost
//	}
//
//	if len(data.Upstream.Scheme) > 0 {
//		result["scheme"] = data.Upstream.Scheme
//	}
//
//	if len(data.Upstream.HashOn) > 0 {
//		result["hash_on"] = data.Upstream.HashOn
//	}
//
//	if len(data.Upstream.Labels) > 0 {
//		result["labels"] = data.Upstream.Labels
//	}
//
//	result["timeout"] = []map[string]interface{}{
//		{
//			"connect": data.Upstream.Timeout.Connect,
//			"read":    data.Upstream.Timeout.Read,
//			"send":    data.Upstream.Timeout.Send,
//		},
//	}
//
//	if data.Upstream.Retries > 0 {
//		result["retries"] = data.Upstream.Retries
//	}
//
//	if data.Upstream.Retries > 0 {
//		result["retry_timeout"] = data.Upstream.RetryTimeout
//	}
//
//	log.Printf("[DEBUG] Upstream map: %v", result)
//	return []map[string]interface{}{result}
//}
//
//func __createRouteObject(d *schema.ResourceData) (Route, error) {
//	err := checkOneOf(d, "uri", "uris")
//	if err != nil {
//		return Route{}, err
//	}
//
//	err = checkOneOf(d, "upstream", "upstream_id")
//	if err != nil {
//		return Route{}, err
//	}
//
//	err = checkOneOfOptional(d, "host", "hosts")
//	if err != nil {
//		return Route{}, err
//	}
//
//	err = checkOneOfOptional(d, "remote_addr", "remote_addrs")
//	if err != nil {
//		return Route{}, err
//	}
//	status := 0
//	if d.Get("is_enabled").(bool) {
//		status = 1
//	}
//
//	route := Route{
//		Name:   d.Get("name").(string),
//		Status: status,
//	}
//
//	if value, exists := d.GetOkExists("desc"); exists {
//		route.Description = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("uri"); exists {
//		route.Uri = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("uris"); exists {
//		var values []string
//		for _, v := range value.([]interface{}) {
//			values = append(values, v.(string))
//		}
//		route.Uris = values
//	}
//
//	if value, exists := d.GetOkExists("host"); exists {
//		route.Host = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("hosts"); exists {
//		var values []string
//		for _, v := range value.([]interface{}) {
//			values = append(values, v.(string))
//		}
//		route.Hosts = values
//	}
//
//	if value, exists := d.GetOkExists("remote_addr"); exists {
//		route.RemoteAddr = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("remote_addrs"); exists {
//		var values []string
//		for _, v := range value.([]interface{}) {
//			values = append(values, v.(string))
//		}
//		route.RemoteAddrs = values
//	}
//
//	if value, exists := d.GetOkExists("methods"); exists {
//
//		var values []string
//		for _, v := range value.([]interface{}) {
//			if !stringContainsInSlice(HttpMethods, v.(string)) {
//				return Route{}, fmt.Errorf("found unknown HTTP method %v, list of supported methods:%v", v, HttpMethods)
//			}
//			values = append(values, v.(string))
//		}
//		route.Methods = values
//	}
//
//	if value, exists := d.GetOkExists("priority"); exists {
//		route.Priority = value.(int)
//	} else {
//		route.Priority = 0
//	}
//
//	if value, exists := d.GetOkExists("enable_websocket"); exists {
//		route.EnableWebsocket = value.(bool)
//	}
//
//	if value, exists := d.GetOkExists("labels"); exists {
//		route.Labels = value.(map[string]interface{})
//	}
//
//	if value, exists := d.GetOkExists("timeout"); exists {
//		route.Timeout = __createTimeOutObject(value)
//	}
//
//	if value, exists := d.GetOkExists("plugins"); exists {
//		plugins, err := __createRouteObjectPlugins(value)
//		if err != nil {
//			return Route{}, err
//		}
//		route.Plugins = plugins
//	}
//
//	log.Printf("[DEBUG] got item for plugins: %v", route.Plugins)
//
//	if value, exists := d.GetOkExists("upstream_id"); exists {
//		route.UpstreamId = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("service_id"); exists {
//		route.ServiceId = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("script"); exists {
//		route.Script = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("plugin_config_id"); exists {
//		route.PluginConfigId = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("filter_func"); exists {
//		route.FilterFunc = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("upstream"); exists {
//		route.Upstream = __createRouteObjectUpstream(value)
//	}
//
//	return route, nil
//}
//
//func __createTimeOutObject(value interface{}) Timeout {
//	values := value.([]interface{})
//	timeout := Timeout{}
//
//	if len(values) == 1 {
//		item := values[0].(map[string]interface{})
//
//		if v := item["connect"]; v != nil {
//			timeout.Connect = v.(int)
//		}
//
//		if v := item["send"]; v != nil {
//			timeout.Send = v.(int)
//		}
//
//		if v := item["read"]; v != nil {
//			timeout.Read = v.(int)
//		}
//
//	}
//
//	return timeout
//}
//
//func __createRouteObjectPlugins(value interface{}) (map[string]interface{}, error) {
//	values := value.([]interface{})
//	plugins := make(map[string]interface{})
//
//	if len(values) == 1 {
//		item := values[0].(map[string]interface{})
//
//		if v := item["custom_jsons"]; v != nil {
//			for _, item := range v.([]interface{}) {
//				r := make(map[string]interface{})
//				err := json.Unmarshal([]byte(item.(string)), &r)
//				if err != nil {
//					return plugins, err
//				}
//				for k, v := range r {
//					plugins[k] = v
//				}
//			}
//		}
//
//		if v := item["proxy_rewrite"]; v != nil {
//			plugins["proxy-rewrite"] = __createPluginsRewriteProxyObject(v)
//		}
//
//		if v := item["ip_restriction"]; v != nil {
//			plugins["ip-restriction"] = __createPluginsIpRestriction(v)
//		}
//	}
//
//	return plugins, nil
//}
//
//func __createPluginsRewriteProxyObject(v interface{}) map[string]interface{} {
//	vCast := v.([]interface{})[0].(map[string]interface{})
//	r := map[string]interface{}{
//		"disable": vCast["disable"].(bool),
//	}
//
//	if itemValue := vCast["scheme"].(string); len(itemValue) > 0 {
//		r["scheme"] = itemValue
//	}
//
//	if itemValue := vCast["uri"].(string); len(itemValue) > 0 {
//		r["uri"] = itemValue
//	}
//
//	if itemValue := vCast["method"].(string); len(itemValue) > 0 {
//		r["method"] = itemValue
//	}
//
//	if itemValue := vCast["regex_uri"].([]interface{}); len(itemValue) > 0 {
//		vRegexUriCast := itemValue[0].(map[string]interface{})
//		r["regex_uri"] = []string{vRegexUriCast["regex"].(string), vRegexUriCast["replacement"].(string)}
//	}
//
//	if itemValue := vCast["host"].(string); len(itemValue) > 0 {
//		r["host"] = itemValue
//	}
//
//	if itemValue := vCast["headers"].(map[string]interface{}); len(itemValue) > 0 {
//		r["headers"] = itemValue
//	}
//
//	return r
//}
//
//func __createPluginsIpRestriction(v interface{}) map[string]interface{} {
//	vCast := v.([]interface{})[0].(map[string]interface{})
//	r := map[string]interface{}{
//		"disable": vCast["disable"].(bool),
//	}
//
//	if itemValue := vCast["whitelist"].([]interface{}); len(itemValue) > 0 {
//		r["whitelist"] = itemValue
//	}
//
//	if itemValue := vCast["blacklist"].([]interface{}); len(itemValue) > 0 {
//		r["blacklist"] = itemValue
//	}
//
//	if itemValue := vCast["message"].(string); len(itemValue) > 0 {
//		r["message"] = itemValue
//	}
//	return r
//}
//
//func __createRouteObjectUpstream(value interface{}) Upstream {
//	values := value.([]interface{})
//	upstream := Upstream{}
//
//	if len(values) == 1 {
//		item := values[0].(map[string]interface{})
//
//		upstream.Type = item["type"].(string)
//
//		if v := item["service_name"].(string); len(v) > 0 {
//			upstream.ServiceName = v
//		}
//
//		if v := item["discovery_type"].(string); len(v) > 0 {
//			upstream.DiscoveryType = v
//		}
//
//		if v := item["name"].(string); len(v) > 0 {
//			upstream.Name = v
//		}
//
//		if v := item["desc"].(string); len(v) > 0 {
//			upstream.Description = v
//		}
//
//		if v := item["scheme"].(string); len(v) > 0 {
//			upstream.Scheme = v
//		}
//
//		if v := item["pass_host"].(string); len(v) > 0 {
//			upstream.PassHost = v
//		}
//
//		if v := item["retries"].(int); v > 0 {
//			upstream.Retries = v
//		}
//
//		if v := item["retry_timeout"].(int); v > 0 {
//			upstream.RetryTimeout = v
//		}
//
//		if v := item["timeout"]; v != nil {
//			upstream.Timeout = __createTimeOutObject(v)
//		}
//
//		if v := item["hash_on"].(string); len(v) > 0 {
//			upstream.HashOn = v
//		}
//	}
//
//	return upstream
//}
//
////- .desc: planned value cty.StringVal("Managed by Terraform") for a non-computed attribute
////      - .filter_func: planned value cty.StringVal("") for a non-computed attribute
////      - .priority: planned value cty.NumberIntVal(0) for a non-computed attribute
////      - .remote_addr: planned value cty.StringVal("") for a non-computed attribute
////      - .script: planned value cty.StringVal("") for a non-computed attribute
////      - .service_id: planned value cty.StringVal("") for a non-computed attribute
////      - .upstream_id: planned value cty.StringVal("") for a non-computed attribute
////      - .enable_websocket: planned value cty.False for a non-computed attribute
////      - .plugin_config_id: planned value cty.StringVal("") for a non-computed attribute
////      - .plugins[0].custom_jsons: planned value cty.ListValEmpty(cty.String) for a non-computed attribute
////      - .plugins[0].ip_restriction[0].blacklist: planned value cty.ListValEmpty(cty.String) for a non-computed attribute
////      - .plugins[0].ip_restriction[0].disable: planned value cty.False for a non-computed attribute
////      - .plugins[0].ip_restriction[0].message: planned value cty.StringVal("") for a non-computed attribute
////      - .plugins[0].proxy_rewrite[0].headers: planned value cty.MapValEmpty(cty.String) for a non-computed attribute
////      - .plugins[0].proxy_rewrite[0].host: planned value cty.StringVal("") for a non-computed attribute
////      - .plugins[0].proxy_rewrite[0].method: planned value cty.StringVal("") for a non-computed attribute
////      - .plugins[0].proxy_rewrite[0].scheme: planned value cty.StringVal("") for a non-computed attribute
////      - .plugins[0].proxy_rewrite[0].uri: planned value cty.StringVal("") for a non-computed attribute
////      - .upstream[0].pass_host: planned value cty.StringVal("pass") for a non-computed attribute
////      - .upstream[0].type: planned value cty.StringVal("roundrobin") for a non-computed attribute
////      - .upstream[0].upstream_host: planned value cty.StringVal("") for a non-computed attribute
////      - .upstream[0].name: planned value cty.StringVal("") for a non-computed attribute
////      - .upstream[0].nodes: planned value cty.StringVal("") for a non-computed attribute
////      - .upstream[0].hash_on: planned value cty.StringVal("vars") for a non-computed attribute
////      - .upstream[0].labels: planned value cty.MapValEmpty(cty.String) for a non-computed attribute
////      - .upstream[0].retry_timeout: planned value cty.NumberIntVal(0) for a non-computed attribute
////      - .upstream[0].scheme: planned value cty.StringVal("http") for a non-computed attribute
////      - .upstream[0].desc: planned value cty.StringVal("") for a non-computed attribute
