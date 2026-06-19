package repository

import (
	"errors"
	"jun3_9/internal/model"
)

type UserRepository struct {
	users map[int]model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: map[int]model.User{
		1:  {ID: 1, Name: "Bob"},
		2:  {ID: 2, Name: "Bill"},
		42: {ID: 42, Name: "Nick"},
	}}
}

var UserNotFoundError = errors.New("User not found")

func (r *UserRepository) ByID(id int) (model.User, error) {
	if user, ok := r.users[id]; ok {
		return user, nil
	}

	return model.User{}, UserNotFoundError
}
