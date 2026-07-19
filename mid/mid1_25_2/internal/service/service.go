package service

import (
	"context"
	"sync"
)

type Repository interface {
	Get(ctx context.Context, id int) error
}

type UserService struct {
	userRepr   Repository
	orderRep   Repository
	profileRep Repository
}

func NewUserService(
	userRep Repository,
	orderRep Repository,
	profileRep Repository,
) *UserService {
	return &UserService{
		userRepr:   userRep,
		orderRep:   orderRep,
		profileRep: profileRep,
	}
}

func (s *UserService) GetUser(parentCtx context.Context, id int) error {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()
	wg := sync.WaitGroup{}
	once := sync.Once{}
	var firstError error

	wg.Go(func() {
		err := s.userRepr.Get(ctx, id)

		if err != nil {
			once.Do(func() {
				firstError = err
				cancel()
			})
		}
	})

	wg.Go(func() {
		err := s.orderRep.Get(ctx, id)

		if err != nil {
			once.Do(func() {
				firstError = err
				cancel()
			})
		}
	})

	wg.Go(func() {
		err := s.profileRep.Get(ctx, id)

		if err != nil {
			once.Do(func() {
				firstError = err
				cancel()
			})
		}
	})

	wg.Wait()

	return firstError
}
