package delivery

import (
	"github.com/labstack/echo/v4"
)

func MapChatRoutes(e *echo.Echo, h ChatHandler, group string) {
	// http.HandleFunc("/ws", ChatConnect)
	e.GET(group, h.ChatConnect())
}
