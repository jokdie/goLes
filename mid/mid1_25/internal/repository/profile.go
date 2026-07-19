package repository

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type ProfileRepository struct{}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (r *ProfileRepository) Get(ctx context.Context, id int) error {
	fmt.Println("start profile")
	select {
	case <-ctx.Done():
		fmt.Println("profile cancel")
		return ctx.Err()
	case <-time.After(3 * time.Second):
		fmt.Println("profile Error")

		return errors.New("profile error")
	}
}
