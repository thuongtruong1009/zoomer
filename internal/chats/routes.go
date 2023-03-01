package chats

import (
	"github.com/labstack/echo/v4"
	"zoomer/internal/middlewares"
)

func MapChatRoutes(chatsGroup *echo.Group, h *Handler, mw *middlewares.MiddlewareManager) {
	chatsGroup.POST("/createRoom", h.CreateRoom(), mw.JWTValidation)
	chatsGroup.GET("/joinRoom/:roomId", h.JoinRoom(), mw.JWTValidation)
	chatsGroup.GET("/getRooms", h.GetRooms(), mw.JWTValidation)
	chatsGroup.GET("/getClients/:roomId", h.GetClients(), mw.JWTValidation)
}