package repository

import (
	"sync"
)

type User struct {
	ID   int
	Name string
}

type UserStore struct {
	users map[int]User
	mu    sync.RWMutex
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[int]User),
	}
}

func (s *UserStore) Add(user User) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[user.ID] = user
}

func (s *UserStore) Get(id int) (User, bool) {
	u, ok := s.users[id]

	return u, ok
}

func (s *UserStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.users)
}
