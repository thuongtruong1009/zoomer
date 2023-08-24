package redis

import "github.com/go-redis/redis/v8"

type RedisAdapter interface {
	getInstance() *redis.Client
	ConnectInstance() *redis.Client
}
