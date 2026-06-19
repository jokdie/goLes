package main

import (
	"jun3_9/internal/config"
	"jun3_9/internal/handler"
	"jun3_9/internal/service"
	"jun3_9/repository"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	rep := repository.NewUserRepository()
	srv := service.NewUserService(rep)
	u := handler.NewUserHandler(srv)
	router := handler.NewRouter(u)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
