package configs

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

type Configuration struct {
	Port                  string `env:"PORT" envDefault:"8080"`
	WsPort                string `env:"WS_PORT" envDefault:"8081"`
	HashSalt              string `env:"HASH_SALT,required"`
	SigningKey            string `env:"SIGNING_KEY,required"`
	TokenTTL              int64  `env:"TOKEN_TTL,required"`
	JwtSecret             string `env:"JWT_SECRET,required"`
	DatabaseConnectionURL string `env:"PG_URI,required"`
	RedisURI              string `env:"REDIS_URI,required"`
	RedisPassword         string `env:"REDIS_PASSWORD,required"`
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
