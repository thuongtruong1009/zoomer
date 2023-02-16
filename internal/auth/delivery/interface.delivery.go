package delivery

import "github.com/labstack/echo/v4"

type Handler interface {
	SignUP() echo.HandlerFunc

	SignIn() echo.HandlerFunc

	// Logout() echo.HandlerFunc
}
