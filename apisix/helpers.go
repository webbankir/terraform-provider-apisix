package apisix

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"reflect"
	"strings"
)

func jsonToStruct(v interface{}) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	jsonPure, err := json.Marshal(v)

	if err != nil {
		return jsonMap, err
	}
	err = json.Unmarshal(jsonPure, &jsonMap)
	if err != nil {
		return jsonMap, err
	}

	return jsonMap, nil
}

func mapKeys(m interface{}) []string {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic(errors.New("not a map"))
	}

	keys := v.MapKeys()
	s := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		s[i] = keys[i].String()
	}

	return s
}

func intToString(m map[string]int, value int) string {
	for k, v := range m {
		if int(v) == value {
			return k
		}
	}
	return ""
}

type AddHeadersRoundtripper struct {
	Headers http.Header
	Nested  http.RoundTripper
}

func (h AddHeadersRoundtripper) RoundTrip(r *http.Request) (*http.Response, error) {
	for k, vs := range h.Headers {
		for _, v := range vs {
			r.Header.Add(k, v)
		}
	}
	return h.Nested.RoundTrip(r)
}

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

func stringContainsInSlice(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
