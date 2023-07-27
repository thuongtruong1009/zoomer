package main

import (
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/internal/app"
)

// @title Zoomer
// @version 2.0
// @description This documentation for Zoomer API
// @termsOfService http://swagger.io/terms/

// @contact.name Tran Nguyen Thuong Truong
// @contact.url https://github.com/thuongtruong1009/zoomer
// @contact.email thuongtruongofficial@gmail.com

// @license.name Apache 2.0
// @license.url https://github.com/thuongtruong1009/zoomer/LICENSE

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
	cfg := configs.NewConfig()

	app.Run(cfg)
}
