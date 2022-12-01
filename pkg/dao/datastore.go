package dao

import "simpleWebTemplate/pkg/model/object"

type UserRepository interface {
	CreateUser(user *object.User) error
	GetUser(id string) (*object.User, error)
	UpdateUser(id string, updatedUserObject *object.User) (*object.User, error)
	DeleteUser(id string) error
}
