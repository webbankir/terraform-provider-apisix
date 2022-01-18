package api_clent

func (client ApiClient) GetGlobalRule(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/global_rules/"+id, nil)
}

func (client ApiClient) CreateGlobalRule(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/global_rules/", &data)
}

func (client ApiClient) UpdateGlobalRule(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/global_rules/"+id+"/__full__", &data)
}

func (client ApiClient) DeleteGlobalRule(id string) (err error) {
	return client.Delete("/global_rules/" + id)
}