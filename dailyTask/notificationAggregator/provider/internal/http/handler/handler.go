package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"provider/internal/model"

	"github.com/go-playground/validator/v10"
)

// вместо провайдера лучше так
//NotificationService
//    │
//    ├── EmailProvider
//    ├── SmsProvider
//    └── PushProvider
//
//type NotificationService struct {
//    providers map[Channel]provider.Provider
//}
//func (s *NotificationService) Send(
//    ctx context.Context,
//    channel Channel,
//    req model.ProviderRequest,
//) error {
//
//    p, ok := s.providers[channel]
//    if !ok {
//        return ErrUnknownProvider
//    }
//
//    return p.Send(ctx, req)
//}
//Тогда хендлер вообще ничего не знает про email/sms/push:
//err := notificationService.Send(
//    ctx,
//    model.Email,
//    req,
//)

// @todo перейти на type Provider interface и удалить сервис что есть
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
	// email provider.Provider,
	// sms   provider.Provider,
	// push  provider.Provider,
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

func respondBadRequest(w http.ResponseWriter, err error) {
	log.Println(err.Error())

	errCode := http.StatusBadRequest
	writeJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Некорректный JSON"})
}

func respondInternal(w http.ResponseWriter, err error) {
	log.Println(err.Error())

	errCode := http.StatusInternalServerError
	writeJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Internal server error"})
}

func decodeRequest(w http.ResponseWriter, r *http.Request) (model.ProviderRequest, error) {
	defer r.Body.Close()
	var req model.ProviderRequest

	// @todo list
	// 1) сделать кейс r.Body = http.MaxBytesReader чтобы не слали 1гб, например ограничение 1мб
	// 2) json.NewDecoder(r.Body).Decode(&req) передеалть на decoder.DisallowUnknownFields()
	// 3) добавить проверку в целом везде на "application/json"
	// 4) в requestId добавить
	// if requestID == "" {
	// 	requestID = uuid.NewString()
	// }
	// 5) все логи переделать на "slog"
	// 6) все errors.New лучше переделать на var ErrProviderUnavailable = errors.New(...) +errors.Is(...)
	// 7) добавить конфиг env и туда 8080 пихнуть
	// 8) добавить серверу таймауты ReadTimeout + WriteTimeout + IdleTimeout + ReadHeaderTimeout
	// 9) структура папок
	//cmd/
	//internal/
	//    app/
	//    config/
	//    logger/
	//    transport/http/
	//        handler/
	//        middleware/
	//        response/
	//        request/
	//    service/
	//    provider/
	//        email/
	//        sms/
	//        push/
	//    repository/
	//    model/
	//    errors/

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("decodeRequest error: %v", err)

		errCode := http.StatusBadRequest
		writeJSON(w, errCode, model.ErrorResponse{Code: errCode, Message: "Некорректный JSON"})

		return model.ProviderRequest{}, err
	}

	return req, nil
}
