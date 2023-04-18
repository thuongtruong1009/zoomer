package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"zoomer/internal/auth/repository"
	"zoomer/pkg/constants"
)

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

		userId, err := mw.authUC.ParseToken(c.Request().Context(), headerParts[1])

		if err != nil {
			status := http.StatusInternalServerError
			if err == constants.ErrInvalidAccessToken {
				status = http.StatusUnauthorized
			}
			return echo.NewHTTPError(status)
		}

		c.Set(repository.CtxUserKey, userId)

		return next(c)
	}
}