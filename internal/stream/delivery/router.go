package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapStreamRoutes(e *echo.Echo, h StreamHandler) {
	e.GET(constants.CreateStreamEndPoint, h.CreateStream())
	e.GET(constants.JoinStreamEndPoint, h.JoinStream())
}
