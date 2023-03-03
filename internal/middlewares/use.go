package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Zoomer/1.0")
		return next(c)
	}
}

func HttpMiddleware(e *echo.Echo) {
	e.Use(serverHeader)
	e.Use(LoggerMiddleware)

	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().Host, "localhost") {
				return true
			}
			return false
		},
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3001", "http://localhost:3002"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderContentLength, echo.HeaderAuthorization},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
}
