package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleWebTemplate/pkg/model"
)

// @Summary Get Users
// @Description List users existing.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.UserResponses
// @Router /users [get]
func (api *API) GetUsers(c *gin.Context) {

	user, err := api.UserService.GetUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			ErrorCode: "1",
			ErrorMessages: fmt.Sprintf("error: %v", err.Error()),
		})
	}
	var userResponse model.UserResponse
	userResponse.ConvertUserToUserResponse(user)

	var userResponses model.UserResponses
	userResponses = append(userResponses, userResponse)

	c.JSON(http.StatusOK, userResponses)
}
