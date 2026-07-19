package service

import (
	"context"
	"mid1_25_2/internal/model"
)

type Repository interface {
	GetUser(parentCtx context.Context, id int) (model.User, error)
}

type UserService struct {
	userRepr Repository
}

func NewUserService(
	userRep Repository,
) *UserService {
	return &UserService{
		userRepr: userRep,
	}
}

func (s *UserService) GetUser(ctx context.Context, id int) (model.User, error) {
	return s.userRepr.GetUser(ctx, id)

}
