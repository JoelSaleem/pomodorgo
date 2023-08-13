package db

import (
	"fmt"

	"github.com/JoelSaleem/pomodorgo/internal/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IRepository interface {
	// ConnectToDB(dbPath string) (*gorm.DB, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(dbPath string) *Repository {
	fmt.Println("\n\n connecting to db \n\n", dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Todo: maybe don't do in prod
	db.AutoMigrate(&models.Task{})

	return &Repository{db: db}
}
