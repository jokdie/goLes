package service

import (
	"jun3_10/internal/dto"
	"jun3_10/internal/model"
)

type UserRepository interface {
	Add(u model.User) model.User
}

type UserService struct {
	repository UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repository: r}
}

func (s *UserService) Create(dto *dto.CreateUserRequest) model.User {
	u := model.User{
		Name: dto.Name,
		Age:  dto.Age,
	}

	u = s.repository.Add(u)

	// return u, err - обычно так будет, но у меня упрощено
	return u
}
