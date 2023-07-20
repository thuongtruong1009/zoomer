package postgres

import (
	"log"
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	// "github.com/sirupsen/logrus"

	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/internal/models"
	// "zoomer/migrations"
)

type postgresStruct struct {
	db *gorm.DB
}

func NewPgAdapter() PgAdapter {
	db := &gorm.DB{}
	return &postgresStruct{db: db}
}

func (pg *postgresStruct) GetInstance(cfg *configs.Configuration) *gorm.DB {
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	pg.SetConnectionPool(db, cfg)

	go func(dsn string) {
		var intervals = []time.Duration{3 * time.Second, 3 * time.Second, 15 * time.Second, 30 * time.Second, 60 * time.Second, 60 * time.Second}

		for {
			time.Sleep(60 * time.Second)
			sqlDB, err := db.DB()
			if err != nil {
				log.Fatal("Error when ping to database: ", err)
			}
			defer sqlDB.Close()

			pong := sqlDB.Ping()
			if pong != nil {
			L:
				for i := 0; i < len(intervals); i++ {
					pong2 := pg.RetryHandler(3, func() (bool, error) {
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

func (pg *postgresStruct) RetryHandler(n int, f func() (bool, error)) error {
	ok, er := f()
	if ok && er == nil {
		return nil
	}
	if n-1 > 0 {
		return pg.RetryHandler(n-1, f)
	}
	return er
}

func (pg *postgresStruct) SetConnectionPool(d *gorm.DB, cfg *configs.Configuration) {
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

func (pg *postgresStruct) Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx := pg.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			log.Println("recover from transaction: ", p)
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Print("rollback from transaction: ", err)
			tx.Rollback()
			panic(err)
		} else {
			err = tx.Commit().Error
		}
	}()

	data, err = txFunc(tx)
	return data, err
}
