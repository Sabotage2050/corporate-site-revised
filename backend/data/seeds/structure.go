// backend/data/seeds/structure.go
package seeds

import (
	"time"
)

// データ定義の共通インターフェース
type Seed interface {
	TableName() string
	ToMap() map[string]interface{}
}

type ForkliftSeed struct {
	Enginetype  string    `csv:"Enginetype"`
	Maker       string    `csv:"Maker"`
	Model       string    `csv:"Model"`
	SerialNo    string    `csv:"SerialNo"`
	Height      float64   `csv:"Height"`
	CT          string    `csv:"CT"`
	Attachment  string    `csv:"Attachment"`
	Year        int       `csv:"Year"`
	HourMeter   float64   `csv:"HourMeter"`
	Application string    `csv:"Application"`
	Fob         int       `csv:"Fob"`
	CreatedAt   time.Time `csv:"CreatedAt"`
	UpdatedAt   time.Time `csv:"UpdatedAt"`
}

func (f ForkliftSeed) TableName() string {
	return "Forklift"
}

func (f ForkliftSeed) ToMap() map[string]interface{} {
	item := map[string]interface{}{
		"Enginetype":  f.Enginetype,
		"SerialNo":    f.SerialNo,
		"Maker":       f.Maker,
		"Model":       f.Model,
		"Height":      f.Height,
		"CT":          f.CT,
		"Attachment":  f.Attachment,
		"Year":        f.Year,
		"HourMeter":   f.HourMeter,
		"Application": f.Application,
		"Fob":         f.Fob,
		"CreatedAt":   f.CreatedAt,
		"UpdatedAt":   f.UpdatedAt,
	}

	return item
}
