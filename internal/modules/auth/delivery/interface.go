package delivery

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	SignUp() echo.HandlerFunc

	SignIn() echo.HandlerFunc

	SignOut() echo.HandlerFunc
}
