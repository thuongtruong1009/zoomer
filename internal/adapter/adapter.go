package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/configs/parameter"
	"github.com/thuongtruong1009/zoomer/db"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	minioAdapter "github.com/thuongtruong1009/zoomer/internal/resources/minio/adapter"
	"github.com/thuongtruong1009/zoomer/internal/server"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
)

func Adapter(cfg *configs.Configuration, paramCfg *parameter.ParameterConfig) *server.Server {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	pgAdapter := postgres.NewPgAdapter(&paramCfg.PostgresConf)
	pgInstance := pgAdapter.ConnectInstance(cfg)

	redisInstance := db.GetRedisInstance(cfg)

	newMinioClient, _ := minioAdapter.RegisterMinioClient(cfg)
	minioAdapter.SetPermission(newMinioClient, constants.BucketName)
	minioAdapter.CreateBucket(newMinioClient, constants.BucketName)
	newMinioAdapter := minioAdapter.NewMinioAdapter(newMinioClient, constants.BucketName, cfg)

	inter := interceptor.NewInterceptor()

	e := echo.New()
	defer e.Close()

	initServer := server.NewServer(e, cfg, paramCfg, pgInstance, redisInstance, newMinioAdapter, logger, inter)

	return initServer
}
