package db

import (
	"fmt"

	"github.com/JoelSaleem/pomodorgo/internal/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(dbPath string) *Repository {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Repository{db: db}
}

func (r *Repository) connectToDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	db.AutoMigrate(&models.Task{})

	return db, nil
}
