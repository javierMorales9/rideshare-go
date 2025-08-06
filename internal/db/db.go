package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return err
}

// Llama esto en main
func MustConnect(dsn string) {
	if err := Connect(dsn); err != nil {
		log.Fatalf("cannot connect db: %v", err)
	}
}
