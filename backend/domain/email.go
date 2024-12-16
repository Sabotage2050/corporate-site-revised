package domain

import "context"

// EmailMessage represents the complete email with all metadata
type EmailMessage struct {
	From        string            `json:"from" validate:"required,email"`
	To          string            `json:"to" validate:"required,email"`
	Subject     string            `json:"subject" validate:"required"`
	TextBody    string            `json:"textBody" validate:"required_without=HTMLBody"`
	HTMLBody    string            `json:"htmlBody" validate:"required_without=TextBody"`
	Attachments []EmailAttachment `json:"attachments" validate:"dive"`
}

// EmailRequest represents the API request without From/To fields
type EmailRequest struct {
	Subject     string            `json:"subject" validate:"required"`
	TextBody    string            `json:"textBody" validate:"required_without=HTMLBody"`
	HTMLBody    string            `json:"htmlBody" validate:"required_without=TextBody"`
	Data        map[string]string `json:"data"`
	Attachments []EmailAttachment `json:"attachments" validate:"dive"`
}

type EmailAttachment struct {
	Filename string `json:"filename" validate:"required"`
	Content  []byte `json:"content" validate:"required"`
	MIMEType string `json:"mimeType" validate:"required"`
}

type EmailResponse struct {
	MessageID string `json:"messageId"`
	Status    string `json:"status"`
}

// EmailClient defines the interface for email operations
type EmailClient interface {
	Send(ctx context.Context, msg *EmailMessage) (*EmailResponse, error)
}
