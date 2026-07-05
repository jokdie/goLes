package main

import (
	"jun4_6/internal/handler"
	"jun4_6/internal/repository"
	"jun4_6/internal/service"
	"log"
	"net/http"
)

func main() {
	repository := repository.NewUserRepository()
	service := service.NewUserService(repository)
	userHandler := handler.NewUserHandler(service)
	router := handler.NewRouter(userHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
