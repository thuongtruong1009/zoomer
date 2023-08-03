package server

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/resources/minio/adapter"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Server struct {
	echo         *echo.Echo
	cfg          *configs.Configuration
	parameterCfg *parameter.ParameterConfig
	pgDB         *gorm.DB
	redisDB      *redis.Client
	minioClient  adapter.ResourceAdapter
	logger       *logrus.Logger
	inter        interceptor.IInterceptor
	notify       chan error
}

func NewServer(e *echo.Echo, cfg *configs.Configuration, parameterCfg *parameter.ParameterConfig, pgDB *gorm.DB, redisDB *redis.Client, minioClient adapter.ResourceAdapter, logger *logrus.Logger, inter interceptor.IInterceptor) IServer {
	s := &Server{
		echo:         e,
		cfg:          cfg,
		parameterCfg: parameterCfg,
		pgDB:         pgDB,
		redisDB:      redisDB,
		minioClient:  minioClient,
		logger:       logger,
		inter:        inter,
		notify:       make(chan error, 1),
	}

	s.start()
	return s
}

func (s *Server) start() {
	function1 := func() {
		httpServer := &http.Server{
			Addr:         ":" + s.cfg.AppPort,
			WriteTimeout: s.parameterCfg.ServerConf.WriteTimeout * time.Second,
			ReadTimeout:  s.parameterCfg.ServerConf.ReadTimeout * time.Second,
		}

		if s.cfg.HttpsMode { // https
			certPath := utils.GetFilePath(constants.CertPath)
			keyPath := utils.GetFilePath(constants.KeyPath)
			configs.TLSConfig(certPath, keyPath)
			if err := s.echo.StartTLS(httpServer.Addr, certPath, keyPath); err != http.ErrServerClosed {
				exceptions.Fatal(constants.ErrorStartHttps, err)
				<-s.notify
			}
		} else { // http
			if err := s.HttpMapServer(); err != nil {
				exceptions.Fatal(constants.ErrorSetupHttpRouter, err)
				<-s.notify
			}

			exceptions.SystemLog(fmt.Sprintf("%s: %s", constants.ServerApiStarted, s.cfg.AppPort))

			if err := s.echo.StartServer(httpServer); err != nil && err != http.ErrServerClosed {
				exceptions.Fatal(constants.ErrorStartHttp, err)
				<-s.notify
			}
		}
	}

	function2 := func() {
		if err2 := s.WsMapServer(); err2 != nil {
			exceptions.Fatal(constants.ErrorSetupSocketRouter, err2)
			<-s.notify
		}
	}

	go func() {
		helpers.Parallelize(function1, function2)
		defer close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
