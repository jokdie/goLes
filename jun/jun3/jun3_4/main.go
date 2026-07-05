package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received")

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Ошибка сервера: ", err)
	}
}
