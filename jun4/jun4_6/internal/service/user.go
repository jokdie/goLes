package service

import (
	"context"
	"jun4_6/internal/model"
)

type repository interface {
	UserByID(ctx context.Context, ID int) (model.User, error)
}

type UserService struct {
	repository repository
}

func NewUserService(r repository) *UserService {
	return &UserService{repository: r}
}

func (s *UserService) GetUser(ctx context.Context, id int) (model.User, error) {
	return s.repository.UserByID(ctx, id)
}
