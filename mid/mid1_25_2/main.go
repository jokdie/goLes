package main

import (
	"log"
	"net/http"
	"test/internal/handler"
	"test/internal/repository"
	"test/internal/service"
)

func main() {
	userRep := repository.NewUserRepository()
	orderRep := repository.NewOrderRepository()
	profileRep := repository.NewProfileRepository()
	s := service.NewUserService(userRep, orderRep, profileRep)
	routerApp := handler.Router(s)

	if err := http.ListenAndServe(":8080", routerApp); err != nil {
		log.Fatal("Fatal ListenAndServe")
	}
}
