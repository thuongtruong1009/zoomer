package main

import (
	"log"
	"github.com/thuongtruong1009/zoomer/configs"
	database "github.com/thuongtruong1009/zoomer/db"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

func main() {
	config := configs.NewConfig()

	db := database.GetPostgresInstance(config)

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	u := &models.User{
		Id: "1",
		Username: "Tom",
		Password: "123456",
		Limit: 1,
	}
	if err := db.Create(u).Error; err != nil {
		log.Fatalf("Failed to seed user data: %v", err)
	}

	log.Println("::: Seeding is done")


}
