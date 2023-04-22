package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
	"zoomer/pkg/constants"
	"zoomer/pkg/interceptor"
)

func RateLimit(e *echo.Echo, inter interceptor.IInterceptor) {
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(ctx echo.Context, err error) error {
			return inter.Error(ctx, http.StatusTooManyRequests, constants.ErrorForbidden, err)
			// return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(ctx echo.Context, identifier string, err error) error {
			return inter.Error(ctx, http.StatusTooManyRequests, constants.ErrorTooManyRequests, err)
			// return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	e.Use(middleware.RateLimiterWithConfig(config))
}
