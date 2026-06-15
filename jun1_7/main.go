package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetByID(id int) (User, error)
}

type MemoryUserRepository struct {
	users map[int]User
}

func (m *MemoryUserRepository) GetByID(id int) (User, error) {
	if u, ok := m.users[id]; ok {
		return u, nil
	}

	return User{}, UserNotFoundError{ID: id}
}

type UserNotFoundError struct {
	ID int
}

func (u UserNotFoundError) Error() string {
	return fmt.Sprintf("User with ID:[%d] not found", u.ID)
}

func main() {
	repo := MemoryUserRepository{
		users: map[int]User{
			1:  {ID: 1, Name: "Alice"},
			2:  {ID: 2, Name: "Bob"},
			43: {ID: 43, Name: "Sam"},
		},
	}

	user, err := repo.GetByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("user = %s\n", user.Name)
	}

	user, err = repo.GetByID(42)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("user = %s\n", user.Name)
	}
}
