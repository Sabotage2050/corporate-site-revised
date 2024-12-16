// internal/infra/db/nosql/interface.go
package nosql

import (
	"context"
)

// NoSQLClient は DynamoDB などの NoSQL データベースの共通インターフェース
type NoSQLClient interface {
	CheckConnection(ctx context.Context) error
	GetItem(ctx context.Context, tableName string, key map[string]interface{}) (map[string]interface{}, error)
	PutItem(ctx context.Context, tableName string, item map[string]interface{}) error
	DeleteItem(ctx context.Context, tableName string, key map[string]interface{}) error
	Query(ctx context.Context, tableName string, keyCondition string, attributeValues map[string]interface{}) ([]map[string]interface{}, error)
	Scan(ctx context.Context, tableName string) ([]map[string]interface{}, error)
}
