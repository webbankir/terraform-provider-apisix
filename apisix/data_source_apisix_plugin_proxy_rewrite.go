package apisix

//func dataSourcePluginProxyRewrite() *schema.Resource {
//	return &schema.Resource{
//		Read: dataSourcePluginProxyRewriteRead,
//		Schema: map[string]*schema.Schema{
//			"disable": {
//				Optional: true,
//				Type:     schema.TypeBool,
//				Default:  false,
//			},
//			"scheme": {
//				Optional:     true,
//				Type:         schema.TypeString,
//				ValidateFunc: validation.StringInSlice([]string{"http", "https"}, false),
//			},
//			"method": {
//				Optional:     true,
//				Type:         schema.TypeString,
//				ValidateFunc: validation.StringInSlice(HttpMethods, false),
//			},
//			"uri": {
//				Optional: true,
//				Type:     schema.TypeString,
//			},
//			"regex_uri": {
//				Optional: true,
//				Type:     schema.TypeList,
//				MaxItems: 1,
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"regex": {
//							Required: true,
//							Type:     schema.TypeString,
//						},
//						"replacement": {
//							Required: true,
//							Type:     schema.TypeString,
//						},
//					},
//				},
//			},
//			"host": {
//				Optional: true,
//				Type:     schema.TypeString,
//			},
//			"headers": {
//				Type:     schema.TypeMap,
//				Optional: true,
//				Elem: &schema.Schema{
//					Type: schema.TypeString,
//				},
//			},
//			"result": {
//				Computed: true,
//				Type:     schema.TypeString,
//			},
//		},
//	}
//}
//
//func dataSourcePluginProxyRewriteRead(d *schema.ResourceData, m interface{}) error {
//	data := PluginProxyRewrite{
//		Disable: d.Get("disable").(bool),
//	}
//
//	if value, exists := d.GetOkExists("scheme"); exists {
//		data.Scheme = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("uri"); exists {
//		data.URI = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("method"); exists {
//		data.Method = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("regex_uri"); exists {
//		v := value.([]interface{})[0].(map[string]interface{})
//		data.RegexUri = []string{v["regex"].(string), v["replacement"].(string)}
//	}
//
//	if value, exists := d.GetOkExists("host"); exists {
//		data.Host = value.(string)
//	}
//
//	if value, exists := d.GetOkExists("headers"); exists {
//		data.Headers = value.(map[string]interface{})
//	}
//
//	f := make(map[string]interface{})
//	f["proxy-rewrite"] = data
//
//	jsonBytes, err := json.Marshal(f)
//	if err != nil {
//		return err
//	}
//
//	d.SetId(fmt.Sprintf("%x", md5.Sum(jsonBytes)))
//
//	return d.Set("result", string(jsonBytes))
//}
//
//type PluginProxyRewrite struct {
//	Disable  bool                   `json:"disable"`
//	Scheme   string                 `json:"scheme,omitempty"`
//	URI      string                 `json:"uri,omitempty"`
//	Method   string                 `json:"method,omitempty"`
//	RegexUri []string               `json:"regex_uri,omitempty"`
//	Host     string                 `json:"host,omitempty"`
//	Headers  map[string]interface{} `json:"headers,omitempty"`
//}
