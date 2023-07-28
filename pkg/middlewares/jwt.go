package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/auth/repository"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"net/http"
	"strings"
)

func (mw *MiddlewareManager) JWTValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return mw.inter.Error(c, http.StatusUnauthorized, constants.ErrorUnauthorized, constants.ErrInvalidAccessToken)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return mw.inter.Error(c, http.StatusUnauthorized, constants.ErrorUnauthorized, constants.ErrInvalidAccessToken)
		}

		if headerParts[0] != "Bearer" {
			return mw.inter.Error(c, http.StatusUnauthorized, constants.ErrorUnauthorized, constants.ErrInvalidAccessToken)
		}

		userId, err := mw.authUC.ParseToken(c.Request().Context(), headerParts[1])

		if err != nil {
			status := http.StatusInternalServerError
			if err == constants.ErrInvalidAccessToken {
				status = http.StatusUnauthorized
			}
			return mw.inter.Error(c, status, constants.ErrorInternalServer, err)
		}

		c.Set(repository.CtxUserKey, userId)

		return next(c)
	}
}

func (mw *MiddlewareManager) CookieValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(constants.CookieKey)
		if err != nil {
			return mw.inter.Error(c, http.StatusUnauthorized, constants.ErrorUnauthorized, constants.ErrInvalidAccessToken)
		}

		userId, err := mw.authUC.ParseToken(c.Request().Context(), cookie.Value)

		if err != nil {
			status := http.StatusInternalServerError
			if err == constants.ErrInvalidAccessToken {
				status = http.StatusUnauthorized
			}
			return mw.inter.Error(c, status, constants.ErrorInternalServer, err)
		}

		c.Set(repository.CtxUserKey, userId)

		return next(c)
	}
}
