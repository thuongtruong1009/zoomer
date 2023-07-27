package app

import (
	"runtime"
	"os"
	"os/signal"
	"syscall"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/db"
	"github.com/thuongtruong1009/zoomer/db/postgres"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/internal/server"
)

func Run(cfg *configs.Configuration) {
	numProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(numProcs)

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	pgAdapter := postgres.NewPgAdapter()
	redisInstance := db.GetRedisInstance(cfg)

	inter := interceptor.NewInterceptor()

	e := echo.New()
	defer e.Close()

	initServer := server.NewServer(e, cfg, pgAdapter, redisInstance, logger, inter)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	signal.Notify(interrupt, os.Kill)

	select {
	case s := <-interrupt:
		logger.Info("Got signal:", s.String())
	case err := <-initServer.Notify():
		logger.Error("Got server error:", err.Error())
	}

	// Shutdown server
	err := initServer.Shutdown()
	if err != nil {
		logger.Error("Error shutting down server: ", err)
	}
}
