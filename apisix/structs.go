package apisix

type RouteObject struct {
	ID              string                 `json:"id,omitempty"`
	CreateTime      int                    `json:"create_time,omitempty"`
	Description     string                 `json:"desc,omitempty"`
	EnableWebsocket bool                   `json:"enable_websocket,omitempty"`
	Host            string                 `json:"host,omitempty"`
	Hosts           []string               `json:"hosts,omitempty"`
	Labels          map[string]interface{} `json:"labels,omitempty"`
	Methods         []string               `json:"methods,omitempty"`
	Name            string                 `json:"name"`
	Priority        int                    `json:"priority,omitempty"`
	Plugins         map[string]interface{} `json:"plugins,omitempty"`
	RemoteAddr      string                 `json:"remote_addr,omitempty"`
	RemoteAddrs     []string               `json:"remote_addrs,omitempty"`
	Status          int                    `json:"status"`
	Timeout         Timeout                `json:"timeout"`
	UpdateTime      int                    `json:"update_time,omitempty"`
	Upstream        Upstream               `json:"upstream,omitempty"`
	UpstreamId      string                 `json:"upstream_id,omitempty"`
	Uri             string                 `json:"uri,omitempty"`
	Uris            []string               `json:"uris,omitempty"`
	Vars            [][]string             `json:"vars,omitempty"`
	ServiceId       string                 `json:"service_id,omitempty"`
	Script          string                 `json:"script,omitempty"`
	PluginConfigId  string                 `json:"plugin_config_id,omitempty"`
	FilterFunc      string                 `json:"filter_func,omitempty"`
}

//nodes	required, can't be used with service_name	Hash table or array. If it is a hash table, the key of the internal element is the upstream machine address list, the format is Address + (optional) Port, where the address part can be IP or domain name, such as 192.168.1.100:80, foo.com:80, etc. The value is the weight of node. If it is an array, each item is a hash table with key host/weight and optional port/priority. The nodes can be empty, which means it is a placeholder and will be filled later. Clients use such an upstream will get 502 response.	192.168.1.100:80
//checks	optional	Configure the parameters of the health check. For details, refer to health-check.
//tls.client_cert	optional	Set the client certificate when connecting to TLS upstream, see below for more details
//tls.client_key	optional	Set the client private key when connecting to TLS upstream, see below for more details

type Upstream struct {
	Type          string                 `json:"type,omitempty"`
	DiscoveryType string                 `json:"discovery_type,omitempty"`
	KeepalivePool KeepalivePool          `json:"keepalive_pool,omitempty"`
	ServiceName   string                 `json:"service_name,omitempty"`
	Retries       int                    `json:"retries,omitempty"`
	RetryTimeout  int                    `json:"retry_timeout,omitempty"`
	Timeout       Timeout                `json:"timeout,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"desc,omitempty"`
	PassHost      string                 `json:"pass_host,omitempty"`
	Scheme        string                 `json:"scheme,omitempty"`
	Labels        map[string]interface{} `json:"labels,omitempty"`
	CreateTime    int                    `json:"create_time,omitempty"`
	UpdateTime    int                    `json:"update_time,omitempty"`
	HashOn        string                 `json:"hash_on,omitempty"`
	Key           string                 `json:"key,omitempty"`
	UpstreamHost  string                 `json:"upstream_host,omitempty"`
}

type KeepalivePool struct {
	Size        int `json:"size,omitempty"`
	IdleTimeout int `json:"idle_timeout,omitempty"`
	Requests    int `json:"requests,omitempty"`
}

type Timeout struct {
	Connect int `json:"connect,omitempty"`
	Send    int `json:"send,omitempty"`
	Read    int `json:"read,omitempty"`
}

var HttpMethods = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"PATCH",
	"HEAD",
	"OPTIONS",
	"CONNECT",
	"TRACE",
}

type SSL struct {
	ID          string   `json:"id,omitempty"`
	Certificate string   `json:"cert,omitempty"`
	PrivateKey  string   `json:"key,omitempty"`
	SNIS        []string `json:"snis,omitempty"`
}

//certs	False	An array of certificate	when you need to configure multiple certificate for the same domain, you can pass extra https certificates (excluding the one given as cert) in this field
//keys	False	An array of private key	https private keys. The keys should be paired with certs above
//client.ca	False	SslCertificate	set the CA certificate which will use to verify client. This feature requires OpenResty 1.19+.
//client.depth	False	SslCertificate	set the verification depth in the client certificates chain, default to 1. This feature requires OpenResty 1.19+.
//labels	False	Match Rules	Key/value pairs to specify attributes	{"version":"v2","build":"16","env":"production"}
//create_time	False	Auxiliary	epoch timestamp in second, will be created automatically if missing	1602883670
//update_time	False	Auxiliary	epoch timestamp in second, will be created automatically if missing	1602883670
//status	False	Auxiliary	enable this SSL, default 1.	1 to enable, 0 to disable
