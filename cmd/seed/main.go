package main

import (
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"log"
)

func main() {
	config := configs.NewConfig()

	db, err := configs.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Seeding is done")


}
