package main

import (
	"log"
	"mid1_25/internal/handler"
	"mid1_25/internal/repository"
	"mid1_25/internal/service"
	"net/http"
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
