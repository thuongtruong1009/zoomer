package interceptor

import (
	"github.com/labstack/echo/v4"
)

type interceptor struct {}

func NewInterceptor() IInterceptor {
	return &interceptor{}
}

func (i *interceptor) Data(c echo.Context, code int, data interface{}) error{
	props := &InterceptorProps{
		Data: data,
		Message: "",
	}

	return c.JSON(code, props)
}

func (i *interceptor) Error(c echo.Context, code int, msg string, err error) error{
	props := &InterceptorProps{
		Data: nil,
		Message: msg + " "+ err.Error(),
	}

	return echo.NewHTTPError(code, props)
}
