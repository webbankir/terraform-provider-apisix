//package main
//
//import (
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
//	"github.com/webbankir/terraform-provider-apisix/apisix"
//)
//
//func main() {
//	plugin.Serve(&plugin.ServeOpts{
//		ProviderFunc: func() *schema.Provider {
//			return apisix.Provider()
//		},
//	})
//}

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
