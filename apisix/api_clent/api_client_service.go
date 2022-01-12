package api_clent

func (client ApiClient) GetService(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/services/"+id, nil)
}

func (client ApiClient) CreateService(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/services/", &data)
}

func (client ApiClient) UpdateService(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/services/"+id+"/__full__", &data)
}

func (client ApiClient) DeleteService(id string) (err error) {
	return client.Delete("/services/" + id)
}
