package apisix

func (client ApiClient) GetRoute(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/routes/"+id, nil)
}

func (client ApiClient) CreateRoute(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/routes/", &data)
}

func (client ApiClient) UpdateRoute(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/routes/"+id, &data)
}

func (client ApiClient) DeleteRoute(id string) (err error) {
	return client.Delete("/routes/" + id)
}
