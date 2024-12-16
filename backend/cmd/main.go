// cmd/main.go
package main

import (
	"corporation-site/infra/db/nosql"
	"corporation-site/infra/email"
	"corporation-site/infra/log"
	"corporation-site/infra/router"
	"corporation-site/infra/server"
	"corporation-site/infra/validation"
	"fmt"
	stdlog "log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		stdlog.Printf("Warning: .env file not found: %v", err)
	}
	println("Debug")
	dir, err := os.Getwd()
	if err != nil {
		stdlog.Fatal(err)
	}
	fmt.Println("Current directory:", dir)

	// 各コンポーネントの初期化
	envVars := os.Environ()

	fmt.Println(envVars)
	app := server.NewConfig().
		Name(os.Getenv("APP_NAME")).
		DbNoSQL(nosql.DynamoDBStorage).
		Email(email.SESProvider).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceCheckApiGin)

	app.Start()
}
