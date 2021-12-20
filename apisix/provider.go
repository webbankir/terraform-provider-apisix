package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/webbankir/terraform-provider-apisix/apisix/api_client"
	"os"
)

//// Provider returns a terraform.ResourceProvider.
//func Provider() *schema.Provider {
//	return &schema.Provider{
//		Schema: map[string]*schema.Schema{
//			"endpoint": &schema.Schema{
//				Type:        schema.TypeString,
//				Required:    true,
//				DefaultFunc: schema.EnvDefaultFunc("APISIX_ENDPOINT", nil),
//			},
//			"api_key": &schema.Schema{
//				Type:        schema.TypeString,
//				Required:    true,
//				DefaultFunc: schema.EnvDefaultFunc("APISIX_API_KEY", nil),
//			},
//		},
//		DataSourcesMap: map[string]*schema.Resource{
//			//"apisix_route":                dataSourceRoute(),
//			"apisix_plugin_proxy_rewrite": dataSourcePluginProxyRewrite(),
//		},
//		ResourcesMap: map[string]*schema.Resource{
//			"apisix_route": resourceRoute(),
//			"apisix_ssl":   resourceSsl(),
//		},
//		ConfigureFunc: func(r *schema.ResourceData) (interface{}, error) {
//			config := New(r.Get("endpoint").(string), r.Get("api_key").(string))
//			return config, nil
//		},
//	}
//}

var stderr = os.Stderr

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *api_client.Client
}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"endpoint": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"api_key": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

type ProviderData struct {
	Endpoint types.String `tfsdk:"endpoint"`
	ApiKey   types.String `tfsdk:"api_key"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	var config ProviderData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Endpoint.Unknown || config.Endpoint.Null || config.Endpoint.Value == "" {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"apisix_ssl":   ResourceSslType{},
		"apisix_route": ResourceRouteType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
