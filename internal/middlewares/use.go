package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strings"
	"zoomer/configs"
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
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "auth")
		},
	}))
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().Host, "localhost") {
				return true
			}
			return false
		},
	}))

	e.Pre(middleware.HTTPSNonWWWRedirect())
	e.Pre(middleware.HTTPSRedirect())

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))
	CheckOrigin(e)

	configs.ProxyConfig(e)
	configs.RateLimit(e)

	e.Use(middleware.Timeout())
}
