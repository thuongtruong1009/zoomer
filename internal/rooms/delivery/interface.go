package delivery

import "github.com/labstack/echo/v4"

type Handler interface {
	GetAll() echo.HandlerFunc
	GetUserRooms() echo.HandlerFunc
	AddRoom() echo.HandlerFunc
}
