package service

import (
	"context"
	"corporation-site/domain"
)

type EmailService interface {
	SendTemplatedEmail(ctx context.Context, req *domain.EmailRequest) (*domain.EmailResponse, error)
}

type emailService struct {
	emailClient domain.EmailClient
}

func NewEmailService(emailClient domain.EmailClient) EmailService {
	return &emailService{
		emailClient: emailClient,
	}
}

func (s *emailService) SendTemplatedEmail(ctx context.Context, req *domain.EmailRequest) (*domain.EmailResponse, error) {

	return s.emailClient.Send(ctx, &domain.EmailMessage{
		Subject:     req.Subject,
		TextBody:    req.TextBody,
		HTMLBody:    req.HTMLBody,
		Attachments: req.Attachments,
	})
}
