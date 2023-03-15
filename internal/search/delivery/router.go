package delivery

import "github.com/labstack/echo/v4"

func MapSearchRoutes(searchGroup *echo.Group, h Handler) {
	searchGroup.POST("/room", h.SearchRoom())
}