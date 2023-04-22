package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
	// "github.com/sirupsen/logrus"

	"zoomer/configs"
	"zoomer/internal/models"
	// "zoomer/migrations"
)

func GetPostgresInstance(cfg *configs.Configuration) *gorm.DB {
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	SetConnectionPool(db, cfg)

	// auto-connect，ping per 60s, re-connect on fail or error with intervels 3s, 3s, 15s, 30s, 60s, 60s ...
	go func(dsn string) {
		var intervals = []time.Duration{3 * time.Second, 3 * time.Second, 15 * time.Second, 30 * time.Second, 60 * time.Second, 60 * time.Second}

		for {
			time.Sleep(60 * time.Second)
			getDb, err := db.DB()
			if err != nil {
				log.Println("Error when ping to database: ", err)
			}

			pong := getDb.Ping()
			if pong != nil {
			L:
				for i := 0; i < len(intervals); i++ {
					pong2 := RetryHandler(3, func() (bool, error) {
						var err error
						db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

						if err != nil {
							log.Println("Error when reconnect to database: ", err)
							return false, err
						}

						log.Println("Reconnect to database successful")
						return true, nil
					})
					if pong2 != nil {
						time.Sleep(intervals[i])
						if i == len(intervals)-1 {
							i--
						}
						continue
					}
					break L
				}
			}
		}
	}(dsn)

	if cfg.AutoMigrate == true {
		if err := db.AutoMigrate(&models.User{}, &models.Room{}); err != nil {
			panic("Error when run migrations")
		}
		log.Println("Migration successful")

		// sqlDB, err := db.DB()
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println("Step here")
		// migrations.RunAutoMigrate(sqlDB, logrus.New())
	}

	return db
}

func RetryHandler(n int, f func() (bool, error)) error {
	ok, er := f()
	if ok && er == nil {
		return nil
	}
	if n-1 > 0 {
		return RetryHandler(n-1, f)
	}
	return er
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
