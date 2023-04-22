package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"zoomer/configs"
	"zoomer/pkg/constants"
	"zoomer/pkg/utils"
)

type Server struct {
	echo   *echo.Echo
	cfg    *configs.Configuration
	db     *gorm.DB
	logger *logrus.Logger
	ready  chan bool
}

func NewServer(e *echo.Echo, cfg *configs.Configuration, db *gorm.DB, logger *logrus.Logger, ready chan bool) *Server {
	return &Server{
		echo: e, cfg: cfg, db: db, logger: logger, ready: ready,
	}
}

func (s *Server) Run() error {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		httpServer := &http.Server{
			Addr:         ":" + s.cfg.HttpPort,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		// https2.0
		if s.cfg.HttpsMode == "true" {
			certPath := utils.GetFilePath(constants.CertPath)
			keyPath := utils.GetFilePath(constants.KeyPath)
			configs.TLSConfig(certPath, keyPath)
			if err := s.echo.StartTLS(httpServer.Addr, certPath, keyPath); err != http.ErrServerClosed {
				s.logger.Fatalln("Error occured when starting the server in HTTPS mode", err)
			}
		}

		// http1.1
		if err := s.echo.StartServer(httpServer); err != nil {
			s.logger.Fatalln("Error occurred while starting the http server: ", err)
		}

		s.logger.Logf(logrus.InfoLevel, "api server is listening on PORT: %s", s.cfg.HttpPort)
		wg.Done()
	}()

	s.logger.Log(logrus.InfoLevel, "Setting up routers")
	if err := s.HttpMapServer(s.echo); err != nil {
		s.logger.Fatalln("Error occurred while setting up routers: ", err)
	}

	go func() {
		WsMapServer(":" + s.cfg.WsPort)
		s.logger.Log(logrus.InfoLevel, "websocket server is starting on :"+s.cfg.WsPort)
		wg.Done()
	}()

	wg.Wait()

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
