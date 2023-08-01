package server

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"time"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/configs/parameter"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/utils"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/internal/resources/minio/adapter"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	echo    *echo.Echo
	cfg     *configs.Configuration
	parameterCfg *parameter.ParameterConfig
	pgDB    *gorm.DB
	redisDB *redis.Client
	minioClient adapter.ResourceAdapter
	logger  *logrus.Logger
	inter   interceptor.IInterceptor
	notify  chan error
}

func NewServer(e *echo.Echo, cfg *configs.Configuration, parameterCfg *parameter.ParameterConfig, pgDB *gorm.DB, redisDB *redis.Client, minioClient adapter.ResourceAdapter, logger *logrus.Logger, inter interceptor.IInterceptor) *Server {
	s := &Server{
		echo:    e,
		cfg:     cfg,
		parameterCfg: parameterCfg,
		pgDB:    pgDB,
		redisDB: redisDB,
		minioClient: minioClient,
		logger:  logger,
		inter:   inter,
		notify:  make(chan error, 1),
	}

	s.start()
	return s
}

func (s *Server) start() {
	function1 := func() {
		httpServer := &http.Server{
			Addr:         ":" + s.cfg.AppPort,
			WriteTimeout: _defaultWriteTimeout,
			ReadTimeout:  _defaultReadTimeout,
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
			if err := s.HttpMapServer(s.echo); err != nil {
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
		if err2 := WsMapServer(s.echo, s.redisDB, s.inter); err2 != nil {
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

func (s *Server) Shutdown() error {
	exceptions.SystemLog(constants.ServerShutdown)
	ctx, cancel := context.WithTimeout(context.Background(), _defaultShutdownTimeout)
	defer cancel()

	exceptions.SystemLog(constants.ServerExitedProperly)
	return s.echo.Shutdown(ctx)
}
