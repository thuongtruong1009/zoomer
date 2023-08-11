package postgres

import (
	"errors"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type postgresStruct struct {
	db       *gorm.DB
	paramCfg *parameter.PostgresConf
}

func NewPgAdapter(paramCfg *parameter.PostgresConf) PgAdapter {
	return &postgresStruct{
		db:       &gorm.DB{},
		paramCfg: paramCfg,
	}
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

	pg.setConnectionPool(db)

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

	if pg.paramCfg.AutoMigrate {
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

func (pg *postgresStruct) setConnectionPool(d *gorm.DB) {
	db, err := d.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(pg.paramCfg.MaxOpenConnection)
	db.SetConnMaxLifetime(pg.paramCfg.MaxLifetimeConnection * time.Second)
	db.SetMaxIdleConns(pg.paramCfg.MaxIdleConnection)
	db.SetConnMaxIdleTime(pg.paramCfg.MaxIdleTimeConnection * time.Second)
}
