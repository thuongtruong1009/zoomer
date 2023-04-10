package middlewares

import (
	"net/http"
	"strings"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"zoomer/internal/auth"
	"zoomer/internal/auth/repository"
)

var caching = cache.New(5*time.Minute, 10*time.Minute)

func (mw *MiddlewareManager) JWTValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		if headerParts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		//check if token is present in the cache
		if _, found := caching.Get(headerParts[1]); found {
			return next(c)
		}

		userId, err := mw.authUC.ParseToken(c.Request().Context(), headerParts[1])

		if err != nil {
			status := http.StatusInternalServerError
			if err == auth.ErrInvalidAccessToken {
				status = http.StatusUnauthorized
			}
			return echo.NewHTTPError(status)
		}

		c.Set(repository.CtxUserKey, userId)

		//cache token
		caching.Set(headerParts[1], userId, cache.DefaultExpiration)

		return next(c)
	}
}
