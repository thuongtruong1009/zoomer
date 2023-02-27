package delivery

import (
	"github.com/labstack/echo/v4"
	"zoomer/internal/middlewares"
)

func MapRoomRoutes(roomGroup *echo.Group, h Handler, mw *middlewares.MiddlewareManager) {
	roomGroup.POST("/", h.AddRoom(), mw.JWTValidation)
	roomGroup.GET("/:userId", h.GetUserRooms(), mw.JWTValidation)
	roomGroup.GET("/", h.GetAll())
}
