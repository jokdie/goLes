package service

import (
	"context"
	"fmt"
	"jun4_2/internal/model"
	"time"
)

type WorkService struct{}

func NewWorkService() *WorkService {
	return &WorkService{}
}

func (s *WorkService) Work(ctx context.Context) (model.SuccessResponse, error) {
	iterations := 5

	for i := 0; i < iterations; i++ {
		select {
		case <-ctx.Done():
			return model.SuccessResponse{}, ctx.Err()
		case <-time.After(1 * time.Second):
			fmt.Printf("Тик %d\n", i)
		}
	}

	return model.SuccessResponse{Status: "completed"}, nil
}
