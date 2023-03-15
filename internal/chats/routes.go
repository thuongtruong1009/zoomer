package chats

import (
	"github.com/labstack/echo/v4"
)

func MapChatRoutes(e *echo.Echo, h *Handler, group string) {
	e.POST(group+"/createRoom", h.CreateRoom())
	e.GET(group+"/joinRoom/:roomId", h.JoinRoom())
	e.GET(group+"/getRooms", h.GetRooms())
	e.GET(group+"/getClients/:roomId", h.GetClients())
	e.GET(group+"/getStats", h.GetStats())
}
