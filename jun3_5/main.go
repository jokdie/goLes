package main

import (
	"encoding/json"
	"fmt"
	"jun3_5/response"
	"net/http"
)

func handleUser(w http.ResponseWriter, r *http.Request) {
	u := response.CreateSuccessResponse() // типа в репе
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handleUser)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}
