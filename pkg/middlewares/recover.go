package middlewares

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
)

func (mwm *MiddlewareManager) RecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				v, _ := err.(error)
				mwm.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, v)
				exceptions.Log(constants.ErrorInternalServer, v)
			}
		}()
		return next(c)
	}
}
