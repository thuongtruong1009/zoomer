package configs

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"log"
	"os"
)

type Configuration struct {
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	HashSalt   string `env:"HASH_SALT,required"`
	SigningKey string `env:"SIGNING_KEY,required"`
	TokenTTL   int64  `env:"TOKEN_TTL,required"`
	JwtSecret  string `env:"JWT_SECRET,required"`
	HttpsMode  bool   `env:"HTTPS_MODE" envDefault:"false"`

	DatabaseConnectionURL string `env:"PG_CONNECTION,required"`

	RedisURI      string `env:"REDIS_URI,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`

	MinIOAccessKey string `env:"MINIO_ACCESS,required"`
	MinIOSecretKey string `env:"MINIO_SECRET,required"`
	MinIOEndpoint  string `env:"MINIO_ENDPOINT,required"`
	MinIOBucket    string `env:"APP_NAME,required"`

	RmqURI string `env:"RMQ_URI,required"`
}

func LoadConfigs(files ...string) *Configuration {
	err := godotenv.Load(files...)
	if err != nil {
		exceptions.Fatal(constants.ErrorLoadEnvFile, err)
	}

	cfg := &Configuration{}

	err = env.Parse(cfg)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return cfg
}

func GetEnvVar(key string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok || len(envVal) == 0 {
		exceptions.Fatal(constants.ErrorEnvKeyNotFound, key)
	}
	return envVal
}
