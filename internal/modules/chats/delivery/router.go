package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapChatRoutes(e *echo.Echo, h ChatHandler) {
	e.GET(constants.ChatConnectEndPoint, h.ChatConnect())
}
