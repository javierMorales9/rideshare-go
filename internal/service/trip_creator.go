package service

import (
	"errors"
	"math/rand"

	"github.com/javierMorales9/rideshare-go/internal/domain/model"
	"gorm.io/gorm"
)

// TripCreator replica la lógica de Rails (driver aleatorio).
type TripCreator struct {
	TripRequestID uint
}

// CreateTrip crea un Trip y lo guarda. Devuelve el Trip o error.
func (tc TripCreator) CreateTrip(tx *gorm.DB) (*model.Trip, error) {
	if tx == nil {
		return nil, errors.New("transaction db is required")
	}

	var req model.TripRequest
	if err := tx.
		Preload("Rider").
		First(&req, tc.TripRequestID).Error; err != nil {
		return nil, err
	}

	driver, err := pickRandomDriver(tx)
	if err != nil {
		return nil, err
	}

	trip := model.Trip{
		TripRequestID: req.ID,
		DriverID:      driver.ID,
	}

	if err := tx.Create(&trip).Error; err != nil {
		return nil, err
	}
	return &trip, nil
}

func pickRandomDriver(tx *gorm.DB) (*model.User, error) {
	var drivers []model.User
	if err := tx.
		Where("type = ?", model.UserTypeDriver).
		Find(&drivers).Error; err != nil {
		return nil, err
	}
	if len(drivers) == 0 {
		return nil, errors.New("no drivers available")
	}
	return &drivers[rand.Intn(len(drivers))], nil
}
