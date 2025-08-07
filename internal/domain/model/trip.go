package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Trip struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TripRequestID uint `json:"trip_request_id"`
	TripRequest   *TripRequest

	DriverID uint  `json:"driver_id"`
	Driver   *User `gorm:"foreignKey:DriverID"`

	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Rating      *int       `json:"rating,omitempty"` // 1–5

	TripPositions []TripPosition `json:"-"`
}

// --- gorm hook: rating sólo si CompletedAt está presente
func (t *Trip) BeforeSave(tx *gorm.DB) error {
	if t.Rating != nil && t.CompletedAt == nil {
		return errors.New("trip must be completed before adding rating")
	}
	return nil
}
