package storage

import "errors"

var ErrUserNotFound = errors.New("User not found")

type User struct {
	ID   int
	Name string
}

type Storage struct {
	users map[int]User
}

func New() Storage {
	m := map[int]User{
		1: {ID: 1, Name: "Arkadiy"},
		2: {ID: 2, Name: "Boris"},
	}

	return Storage{
		users: m,
	}
}

func (s Storage) GetByID(id int) (User, error) {
	if u, ok := s.users[id]; ok {
		return u, nil
	}

	return User{}, ErrUserNotFound
}
