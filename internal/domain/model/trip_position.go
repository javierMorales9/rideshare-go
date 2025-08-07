package model

import "time"

type TripPosition struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	TripID    uint
	Trip      *Trip `gorm:"foreignKey:TripID"`

	Latitude  float64
	Longitude float64
}
