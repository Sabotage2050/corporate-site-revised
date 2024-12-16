// infra/email/ses/config.go
package ses

import (
	"corporation-site/infra/aws"
	"fmt"
	"os"
)

type Config struct {
	SharedConfig *aws.SharedConfig
	FromAddress  string
	ToAddress    string
}

func NewConfig() *Config {
	fromAddr := os.Getenv("SES_FROM_ADDRESS")
	toAddr := os.Getenv("SES_TO_ADDRESS")

	fmt.Println(fromAddr)
	fmt.Println(toAddr)
	// 環境変数の検証とログ出力
	if fromAddr == "" || toAddr == "" {
		errMsg := fmt.Sprintf(
			"Required environment variables not set. SES_FROM_ADDRESS='%s', SES_TO_ADDRESS='%s'",
			fromAddr,
			toAddr,
		)
		// CloudWatchに出力されるようにfmt.Printlnを使用
		fmt.Println(errMsg)
		return nil
	}

	return &Config{
		SharedConfig: aws.NewSharedConfig(),
		FromAddress:  os.Getenv("SES_FROM_ADDRESS"),
		ToAddress:    os.Getenv("SES_TO_ADDRESS"),
	}
}
