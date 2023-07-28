package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/db"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/internal/server"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func Run(cfg *configs.Configuration) {
	numProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(numProcs)

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	pgAdapter := postgres.NewPgAdapter()
	pgInstance := pgAdapter.ConnectInstance(cfg)
	redisInstance := db.GetRedisInstance(cfg)

	inter := interceptor.NewInterceptor()

	e := echo.New()
	defer e.Close()

	initServer := server.NewServer(e, cfg, pgInstance, redisInstance, logger, inter)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(interrupt, os.Kill)

	select {
	case s := <-interrupt:
		logger.Info("Got terminate signal: ", s.String())
	case err := <-initServer.Notify():
		logger.Error("Got server error: ", err.Error())
	}

	if err := initServer.Shutdown(); err != nil {
		logger.Errorf("Error shutting down server: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
