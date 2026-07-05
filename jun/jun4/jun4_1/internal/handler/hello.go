package handler

import (
	"encoding/json"
	"jun4_1/internal/model"
	"jun4_1/internal/requestid"
	"log"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Write response error: %v", err)
	}
}

func (h *HelloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	requestID := requestid.GetRequestID(r.Context())
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	writeJSON(w, model.SuccessResponse{
		Message:   "Hello, " + name,
		RequestId: requestID,
	})
}
