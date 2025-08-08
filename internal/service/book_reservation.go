package service

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/javierMorales9/rideshare-go/internal/db"
	"github.com/javierMorales9/rideshare-go/internal/domain/model"
)

type BookReservationParams struct {
	VehicleID       uint
	RiderID         uint
	StartLocationID uint
	EndLocationID   uint
	StartsAt        time.Time
	EndsAt          time.Time
}

var (
	ErrVehicleNotFound = errors.New("vehicle not found")
	ErrRiderNotFound   = errors.New("rider not found")
	ErrLocationMissing = errors.New("start or end location not found")
	ErrInvalidWindow   = errors.New("ends_at must be after starts_at")
)

func BookReservation(p BookReservationParams) (*model.TripRequest, *model.VehicleReservation, error) {
	if !p.EndsAt.After(p.StartsAt) {
		return nil, nil, ErrInvalidWindow
	}

	var (
		v  model.Vehicle
		r  model.User
		sl model.Location
		el model.Location

		tr model.TripRequest
		vr model.VehicleReservation
	)

	// Cargamos entidades base
	if err := db.DB.First(&v, p.VehicleID).Error; err != nil {
		return nil, nil, ErrVehicleNotFound
	}
	if err := db.DB.First(&r, p.RiderID).Error; err != nil {
		return nil, nil, ErrRiderNotFound
	}
	if err := db.DB.First(&sl, p.StartLocationID).Error; err != nil {
		return nil, nil, ErrLocationMissing
	}
	if err := db.DB.First(&el, p.EndLocationID).Error; err != nil {
		return nil, nil, ErrLocationMissing
	}

	// Transacción
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		tr = model.TripRequest{
			RiderID:         r.ID,
			StartLocationID: sl.ID,
			EndLocationID:   el.ID,
		}
		if err := tx.Create(&tr).Error; err != nil {
			return err
		}

		vr = model.VehicleReservation{
			VehicleID:     v.ID,
			TripRequestID: tr.ID,
			StartsAt:      p.StartsAt,
			EndsAt:        p.EndsAt,
		}
		if err := tx.Create(&vr).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}
	return &tr, &vr, nil
}
