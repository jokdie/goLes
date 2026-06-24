package repository

import (
	"context"
	"log"
	"time"
)

type TestRepository struct{}

func NewRepository() *TestRepository {
	return &TestRepository{}
}

func (r *TestRepository) Dododo(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Second):
		log.Println("оч долгая работа")

		return nil
	}
}
