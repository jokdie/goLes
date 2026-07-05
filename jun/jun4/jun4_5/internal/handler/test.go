package handler

import (
	"context"
	"encoding/json"
	"jun4_5/internal/model"
	"log"
	"net/http"
)

type service interface {
	TestSrcDo(ctx context.Context) model.SuccessResponse
}

type TestHandler struct {
	service service
}

func NewTestHandler(s service) *TestHandler {
	return &TestHandler{service: s}
}

func (h *TestHandler) Test(w http.ResponseWriter, r *http.Request) {
	response := h.service.TestSrcDo(r.Context())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("encoder error:", err)
	}
}
