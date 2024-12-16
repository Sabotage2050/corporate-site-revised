// infra/repository/nosql/forklift/forklift_repository.go
package forklift

import (
	"context"
	"corporation-site/domain"
	"corporation-site/infra/db/nosql"
	"fmt"
	"time"
)

type forkliftRepository struct {
	client nosql.NoSQLClient
}

// NewForkliftRepository creates a new instance of ForkliftRepository
func NewForkliftRepository(client nosql.NoSQLClient) (domain.ForkliftRepository, error) {
	result := &forkliftRepository{
		client: client,
	}
	return result, nil
}

func (r *forkliftRepository) GetByEngineType(ctx context.Context, engineType string) ([]domain.Forklift, error) {
	results, err := r.client.Scan(ctx, "Forklift")
	if err != nil {
		return nil, fmt.Errorf("failed to scan forklifts: %w", err)
	}

	var filteredResults []map[string]interface{}
	for _, item := range results {
		if item["Enginetype"].(string) == engineType {
			filteredResults = append(filteredResults, item)
		}
	}

	return convertToForkliftList(filteredResults)
}

func (r *forkliftRepository) GetByEngineTypeModelSerial(ctx context.Context, engineType, model, serial string) (*domain.Forklift, error) {
	results, err := r.client.Scan(ctx, "Forklift")
	if err != nil {
		return nil, fmt.Errorf("failed to scan forklifts: %w", err)
	}

	for _, item := range results {
		if item["Enginetype"].(string) == engineType && item["Model"].(string) == model && item["SerialNo"].(string) == serial {
			return mapToForklift(item)
		}
	}

	return nil, fmt.Errorf("forklift not found")
}

func convertToForkliftList(items []map[string]interface{}) ([]domain.Forklift, error) {
	forklifts := make([]domain.Forklift, 0, len(items))
	for _, item := range items {
		forklift, err := mapToForklift(item)
		if err != nil {
			return nil, err
		}
		forklifts = append(forklifts, *forklift)
	}
	return forklifts, nil
}

func mapToForklift(item map[string]interface{}) (*domain.Forklift, error) {
	createdAt, err := parseTime(item["CreatedAt"])
	if err != nil {
		return nil, fmt.Errorf("invalid CreatedAt: %w", err)
	}

	updatedAt, err := parseTime(item["UpdatedAt"])
	if err != nil {
		return nil, fmt.Errorf("invalid UpdatedAt: %w", err)
	}

	return &domain.Forklift{
		Enginetype:  getStringValue(item["Enginetype"]),
		Maker:       getStringValue(item["Maker"]),
		Model:       getStringValue(item["Model"]),
		SerialNo:    getStringValue(item["SerialNo"]),
		Height:      getFloatValue(item["Height"]),
		Ct:          getStringValue(item["Ct"]),
		Attachment:  getStringValue(item["Attachment"]),
		Year:        getIntValue(item["Year"]),
		HourMeter:   getFloatValue(item["HourMeter"]),
		Application: getStringValue(item["Application"]),
		Fob:         getIntValue(item["Fob"]),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func getStringValue(value interface{}) string {
	if value == nil {
		return ""
	}
	return value.(string)
}

func getFloatValue(value interface{}) float64 {
	if value == nil {
		return 0.0
	}
	return value.(float64)
}

func getIntValue(value interface{}) int {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case int:
		return v
	case float64:
		return int(v)
	default:
		panic(fmt.Sprintf("unsupported type for int conversion: %T", value))
	}
}

func parseTime(v interface{}) (time.Time, error) {
	switch t := v.(type) {
	case string:
		return time.Parse(time.RFC3339, t)
	case time.Time:
		return t, nil
	default:
		return time.Time{}, fmt.Errorf("unsupported time format: %v", v)
	}
}
