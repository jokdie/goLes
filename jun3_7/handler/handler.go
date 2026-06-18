package handler

import (
	"encoding/json"
	"jun3_7/repository"
	"jun3_7/validator"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Success bool `json:"success"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func errorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Error:   message,
	}); err != nil {
		log.Printf("Write ErrorResponse error: %v", err)
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var u repository.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("Decode User error: %v", err)
		errorResponse(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if err := validator.ValidateUser(u); err != nil {
		errorResponse(w, http.StatusUnprocessableEntity, "Validation failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(SuccessResponse{Success: true}); err != nil {
		log.Printf("Write SuccessResponse error: %v", err)
	}
}
