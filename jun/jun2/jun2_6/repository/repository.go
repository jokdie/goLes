package repository

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	id   int
	name string
}

func (u User) ID() int {
	return u.id
}

func (u User) Name() string {
	return u.name
}

type Repository struct {
	users map[int]User
}

func New() Repository {
	m := map[int]User{}

	return Repository{users: m}
}

func (r *Repository) AddUser(id int, name string) {
	r.users[id] = User{id: id, name: name}
}

func (r Repository) GetByID(id int) (User, error) {
	if u, ok := r.users[id]; ok {
		return u, nil
	}

	err := fmt.Errorf("User with ID=%d: %w", id, ErrUserNotFound)

	return User{}, err
}
