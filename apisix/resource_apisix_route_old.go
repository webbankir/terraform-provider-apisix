package apisix

//func resourceRoute() *schema.Resource {
//	return &schema.Resource{
//		Create: resourceRouteCreate,
//		Read:   resourceRouteRead,
//		Update: resourceRouteUpdate,
//		Delete: resourceRouteDelete,
//		Importer: &schema.ResourceImporter{
//			State: schema.ImportStatePassthrough,
//		},
//		Schema: schemeRouterResource,
//	}
//}
//
//func resourceRouteCreate(d *schema.ResourceData, m interface{}) error {
//
//	route, err := __createRouteObject(d)
//	if err != nil {
//		return err
//	}
//	result, err := m.(ApiClient).CreateRoute(route)
//
//	if err != nil {
//		return err
//	}
//	d.SetId(result.ID)
//	return __updateRouteResource(d, result)
//}
//
////vars	False	Match Rules	A list of one or more [var, operator, val] elements, like this: [[var, operator, val], [var, operator, val], ...]]. For example: ["arg_name", "==", "json"] means that the current request parameter name is json. The var here is consistent with the internal variable name of Nginx, so you can also use request_uri, host, etc. For more details, see lua-resty-expr	[["arg_name", "==", "json"], ["arg_age", ">", 18]]
////upstream	False	Upstream	Enabled Upstream configuration, see Upstream for more
//
//func resourceRouteRead(d *schema.ResourceData, m interface{}) error {
//	id := d.Id()
//	route, err := m.(ApiClient).GetRoute(id)
//	if err != nil {
//		return err
//	}
//	return __updateRouteResource(d, route)
//}
//
//func resourceRouteUpdate(d *schema.ResourceData, m interface{}) error {
//	id := d.Id()
//
//	route, err := __createRouteObject(d)
//	if err != nil {
//		return err
//	}
//	_, err = m.(ApiClient).UpdateRoute(id, route)
//	if err != nil {
//		return err
//	}
//	log.Printf("[DEBUG] WMWWMWM got error: %v", err)
//	_ = d.Set("last_updated", time.Now().Format(time.RFC850))
//	return resourceRouteRead(d, m)
//}
//
//func resourceRouteDelete(d *schema.ResourceData, m interface{}) error {
//	return m.(ApiClient).DeleteRoute(d.Id())
//}
