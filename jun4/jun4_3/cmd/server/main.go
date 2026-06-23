package main

import (
	"jun4_3/internal/handler"
	"jun4_3/internal/service"
	"log"
	"net/http"
)

func main() {
	s := service.NewWorkService()
	h := handler.NewWorkHandler(s)
	router := handler.NewRouter(h)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
