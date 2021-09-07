package pkg

import "flag"

type ReadConfig interface {
	Read() map[string]interface{}
}

type ReadCMD struct {
}

func (this *ReadCMD) Read() map[string]interface{} {
	var (
		address    string
		ttl        int64
		serverUrl  string
		port       int
		serverName string
	)

	flag.StringVar(&address, "address", "http://127.0.0.1:2379", "http://127.0.0.1:2379")
	flag.Int64Var(&ttl, "ttl", 60, "input int type number for ttl")
	flag.StringVar(&serverUrl, "server_url", "127.0.0.1", "127.0.0.1")
	flag.IntVar(&port, "port", 8080, "input int type number for port")
	flag.StringVar(&serverName, "server_name", "server", "server")

	return map[string]interface{}{
		"address": address,
		"ttl": ttl,
		"server_url": serverUrl,
		"port": port,
		"serverName": serverName,
	}
}
