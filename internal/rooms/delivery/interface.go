package delivery

import "github.com/labstack/echo/v4"

type Handler interface {
	GetAll() echo.HandlerFunc
	GetUserRooms() echo.HandlerFunc
	AddRoom() echo.HandlerFunc
	//sync to redis
	ChatHistoryHandler() echo.HandlerFunc
	ContactListHandler() echo.HandlerFunc
	CreateFetchChatBetweenIndexHandler() echo.HandlerFunc
}
