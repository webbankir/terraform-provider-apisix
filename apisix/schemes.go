package apisix

//var timeOutSchema = map[string]*schema.Schema{
//
//	"connect": {
//		Required: true,
//		Type:     schema.TypeInt,
//	},
//	"send": {
//		Required: true,
//		Type:     schema.TypeInt,
//	},
//	"read": {
//		Required: true,
//		Type:     schema.TypeInt,
//	},
//}
//
//var schemePluginProxyRewrite = map[string]*schema.Schema{
//	"disable": {
//		Optional: true,
//		Type:     schema.TypeBool,
//		Default:  false,
//	},
//	"scheme": {
//		Optional:     true,
//		Type:         schema.TypeString,
//		ValidateFunc: validation.StringInSlice([]string{"http", "https"}, false),
//	},
//	"method": {
//		Optional:     true,
//		Type:         schema.TypeString,
//		ValidateFunc: validation.StringInSlice([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS", "MKCOL", "COPY", "MOVE", "PROPFIND", "PROPFIND", "LOCK", "UNLOCK", "PATCH", "TRACE"}, false),
//	},
//	"uri": {
//		Optional: true,
//		Type:     schema.TypeString,
//	},
//	"regex_uri": {
//		Optional: true,
//		Type:     schema.TypeList,
//		MaxItems: 1,
//		Elem: &schema.Resource{
//			Schema: map[string]*schema.Schema{
//				"regex": {
//					Required: true,
//					Type:     schema.TypeString,
//				},
//				"replacement": {
//					Required: true,
//					Type:     schema.TypeString,
//				},
//			},
//		},
//	},
//	"host": {
//		Optional: true,
//		Type:     schema.TypeString,
//	},
//	"headers": {
//		Type:     schema.TypeMap,
//		Optional: true,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//	},
//}
//
//var schemePluginIpRestriction = map[string]*schema.Schema{
//	"disable": {
//		Optional: true,
//		Type:     schema.TypeBool,
//		Default:  false,
//	},
//	"whitelist": {
//		Optional: true,
//		Type:     schema.TypeList,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//	},
//	"blacklist": {
//		Optional: true,
//		Type:     schema.TypeList,
//		Elem: &schema.Schema{
//			Type: schema.TypeString,
//		},
//	},
//	"message": {
//		Optional: true,
//		Type:     schema.TypeString,
//	},
//}
