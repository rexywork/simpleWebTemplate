package postgres

import "simpleWebTemplate/pkg/model/object"

type MockUserDAO struct {

}

func NewMockUserDAO() *MockUserDAO{
	return &MockUserDAO{
	}
}

func (userDAO *MockUserDAO) CreateUser(userObject *object.User) error {
	return nil
}

func (userDAO *MockUserDAO) GetUser(id string) (*object.User, error) {
	return &object.User{
		ID:   "Test",
		Name: "Test",
	}, nil
}

func (userDAO *MockUserDAO) UpdateUser(id string, updatedUserObject *object.User) (*object.User, error) {
	return nil, nil
}

func (userDAO *MockUserDAO) DeleteUser(id string) error {
	return nil
}