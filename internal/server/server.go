package server

import (
	"context"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/utils"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	echo   *echo.Echo
	cfg    *configs.Configuration
	pgDB     postgres.PgAdapter
	redisDB  *redis.Client
	logger *logrus.Logger
	inter  interceptor.IInterceptor
	notify chan error

}

func NewServer(e *echo.Echo, cfg *configs.Configuration, pgDB postgres.PgAdapter, redisDB *redis.Client, logger *logrus.Logger,inter interceptor.IInterceptor) *Server {
	s := &Server{
		echo: e,
		cfg: cfg,
		pgDB: pgDB,
		redisDB: redisDB,
		logger: logger,
		inter: inter,
		notify: make(chan error, 1),
	}

	s.start()
	return s
}

func (s *Server) start() {
	function1 := func() {
		httpServer := &http.Server{
			Addr:         ":" + s.cfg.AppPort,
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}

		if s.cfg.HttpsMode == "true" {	// https
			certPath := utils.GetFilePath(constants.CertPath)
			keyPath := utils.GetFilePath(constants.KeyPath)
			configs.TLSConfig(certPath, keyPath)
			if err := s.echo.StartTLS(httpServer.Addr, certPath, keyPath); err != http.ErrServerClosed {
				s.logger.Fatalln("Error occured when starting the server in HTTPS mode", err)
			}
		} else { // http
			err := s.HttpMapServer(s.echo)
			if err != nil {
				s.logger.Fatalln("Error occurred while setting up http routers: ", err)
			}

			s.logger.Logf(logrus.InfoLevel, "::: Api server is listening on PORT: %s", s.cfg.AppPort)

			if err := s.echo.StartServer(httpServer); err != nil {
				s.logger.Fatalln("Error occurred while starting the http server: ", err)
			}
		}
	}


	function2 := func() {
		if err2 := WsMapServer(s.echo, s.redisDB, s.inter); err2 != nil {
			s.logger.Fatalln("Error occurred while setting up websocket routers: ", err2)
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
	s.logger.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), _defaultShutdownTimeout)
	defer cancel()

	s.logger.Fatalln("Server is exited properly")
	return s.echo.Shutdown(ctx)
}
