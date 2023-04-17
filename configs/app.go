package configs

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

type Configuration struct {
	HttpPort              string `env:"HTTP_PORT" envDefault:"8080"`
	WsPort                string `env:"WS_PORT" envDefault:"8081"`
	HashSalt              string `env:"HASH_SALT,required"`
	SigningKey            string `env:"SIGNING_KEY,required"`
	TokenTTL              int64  `env:"TOKEN_TTL,required"`
	JwtSecret             string `env:"JWT_SECRET,required"`
	AutoMigrate		   bool   `env:"AUTO_MIGRATE" envDefault:"true"`

	DatabaseConnectionURL string `env:"PG_URI,required"`
	MaxOpenConnection     int    `env:"PG_MAX_OPEN_CONN" envDefault:"20"`
	MaxIdleConnection     int    `env:"PG_MAX_IDLE_CONN" envDefault:"20"`
	MaxLifetimeConnection int    `env:"PG_MAX_LIFETIME_CONN" envDefault:"20"`
	MaxIdleTimeConnection int    `env:"PG_MAX_IDLE_TIME_CONN" envDefault:"20"`

	RedisURI              string `env:"REDIS_URI,required"`
	RedisPassword         string `env:"REDIS_PASSWORD,required"`

	MinIOAccessKey        string `env:"MINIO_ACCESS,required"`
	MinIOSecretKey        string `env:"MINIO_SECRET,required"`
	MinIOEndpoint         string `env:"MINIO_ENDPOINT,required"`
	MinIOBucket           string `env:"MINIO_BUCKET,required"`
}

func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...)

	if err != nil {
		log.Fatal("Unable to Load the .env file\n", err)
	}

	cfg := Configuration{}

	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
