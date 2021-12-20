package apisix

//func dataSourceRoute() *schema.Resource {
//	return &schema.Resource{
//		Read: dataSourceAccountRead,
//		Schema: map[string]*schema.Schema{
//			"id": {
//				Required: true,
//				Type:     schema.TypeString,
//			},
//			"uri": {
//				Computed:      true,
//				Type:          schema.TypeString,
//				ConflictsWith: []string{"uris"},
//			},
//			"uris": {
//				Computed:      true,
//				Type:          schema.TypeList,
//				ConflictsWith: []string{"uri"},
//				Elem: &schema.Schema{
//					Type: schema.TypeString,
//				},
//			},
//			"name": {
//				Computed: true,
//				Type:     schema.TypeString,
//			},
//			"description": {
//				Computed: true,
//				Type:     schema.TypeString,
//			},
//
//			"methods": {
//				Computed: true,
//				Type:     schema.TypeList,
//				Elem: &schema.Schema{
//					Type: schema.TypeString,
//				},
//			},
//			"is_online": {
//				Computed: true,
//				Type:     schema.TypeBool,
//			},
//
//			"timeout": {
//				Computed: true,
//				Type:     schema.TypeMap,
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"connect": {
//							Computed: true,
//							Type:     schema.TypeInt,
//						},
//						"send": {
//							Computed: true,
//							Type:     schema.TypeInt,
//						},
//						"read": {
//							Computed: true,
//							Type:     schema.TypeInt,
//						},
//					},
//				},
//			},
//			"upstream": {
//				Computed: true,
//				Type:     schema.TypeMap,
//
//				//type Upstream struct {
//				//	KeepalivePool KeepalivePool          `json:"keepalive_pool,omitempty"`
//				//	ServiceName   string                 `json:"service_name,omitempty"`
//				//	Retries       int                    `json:"retries,omitempty"`
//				//	RetryTimeout  int                    `json:"retry_timeout,omitempty"`
//				//	Timeout       Timeout                `json:"timeout,omitempty"`
//				//	Name          string                 `json:"name,omitempty"`
//				//	Description   string                 `json:"desc,omitempty"`
//				//	PassHost      bool                   `json:"pass_host,omitempty"`
//				//	Scheme        string                 `json:"scheme,omitempty"`
//				//	Labels        map[string]interface{} `json:"labels,omitempty"`
//				//	CreateTime    int                    `json:"create_time,omitempty"`
//				//	UpdateTime    int                    `json:"update_time,omitempty"`
//				//	HashOn        string                 `json:"hash_on,omitempty"`
//				//	Key           string                 `json:"key,omitempty"`
//				//	UpstreamHost  string                 `json:"upstream_host,omitempty"`
//				//}
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"type": {
//							Computed: true,
//							Type:     schema.TypeString,
//						},
//						"discovery_type": {
//							Computed: true,
//							Type:     schema.TypeString,
//						},
//						//"send": {
//						//	Computed: true,
//						//	Type:     schema.TypeInt,
//						//},
//						//"read": {
//						//	Computed: true,
//						//	Type:     schema.TypeInt,
//						//},
//					},
//				},
//			},
//		},
//	}
//}
//
//func dataSourceAccountRead(d *schema.ResourceData, m interface{}) error {
//
//	routeId := d.Get("id").(string)
//
//	route, err := m.(ApiClient).GetRouteDetails(routeId)
//	if err != nil {
//		return err
//	}
//
//	d.SetId(routeId)
//
//	if err := d.Set("uri", route.Uri); err != nil {
//		return err
//	}
//
//	if err := d.Set("name", route.Name); err != nil {
//		return err
//	}
//
//	if err := d.Set("description", route.Description); err != nil {
//		return err
//	}
//
//	if err := d.Set("uris", route.Uris); err != nil {
//		return err
//	}
//
//	if err := d.Set("methods", route.Methods); err != nil {
//		return err
//	}
//
//	if err := d.Set("timeout", route.Timeout); err != nil {
//		return err
//	}
//
//	if err := d.Set("is_online", route.Status != 0); err != nil {
//		return err
//	}
//
//	jsonMap, err := jsonToStruct(route.Upstream)
//	if err != nil {
//		return err
//	}
//
//	return d.Set("upstream", &jsonMap)
//}
