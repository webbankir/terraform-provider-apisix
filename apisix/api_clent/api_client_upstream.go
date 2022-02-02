package api_clent

func (client ApiClient) GetUpstream(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/upstreams/"+id, nil)
}

func (client ApiClient) CreateUpstream(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/upstreams/", &data)
}

func (client ApiClient) UpdateUpstream(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PATCH", "/upstreams/"+id+"/__patch_terraform_plugin_apisix__", &data)
}

func (client ApiClient) DeleteUpstream(id string) (err error) {
	return client.Delete("/upstreams/" + id)
}
