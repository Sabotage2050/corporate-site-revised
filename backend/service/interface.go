// internal/service/forklift_service.go
package service

import (
	"context"
	"corporation-site/domain"
)

type ForkliftService interface {
	GetForkliftsByType(ctx context.Context, engineType string) ([]domain.Forklift, error)
	GetForkliftByEngineTypeModelSerial(ctx context.Context, engineType, model, serial string) (*domain.Forklift, error)
}
