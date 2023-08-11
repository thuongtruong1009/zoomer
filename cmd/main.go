package main

import (
	"github.com/thuongtruong1009/zoomer/infrastructure/app"
	_"github.com/thuongtruong1009/zoomer/pkg/constants"
	"os"
)

//	@title			constants.AppName
//	@version		constants.AppVersion
//	@description	constants.AppDescription
//	@termsOfService	constants.AppTermsOfService

//	@contact.name	constants.AppContactName
//	@contact.url	constants.AppContactURL
//	@contact.email	constants.AppContactEmail

//	@license.name	constants.AppLicenseName
//	@license.url	constants.AppLicenseURL

//	@host		constants.AppHost
//	@BasePath	constants.ApiGroup

//	@securityDefinitions	bearerAuth
//	@in						header
//	@name					constants.BearerHeader
//	@description			Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.

//	@securityDefinitions.apikey	XFirebaseBearer
//	@in							header
//	@name						Authorization
//	@description				Enter the token with the `Bearer ` prefix, e.g. `Bearer jwt_token_string`.

// func init() {
// 	echo.SetMode(echo.ReleaseMode)
// }

func main() {
	app.Run()
	defer os.Exit(0)
}
