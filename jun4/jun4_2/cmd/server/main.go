package main

import (
	"jun4_2/internal/handler"
	"jun4_2/internal/service"
	"log"
	"net/http"
)

func main() {
	s := service.NewWorkService()
	h := handler.NewWorkHandler(s)
	r := handler.NewRouter(h)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
