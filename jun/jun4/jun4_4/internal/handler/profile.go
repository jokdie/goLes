package handler

import (
	"encoding/json"
	"jun4_4/internal/userid"
	"net/http"
)

type SuccessResponse struct {
	UserID string `json:"user_id"`
}

type ProfileHandler struct{}

func NewProfileHandle() *ProfileHandler {
	return &ProfileHandler{}
}

func (h *ProfileHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userId := userid.GetUserIDCtxKey(r.Context())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{UserID: userId})
}
