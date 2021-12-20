package apisix

import (
	"encoding/json"
	"fmt"
	"log"
)

func (client ApiClient) GetSsl(id string) (acc SSL, err error) {
	item := SSL{}
	statusCode, body, err := client.Get("/ssl/" + id)

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

func (client ApiClient) CreateSsl(sslItem SSL) (acc SSL, err error) {
	item := SSL{}
	statusCode, body, err := client.Post("/ssl", sslItem)

	if err != nil {
		return
	}

	if statusCode >= 400 {
		return item, nil
	}

	log.Printf("[DEBUG] Result of creating SSL: %#v", body)

	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err)
		return item, err
	}

	return item, nil
}

func (client ApiClient) UpdateSsl(id string, ssl SSL) (acc SSL, err error) {
	item := SSL{}
	statusCode, body, err := client.Put("/ssl/"+id, ssl)

	if statusCode >= 400 {
		return item, nil
	}

	if err != nil {
		return
	}

	if err := json.Unmarshal(body, &item); err != nil {
		fmt.Println(err)
		return item, err
	}

	return item, nil
}

func (client ApiClient) DeleteSsl(id string) (err error) {
	return client.Delete("/ssl/" + id)
}
