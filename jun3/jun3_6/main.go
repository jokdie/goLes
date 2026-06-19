package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var u CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("Add new user: %+v\n", u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(SuccessResponse{Success: true}); err != nil {
		log.Printf("Write response error: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handleUsers)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Ошибка сервера: ", err)
	}
}
