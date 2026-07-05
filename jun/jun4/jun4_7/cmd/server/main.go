package main

import (
	"jun4_7/internal/handler"
	"jun4_7/internal/repository"
	"jun4_7/internal/service"
	"log"
	"net/http"
)

func main() {
	r := repository.NewRepository()
	s := service.NewTestService(r)
	h := handler.NewTestHandler(s)
	router := handler.NewRouter(h)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
