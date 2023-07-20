package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/utils"
)

type Server struct {
	echo   *echo.Echo
	cfg    *configs.Configuration
	pgDB     postgres.PgAdapter
	redisDB  *redis.Client
	logger *logrus.Logger
	ready  chan bool
	inter  interceptor.IInterceptor
}

func NewServer(e *echo.Echo, cfg *configs.Configuration, pgDB postgres.PgAdapter, redisDB *redis.Client, logger *logrus.Logger, ready chan bool, inter interceptor.IInterceptor) *Server {
	return &Server{
		echo: e,
		cfg: cfg,
		pgDB: pgDB,
		redisDB: redisDB,
		logger: logger,
		ready: ready,
		inter: inter,
	}
}

func (s *Server) Run() error {
	function1 := func() {
		httpServer := &http.Server{
			Addr:         ":" + s.cfg.HttpPort,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		if s.cfg.HttpsMode == "true" {	// https
			certPath := utils.GetFilePath(constants.CertPath)
			keyPath := utils.GetFilePath(constants.KeyPath)
			configs.TLSConfig(certPath, keyPath)
			if err := s.echo.StartTLS(httpServer.Addr, certPath, keyPath); err != http.ErrServerClosed {
				s.logger.Fatalln("Error occured when starting the server in HTTPS mode", err)
			}
		}else { // http
			if err := s.echo.StartServer(httpServer); err != nil {
				s.logger.Fatalln("Error occurred while starting the http server: ", err)
			}
		}

		s.logger.Logf(logrus.InfoLevel, "api server is listening on PORT: %s", s.cfg.HttpPort)

		if err := s.HttpMapServer(s.echo); err != nil {
			s.logger.Fatalln("Error occurred while setting up http routers: ", err)
		}
	}

	function2 := func(){
		if err2 := WsMapServer(echo.New(), s.redisDB, s.inter, ":8081"); err2 != nil {
			s.logger.Fatalln("Error occurred while setting up websocket routers: ", err2)
		}
	}

	Parallelize(function1, function2)

	if s.ready != nil {
		s.ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	signal.Notify(quit, os.Kill)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	s.logger.Fatalln("Server is exited properly")
	return s.echo.Server.Shutdown(ctx)
}

func Parallelize(functions ...func()) {
    var waitGroup sync.WaitGroup

	ch := make(chan struct{}, 2)

    waitGroup.Add(len(functions))

    defer waitGroup.Wait()

    for _, function := range functions {
		ch <- struct{}{}
		go func(copyFunc func()) {
			defer func() {
				<-ch
				waitGroup.Done()
			}()
			copyFunc()
		}(function)
	}
}
