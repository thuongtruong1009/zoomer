package delivery

import "github.com/labstack/echo/v4"

type IUserHandler interface {
	GetUserByIdOrName() echo.HandlerFunc
}
