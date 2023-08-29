package postgres

import (
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"gorm.io/gorm/logger"
)

type postgresStruct struct {
	db       *gorm.DB
	cfg 	*configs.Configuration
	paramCfg *parameter.PostgresConf
}

func NewPgAdapter(cfg *configs.Configuration, paramCfg *parameter.PostgresConf) PgAdapter {
	return &postgresStruct{
		db:       &gorm.DB{},
		cfg: cfg,
		paramCfg: paramCfg,
	}
}

func (pg *postgresStruct) getInstance(uri string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		exceptions.Fatal(constants.ErrorPostgresConnectionFailed, err)
	}

	return db
}

func (pg *postgresStruct) ConnectInstance() *gorm.DB {
	dsn := pg.cfg.DatabaseConnectionURL
	db := pg.getInstance(dsn)

	pg.setConnectionPool(db)

	go func(dsn string) {
		var intervals = []time.Duration{3 * time.Second, 3 * time.Second, 15 * time.Second, 30 * time.Second, 60 * time.Second, 60 * time.Second}

		for {
			time.Sleep(helpers.DurationSecond(pg.paramCfg.RetryDelay))
			sqlDB, err := db.DB()
			if err != nil {
				exceptions.Fatal(constants.ErrorPostgresGetResponse, err)
			}

			pong := sqlDB.Ping()
			if pong != nil {
			L:
				for i := 0; i < len(intervals); i++ {
					pong2 := pg.retryHandler(pg.paramCfg.RetryAttempts, func() (bool, error) {
						db = pg.getInstance(dsn)

						if db == nil {
							return false, constants.ErrorPostgresReconnect
						}

						exceptions.SystemLog(constants.PostgresConnectionSuccessful)
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
			exceptions.Panic(constants.ErrorPostgresAutoMigration, err)
		}
		exceptions.SystemLog(constants.PostgresAutoMigrationSuccessful)

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
	db.SetMaxIdleConns(pg.paramCfg.MaxIdleConnection)
	db.SetMaxOpenConns(pg.paramCfg.MaxOpenConnection)
	db.SetConnMaxLifetime(helpers.DurationSecond(pg.paramCfg.MaxLifetimeConnection))
	db.SetConnMaxIdleTime(helpers.DurationSecond(pg.paramCfg.MaxIdleTimeConnection))
}
