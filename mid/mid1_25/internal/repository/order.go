package repository

import (
	"context"
	"fmt"
	"time"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Get(ctx context.Context, id int) error {
	fmt.Println("start order")
	select {
	case <-ctx.Done():
		fmt.Println("order cancel")
		return ctx.Err()
	case <-time.After(8 * time.Second):
		fmt.Println("order OK")
		return nil
	}
}
