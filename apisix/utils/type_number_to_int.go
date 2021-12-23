package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
)

func TypeNumberToInt(v types.Number) int {
	if v.Null || v.Value == nil {
		return 0
	}
	nv, _ := strconv.Atoi(v.Value.String())
	return nv
}
