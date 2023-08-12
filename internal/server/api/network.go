package api

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/modules/resources/minio/adapter"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"gorm.io/gorm"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/pkg/utils"
	"net/http"
	"time"
	"context"
)

type Api struct {
	Echo         *echo.Echo
	Cfg          *configs.Configuration
	ParameterCfg *parameter.ParameterConfig
	PgDB         *gorm.DB
	RedisDB      *redis.Client
	MinioClient  adapter.ResourceAdapter
	Logger       *logrus.Logger
	Inter        interceptor.IInterceptor
}

type IApi interface {
	HttpApi() error
	WsApi() error
	Start(chan error)
	Stop(context.Context) error
}

type Options func(opts *Api) error

func NewApi(e *echo.Echo, cfg *configs.Configuration, parameterCfg *parameter.ParameterConfig, pgDB *gorm.DB, redisDB *redis.Client, minioClient adapter.ResourceAdapter, logger *logrus.Logger, inter interceptor.IInterceptor) IApi {

	s := &Api{
		Echo:         e,
		Cfg:          cfg,
		ParameterCfg: parameterCfg,
		PgDB:         pgDB,
		RedisDB:      redisDB,
		MinioClient:  minioClient,
		Logger:       logger,
		Inter:        inter,
	}

	return s
}

func (s *Api) Start(c chan error) {
	function1 := func() {
		httpServer := &http.Server{
			Addr:         ":" + s.Cfg.AppPort,
			WriteTimeout: s.ParameterCfg.ServerConf.WriteTimeout * time.Second,
			ReadTimeout:  s.ParameterCfg.ServerConf.ReadTimeout * time.Second,
		}

		if s.Cfg.HttpsMode { // https
			certPath := utils.GetFilePath(constants.CertPath)
			keyPath := utils.GetFilePath(constants.KeyPath)
			configs.TLSConfig(certPath, keyPath)
			if err := s.Echo.StartTLS(httpServer.Addr, certPath, keyPath); err != http.ErrServerClosed {
				exceptions.Fatal(constants.ErrorStartHttps, err)
				c <- constants.ErrorStartHttps
			}
		} else { // http
			if err := s.HttpApi(); err != nil {
				exceptions.Fatal(constants.ErrorSetupHttpRouter, err)
				c <- constants.ErrorSetupHttpRouter
			}

			exceptions.SystemLog(fmt.Sprintf("%s: %s", constants.ServerApiStarted, s.Cfg.AppPort))

			if err := s.Echo.StartServer(httpServer); err != nil && err != http.ErrServerClosed {
				exceptions.Fatal(constants.ErrorStartHttp, err)
				c <- constants.ErrorStartHttp
			}
		}
	}

	function2 := func() {
		if err2 := s.WsApi(); err2 != nil {
			exceptions.Fatal(constants.ErrorSetupSocketRouter, err2)
			c <- constants.ErrorSetupSocketRouter
		}
	}

	helpers.Parallelize(function1, function2)
	c <- nil
}

func (s *Api) Stop(ctx context.Context) error {
	s.RedisDB.Close()
	s.RedisDB.Shutdown(ctx)
	return s.Echo.Shutdown(ctx)
}
