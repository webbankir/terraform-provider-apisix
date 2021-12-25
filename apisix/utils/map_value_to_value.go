package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"log"
	"math/big"
)

func MapValueToValue(sMap map[string]interface{}, mapKey string, value interface{}) {
	v := sMap[mapKey]

	switch d := value.(type) {
	case *types.String:
		if v == nil {
			d.Null = true
			log.Printf("[DEBUG] JOPAAAAAAAAAAA")
		} else {
			d.Value = v.(string)
		}

	case *types.Number:
		if v == nil {
			d.Null = true
		} else {
			d.Value = big.NewFloat(v.(float64))
		}
	case *types.Bool:
		if v == nil {
			d.Null = true
		} else {
			d.Value = v.(bool)
		}
	//case types.List:
	//	switch d.ElemType {
	//	case types.StringType:
	//		if !d.Null {
	//			var values []string
	//			for _, v := range d.Elems {
	//				values = append(values, v.(types.String).Value)
	//			}
	//			dMap[mapKey] = values
	//		} else if bindAsNil {
	//			dMap[mapKey] = nil
	//		}
	//	default:
	//		panic("WTF")
	//	}

	//if v := jsonMap["labels"]; v != nil {

	//	} else {
	//		newState.Labels = types.Map{Null: true}
	//	}
	case *types.Map:
		if v == nil {
			d.Null = true
		} else {
			values := make(map[string]attr.Value)
			for key, value := range v.(map[string]interface{}) {
				switch dd := value.(type) {
				case string:
					d.ElemType = types.StringType
					values[key] = types.String{Value: dd}
				default:
					panic("WTF")
				}

			}
			d.Elems = values
		}

	default:
		panic(d)
	}

	log.Printf("[DEBUG] dd: %v %v", mapKey, value)
}
