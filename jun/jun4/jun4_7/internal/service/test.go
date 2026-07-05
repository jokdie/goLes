package service

import (
	"context"
)

type repository interface {
	Dododo(ctx context.Context) error
}

type TestService struct {
	repository repository
}

func NewTestService(r repository) *TestService {
	return &TestService{repository: r}
}

func (s *TestService) Do(ctx context.Context) error {
	return s.repository.Dododo(ctx)
}
