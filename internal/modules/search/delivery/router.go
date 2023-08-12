package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapSearchRoutes(searchGroup *echo.Group, h Handler) {
	searchGroup.POST(constants.SearchRoomEndPoint, h.SearchRoom())
}
