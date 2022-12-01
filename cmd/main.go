package main

import (
	"simpleWebTemplate/internal/server"
)


func main() {
	srv := server.NewServer()
	srv.Serve()
}
