package api_clent

func (client ApiClient) GetConsumer(username string) (map[string]interface{}, error) {
	return client.RunObject("GET", "/consumers/"+username, nil)
}

func (client ApiClient) CreateConsumer(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/consumers/", &data)
}

func (client ApiClient) UpdateConsumer(data map[string]interface{}) (map[string]interface{}, error) {
	return client.RunObject("PUT", "/consumers/", &data)
}

func (client ApiClient) DeleteConsumer(username string) (err error) {
	return client.Delete("/consumers/" + username)
}
