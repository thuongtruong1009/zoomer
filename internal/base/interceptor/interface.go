package interceptor

import (
	"github.com/labstack/echo/v4"
)

type IInterceptor interface {
	Data(c echo.Context, code int, data interface{}) error

	Error(c echo.Context, code int, msg error, err error) error
}

type InterceptorProps struct {
	Data interface{} `json:"data"`
	Message interface{} `json:"message"`
}
