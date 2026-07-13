package email

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"provider/internal/model"
	"provider/internal/requestid"
	"time"
)

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Send(
	ctx context.Context,
	req model.ProviderRequest,
) error {

	log.Printf(
		"[Email] request=%s user=%d",
		requestid.GetRequestID(ctx),
		req.UserID,
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
