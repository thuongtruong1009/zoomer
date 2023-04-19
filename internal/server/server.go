package server

import (
	"context"
	"fmt"
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
	"zoomer/pkg/load_balancer"
	// "log"
	// "golang.org/x/net/http2"
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
		Addr:         ":" + s.cfg.HttpPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", s.cfg.HttpPort)

		// https
		if s.cfg.HttpsMode == "true" {
			if err := s.echo.StartTLS(httpServer.Addr, constants.CertPath, constants.KeyPath); err != http.ErrServerClosed {
				s.logger.Fatalln("Error occured when starting the server in HTTPS mode", err)
			}
		}

		// http1.1
		// if err := s.echo.StartServer(httpServer); err != nil {
		// 	s.logger.Fatalln("Error occurred while starting the http server: ", err)
		// }
		loadBalancer(s)
	}()

	s.logger.Log(logrus.InfoLevel, "Setting up routers")
	if err := s.HttpMapServer(s.echo); err != nil {
		return err
	}

	WsMapServer(":" + s.cfg.WsPort)
	fmt.Println("websocket server is starting on :" + s.cfg.WsPort)

	if s.ready != nil {
		s.ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	signal.Notify(quit, os.Kill)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)

	defer shutdown()

	s.logger.Fatalln("Server is exited properly")
	return s.echo.Server.Shutdown(ctx)
}

func http11ApiStart(s *Server, port string) {
	if err := s.echo.StartServer(&http.Server{
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}); err != nil {
		s.logger.Fatalln("Error occurred while starting the http server: ", err)
	}
}

func loadBalancer(s *Server) {
	wg := new(sync.WaitGroup)
	wg.Add(5)

	go func() {
		load_balancer.LoadBalancer()
		wg.Done()
	}()

	loadBalancerPorts := [3]string{"8082", "8083", "8084"}

	for i := 0; i < len(loadBalancerPorts); i++ {
		go func() {
			http11ApiStart(s, loadBalancerPorts[i])
			wg.Done()
		}()
	}
	wg.Wait()
}
