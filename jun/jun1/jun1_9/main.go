package main

import "fmt"

type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Printf("[LOG] %s", message)
}

type UserService struct {
	Logger
}

func (s UserService) CreateUser(name string) {
	s.Log(fmt.Sprintf("creating user: <%s>", name))
}

func main() {
	service := UserService{}

	service.CreateUser("Alice")
	service.CreateUser("Miromir")
	service.CreateUser("Rostislav")
	service.CreateUser("Svyatozar")
	service.Log("hello")
}
