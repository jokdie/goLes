package service

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"provider/internal/model"
	"provider/internal/requestid"
	"time"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendEmail(ctx context.Context, req model.ProviderRequest) error {
	log.Printf(
		"[EmailService] лог для тестирования.\n[X-GUID]: %v\n[User]: %d\n[Message]: %s\n",
		requestid.GetRequestID(ctx),
		req.UserID,
		req.Message,
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
