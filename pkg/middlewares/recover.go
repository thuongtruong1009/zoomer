package middlewares

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
)

func RecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		base := interceptor.NewInterceptor()

		defer func() {
			if err := recover(); err != nil {
				v, _ := err.(error)
				base.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, v)
			}
		}()
		return next(c)
	}
}
