package service

import (
	"context"
	"errors"

	"github.com/javierMorales9/rideshare-go/internal/db"
	"github.com/javierMorales9/rideshare-go/internal/domain/model"
	"gorm.io/gorm"
)

// CreateOrGetLocation busca por dirección (case-insensitive); si no existe la crea.
func CreateOrGetLocation(ctx context.Context, addr, state string) (*model.Location, error) {
	var loc model.Location
	err := db.DB.WithContext(ctx).
		Where("lower(address) = lower(?)", addr).
		First(&loc).Error
	if err == nil {
		return &loc, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	loc = model.Location{Address: addr, State: state}
	return &loc, db.DB.Create(&loc).Error
}
