package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description get user
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /v1/users [get]
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, User{
		Name: "Test",
	})
}
