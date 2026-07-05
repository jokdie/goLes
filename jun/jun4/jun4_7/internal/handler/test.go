package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"jun4_7/internal/model"
	"net/http"
	"time"
)

const (
	timeoutMode = "timeout"
	cancelMode  = "cancel"
)

type service interface {
	Do(ctx context.Context) error
}

type TestHandler struct {
	service service
}

func NewTestHandler(s service) *TestHandler {
	return &TestHandler{service: s}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println("Ошибка записи ответа:", err)
	}
}

func (h *TestHandler) Do(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var cancel context.CancelFunc

	mode := r.URL.Query().Get("mode")

	switch mode {
	case timeoutMode:
		ctx, cancel = context.WithTimeout(r.Context(), 2*time.Second)
	case cancelMode:
		ctx, cancel = context.WithCancel(r.Context())
	default:
		writeJSON(w, http.StatusBadRequest, model.ErrorResponse{Error: "invalid mode"})
		return
	}

	defer cancel()

	if err := h.service.Do(ctx); err != nil {
		writeJSON(w, http.StatusRequestTimeout, model.ErrorResponse{Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, nil)
}
