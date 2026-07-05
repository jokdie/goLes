package service

import (
	"context"
	"jun4_5/internal/model"
)

type repository interface {
	TestRepDo(ctx context.Context) model.SuccessResponse
}

type TestService struct {
	repository repository
}

func NewTestService(r repository) *TestService {
	return &TestService{repository: r}
}

func (s *TestService) TestSrcDo(ctx context.Context) model.SuccessResponse {
	return s.repository.TestRepDo(ctx)
}
