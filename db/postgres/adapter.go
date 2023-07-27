package postgres

import (
	"log"
	"time"
	"errors"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type postgresStruct struct {
	db *gorm.DB
}

func NewPgAdapter() PgAdapter {
	return &postgresStruct{db: &gorm.DB{}}
}

func (pg *postgresStruct) getInstance(uri string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func (pg *postgresStruct) ConnectInstance(cfg *configs.Configuration) *gorm.DB {
	dsn := cfg.DatabaseConnectionURL
	db := pg.getInstance(dsn)

	pg.setConnectionPool(db, cfg)

	go func(dsn string) {
		var intervals = []time.Duration{3 * time.Second, 3 * time.Second, 15 * time.Second, 30 * time.Second, 60 * time.Second, 60 * time.Second}

		for {
			time.Sleep(60 * time.Second)
			sqlDB, err := db.DB()
			if err != nil {
				log.Fatal("Error when ping to database: ", err)
			}

			pong := sqlDB.Ping()
			if pong != nil {
			L:
				for i := 0; i < len(intervals); i++ {
					pong2 := pg.retryHandler(3, func() (bool, error) {
						db = pg.getInstance(dsn)

						if db == nil {
							return false, errors.New("Error when reconnect to database")
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

	if cfg.AutoMigrate {
		if err := db.AutoMigrate(&models.User{}, &models.Room{}); err != nil {
			panic("Error when run auto migrations")
		}
		log.Println("Auto migration successful")

		// sqlString := fmt.Sprintf("CREATE TABLE IF NOT EXISTS users(%s);", db.Migrator().CurrentDatabase().Migrator().GetTable(&User{}))
		// fmt.Println(sqlString)
	}

	return db
}

func (pg *postgresStruct) retryHandler(n int, f func() (bool, error)) error {
	ok, er := f()
	if ok && er == nil {
		return nil
	}
	if n-1 > 0 {
		return pg.retryHandler(n-1, f)
	}
	return er
}

func (pg *postgresStruct) setConnectionPool(d *gorm.DB, cfg *configs.Configuration) {
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
