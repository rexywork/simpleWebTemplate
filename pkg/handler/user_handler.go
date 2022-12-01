package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
}

// @Summary Get User
// @Description get user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /users [get]
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, User{
		Name: "Test",
	})
}
