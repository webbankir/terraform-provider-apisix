package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ValueToMap(value interface{}, dMap map[string]interface{}, mapKey string, bindAsNil bool) {
	switch d := value.(type) {
	case types.String:
		if !d.Null {
			dMap[mapKey] = d.Value
		} else if bindAsNil {
			dMap[mapKey] = nil
		}
	case types.Number:
		if !d.Null {
			dMap[mapKey] = TypeNumberToInt(d)
		} else if bindAsNil {
			dMap[mapKey] = nil
		}
	case types.Bool:
		if !d.Null {
			dMap[mapKey] = d.Value
		} else if bindAsNil {
			dMap[mapKey] = nil
		}
	case types.List:
		switch d.ElemType {
		case types.StringType:
			if !d.Null {
				var values []string
				for _, v := range d.Elems {
					values = append(values, v.(types.String).Value)
				}
				dMap[mapKey] = values
			} else if bindAsNil {
				dMap[mapKey] = nil
			}
		default:
			panic("WTF")
		}
	case types.Map:
		switch d.ElemType {
		case types.StringType:
			if !d.Null {
				values := make(map[string]interface{})
				for k, v := range d.Elems {
					values[k] = v.(types.String).Value
				}
				dMap[mapKey] = values
			} else if bindAsNil {
				dMap[mapKey] = nil
			}
		default:
			panic("WTF")
		}
	default:
		panic("WTF")
	}
}
