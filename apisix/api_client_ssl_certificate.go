package apisix

func (client ApiClient) GetSslCertificate(id string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/ssl/"+id, nil)
}

func (client ApiClient) CreateSslCertificate(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("POST", "/ssl/", &data)
}

func (client ApiClient) UpdateSslCertificate(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/ssl/"+id, &data)
}

func (client ApiClient) DeleteSslCertificate(id string) (err error) {
	return client.Delete("/ssl/" + id)
}
