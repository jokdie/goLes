package main

import (
	"fmt"
	"jun5_7/repository"
)

func main() {
	s := repository.NewUserStore()
	u := repository.User{ID: 123, Name: "Bob"}

	s.Add(u)
	bob, ok := s.Get(123)
	fmt.Println("Bob: ", bob, ok)
	bob, ok = s.Get(1423)
	fmt.Println("Bob: ", bob, ok)

	fmt.Println("Result: ", s, s.Count())
}
