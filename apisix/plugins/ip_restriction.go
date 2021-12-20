package plugins

import "github.com/hashicorp/terraform-plugin-framework/types"

type IpRestrictionType struct {
	Disable   types.Bool   `tfsdk:"disable"`
	Message   types.String `tfsdk:"message"`
	WhiteList types.List   `tfsdk:"whitelist"`
	BlackList types.List   `tfsdk:"blacklist"`
}

func (s IpRestrictionType) validate() error { return nil }

func (s IpRestrictionType) EncodeToMap(m map[string]interface{}) {
	pluginValue := map[string]interface{}{
		"disable": s.Disable.Value,
	}

	if !s.BlackList.Null {
		var values []string
		for _, v := range s.BlackList.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["blacklist"] = values
	}

	if !s.WhiteList.Null {
		var values []string
		for _, v := range s.WhiteList.Elems {
			values = append(values, v.(types.String).Value)
		}
		pluginValue["whitelist"] = values
	}

	if !s.Message.Null {
		pluginValue["message"] = s.Message.Value
	}

	m["ip-restriction"] = pluginValue
}
