package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleWebTemplate/pkg/model"
)

// @Summary Get User
// @Description get user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.UserResponse
// @Router /users [get]
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserResponse{
		ID: "Test",
		Name: "Test",
	})
}
