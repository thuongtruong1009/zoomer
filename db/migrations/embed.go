package migrations

import (
	"errors"
	"time"
	"log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/thuongtruong1009/zoomer/configs"
)

const (
	_defaultAttempts = 20
	_defaultTimeout = time.Second
)

func init() {
	databaseURL := configs.LookupEnv("PG_URI") + "?sslmode=disable"

	var (
		attempts = _defaultAttempts
		err error
		m *migrate.Migrate
	)

	log.Println("Migration - start")

	for attempts > 0 {
		m, err = migrate.New("file://db/migrations", databaseURL)
		if err == nil {
			break
		}

		log.Println("Migration - trying to reconnect, attempts left:", attempts, "err:", err.Error())
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalln("Migration - error connect to db", "err: ", err.Error())
	}

	err = m.Up()
	defer m.Close()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalln("Migration - error migrate up", "err: ", err.Error())
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("Migration - up no change")
	}

	err = m.Down()
	if err != nil {
		log.Fatalln("Migration - error migrate down", "err: ", err.Error())
	}

	dbversion, dirty, err := m.Version()
	if err != nil {
		log.Fatalln("auto migration - error get version", "err", err.Error())
	}

	log.Printf("Currentdb version: %d, dirty: %t", dbversion, dirty)

	log.Println("Migration - success")
}
