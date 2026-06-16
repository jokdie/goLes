package main

import (
	"errors"
	"fmt"
	"jun2_6/repository"
)

func main() {
	r := repository.New()

	r.AddUser(1, "Boris")
	r.AddUser(2, "Stepan")

	u, err := r.GetByID(3)
	if errors.Is(err, repository.ErrUserNotFound) {
		fmt.Println(err)
		fmt.Println("user not found")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ID: %d, Name: %s", u.ID(), u.Name())
	}
}
