package repository

import (
	"context"
	"fmt"
	"time"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Get(ctx context.Context, id int) error {
	fmt.Println("start user")
	select {
	case <-ctx.Done():
		fmt.Println("user cancel")
		return ctx.Err()
	case <-time.After(2 * time.Second):
		fmt.Println("user OK")
		return nil
	}
}
