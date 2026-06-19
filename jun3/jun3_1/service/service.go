package service

import (
	"fmt"
	"jun3_1/logger"
)

type UserService struct {
	logger logger.Logger
}

func (s UserService) CreateUser(name string) {
	s.logger.Log("User created: " + name)
}

type AdminService struct {
	logger.Logger
}

func (s AdminService) CreateAdmin(name string) {
	s.Log("Admin created: " + name)
}

func (s AdminService) Log(message string) {
	fmt.Printf("[AdminService][LOG] %s\n", message)
}
