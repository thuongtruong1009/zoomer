package delivery

import (
	"github.com/labstack/echo/v4"
	"zoomer/internal/middlewares"
	"zoomer/internal/rooms"
)

func MapAuthRoutes(roomGroup *echo.Group, h rooms.Handler, mw *middlewares.MiddlewareManager) {
	roomGroup.POST("/", h.AddRoom(), mw.JWTValidation)
	roomGroup.GET("/:userId", h.GetUserRooms(), mw.JWTValidation)
	roomGroup.GET("/", h.GetAll(), mw.JWTValidation)
}
