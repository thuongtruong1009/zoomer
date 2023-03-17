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

	db := db.GetPostgresInstance(cfg, true)
	defer db.Close()

	configs.SetConnectionPool(db, cfg)
	
	s := server.NewServer(cfg, db, logrus.New(), nil)
	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	log.Println("Starting api server")
}
