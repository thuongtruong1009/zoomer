package delivery

import (
	"github.com/labstack/echo/v4"
)

type StreamHandler interface {
	CreateStream() echo.HandlerFunc
	JoinStream() echo.HandlerFunc
}
