package api

import (
	"github.com/gin-gonic/gin"
	"simpleWebTemplate/pkg/handler"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("", handler.GetUser)
}