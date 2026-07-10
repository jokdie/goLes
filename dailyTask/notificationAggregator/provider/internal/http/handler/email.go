package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"provider/internal/model"
)

func (h *Handler) email(w http.ResponseWriter, r *http.Request) {
	var req model.ProviderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Decode ProviderRequest error: %v", err)

		errCode := http.StatusBadRequest
		writeJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Некорректный JSON"})

		return
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("Validate ProviderRequest error: %v", err)

		errCode := http.StatusBadRequest
		writeJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Некорректный JSON"})

		return
	}

	if err := h.emailService.SendEmail(r.Context(), req); err != nil {
		log.Printf("SendEmail error: %v", err)

		errCode := http.StatusInternalServerError
		writeJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Internal server error"})
	}

	w.WriteHeader(http.StatusNoContent)
}
