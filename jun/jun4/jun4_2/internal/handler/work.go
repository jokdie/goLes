package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"jun4_2/internal/model"
	"net/http"
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

func (h *WorkHandler) Work(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.s.Work(ctx)
	if err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			fmt.Println("Отмена запроса:", err)

			return
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
