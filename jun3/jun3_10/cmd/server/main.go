package main

import (
	"jun3_10/internal/config"
	"jun3_10/internal/handler"
	"jun3_10/internal/middleware"
	"jun3_10/internal/repository"
	"jun3_10/internal/validator"
	"jun3_10/service"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	uRepository := repository.NewUserRepository()
	uService := service.NewUserService(uRepository)
	uValidator := validator.NewUserValidator()
	uHandler := handler.NewUserHandler(uService, uValidator)
	router := handler.NewRouter(uHandler)
	router = middleware.LoggingMiddleware(router)
	router = middleware.RequestIDMiddleware(router)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Ошибка запуска сервера", err)
	}
}
