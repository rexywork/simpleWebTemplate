package service

import (
	"simpleWebTemplate/pkg/dao"
	"simpleWebTemplate/pkg/model/object"
)

type UserService struct {
	UserRepository dao.UserRepository
}

func (u *UserService) GetUser(id string) (*object.User, error){
	user, err := u.UserRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}