package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/webbankir/terraform-provider-apisix/apisix"
)

func main() {
	_ = tfsdk.Serve(context.Background(), apisix.New, tfsdk.ServeOpts{
		Name: "apisix",
	})
}
