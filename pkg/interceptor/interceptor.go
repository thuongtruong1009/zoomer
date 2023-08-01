package interceptor

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
)

type interceptor struct{}

func NewInterceptor() IInterceptor {
	return &interceptor{}
}

func (i *interceptor) Data(c echo.Context, code int, data interface{}) error {
	props := &InterceptorSuccessProps{
		Status: code,
		Data:   data,
	}

	return c.JSON(code, props)
}

func (i *interceptor) Error(c echo.Context, code int, msg error, err error) error {
	props := &InterceptorErrorProps{
		Status:  code,
		Message: msg.Error(),
		Error:   err.Error(),
	}

	exceptions.Log(msg, err)

	return echo.NewHTTPError(code, props)
}
