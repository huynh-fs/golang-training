package database

import (
	"log"

	"github.com/huynh-fs/gin-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(model.Todo{}, model.User{}, model.RefreshToken{})

	log.Println("Connected to database successfully")
	return db, nil
}