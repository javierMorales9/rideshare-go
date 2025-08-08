package model

import "time"

type VehicleReservation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	VehicleID uint     `json:"vehicle_id"`
	Vehicle   *Vehicle `gorm:"foreignKey:VehicleID" json:"-"`

	TripRequestID uint         `json:"trip_request_id"`
	TripRequest   *TripRequest `gorm:"foreignKey:TripRequestID" json:"-"`

	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}
