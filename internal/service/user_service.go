package service

import (
	"errors"

	"github.com/javierMorales9/rideshare-go/internal/db"
	"github.com/javierMorales9/rideshare-go/internal/domain/model"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

func RegisterUser(u *model.User, rawPassword string) error {
	if err := u.SetPassword(rawPassword); err != nil {
		return err
	}
	return db.DB.Create(u).Error
}

func Authenticate(email, rawPassword string) (*model.User, error) {
	var u model.User
	if err := db.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, ErrInvalidCredentials
	}
	if !u.CheckPassword(rawPassword) {
		return nil, ErrInvalidCredentials
	}
	return &u, nil
}
