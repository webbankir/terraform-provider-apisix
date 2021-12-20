package apisix

import (
	"encoding/json"
	"fmt"
	"log"
)

func (client ApiClient) GetRoute(id string) (output RouteObject, err error) {
	item := RouteObject{}
	statusCode, body, err := client.Get("/routes/" + id)

	if err != nil {
		return item, err
	}

	if statusCode >= 400 {
		return item, nil
	}

	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err)
		return item, err
	}

	return item, nil
}

func (client ApiClient) CreateRoute(data RouteObject) (output RouteObject, err error) {
	item := RouteObject{}
	statusCode, body, err := client.Post("/routes", data)

	if err != nil {
		return item, err
	}

	if statusCode >= 400 {
		return item, fmt.Errorf("got error: %v", string(body))
	}

	log.Printf("[DEBUG] Result of creating SSL: %#v", body)

	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err)
		return item, err
	}

	return item, nil
}

func (client ApiClient) UpdateRoute(id string, data RouteObject) (output RouteObject, err error) {
	item := RouteObject{}
	statusCode, body, err := client.Put("/routes/"+id, data)

	if err != nil {
		return item, err
	}

	if statusCode >= 400 {
		return item, nil
	}

	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err)
		return item, err
	}

	return item, nil
}

func (client ApiClient) DeleteRoute(id string) (err error) {
	return client.Delete("/routes/" + id)
}
