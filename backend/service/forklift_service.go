package service

import (
	"context"
	"corporation-site/domain"
)

type forkliftService struct {
	repo domain.ForkliftRepository
}

func NewForkliftService(repo domain.ForkliftRepository) ForkliftService {
	return &forkliftService{repo: repo}
}

func (s *forkliftService) GetForkliftsByType(ctx context.Context, engineType string) ([]domain.Forklift, error) {
	return s.repo.GetByEngineType(ctx, engineType)
}

func (s *forkliftService) GetForkliftByEngineTypeModelSerial(ctx context.Context, engineType, model, serial string) (*domain.Forklift, error) {
	return s.repo.GetByEngineTypeModelSerial(ctx, engineType, model, serial)
}
