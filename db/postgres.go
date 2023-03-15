package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"zoomer/configs"
	"zoomer/internal/models"
)

func GetPostgresInstance(cfg *configs.Configuration, migrate bool) *gorm.DB {
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if migrate {
		db.AutoMigrate(&models.User{}, &models.Room{})
		if err != nil {
			panic("Error when run migrations")
		}
	}
	return db
}
