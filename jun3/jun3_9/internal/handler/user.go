package handler

import (
	"encoding/json"
	"errors"
	"jun3_9/internal/model"
	"jun3_9/repository"
	"log"
	"net/http"
	"strconv"
)

type UserService interface {
	ByID(id int) (model.User, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка записи json: %s", err)
	}
}

func (u *UserHandler) User(w http.ResponseWriter, r *http.Request) {
	rawId := r.PathValue("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			model.ErrorResponse{Error: "Invalid user ID"},
		)
		return
	}

	user, err := u.service.ByID(id)

	if errors.Is(err, repository.UserNotFoundError) {
		writeJSON(
			w,
			http.StatusNotFound,
			model.ErrorResponse{Error: err.Error()},
		)
		return
	}
	if err != nil {
		writeJSON(
			w,
			http.StatusInternalServerError,
			model.ErrorResponse{Error: "Internal server error"},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		model.SuccessResponse{ID: user.ID},
	)
}
