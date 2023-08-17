package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"net/http"
)

func (mwm *MiddlewareManager) RateLimit(pmt parameter.MiddlewareConf) {
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      pmt.RateLimit.Rate,
				Burst:     pmt.RateLimit.Burst,
				ExpiresIn: helpers.DurationSecond(pmt.RateLimit.ExpiresIn),
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(ctx echo.Context, err error) error {
			return mwm.inter.Error(ctx, http.StatusTooManyRequests, constants.ErrorForbidden, err)
		},
		DenyHandler: func(ctx echo.Context, identifier string, err error) error {
			return mwm.inter.Error(ctx, http.StatusTooManyRequests, constants.ErrorTooManyRequests, err)
		},
	}

	mwm.e.Use(middleware.RateLimiterWithConfig(config))
}
