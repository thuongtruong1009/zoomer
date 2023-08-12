package middlewares

import (
	"github.com/labstack/echo/v4"
)

func (mwm *MiddlewareManager) HttpHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Zoomer/1.1")
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set("X-XSS-Protection", "1; mode=block")
		c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Response().Header().Set("Access-Control-Max-Age", "3600")
		return next(c)
	}
}

func (mwm *MiddlewareManager) WsHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type, Authorization")
		c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, "GET, POST, PUT, DELETE")
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		c.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
		return next(c)
	}
}
