package migrations

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
)

//go:embed sql/*
var fs embed.FS

type MLog struct {
	log *logrus.Entry
}

func (MLog) Verbose() bool {
	return false
}

func (m *MLog) Printf(format string, v ...interface{}) {
	m.log.Infof(format, v...)
}

func (m *MLog) Errorf(format string, v ...interface{}) {
	m.log.Errorf(format, v...)
}

func RunAutoMigrate(db *sql.DB, log *logrus.Logger) {
	d, err := iofs.New(fs, "sql")
	if err != nil {
		log.Fatalln("auto migration - new iofs", "err", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln("auto migration - new postgres driver", "err", err.Error())
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		log.Fatalln("auto migration - new migrate", "err", err.Error())
	}

	defer m.Close()
	m.Log = &MLog{
		log: logrus.NewEntry(log),
	}
	// 	err = m.Down()
	// if err != nil && err != migrate.ErrNoChange {
	//     log.Fatalln("auto migration - error migrate down", "err", err.Error())
	// }

	// err = m.Up()
	// if err != nil && err != migrate.ErrNoChange {
	//     log.Fatalln("auto migration - error migrate up", "err", err.Error())
	// }
	// 	log.Info("step 6")

	// 	dbversion, dirty, err := m.Version()
	// 	if err != nil {
	// 		log.Fatalln("auto migration - error get version", "err", err.Error())
	// 	}

	// log.Infof("Currentdb version: %d, dirty: %t", dbversion, dirty)
}

// func demo1() {
// 	//_ "github.com/golang-migrate/migrate/v4/source/file"
// 	driver, err := postgres.WithInstance(db, &postgres.Config{})
// 	if err != nil {
// 		log.Fatalf("cannot create postgres driver %v", err)
// 	}

// 	m, err := migrate.NewWithDatabaseInstance("file://migrations/", "postgres", driver)
// 	if err != nil {
// 		log.Fatalf("cannot create migrations instance %v", err)
// 	}
// 	if err = m.Up(); err != nil && err.Error() != "no change" {
// 		log.Fatalf("error running migrations %v", err)
// 	}
// }
