package domain

import (
	"context"
	"time"
)

// Forklift represents the main forklift entity
type Forklift struct {
	Enginetype  string    `json:"enginetype" validate:"required"`
	Maker       string    `json:"maker" validate:"required"`
	Model       string    `json:"model" validate:"required"`
	SerialNo    string    `json:"serialNo" validate:"required"`
	Height      float64   `json:"height" validate:"required"`
	Ct          string    `json:"ct" validate:"required"`
	Attachment  string    `json:"attachment" validate:"required"`
	Year        int       `json:"year" validate:"required"`
	HourMeter   float64   `json:"hourMeter" validate:"required"`
	Application string    `json:"application" validate:"required"`
	Fob         int       `json:"fob" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"` // デフォルト値を含む
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"` // デフォルト値を含む
}

// ForkliftRequest represents the API request for creating or updating a forklift
type ForkliftRequest struct {
	Enginetype  string  `json:"enginetype" validate:"required"`
	Maker       string  `json:"maker" validate:"required"`
	Model       string  `json:"model" validate:"required"`
	SerialNo    string  `json:"serialNo" validate:"required"`
	Height      float64 `json:"height" validate:"required"`
	Ct          string  `json:"ct" validate:"required"`
	Attachment  string  `json:"attachment" validate:"required"`
	Year        int     `json:"year" validate:"required"`
	HourMeter   float64 `json:"hourMeter" validate:"required"`
	Application string  `json:"application" validate:"required"`
	Fob         float64 `json:"fob" validate:"required"`
}

// ForkliftResponse represents the API response for a forklift
type ForkliftResponse struct {
	Forklift Forklift `json:"forklift"`
}

// Default values for CreatedAt and UpdatedAt
func NewForklift() *Forklift {
	now := time.Now()
	return &Forklift{
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// ForkliftRepository defines the methods for interacting with the forklift data
type ForkliftRepository interface {
	GetByEngineType(ctx context.Context, engineType string) ([]Forklift, error)
	GetByEngineTypeModelSerial(ctx context.Context, engineType, model, serial string) (*Forklift, error)
}
