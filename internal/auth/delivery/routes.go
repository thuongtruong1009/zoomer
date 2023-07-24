package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapAuthRoutes(authGroup *echo.Group, h AuthHandler) {
	authGroup.POST(constants.SignUpEndPoint, h.SignUp())
	authGroup.POST(constants.SignInEndPoint, h.SignIn())
	authGroup.POST(constants.SignOutEndPoint, h.SignOut())
}
