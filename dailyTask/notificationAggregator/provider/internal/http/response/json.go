package response

import (
	"encoding/json"
	"log"
	"net/http"
	"provider/internal/model"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка записи json: %s", err)
	}
}

func BadRequest(w http.ResponseWriter, err error) {
	log.Println(err.Error())

	errCode := http.StatusBadRequest
	WriteJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Некорректный JSON"})
}

func Internal(w http.ResponseWriter, err error) {
	log.Println(err.Error())

	errCode := http.StatusInternalServerError
	WriteJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Internal server error"})
}
