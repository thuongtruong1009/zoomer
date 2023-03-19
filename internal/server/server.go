package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zoomer/configs"
)

type Server struct {
	echo   *echo.Echo
	cfg    *configs.Configuration
	db     *gorm.DB
	logger *logrus.Logger
	ready  chan bool
}

func NewServer(cfg *configs.Configuration, db *gorm.DB, logger *logrus.Logger, ready chan bool) *Server {
	return &Server{
		echo: echo.New(), cfg: cfg, db: db, logger: logger, ready: ready,
	}
}

func (s *Server) Run() error {
	httpServer := &http.Server{
		Addr:         ":" + s.cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		// s.echo.Logger.Fatal(e.Start(":" + port))
		s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", s.cfg.Port)

		// http1.1
		// if err := s.echo.StartServer(httpServer); err != nil {
		// 	s.logger.Fatalln("Error starting server: ", err)
		// }

		// http2
		h2c := &http2.Server{
			MaxConcurrentStreams: 250,
			MaxReadFrameSize:     1048576,
			IdleTimeout:          10 * time.Second,
		}
		if err := e.StartH2CServer(":8080", h2c); err != http.ErrServerClosed {
			log.Fatal(err)
		}

		// https
		// if err := s.echo.StartTLS(":8080", ".docker/nginx/cert.pem", ".docker/nginx/key.pem"); err != http.ErrServerClosed {
		// 	log.Fatal(err)
		//   }
	}()

	if err := s.HttpMapHandlers(s.echo); err != nil {
		return err
	}

	go WsMapHandlers(s.cfg.WsPort)

	if s.ready != nil {
		s.ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)

	defer shutdown()

	s.logger.Fatalln("Server is exited properly")
	return s.echo.Server.Shutdown(ctx)
}
