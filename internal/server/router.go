package server

import (
	"github.com/gin-gonic/gin"
	v1 "simpleWebTemplate/pkg/api/v1"
	v2 "simpleWebTemplate/pkg/api/v2"
)


func NewRouter() *gin.Engine{
	router := gin.Default()
	v1.AddV1UserRouter(router)
	v2.AddV2UserRouter(router)

	return router
}