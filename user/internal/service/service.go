package service

import "user/internal/models"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Create(id, name, email string) (string, error) {
	return "created", nil
}

func (u *UserService) Get(id string) (models.User, error) {
	return models.User{
		Id:    id,
		Name:  "qwer",
		Email: "asdf",
	}, nil
}
