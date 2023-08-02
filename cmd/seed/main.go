package main

import (
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/configs/parameter"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	cfg := configs.LoadConfigs(constants.EnvConfPath)
	paramCfg := parameter.LoadParameterConfigs(constants.ParamConfPath)
	pgAdapter := postgres.NewPgAdapter(&paramCfg.PostgresConf)
	pgInstance := pgAdapter.ConnectInstance(cfg)


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
		Id:       "01",
		Username: "user01",
		Password: "password01",
		Limit:    1,
	}
	if err := db.Create(u).Error; err != nil {
		log.Fatalf("Failed to seed user data: %v", err)
	}
}

func seedRoom(db *gorm.DB) {
	r := &models.Room{
		Id:          "01",
		Name:        "Room 01",
		Description: "Room 01",
		Category:    "Room 01",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   "01",
	}
	if err := db.Create(r).Error; err != nil {
		log.Fatalf("Failed to seed room data: %v", err)
	}
}
