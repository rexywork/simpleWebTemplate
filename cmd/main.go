package main

import (
	config "simpleWebTemplate/config"
	"simpleWebTemplate/pkg/server"
)


func main() {
	// initialize server config
	config := &config.Config{}
	config.Default()

	// load config from environment variables
	err := config.LoadFromEnvironmentVariables()
	if err != nil {
		panic("error when loading environment variables for config")
	}

	srv := server.NewServer(config)
	srv.Serve()
}
