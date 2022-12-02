package main

import (
	config "simpleWebTemplate/config"
	"simpleWebTemplate/pkg/server"
)


func main() {
	// initialize server config
	config := &config.Config{}
	config.Default()
	srv := server.NewServer(config)
	srv.Serve()
}
