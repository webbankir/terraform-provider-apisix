package plugins

import "github.com/hashicorp/terraform-plugin-framework/types"

type ProxyRewriteRegexUriType struct {
	Regex       types.String `tfsdk:"regex"`
	Replacement types.String `tfsdk:"replacement"`
}

type ProxyRewriteType struct {
	Disable  types.Bool                `tfsdk:"disable"`
	Scheme   types.String              `tfsdk:"scheme"`
	Method   types.String              `tfsdk:"method"`
	Uri      types.String              `tfsdk:"uri"`
	Host     types.String              `tfsdk:"host"`
	Headers  types.Map                 `tfsdk:"headers"`
	RegexUri *ProxyRewriteRegexUriType `tfsdk:"regex_uri"`
}

func (s ProxyRewriteType) validate() error { return nil }

func (s ProxyRewriteType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	if !s.Scheme.Null {
		pluginValue["scheme"] = s.Scheme.Value
	}

	if !s.Uri.Null {
		pluginValue["uri"] = s.Uri.Value
	}

	if !s.Headers.Null {
		values := make(map[string]interface{})
		for k, v := range s.Headers.Elems {
			values[k] = v.(types.String).Value
		}
		pluginValue["headers"] = values
	}

	if !s.Host.Null {
		pluginValue["host"] = s.Host.Value
	}

	if !s.Method.Null {
		pluginValue["method"] = s.Method.Value
	}
	if s.RegexUri != nil {
		pluginValue["regex_uri"] = []string{s.RegexUri.Regex.Value, s.RegexUri.Replacement.Value}
	}

	m["proxy_rewrite"] = pluginValue
}
