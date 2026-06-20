package repository

import "jun3_10/internal/model"

type UserRepository struct {
	users  map[int]model.User
	freeID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  map[int]model.User{},
		freeID: 1,
	}
}

func (r *UserRepository) Add(u model.User) model.User {
	u.ID = r.freeID
	r.freeID++

	r.users[u.ID] = u

	// return err - обычно так будет, но у меня упрощено
	return u
}
