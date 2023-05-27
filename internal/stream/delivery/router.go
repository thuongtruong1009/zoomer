package delivery

import (
	"github.com/labstack/echo/v4"
)

func MapStreamRoutes(e *echo.Echo, h StreamHandler, group string) {
	e.GET( /*group + */ "/create", h.CreateStream())
	e.GET( /*group + */ "/join", h.JoinStream())
}
