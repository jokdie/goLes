package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"provider/internal/http/response"
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

// @todo удалить все writeJSON из метода decodeRequest
func decodeRequest(w http.ResponseWriter, r *http.Request) (model.ProviderRequest, error) {
	var req model.ProviderRequest

	const maxBodySize = 1 << 20 // 1 MB

	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		var maxBytesErr *http.MaxBytesError

		switch {
		case errors.As(err, &maxBytesErr):
			errCode := http.StatusRequestEntityTooLarge

			response.WriteJSON(w, errCode, model.ErrorResponse{
				Code:    errCode,
				Message: "Request body too large",
			})

		default:
			errCode := http.StatusBadRequest

			response.WriteJSON(w, errCode, model.ErrorResponse{
				Code:    errCode,
				Message: "Некорректный JSON",
			})
		}

		return model.ProviderRequest{}, err
	}

	if decoder.Decode(&struct{}{}) != io.EOF {
		errCode := http.StatusBadRequest

		response.WriteJSON(w, errCode, model.ErrorResponse{
			Code:    errCode,
			Message: "Разрешен только один объект JSON",
		})

		return model.ProviderRequest{}, errors.New("request body must contain only one JSON object")
	}

	return req, nil
	// @todo list
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
}
