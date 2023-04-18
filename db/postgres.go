package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"

	"zoomer/configs"
	"zoomer/internal/models"
)

func GetPostgresInstance(cfg *configs.Configuration, migrate bool) *gorm.DB {
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	if migrate {
		db.AutoMigrate(&models.User{}, &models.Room{})
		if err != nil {
			panic("Error when run migrations")
		}
		log.Println("Migration successful")
	}
	return db
}

func SetConnectionPool(d *gorm.DB, cfg *configs.Configuration) {
	maxOpen := cfg.MaxOpenConnection
	maxLifetime := cfg.MaxLifetimeConnection
	maxIdleConn := cfg.MaxIdleConnection
	maxIdleTime := cfg.MaxIdleTimeConnection

	db, err := d.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(int(maxOpen))
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	db.SetMaxIdleConns(int(maxIdleConn))
	db.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Second)
}
