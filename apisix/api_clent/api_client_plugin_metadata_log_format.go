package api_clent

func (client ApiClient) GetPluginMetadataLogFormat(pluginName string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/plugin_metadata/"+pluginName, nil)
}

func (client ApiClient) CreatePluginMetadataLogFormat(pluginName string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/plugin_metadata/"+pluginName, &data)
}

func (client ApiClient) UpdatePluginMetadataLogFormat(pluginName string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/plugin_metadata/"+pluginName, &data)
}

func (client ApiClient) DeletePluginMetadataLogFormat(pluginName string) (err error) {
	return client.Delete("/plugin_metadata/" + pluginName)
}
