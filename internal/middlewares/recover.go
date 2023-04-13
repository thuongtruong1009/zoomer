package middlewares

import (
	"errors"
	"github.com/labstack/echo/v4"
)

func RecoverMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				// c.Logger().Error(err)
				friendlyErrorToClient := errors.New("Error occured in our own server. Sorry")
				if v, ok := err.(error); ok {
					base.Error(c, http.StatusInternalServerError,
						friendlyErrorToClient, v.Error())
				} else {
					base.Error(c, http.StatusInternalServerError,
						friendlyErrorToClient, err.(string))
				}
			}
		}()
		return next(c)
	}
}
