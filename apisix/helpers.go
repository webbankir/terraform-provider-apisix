package apisix

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func checkOneOf(d *schema.ResourceData, keys ...string) error {
	var gotKey bool
	for _, key := range keys {
		_, ok := d.GetOk(key)

		if ok {
			if gotKey {
				return fmt.Errorf("only one of %s can be provided", getJoinedKeys(keys))
			}

			gotKey = true
		}
	}

	if !gotKey {
		return fmt.Errorf("one of %s should be provided", getJoinedKeys(keys))
	}

	return nil
}

func checkOneOfOptional(d *schema.ResourceData, keys ...string) error {
	var gotKey bool
	for _, key := range keys {
		_, ok := d.GetOk(key)

		if ok {
			if gotKey {
				return fmt.Errorf("only one of %s can be provided", getJoinedKeys(keys))
			}

			gotKey = true
		}
	}
	return nil
}

func getJoinedKeys(keys []string) string {
	return "`" + strings.Join(keys, "`, `") + "`"
}
