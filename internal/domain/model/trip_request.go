package model

import "time"

type TripRequest struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	RiderID uint  `json:"rider_id"`
	Rider   *User `gorm:"foreignKey:RiderID"`

	StartLocationID uint      `json:"start_location_id"`
	StartLocation   *Location `gorm:"foreignKey:StartLocationID"`

	EndLocationID uint      `json:"end_location_id"`
	EndLocation   *Location `gorm:"foreignKey:EndLocationID"`

	Trip *Trip
}
