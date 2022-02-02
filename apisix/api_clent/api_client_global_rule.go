package api_clent

func (client ApiClient) GetGlobalRule(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/global_rules/"+id, nil)
}

func (client ApiClient) CreateGlobalRule(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/global_rules/"+id, &data)
}

func (client ApiClient) UpdateGlobalRule(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/global_rules/"+id+"/__patch_terraform_plugin_apisix__", &data)
}

func (client ApiClient) DeleteGlobalRule(id string) (err error) {
	return client.Delete("/global_rules/" + id)
}
