package middlewares

import (
	"strings"
	"zoomer/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Zoomer/1.0")
		c.Response().Header().Set("X-Frame-Options", "SAMEORIGIN")
		c.Response().Header().Set("X-XSS-Protection", "1; mode=block")
		c.Response().Header().Set("X-Content-Type-Options", "nosniff")
		c.Response().Header().Set("Referrer-Policy", "same-origin")
		c.Response().Header().Set("Content-Security-Policy", "default-src 'self'")
		c.Response().Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		c.Response().Header().Set("X-Download-Options", "noopen")
		c.Response().Header().Set("X-Content-Security-Policy", "default-src 'self'")
		c.Response().Header().Set("X-WebKit-CSP", "default-src 'self'")
		c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Response().Header().Set("X-Content-Type-Options", "nosniff")
		c.Response().Header().Set("Access-Control-Max-Age", "3600")

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

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	HttpCORS(e)

	// configs.ProxyConfig(e)
	configs.RateLimit(e)

	e.Use(middleware.Timeout())
}
