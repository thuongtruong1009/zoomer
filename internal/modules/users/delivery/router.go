package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapUserRoutes(userGroup *echo.Group, h IUserHandler) {
	userGroup.GET(constants.GetUserByIdOrName, h.GetUserByIdOrName())

	userGroup.GET("/search", h.SearchUser())
}
