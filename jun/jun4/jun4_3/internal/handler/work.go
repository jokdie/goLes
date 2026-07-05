package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"jun4_3/internal/model"
	"net/http"
	"time"
)

type Work interface {
	Work(ctx context.Context) (model.SuccessResponse, error)
}

type WorkHandler struct {
	s Work
}

func NewWorkHandler(s Work) *WorkHandler {
	return &WorkHandler{s: s}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *WorkHandler) Work(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	data, err := h.s.Work(ctx)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			writeJSON(w, http.StatusRequestTimeout, model.ErrorResponse{Error: "timeout"})
			return
		case errors.Is(err, context.Canceled):
			fmt.Println("Клиент отменил запрос")
			return
		default:
			writeJSON(w, http.StatusInternalServerError, model.ErrorResponse{Error: "Internal server error"})
			return
		}
	}

	writeJSON(w, http.StatusOK, data)
}
