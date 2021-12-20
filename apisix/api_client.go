package apisix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//func New(endpoint string, apiKey string) ApiClient {

//}

func getCl() ApiClient {
	apiClient := http.DefaultClient
	headers := make(http.Header, 0)
	headers.Add("X-API-KEY", "edd1c9f034335f136f87ad84b625c8f1")
	apiClient.Transport = AddHeadersRoundtripper{
		Headers: headers,
		Nested:  http.DefaultTransport,
	}
	return ApiClient{
		Endpoint: "http://172.16.104.3/apisix/admin",
		HTTP:     http.DefaultClient,
	}
}

type ApiClient struct {
	Endpoint string
	HTTP     *http.Client
}

func parseHttpResult(res *http.Response, body []byte) (int, []byte, error) {
	log.Printf("[DEBUG] Got response: %#v", res)
	log.Printf("[DEBUG] Got body: %#v", string(body))

	var result map[string]interface{}
	err := json.Unmarshal(body, &result)

	if err != nil {
		return 0, []byte{}, err
	}

	if res.StatusCode >= 400 {
		return res.StatusCode, []byte(result["error_msg"].(string)), fmt.Errorf("can't make request, cause: %v", result["error_msg"].(string))
	}

	node := result["node"].(map[string]interface{})
	value := node["value"].(map[string]interface{})
	item, err := json.Marshal(value)
	return res.StatusCode, item, err
}

func (client ApiClient) Get(path string) (int, []byte, error) {
	url := client.Endpoint + path
	res, err := client.HTTP.Get(url)

	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	return parseHttpResult(res, body)
}

func (client ApiClient) Post(path string, obj interface{}) (int, []byte, error) {
	apiUrl := client.Endpoint + path
	jsonString, err := json.Marshal(obj)

	if err != nil {
		return 0, []byte{}, err
	}

	res, err := client.HTTP.Post(apiUrl, "application/json; charset=utf-8", bytes.NewReader(jsonString))

	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	return parseHttpResult(res, body)
}

func (client ApiClient) Put(path string, obj interface{}) (int, []byte, error) {
	apiUrl := client.Endpoint + path
	jsonString, err := json.Marshal(obj)
	log.Printf("[DEBUG] SEND -> %v", string(jsonString))
	if err != nil {
		return 0, []byte{}, err
	}

	req, err := http.NewRequest("PUT", apiUrl, bytes.NewReader(jsonString))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.HTTP.Do(req)

	if err != nil {
		return 0, nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	return parseHttpResult(res, body)
}

func (client ApiClient) Delete(path string) error {
	apiUrl := client.Endpoint + path

	req, err := http.NewRequest("DELETE", apiUrl, nil)
	if err != nil {
		return err
	}
	res, err := client.HTTP.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	return err
}
