package ses

import (
	"context"
	"corporation-site/domain"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type Client struct {
	client *ses.Client
	config *Config
}

func NewSESClient(ctx context.Context, cfg *Config) (domain.EmailClient, error) {

	awsCfg, err := cfg.SharedConfig.LoadAWSConfig(ctx)
	if err != nil {
		fmt.Printf("AWS Config Error: %v\n", err) // エラー内容を出力
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &Client{
		client: ses.NewFromConfig(awsCfg),
		config: cfg,
	}, nil
}

func (c *Client) Send(ctx context.Context, msg *domain.EmailMessage) (*domain.EmailResponse, error) {
	// 設定からFromとToを取得
	msg.From = c.config.FromAddress
	if msg.To == "" {
		msg.To = c.config.ToAddress
	}

	message := &types.Message{
		Subject: &types.Content{
			Data:    &msg.Subject,
			Charset: aws.String("UTF-8"),
		},
		Body: &types.Body{},
	}

	if msg.TextBody != "" {
		message.Body.Text = &types.Content{
			Data:    &msg.TextBody,
			Charset: aws.String("UTF-8"),
		}
	}

	if msg.HTMLBody != "" {
		message.Body.Html = &types.Content{
			Data:    &msg.HTMLBody,
			Charset: aws.String("UTF-8"),
		}
	}

	result, err := c.client.SendEmail(ctx, &ses.SendEmailInput{
		Source: &msg.From,
		Destination: &types.Destination{
			ToAddresses: []string{msg.To}, // 単一のアドレスをスライスに変換
		},
		Message: message,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %w", err)
	}

	return &domain.EmailResponse{
		MessageID: *result.MessageId,
		Status:    "sent",
	}, nil
}
