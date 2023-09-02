package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/middlewares"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapRoomRoutes(roomGroup *echo.Group, h Handler, mw *middlewares.AuthMiddleware) {
	roomGroup.POST(constants.CreateRoomEndPoint, h.AddRoom(), mw.JWTValidation)
	roomGroup.GET(constants.GetRoomsOfUserEndPoint, h.GetUserRooms(), mw.JWTValidation)
	roomGroup.GET(constants.GetAllRoomsEndPoint, h.GetAll())

	//sync to redis
	roomGroup.GET(constants.GetChatHistoryEndPoint, h.ChatHistoryHandler())
	roomGroup.GET(constants.GetContactListEndPoint, h.ContactListHandler())
}
