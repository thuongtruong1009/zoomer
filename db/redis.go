package db

import (
	"context"
	"log"
	"time"
	"fmt"
	"github.com/go-redis/redis/v8"
	"zoomer/configs"
)

type RedisClient struct {
	client *redis.Client
}


func NewRedisClient(cfg *configs.Configuration) *RedisClient {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := conn.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("Redis connection failed:  %w", err))
	}

	log.Println("Redis connection successful", pong)

	return &RedisClient{conn}
}
