package main

import "fmt"

type Auditor struct{}

func (Auditor) Audit(action string) {
	fmt.Printf("AUDIT: <%s>\n", action)
}

type UserService struct {
	Auditor
}

func (s UserService) CreateUser(name string) {
	s.Audit(name)
	s.Auditor.Audit(name)
}

func (UserService) Audit(action string) {
	fmt.Printf("USER SERVICE AUDIT: <%s>\n", action)
}

func main() {
	service := UserService{}

	service.CreateUser("Gomer")
}
