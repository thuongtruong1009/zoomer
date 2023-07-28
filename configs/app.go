package configs

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type appConfig struct {
	AppPort     string `env:"APP_PORT" envDefault:"8080"`
	HashSalt    string `env:"HASH_SALT,required"`
	SigningKey  string `env:"SIGNING_KEY,required"`
	TokenTTL    int64  `env:"TOKEN_TTL,required"`
	JwtSecret   string `env:"JWT_SECRET,required"`
	AutoMigrate bool   `env:"AUTO_MIGRATE" envDefault:"true"`
	HttpsMode   bool   `env:"HTTPS_MODE" envDefault:"false"`
}

type postgresConfig struct {
	DatabaseConnectionURL string `env:"PG_URI,required"`
	MaxOpenConnection     int    `env:"PG_MAX_OPEN_CONN" envDefault:"20"`
	MaxIdleConnection     int    `env:"PG_MAX_IDLE_CONN" envDefault:"20"`
	MaxLifetimeConnection int    `env:"PG_MAX_LIFETIME_CONN" envDefault:"20"`
	MaxIdleTimeConnection int    `env:"PG_MAX_IDLE_TIME_CONN" envDefault:"20"`
}

type redisConfig struct {
	RedisURI      string `env:"REDIS_URI,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
}

type minioConfig struct {
	MinIOAccessKey string `env:"MINIO_ACCESS,required"`
	MinIOSecretKey string `env:"MINIO_SECRET,required"`
	MinIOEndpoint  string `env:"MINIO_ENDPOINT,required"`
	MinIOBucket    string `env:"MINIO_BUCKET,required"`
}

type rmqConfig struct {
	RmqURI string `env:"RMQ_URI,required"`
}

type Configuration struct {
	appConfig
	postgresConfig
	redisConfig
	minioConfig
	rmqConfig
}

func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...)

	if err != nil {
		log.Fatal("Unable to Load the .env file\n", err)
	}

	cfg := &Configuration{}

	err = env.Parse(cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return cfg
}

func GetEnvKey(key string) string {
	return os.Getenv(key)
}

func LookupEnv(key string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok || len(envVal) == 0 {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return envVal
}
