package db

import (
	"context"
	// "crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"github.com/thuongtruong1009/zoomer/configs"
)

var RedisClient *redis.Client

func GetRedisInstance(cfg *configs.Configuration) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPassword,
		DB:       0,
		// DialTimeout:  100 * time.Millisecond,
		// ReadTimeout:  100 * time.Millisecond,
		// WriteTimeout: 100 * time.Millisecond,
		// TLSConfig:    &tls.Config{MinVersion: tls.VersionTLS12},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := conn.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("Redis connection failed:  %w", err))
	}

	log.Println("Redis connection successful", pong)

	RedisClient = conn

	return RedisClient
}
