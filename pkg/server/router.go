package server

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	router := gin.Default()
	return router
}

func (server *Server) RegisterRouter() {
	server.ApiRoute.AddV1UserRouter(server.Handler)
	server.ApiRoute.AddV2UserRouter(server.Handler)
}