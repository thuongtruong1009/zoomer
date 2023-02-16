package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"zoomer/configs"
	"zoomer/internal/models"
)

func GetPostgresInstance(cfg *config.Configuration, migrate bool) {
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database ", err)
	}

	if migrate {
		db.AutoMigrate(&models.User{}, &models.Todo{})
		if err != nil {
			panic("Error when run migrations")
		}
	}
	return db
}
