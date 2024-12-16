package email

import (
	"context"
	"corporation-site/domain"
	"corporation-site/infra/email/ses"
	"errors"
	"fmt"
)

var (
	ErrInvalidEmailProvider = errors.New("invalid email provider")
)

const (
	SESProvider int = iota
	SMTPProvider
)

// NewEmailClient creates a new email client implementation
func NewEmailClientFactory(provider int) (domain.EmailClient, error) {
	ctx := context.Background()

	switch provider {
	case SESProvider:
		config := ses.NewConfig()
		client, err := ses.NewSESClient(ctx, config)
		if err != nil {
			fmt.Printf("error %v", err)
			return nil, fmt.Errorf("failed to create SES client: %w", err)
		}
		return client, nil
	default:
		return nil, ErrInvalidEmailProvider
	}
}
