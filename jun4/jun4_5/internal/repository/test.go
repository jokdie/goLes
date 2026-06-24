package repository

import (
	"context"
	"jun4_5/internal/model"
	"jun4_5/internal/requestid"
)

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (r *TestRepository) TestRepDo(ctx context.Context) model.SuccessResponse {
	requestID := requestid.GetRequestID(ctx)

	return model.SuccessResponse{RequestID: requestID}
}
