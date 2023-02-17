package main

import (
	"log"
	"zoomer/configs"
	"zoomer/db"
	"zoomer/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting api server")

	cfg := configs.NewConfig()

	db := db.GetPostgresInstance(cfg, true)

	s := server.NewServer(cfg, db, logrus.New(), nil)

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}