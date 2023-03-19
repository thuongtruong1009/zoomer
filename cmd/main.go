package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"zoomer/configs"
	"zoomer/db"
	"zoomer/internal/server"
)

func main() {
	cfg := configs.NewConfig()

	instance := db.GetPostgresInstance(cfg, true)
	defer instance.Close()

	db.SetConnectionPool(instance, cfg)

	s := server.NewServer(cfg, instance, logrus.New(), nil)
	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	log.Println("Starting api server")
}
