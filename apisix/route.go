package apisix

////filter_func	False	Match Rules	User-defined filtering function. You can use it to achieve matching requirements for special scenarios. This function accepts an input parameter named vars by default, which you can use to get Nginx variables.	function(vars) return vars["arg_name"] == "json" end
////plugins	False	Plugin	See Plugin for more
////script	False	Script	See Script for more
////upstream	False	Upstream	Enabled Upstream configuration, see Upstream for more
////upstream_id	False	Upstream	Enabled upstream id, see Upstream for more
////service_id	False	Service	Bound Service configuration, see Service for more
////plugin_config_id	False, can't be used with script	Plugin	Bound plugin config object, see Plugin Config for more
//
//func (client ApiClient) GetRouteDetails(routeId string) (acc Route, err error) {
//	//body, err := client.Get("/routes/" + routeId)
//	//
//	//if err != nil {
//	//	return
//	//}
//	//
//	//if err != nil {
//	//	// do error check
//	//	fmt.Println(err)
//	//}
//	//
//	//route := Route{}
//	//if err := json.Unmarshal(body, &route); err != nil {
//	//	// do error check
//	//	fmt.Println(err)
//	//}
//
//	return Route{}, nil
//}
