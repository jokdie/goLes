package repository

import (
	"context"
	"fmt"
	"mid1_25_2/internal/model"
	"time"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (model.User, error) {
	fmt.Println("start user")
	select {
	case <-ctx.Done():
		fmt.Println("user cancel")
		return model.User{}, ctx.Err()
	case <-time.After(1 * time.Second):
		fmt.Println("user OK")
		return model.User{ID: id}, nil
	}
}
