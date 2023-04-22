package delivery

import (
	"github.com/labstack/echo/v4"
	"zoomer/pkg/middlewares"
)

func MapRoomRoutes(roomGroup *echo.Group, h Handler, mw *middlewares.MiddlewareManager) {
	roomGroup.POST("/", h.AddRoom(), mw.JWTValidation)
	roomGroup.GET("/:userId", h.GetUserRooms(), mw.JWTValidation)
	roomGroup.GET("/", h.GetAll())
	//sync to redis

	h.CreateFetchChatBetweenIndexHandler()

	roomGroup.GET("/chat-history", h.ChatHistoryHandler())

	roomGroup.GET("/contact-list", h.ContactListHandler())
}
