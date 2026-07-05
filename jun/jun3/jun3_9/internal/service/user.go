package service

import (
	"jun3_9/internal/model"
)

type UserRepository interface {
	ByID(id int) (model.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (s *UserService) ByID(id int) (model.User, error) {
	user, err := s.repository.ByID(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
