// tests/common/setup/server.go
package setup

import (
	"context"
	"corporation-site/data/seeds"
	"corporation-site/infra/db/nosql"
	"corporation-site/infra/email"
	"corporation-site/infra/log"
	"corporation-site/infra/router"
	"corporation-site/infra/validation"
	"fmt"
	stdlog "log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	TestServer *httptest.Server
	once       sync.Once
	dbClient   nosql.NoSQLClient
)

func SetupTestServer() error {
	var setupErr error
	once.Do(func() {
		setupErr = setupTestEnvironment()
	})
	return setupErr
}

func setupTestEnvironment() error {
	if os.Getenv("RUN_ENV") == "local" {
		if err := godotenv.Overload("../../.env.test"); err != nil {
			stdlog.Printf("Warning: .env.test file not found: %v", err)
		}
	}

	var err error
	// DBクライアントの作成
	dbClient, err = nosql.NewNosqlDBFactory(nosql.DynamoDBStorage)
	if err != nil {
		return fmt.Errorf("failed to create DB client: %w", err)
	}

	// メールクライアントの作成
	emailClient, err := email.NewEmailClientFactory(email.SESProvider)
	if err != nil {
		return fmt.Errorf("failed to create DB client: %w", err)
	}

	// ロガーの作成（追加）
	logger, err := log.NewLoggerFactory(log.InstanceZapLogger)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}

	// バリデーターの作成
	validator, err := validation.NewValidatorFactory(validation.InstanceGoPlayground)
	if err != nil {
		return fmt.Errorf("failed to create validator: %w", err)
	}

	// ポート番号の設定
	portStr := os.Getenv("APP_PORT")
	portNum, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid port number: %w", err)
	}

	port := router.Port(portNum)
	// サーバーファクトリーにロガーを追加
	server, err := router.NewWebServerFactory(
		router.InstanceCheckApiGin,
		port,
		dbClient,
		emailClient,
		logger,
		validator,
	)
	if err != nil {
		return err
	}

	if err := server.SetupRouter(); err != nil {
		return err
	}

	TestServer = httptest.NewServer(server.GetHandler())
	return nil
}

func GetTestDBClient() nosql.NoSQLClient {
	return dbClient
}

func reloadTestDataForTable(tableName string) error {
	dataDir := filepath.Join("..", "..", "data", "seeds", "csv", "test")
	loader := seeds.NewDataLoader(dataDir, seeds.Test)

	ctx := context.Background()

	switch tableName {
	case "Forklift":
		var forkliftSeeds []seeds.ForkliftSeed
		if err := loader.LoadData(&forkliftSeeds); err != nil {
			return fmt.Errorf("failed to load forklift seeds: %w", err)
		}

		for _, seed := range forkliftSeeds {
			if err := dbClient.PutItem(ctx, seed.TableName(), seed.ToMap()); err != nil {
				return fmt.Errorf("failed to insert forklift seed data: %w", err)
			}
		}
	default:
		return fmt.Errorf("unknown table: %s", tableName)
	}

	return nil
}

func TearDownTestServer() {
	if TestServer != nil {
		TestServer.Close()
	}
}

// tests/common/setup/server.go

func ResetTestDatabase(tables ...string) error {
	ctx := context.Background()

	// 指定されたテーブルのデータをクリア
	for _, table := range tables {
		results, err := dbClient.Scan(ctx, table)
		if err != nil {
			return fmt.Errorf("failed to scan table %s: %w", table, err)
		}

		for _, item := range results {
			// 複合キーを使用してアイテムを削除
			key := map[string]interface{}{
				"Enginetype": item["Enginetype"], // 修正: Type -> Enginetype
				"SerialNo":   item["SerialNo"],   // 修正: ID -> SerialNo
			}
			if err := dbClient.DeleteItem(ctx, table, key); err != nil {
				return fmt.Errorf("failed to delete item in table %s: %w", table, err)
			}
		}

		// テーブルをクリアした後、該当するテストデータを再ロード
		if err := reloadTestDataForTable(table); err != nil {
			return fmt.Errorf("failed to reload test data for table %s: %w", table, err)
		}
	}

	return nil
}

func TearDownTestDatabase() error {
	ctx := context.Background()
	tables := []string{"Forklift"} // 必要なテーブルを追加

	// 全テーブルをクリア
	for _, table := range tables {
		results, err := dbClient.Scan(ctx, table)
		if err != nil {
			return fmt.Errorf("failed to scan table %s: %w", table, err)
		}

		for _, item := range results {
			// 複合キーを使用してアイテムを削除
			key := map[string]interface{}{
				"Enginetype": item["Enginetype"], // 修正: Type -> Enginetype
				"SerialNo":   item["SerialNo"],
			}
			if err := dbClient.DeleteItem(ctx, table, key); err != nil {
				return fmt.Errorf("failed to delete item in table %s: %w", table, err)
			}
		}
	}

	return nil
}
