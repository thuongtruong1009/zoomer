// +build prod

package main

import (
	"os"
	"github.com/thuongtruong1009/zoomer/infrastructure/app"
)

//	@title			Zoomer
//	@version		1.1
//	@description	The HTTP documentation for Zoomer API
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Tran Nguyen Thuong Truong
//	@contact.url	https://github.com/thuongtruong1009/zoomer
//	@contact.email	mailto:thuongtruongofficial@gmail.com

//	@license.name	Apache 2.0
//	@license.url	https://github.com/thuongtruong1009/zoomer/LICENSE

//	@schemes	https
//	@host		localhost:80
//	@BasePath	/api

//	@securityDefinitions	bearerAuth
//	@in						header
//	@name					Authorization
//	@description			Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.

//	@securityDefinitions.apikey	XFirebaseBearer
//	@in							header
//	@name						Authorization
//	@description				Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.

func main() {
	app.Run()
	defer os.Exit(0)
}
