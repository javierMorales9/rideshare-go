package model

import "time"

const (
	VehicleStatusDraft     = "draft"
	VehicleStatusPublished = "published"
)

type Vehicle struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name   string `gorm:"size:255;uniqueIndex;not null" json:"name"`
	Status string `gorm:"size:20;not null"             json:"status"`

	Reservations []VehicleReservation `json:"-"`
}

func (v Vehicle) IsPublished() bool { return v.Status == VehicleStatusPublished }
