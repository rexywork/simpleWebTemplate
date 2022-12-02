package api

import (
	"github.com/gin-gonic/gin"
)

func (api *API) AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("", api.GetUser)
}