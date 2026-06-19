package handler

import (
	"encoding/json"
	"jun3_8/internal/model"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка записи json: %s", err)
	}
}

type HelloHandler struct {
	// место для инъекций зависимости сервиса с бизнес логикой, за которым уже мб -> репа -> бд
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		writeJSON(
			w,
			http.StatusBadRequest,
			model.ErrorResponse{Error: "Name is required"},
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		model.SuccessResponse{Message: "Hello, " + name},
	)
}
