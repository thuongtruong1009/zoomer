package main

import (
	"log"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/db"
	"github.com/thuongtruong1009/zoomer/internal/app"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
)

// @title Echo REST API
// @version 1.0
// @description This documentation for Echo REST server.
// @termsOfService http://swagger.io/terms/

// @contact.name Dzung Tran
// @contact.url https://docs.api.com/support
// @contact.email support@api.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey  XUserEmailAuth
// @in                          header
// @name                        X-User-Email
// @description					This method just enabled for local development

// @securityDefinitions.apikey  XFirebaseBearer
// @in                          header
// @name                        Authorization
// @description					Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.

// func init() {
// 	echo.SetMode(echo.ReleaseMode)
// }

func main() {
	e := echo.New()
	defer e.Close()

	cfg := configs.NewConfig()

	pgInstance := db.GetPostgresInstance(cfg)
	redisInstance := db.GetRedisInstance(cfg)

	inter := interceptor.NewInterceptor()

	s := app.NewServer(e, cfg, pgInstance, redisInstance, logrus.New(), nil, inter)

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
		os.Exit(0)
	}

	log.Println("Starting server")
}
