package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
)

type AuthHandler interface {
	SignUp() echo.HandlerFunc

	SignIn() echo.HandlerFunc

	SignOut() echo.HandlerFunc

	writeCookie(ctx echo.Context, cookie *presenter.SetCookie)

	ForgotPassword() echo.HandlerFunc

	ResetPassword() echo.HandlerFunc
}
