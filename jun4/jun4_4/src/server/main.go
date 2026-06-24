package main

import (
	"jun4_4/internal/handler"
	"jun4_4/internal/middleware"
	"log"
	"net/http"
)

func main() {
	h := handler.NewProfileHandle()
	r := handler.NewRouter(h)
	r = middleware.RequestIDMiddleware(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
