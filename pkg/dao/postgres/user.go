package postgres

import (
	"gorm.io/gorm"
	"simpleWebTemplate/pkg/model/object"
)

type UserDAO struct {
	DB *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO{
	return &UserDAO{
		DB: db,
	}
}



type User struct {
	ID string `gorm:"primaryKey;type:varchar(32)"`
	Name string `gorm:"type:varchar(256)"`
}

func (user *User) convertFromUserObject(userObject *object.User) {
	user.ID = userObject.ID
	user.Name = userObject.Name
}

func (user *User) convertToUserObject() *object.User {
	return &object.User{
		ID: user.ID,
		Name: user.Name,
	}
}

func (userDAO *UserDAO) CreateUser(userObject *object.User) error {
	return nil
}

func (userDAO *UserDAO) GetUser(id string) (*object.User, error) {
	return nil, nil
}

func (userDAO *UserDAO) UpdateUser(id string, updatedUserObject *object.User) (*object.User, error) {
	return nil, nil
}

func (userDAO *UserDAO) DeleteUser(id string) error {
	return nil
}

