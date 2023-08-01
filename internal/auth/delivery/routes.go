package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/middlewares"
)

func MapAuthRoutes(authGroup *echo.Group, h AuthHandler, mw *middlewares.AuthMiddleware) {
	authGroup.POST(constants.SignUpEndPoint, h.SignUp())
	authGroup.POST(constants.SignInEndPoint, h.SignIn())
	authGroup.POST(constants.SignOutEndPoint, h.SignOut(), mw.JWTValidation)
}
