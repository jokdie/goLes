package repository

import (
	"context"
	"jun4_6/internal/model"
	"time"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) UserByID(ctx context.Context, id int) (model.User, error) {
	select {
	case <-ctx.Done():
		return model.User{}, ctx.Err()
	case <-time.After(5 * time.Second):
		return model.User{ID: id}, nil
	}
}
