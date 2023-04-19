package middlewares

import (
	"strings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"zoomer/pkg/interceptor"
)

func HttpMiddleware(e *echo.Echo, inter interceptor.IInterceptor) {
	e.Use(HttpHeader)
	e.Use(LoggerMiddleware)
	e.Use(RecoveryMiddleware)
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Timeout())
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyLimit("2M"))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "auth")
		},
	}))

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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderContentLength, echo.HeaderAuthorization, echo.HeaderAccessControlAllowHeaders},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.HEAD},
	}))

	RateLimit(e, inter)
}
