package service

import (
	"errors"
	"github.com/javierMorales9/rideshare-go/internal/db"
	"github.com/javierMorales9/rideshare-go/internal/domain/model"
	"math/rand"
)

// TripCreator replica la lógica de Rails (driver aleatorio).
type TripCreator struct {
	TripRequestID uint
}

// CreateTrip crea un Trip y lo guarda. Devuelve el Trip o error.
func (tc TripCreator) CreateTrip() (*model.Trip, error) {
	var req model.TripRequest
	if err := db.DB.
		Preload("Rider").
		First(&req, tc.TripRequestID).Error; err != nil {
		return nil, err
	}

	driver, err := pickRandomDriver()
	if err != nil {
		return nil, err
	}

	trip := model.Trip{
		TripRequestID: req.ID,
		DriverID:      driver.ID,
	}

	if err := db.DB.Create(&trip).Error; err != nil {
		return nil, err
	}
	return &trip, nil
}

func pickRandomDriver() (*model.User, error) {
	var drivers []model.User
	if err := db.DB.
		Where("type = ?", model.UserTypeDriver).
		Find(&drivers).Error; err != nil {
		return nil, err
	}
	if len(drivers) == 0 {
		return nil, errors.New("no drivers available")
	}
	return &drivers[rand.Intn(len(drivers))], nil
}
