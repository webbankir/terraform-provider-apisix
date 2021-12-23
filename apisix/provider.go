package apisix

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"os"
)

var stderr = os.Stderr

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *ApiClient
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
		"apisix_ssl_certificate": ResourceSslCertificateType{},
		"apisix_route":           ResourceRouteType{},
		"apisix_upstream":        ResourceUpstreamType{},
		"apisix_stream_route":    ResourceStreamRouteType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
