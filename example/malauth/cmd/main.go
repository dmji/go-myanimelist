package main

import "github.com/koding/tunnel"

func main() {
	cfg := &tunnel.ClientConfig{
		Identifier: "1234",
		ServerAddr: "proxy.dmji.serv00.net",
	}

	client, err := tunnel.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	client.Start()
}
