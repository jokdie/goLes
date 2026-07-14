package email

import (
	"context"
	"errors"
	"log/slog"
	"math/rand"
	"provider/internal/model"
	"provider/internal/requestid"
	"time"
)

type Provider struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Provider {
	return &Provider{logger}
}

func (p *Provider) Send(
	ctx context.Context,
	req model.ProviderRequest,
) error {
	p.logger.Info(
		"[Email] request.",
		slog.String("X-GUID", requestid.Get(ctx)),
		slog.Int("User", req.UserID),
		slog.String("Message", req.Message),
	)

	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-time.After(250 * time.Millisecond):
		if rand.Intn(100) > 85 {
			return errors.New("email failed")
		}

		return nil
	}
}
