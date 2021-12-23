package apisix

func (client ApiClient) GetStreamRoute(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/stream_routes/"+id, nil)
}

func (client ApiClient) CreateStreamRoute(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/stream_routes/", &data)
}

func (client ApiClient) UpdateStreamRoute(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/stream_routes/"+id, &data)
}

func (client ApiClient) DeleteStreamRoute(id string) (err error) {
	return client.Delete("/stream_routes/" + id)
}
