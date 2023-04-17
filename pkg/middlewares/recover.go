package middlewares

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"zoomer/constants"
	"zoomer/internal/base/interceptor"
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
