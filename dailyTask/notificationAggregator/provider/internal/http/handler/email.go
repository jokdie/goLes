package handler

import (
	"fmt"
	"net/http"
	"provider/internal/http/response"
)

func (h *Handler) email(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRequest(w, r)

	if err != nil {
		err = fmt.Errorf("decode ProviderRequest error: %v", err)
		response.BadRequest(w, err)

		return
	}

	if err := h.validate.Struct(req); err != nil {
		err = fmt.Errorf("validate ProviderRequest error: %v", err)
		response.BadRequest(w, err)

		return
	}

	if err := h.emailService.SendEmail(r.Context(), req); err != nil {
		err = fmt.Errorf("SendEmail error: %v", err)
		response.Internal(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
