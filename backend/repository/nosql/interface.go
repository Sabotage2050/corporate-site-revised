package nosql

import (
	"context"
	api "corporation-site/infra/api/forklift"
	"errors"
)

var ErrForkliftNotFound = errors.New("forklift not found")

type ForkliftRepository interface {
	GetAll(ctx context.Context) ([]api.Forklift, error)
	GetByID(ctx context.Context, id string) (*api.Forklift, error)
	GetByType(ctx context.Context, forkliftType string) ([]api.Forklift, error)
}
