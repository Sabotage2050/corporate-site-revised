package aws

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// 環境の定数定義
const (
	EnvLocal   = "local"
	EnvTest    = "test"
	EnvDev     = "dev"
	EnvCI      = "ci"
	EnvStaging = "staging"
	EnvProd    = "prod"
)

// SharedConfig は AWS サービス共通の設定を保持する
type SharedConfig struct {
	Region           string
	Profile          string
	RetryMaxAttempts int
	TimeoutDuration  time.Duration
}

// NewSharedConfig は新しい共有設定を作成する
func NewSharedConfig() *SharedConfig {
	return &SharedConfig{
		Region:           os.Getenv("AWS_REGION"),
		Profile:          os.Getenv("AWS_PROFILE"),
		RetryMaxAttempts: 3,
		TimeoutDuration:  20 * time.Second,
	}
}

func (c *SharedConfig) LoadAWSConfig(ctx context.Context) (aws.Config, error) {
	env := os.Getenv("RUN_ENV")
	if env == "" {
		env = EnvLocal
	}

	fmt.Printf("Loading AWS Config for environment: %s\n", env)
	fmt.Printf("AWS_REGION: %s\n", os.Getenv("AWS_REGION"))

	var optFns []func(*config.LoadOptions) error

	// 環境に応じた認証設定
	switch env {
	case EnvStaging, EnvProd:
		// ECS環境では特別な設定は不要
		// コンテナの認証情報が自動的に使用される
		optFns = append(optFns,
			config.WithRegion(c.Region),
			config.WithRetryMaxAttempts(c.RetryMaxAttempts),
			config.WithRetryMode(aws.RetryModeStandard),
		)
	default:
		// ローカル環境など
		if c.Profile != "" {
			optFns = append(optFns, config.WithSharedConfigProfile(c.Profile))
		}
		optFns = append(optFns,
			config.WithRegion(c.Region),
			config.WithRetryMaxAttempts(c.RetryMaxAttempts),
			config.WithRetryMode(aws.RetryModeStandard),
		)
	}

	// HTTP クライアントの設定
	optFns = append(optFns, config.WithHTTPClient(&http.Client{
		Timeout: c.TimeoutDuration,
	}))

	cfg, err := config.LoadDefaultConfig(ctx, optFns...)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// 認証情報の検証（デバッグ情報を追加）
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to retrieve credentials: %w", err)
	}

	// 認証情報が取得できたことを確認
	fmt.Printf("Successfully retrieved credentials. Provider: %s\n", creds.Source)

	return cfg, nil
}
