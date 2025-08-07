package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Location struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Address   string   `gorm:"uniqueIndex;not null" json:"address"`
	State     string   `gorm:"size:2;not null"     json:"state"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

// --- hooks ------------------------------------------------------------

// BeforeCreate / BeforeUpdate: simple validación state == 2 letras
func (l *Location) BeforeSave(tx *gorm.DB) error {
	if len(l.State) != 2 {
		return fmt.Errorf("state must be 2 characters")
	}
	return nil
}

// AfterSave: STUB de geocodificación
func (l *Location) AfterSave(tx *gorm.DB) error {
	// TODO: llamar a servicio externo sólo si lat/lon son nil
	return nil
}
