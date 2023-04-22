package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"log"
	"zoomer/configs"
	"zoomer/db"
	"zoomer/internal/server"
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

// @BasePath /

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
	cfg := configs.NewConfig()
	instance := db.GetPostgresInstance(cfg)

	s := server.NewServer(echo.New(), cfg, instance, logrus.New(), nil)

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	log.Println("Starting api server")
}
