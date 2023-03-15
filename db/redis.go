package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"zoomer/configs"
)

var RedisClient *redis.Client

func InitialiseRedis(cfg *configs.Configuration) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	pong, err := conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis connection failed", err)
	}

	log.Println("Redis connection successful", pong)

	RedisClient = conn
	return RedisClient
}
