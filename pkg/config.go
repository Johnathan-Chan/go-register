package pkg

type EtcdConfig struct {
	Address    string `json:"address"`
	Ttl        int64  `json:"ttl"`
	ServerUrl  string `json:"server_url"`
	Port       int    `json:"port"`
	ServerName string `json:"server_name"`
}
