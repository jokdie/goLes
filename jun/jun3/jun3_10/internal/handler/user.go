package handler

import (
	"encoding/json"
	"fmt"
	"jun3_10/internal/dto"
	"jun3_10/internal/model"
	"net/http"
)

type UserService interface {
	Create(dto *dto.CreateUserRequest) model.User
}

type UserValidator interface {
	Validate(dto *dto.CreateUserRequest) error
}

type UserHandler struct {
	service   UserService
	validator UserValidator
}

func NewUserHandler(s UserService, v UserValidator) *UserHandler {
	return &UserHandler{
		service:   s,
		validator: v,
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Ошибка записи ответа:", err)
	}
}

func (h *UserHandler) User(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeJSON(w, http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request data"})
		return
	}

	if err := h.validator.Validate(&dto); err != nil {
		writeJSON(w, http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	u := h.service.Create(&dto)
	writeJSON(w, http.StatusCreated, u)
}
