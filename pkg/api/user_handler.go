package api

import (
	"fmt"
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
func (api *API) GetUser(c *gin.Context) {

	user, err := api.UserService.GetUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			ErrorCode: "1",
			ErrorMessages: fmt.Sprintf("error: %v", err.Error()),
		})
	}
	var userResponse model.UserResponse
	userResponse.ConvertUserToUserResponse(user)
	c.JSON(http.StatusOK, userResponse)
}
