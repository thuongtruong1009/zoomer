package delivery

import (
	"github.com/labstack/echo/v4"
)

func MapAuthRoutes(authGroup *echo.Group, h AuthHandler) {
	authGroup.POST("/signup", h.SignUp())
	authGroup.POST("/signin", h.SignIn())
	authGroup.POST("/signout", h.SignOut())
}
