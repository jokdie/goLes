package service

import (
	"context"
	"errors"
	"testing"
)

type fakeRepository struct {
	getFunc func(ctx context.Context, id int) error
}

func (f *fakeRepository) Get(ctx context.Context, id int) error {
	return f.getFunc(ctx, id)
}

func getRepWithoutError() *fakeRepository {
	return &fakeRepository{
		getFunc: func(ctx context.Context, id int) error {
			return nil
		},
	}
}

func TestUserService_GetUser_Success(t *testing.T) {
	userRep := getRepWithoutError()
	orderRep := getRepWithoutError()
	profileRep := getRepWithoutError()

	service := NewUserService(userRep, orderRep, profileRep)

	err := service.GetUser(t.Context(), 123)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserService_GetUser_UserRepositoryError(t *testing.T) {
	UserError := errors.New("db error")

	userRep := &fakeRepository{
		getFunc: func(ctx context.Context, id int) error {
			return UserError
		},
	}
	orderRep := getRepWithoutError()
	profileRep := getRepWithoutError()

	service := NewUserService(userRep, orderRep, profileRep)

	err := service.GetUser(t.Context(), 123)
	if !errors.Is(err, UserError) {
		t.Fatalf("expected %v, got %v", UserError, err)
	}
}

func TestUserService_GetUser_OrderRepositoryError(t *testing.T) {
	OrderError := errors.New("db error")

	userRep := getRepWithoutError()
	orderRep := &fakeRepository{
		getFunc: func(ctx context.Context, id int) error {
			return OrderError
		},
	}
	profileRep := getRepWithoutError()

	service := NewUserService(userRep, orderRep, profileRep)

	err := service.GetUser(t.Context(), 123)
	if !errors.Is(err, OrderError) {
		t.Fatalf("expected %v, got %v", OrderError, err)
	}
}

func TestUserService_GetUser_ProfileRepositoryError(t *testing.T) {
	ProfileError := errors.New("db error")

	userRep := getRepWithoutError()
	orderRep := getRepWithoutError()
	profileRep := &fakeRepository{
		getFunc: func(ctx context.Context, id int) error {
			return ProfileError
		},
	}

	service := NewUserService(userRep, orderRep, profileRep)

	err := service.GetUser(t.Context(), 123)
	if !errors.Is(err, ProfileError) {
		t.Fatalf("expected %v, got %v", ProfileError, err)
	}
}
