package main

import (
	"jun3_7/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler.HandleUsers)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Ошибка сервера: ", err)
	}
}
