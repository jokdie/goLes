package main

import (
	"jun3_8/internal/config"
	"jun3_8/internal/handler"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	helloHandler := handler.NewHelloHandler()
	router := handler.NewRouter(helloHandler)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
