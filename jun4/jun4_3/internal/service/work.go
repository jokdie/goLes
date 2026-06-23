package service

import (
	"context"
	"fmt"
	"jun4_3/internal/model"
	"time"
)

type WorkService struct{}

func NewWorkService() *WorkService {
	return &WorkService{}
}

func (s *WorkService) Work(ctx context.Context) (model.SuccessResponse, error) {
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return model.SuccessResponse{}, ctx.Err()
		case <-time.After(1 * time.Second):
			fmt.Printf("Тик %d\n", i)
		}
	}

	return model.SuccessResponse{Status: "completed"}, nil
}
