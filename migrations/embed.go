package migrations

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
)

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

func RunAutoMigration(db *sql.DB, log *logrus.Logger) {
	d, err := iofs.New(fs, "sql")
	if err != nil {
		log.Fatalln("auto migration - new iofs", "err", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln("auto migration - new postgres driver", "err", err.Error())
	}

	// m, err := migrate.NewWithSourceInstance("iofs", iofs.New(fs, "migrations"), "postgres", driver)
	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		log.Fatalln("auto migration - new migrate", "err", err.Error())
	}

	defer m.Close()
	m.Log = &MLog{
		// log: logger.Log()
		log: logrus.NewEntry(log),
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalln("auto migration - error migrate up", "err", err.Error())
	}
	dbversion, dirty, err := m.Version()
	if err != nil {
		log.Fatalln("auto migration - error get version", "err", err.Error())
	}

	log.Infof("auto migration - db version: %d, dirty: %t", dbversion, dirty)
}
