// internal/infra/db/nosql/factory.go
package nosql

import (
	"context"
	"corporation-site/infra/db/nosql/dynamo"
	"errors"
	"fmt"
)

var (
	ErrInvalidDBInstance = errors.New("invalid database instance")
)

const (
	DynamoDBStorage int = iota
	TinyDBStorage
)

func NewNosqlDBFactory(instance int) (NoSQLClient, error) {
	ctx := context.Background()

	switch instance {
	case DynamoDBStorage:
		config := dynamo.NewConfig()
		client, err := dynamo.NewDynamoDBClient(ctx, config)
		if err != nil {
			return nil, fmt.Errorf("failed to create DynamoDB client: %w", err)
		}
		return client, nil
	default:
		return nil, ErrInvalidDBInstance
	}
}
