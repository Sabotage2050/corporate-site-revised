// internal/infra/db/nosql/dynamo/client.go
package dynamo

import (
	"context"
	"corporation-site/data/seeds"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Client struct {
	*dynamodb.Client
	tableNames map[string]string
}

func NewDynamoDBClient(ctx context.Context, cfg *Config) (*Client, error) {
	awsCfg, err := cfg.SharedConfig.LoadAWSConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := &Client{
		Client: dynamodb.NewFromConfig(awsCfg),
		tableNames: map[string]string{
			cfg.TableName: cfg.TablePrefix + cfg.TableName,
		},
	}

	// テーブルの存在確認と作成
	for _, tableName := range client.tableNames {
		_, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})
		if err != nil {
			if err := client.createTable(ctx, tableName); err != nil {
				return nil, fmt.Errorf("failed to create table %s: %w", tableName, err)
			}
		}
	}

	// テストデータのロード（環境変数に基づいて）
	if os.Getenv("RUN_ENV") == "local" || os.Getenv("RUN_ENV") == "staging" {
		if err := client.loadInitialData(ctx); err != nil {
			return nil, fmt.Errorf("failed to load initial data: %w", err)
		}
	}

	return client, nil
}

// 論理テーブル名を物理テーブル名に解決するヘルパーメソッド
func (c *Client) resolveTableName(logicalName string) string {
	if physicalName, exists := c.tableNames[logicalName]; exists {
		return physicalName
	}
	return logicalName // 該当する物理テーブル名がない場合は論理名をそのまま返す
}

// 新しく追加するメソッド
func (c *Client) loadInitialData(ctx context.Context) error {
	// 環境に応じてディレクトリを選択
	env := "dev"
	if os.Getenv("RUN_ENV") == "test" {
		env = "test"
	}

	dataDir := filepath.Join("data", "seeds", "csv", env)
	loader := seeds.NewDataLoader(dataDir, seeds.Test)

	// 全フォークリフトデータをロード
	forkliftSeeds, err := loader.LoadAllForkliftData()
	if err != nil {
		return fmt.Errorf("failed to load forklift seeds: %w", err)
	}

	// 既存データのクリア
	results, err := c.Scan(ctx, "Forklift")
	if err != nil {
		return fmt.Errorf("failed to scan existing data: %w", err)
	}

	for _, item := range results {
		key := map[string]interface{}{
			"Enginetype": item["Enginetype"],
			"SerialNo":   item["SerialNo"],
		}
		if err := c.DeleteItem(ctx, "Forklift", key); err != nil {
			return fmt.Errorf("failed to delete existing item: %w", err)
		}
	}

	// テストデータの挿入
	for _, seed := range forkliftSeeds {
		if err := c.PutItem(ctx, seed.TableName(), seed.ToMap()); err != nil {
			return fmt.Errorf("failed to insert seed data: %w", err)
		}
	}

	return nil
}

func (c *Client) createTable(ctx context.Context, tableName string) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Enginetype"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SerialNo"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Enginetype"),
				KeyType:       types.KeyTypeHash, // Partition Key
			},
			{
				AttributeName: aws.String("SerialNo"),
				KeyType:       types.KeyTypeRange, // Sort Key
			},
		},
		BillingMode: types.BillingModePayPerRequest,
		TableName:   aws.String(tableName),
	}

	_, err := c.Client.CreateTable(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	waiter := dynamodb.NewTableExistsWaiter(c.Client)
	err = waiter.Wait(ctx,
		&dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
		2*time.Minute,
	)
	if err != nil {
		return fmt.Errorf("failed to wait for table creation: %w", err)
	}

	return nil
}

func (c *Client) CheckConnection(ctx context.Context) error {
	// いずれかのテーブルに対してDescribeTableを実行してチェック
	for _, tableName := range c.tableNames {
		_, err := c.DescribeTable(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})
		if err == nil {
			return nil // 1つでも成功すれば接続OK
		}
	}
	return fmt.Errorf("failed to connect to DynamoDB")
}

func (c *Client) GetItem(ctx context.Context, tableName string, key map[string]interface{}) (map[string]interface{}, error) {
	av, err := attributevalue.MarshalMap(key)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal key: %w", err)
	}

	physicalTableName := c.resolveTableName(tableName)
	result, err := c.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(physicalTableName),
		Key:       av,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	if result.Item == nil {
		return nil, nil
	}

	var item map[string]interface{}
	if err := attributevalue.UnmarshalMap(result.Item, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal item: %w", err)
	}

	return item, nil
}

func (c *Client) PutItem(ctx context.Context, tableName string, item map[string]interface{}) error {

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %w", err)
	}

	physicalTableName := c.resolveTableName(tableName)
	_, err = c.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(physicalTableName),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}

func (c *Client) DeleteItem(ctx context.Context, tableName string, key map[string]interface{}) error {
	av, err := attributevalue.MarshalMap(key)
	if err != nil {
		return fmt.Errorf("failed to marshal key: %w", err)
	}

	physicalTableName := c.resolveTableName(tableName)
	_, err = c.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(physicalTableName),
		Key:       av,
	})
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}

func (c *Client) Query(ctx context.Context, tableName string, keyCondition string, attributeValues map[string]interface{}) ([]map[string]interface{}, error) {
	physicalTableName := c.resolveTableName(tableName)

	// Extract Expression Attribute Names
	expressionAttributeNames := make(map[string]string)
	for k, v := range attributeValues {
		if k[0] == '#' {
			expressionAttributeNames[k] = v.(string)
			delete(attributeValues, k)
		}
	}

	// Marshal remaining values for Expression Attribute Values
	avs, err := attributevalue.MarshalMap(attributeValues)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal attribute values: %w", err)
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(physicalTableName),
		KeyConditionExpression:    aws.String(keyCondition),
		ExpressionAttributeValues: avs,
	}

	// Add ExpressionAttributeNames if any
	if len(expressionAttributeNames) > 0 {
		input.ExpressionAttributeNames = expressionAttributeNames
	}

	result, err := c.Client.Query(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var items []map[string]interface{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal query results: %w", err)
	}

	return items, nil
}

// Scanメソッドにログを追加
func (c *Client) Scan(ctx context.Context, tableName string) ([]map[string]interface{}, error) {
	physicalTableName := c.resolveTableName(tableName)
	fmt.Printf("Scanning table - Logical: %s, Physical: %s\n", tableName, physicalTableName)

	input := &dynamodb.ScanInput{
		TableName: aws.String(physicalTableName),
	}
	result, err := c.Client.Scan(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute scan: %w", err)
	}

	fmt.Printf("Scan result - Items count: %d\n", len(result.Items))

	var items []map[string]interface{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal scan results: %w", err)
	}

	return items, nil
}
