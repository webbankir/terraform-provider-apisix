package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func StringContainsInSlice(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func NumberContainsInSlice(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsAttributeIsNull(attribute attr.Value) bool {
	switch dd := attribute.(type) {
	case types.Object:
		return dd.Null
	case types.Map:
		return dd.Null
	case types.String:
		return dd.Null
	case types.Bool:
		return dd.Null
	case types.Number:
		return dd.Null
	case types.Set:
		return dd.Null
	case types.List:
		return dd.Null
	default:
		return true
	}
}
