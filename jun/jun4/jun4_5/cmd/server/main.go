package main

import (
	"jun4_5/internal/handler"
	"jun4_5/internal/middleware"
	"jun4_5/internal/repository"
	"jun4_5/internal/service"
	"log"
	"net/http"
)

func main() {
	rep := repository.NewTestRepository()
	src := service.NewTestService(rep)
	testHandler := handler.NewTestHandler(src)
	router := handler.NewRouter(testHandler)
	router = middleware.RequestIDMiddleware(router)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("fatal: ", err)
	}
}
