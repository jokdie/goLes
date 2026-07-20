package service

import (
	"context"
	"errors"
	"testing"
)

type fakeRepository struct {
	ctx     context.Context
	id      int
	called  bool
	getFunc func(ctx context.Context, id int) error
}

func (f *fakeRepository) Get(ctx context.Context, id int) error {
	f.ctx = ctx
	f.called = true
	f.id = id

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

	id := 123

	err := service.GetUser(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if !userRep.called {
		t.Fatal("userRep does not used")
	}
	if !orderRep.called {
		t.Fatal("orderRep does not used")
	}
	if !profileRep.called {
		t.Fatal("profileRep does not used")
	}
	if userRep.id != id {
		t.Fatalf("userRep != %d", id)
	}
	if orderRep.id != id {
		t.Fatalf("orderRep != %d", id)
	}
	if profileRep.id != id {
		t.Fatalf("profileRep != %d", id)
	}
	if userRep.ctx == nil {
		t.Fatal("userRep - context was not passed")
	}
	if orderRep.ctx == nil {
		t.Fatal("orderRep - context was not passed")
	}
	if profileRep.ctx == nil {
		t.Fatal("profileRep - context was not passed")
	}
}

func TestUserService_GetUser_UserRepositoryError(t *testing.T) {
	userError := errors.New("db error")

	userRep := &fakeRepository{
		getFunc: func(ctx context.Context, id int) error {
			return userError
		},
	}
	orderRep := getRepWithoutError()
	profileRep := getRepWithoutError()

	service := NewUserService(userRep, orderRep, profileRep)

	id := 123
	err := service.GetUser(t.Context(), id)
	if !errors.Is(err, userError) {
		t.Fatalf("expected %v, got %v", userError, err)
	}
	if !userRep.called {
		t.Fatal("user repository was not called")
	}
	if userRep.id != id {
		t.Fatalf("expected id %d, got %d", id, userRep.id)
	}
	if userRep.ctx == nil {
		t.Fatal("userRep - context was not passed")
	}
}

func TestUserService_GetUser_OrderRepositoryError(t *testing.T) {
	orderError := errors.New("db error")

	userRep := getRepWithoutError()
	orderRep := &fakeRepository{
		getFunc: func(ctx context.Context, id int) error {
			return orderError
		},
	}
	profileRep := getRepWithoutError()

	service := NewUserService(userRep, orderRep, profileRep)

	id := 123
	err := service.GetUser(t.Context(), id)
	if !errors.Is(err, orderError) {
		t.Fatalf("expected %v, got %v", orderError, err)
	}
	if !orderRep.called {
		t.Fatal("order repository was not called")
	}
	if orderRep.id != id {
		t.Fatalf("expected id %d, got %d", id, orderRep.id)
	}
	if orderRep.ctx == nil {
		t.Fatal("orderRep - context was not passed")
	}
}

func TestUserService_GetUser_ProfileRepositoryError(t *testing.T) {
	profileError := errors.New("db error")

	userRep := getRepWithoutError()
	orderRep := getRepWithoutError()
	profileRep := &fakeRepository{
		getFunc: func(ctx context.Context, id int) error {
			return profileError
		},
	}

	service := NewUserService(userRep, orderRep, profileRep)

	id := 123
	err := service.GetUser(t.Context(), id)
	if !errors.Is(err, profileError) {
		t.Fatalf("expected %v, got %v", profileError, err)
	}
	if !profileRep.called {
		t.Fatal("profile repository was not called")
	}
	if profileRep.id != id {
		t.Fatalf("expected id %d, got %d", id, profileRep.id)
	}
	if profileRep.ctx == nil {
		t.Fatal("profileRep - context was not passed")
	}
}
