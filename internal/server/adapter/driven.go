package adapter

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/db/redis"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	minioAdapter "github.com/thuongtruong1009/zoomer/internal/modules/resources/minio/adapter"
	"github.com/thuongtruong1009/zoomer/internal/server/api"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
)

type IAdapter interface {
	Notify() <-chan error
	Shutdown(ctx context.Context) error
}

type Adapter struct {
	Api api.IApi
}

type Options func(opts *api.IApi) error

var signal = make(chan error, 1)

func NewAdapter(cfg *configs.Configuration, paramCfg *parameter.ParameterConfig, opts ...Options) IAdapter {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	pgAdapter := postgres.NewPgAdapter(cfg, &paramCfg.PostgresConf)
	pgInstance := pgAdapter.ConnectInstance()

	redisAdapter := redis.NewRedisAdapter(cfg, &paramCfg.RedisConf)
	redisInstance := redisAdapter.ConnectInstance()

	newMinioClient, _ := minioAdapter.RegisterMinioClient(cfg)
	minioAdapter.SetPermission(newMinioClient, constants.BucketName)
	minioAdapter.CreateBucket(newMinioClient, constants.BucketName)
	newMinioAdapter := minioAdapter.NewMinioAdapter(newMinioClient, constants.BucketName, cfg)

	inter := interceptor.NewInterceptor()

	e := echo.New()
	defer e.Close()

	initServer := api.NewApi(e, cfg, paramCfg, pgInstance, redisInstance, newMinioAdapter, logger, inter)
	for _, opt := range opts {
		err := opt(&initServer)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"error": err,
			}).Error("Error when init server")
			return nil
		}
	}
	go initServer.Start(signal)
	return &Adapter{
		Api: initServer,
	}
}

func (s *Adapter) Notify() <-chan error {
	return signal
}

func (s *Adapter) Shutdown(ctx context.Context) error {
	return s.Api.Stop(ctx)
}
