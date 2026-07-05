package handler

import (
	"context"
	"encoding/json"
	"jun4_6/internal/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

type service interface {
	GetUser(ctx context.Context, id int) (model.User, error)
}

type UserHandler struct {
	service service
}

func NewUserHandler(s service) *UserHandler {
	return &UserHandler{service: s}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка записи json: %s", err)
	}
}

func (h *UserHandler) User(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	rawCancel := r.URL.Query().Get("cancel")
	if Cancel, err := strconv.ParseBool(rawCancel); err == nil && Cancel {
		time.AfterFunc(2*time.Second, func() {
			cancel()
		})
	}

	user, err := h.service.GetUser(ctx, id)

	if err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			model.ErrorResponse{Error: err.Error()},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		model.SuccessResponse{User: user},
	)
}
