package model

import "simpleWebTemplate/pkg/model/object"

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (userResponse *UserResponse) ConvertUserToUserResponse(user *object.User) {
	userResponse.ID = user.ID
	userResponse.Name = user.Name
}