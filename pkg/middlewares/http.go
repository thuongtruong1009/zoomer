package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strings"
)

func (mwm *MiddlewareManager) HttpMiddleware() {
	mwm.e.Logger.SetLevel(log.INFO)

	mwm.e.Use(mwm.HttpHeader)
	mwm.e.Use(mwm.RequestMiddleware)
	mwm.e.Use(mwm.RecoveryMiddleware)

	mwm.e.Use(middleware.Secure())
	mwm.e.Use(middleware.Logger())
	mwm.e.Use(middleware.Timeout())
	mwm.e.Use(middleware.RequestID())
	mwm.e.Use(middleware.BodyLimit("2M"))
	mwm.e.Use(middleware.Recover())

	mwm.e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			skips := [2]string{"auth", "docs"}
			for _, skip := range skips {
				if strings.Contains(c.Request().URL.Path, skip) {
					return true
				}
			}
			return false
		},
	}))

	mwm.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().Host, "localhost") {
				return true
			}
			return false
		},
	}))

	mwm.e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		LogLevel:          log.ERROR,
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogErrorFunc:      nil,
	}))

	mwm.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderContentLength, echo.HeaderAuthorization, echo.HeaderAccessControlAllowHeaders},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.HEAD},
	}))

	mwm.RateLimit(mwm.paramCfg.RateLimitConf)
}
