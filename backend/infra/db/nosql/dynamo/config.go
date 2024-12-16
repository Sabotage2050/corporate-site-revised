// internal/infra/db/nosql/dynamo/config.go
package dynamo

import (
	"corporation-site/infra/aws"
	"os"
)

type Config struct {
	SharedConfig *aws.SharedConfig
	TablePrefix  string
	TableName    string
}

func NewConfig() *Config {
	return &Config{
		SharedConfig: aws.NewSharedConfig(),
		TablePrefix:  os.Getenv("APP_NAME") + "-" + os.Getenv("RUN_ENV") + "-",
		TableName:    os.Getenv("DYNAMODB_TABLE_NAME"),
	}
}
