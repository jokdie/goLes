package main

import (
	"jun4_1/internal/config"
	"jun4_1/internal/handler"
	"jun4_1/internal/middleware"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	hHello := handler.NewHelloHandler()
	router := handler.NewRouter(hHello)
	router = middleware.LoggingMiddleware(router)
	router = middleware.RequestIDMiddleware(router)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
