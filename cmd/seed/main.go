package main

import (
	"log"
	"gorm.io/gorm"
	"time"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

func main() {
	config := configs.NewConfig()

	pg := postgres.NewPgAdapter()
	db := pg.ConnectInstance(config)

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	helpers.Parallelize(
		func() {
			seedUser(db)
		},
		func() {
			seedRoom(db)
		},
	)

	log.Println("::: Seeding is done")
}

func seedUser(db *gorm.DB) {
	u := &models.User{
		Id: "1",
		Username: "Tom",
		Password: "123456",
		Limit: 1,
	}
	if err := db.Create(u).Error; err != nil {
		log.Fatalf("Failed to seed user data: %v", err)
	}
}

func seedRoom(db *gorm.DB) {
	r := &models.Room{
		Id: "1",
		Name: "Room 1",
		Description: "Room 1",
		Category: "Room 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "1",
		}
	if err := db.Create(r).Error; err != nil {
		log.Fatalf("Failed to seed room data: %v", err)
	}
}
