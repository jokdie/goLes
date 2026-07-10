package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"provider/internal/model"

	"github.com/go-playground/validator/v10"
)

type EmailService interface {
	SendEmail(ctx context.Context, req model.ProviderRequest) error
}

type Handler struct {
	validate     *validator.Validate
	emailService EmailService
}

func NewHandler(
	validate *validator.Validate,
	emailService EmailService,
) *Handler {
	return &Handler{
		validate:     validate,
		emailService: emailService,
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка записи json: %s", err)
	}
}
