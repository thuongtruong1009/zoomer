package delivery

import "github.com/labstack/echo/v4"

type Handler interface {
	SearchRoom() echo.HandlerFunc
}