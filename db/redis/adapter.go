package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
)

type redisStruct struct {
	redis *redis.Client
	cfg  *configs.Configuration
	paramCfg *parameter.RedisConf
}

func NewRedisAdapter(cfg *configs.Configuration, paramCfg *parameter.RedisConf) RedisAdapter {
	return &redisStruct{
		cfg: cfg,
		paramCfg: paramCfg,
	}
}

func (rd *redisStruct) getInstance() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rd.cfg.RedisURI,
		Password: rd.cfg.RedisPassword,
		DB:       rd.paramCfg.DB,
		MinIdleConns: rd.paramCfg.MinIdleConns,
		PoolSize: rd.paramCfg.PoolSize,
		PoolTimeout:  helpers.DurationSecond(rd.paramCfg.PoolTimeout),
	})
}

func (rd *redisStruct) ConnectInstance() *redis.Client {
	conn := rd.getInstance()

	ctx, cancel := context.WithTimeout(context.Background(), helpers.DurationSecond(rd.paramCfg.IdleTimeout))
	defer cancel()

	pong, err := conn.Ping(ctx).Result()
	if err != nil {
		exceptions.Fatal(constants.ErrorRedisConnectionFailed, err)
	}

	exceptions.SystemLog(fmt.Sprintf("%s: %s", constants.RedisConnectionSuccessful, pong))

	return conn
}


