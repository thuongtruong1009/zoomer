package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"zoomer/db"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Delete(key string) error
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisClient) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisClient) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}