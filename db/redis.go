package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"zoomer/configs"
)

var RedisClient *redis.Client

func GetRedisInstance() *redis.Client {
	fmt.Println("Redis connection successful", cfg.RedisURI, cfg.RedisPassword)
	conn := redis.NewClient(&redis.Options{
		Addr:     configs.Configuration.RedisURI,
		Password: configs.Configuration.RedisPassword,
		DB:       0,
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
