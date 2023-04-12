package delivery

import (
	"github.com/labstack/echo/v4"
)

type ChatHandler interface {
	ChatConnect() echo.HandlerFunc
}
