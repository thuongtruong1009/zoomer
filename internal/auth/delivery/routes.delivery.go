package delivery

import (
	"github.com/labstack/echo/v4"
	"zoomer/internal/auth"
)

func MapAuthRoutes(authGroup *echo.Group, h auth.Handler) {
	authGroup.POST("/signup", h.SignUp())
	authGroup.POST("/signin", h.SignIn())
}
