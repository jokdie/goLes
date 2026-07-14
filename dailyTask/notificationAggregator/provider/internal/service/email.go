package service

import (
	"context"
	"errors"
	"log/slog"
	"math/rand"
	"provider/internal/model"
	"provider/internal/requestid"
	"time"
)

type EmailService struct {
	logger *slog.Logger
}

func NewEmailService(logger *slog.Logger) *EmailService {
	return &EmailService{logger}
}

func (s *EmailService) SendEmail(ctx context.Context, req model.ProviderRequest) error {
	s.logger.Info(
		"[EmailService] лог для тестирования.",
		slog.String("X-GUID", requestid.Get(ctx)),
		slog.Int("User", req.UserID),
		slog.String("Message", req.Message),
	)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(250 * time.Millisecond):
		minV := 1
		maxV := 100
		randomNumber := rand.Intn(maxV-minV+1) + minV
		if randomNumber > 85 {
			return errors.New("email service is failed")
		}

		return nil
	}
}
