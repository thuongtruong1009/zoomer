package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/repository"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	authUC usecase.UseCase
	inter  interceptor.IInterceptor
}

func NewAuthMiddleware(authUC usecase.UseCase, inter interceptor.IInterceptor) *AuthMiddleware {
	return &AuthMiddleware{authUC, inter}
}

func (mw *AuthMiddleware) JWTValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(constants.BearerHeader)
		if authHeader == "" {
			return mw.inter.Error(c, http.StatusUnauthorized, constants.ErrorUnauthorized, constants.ErrInvalidAccessToken)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return mw.inter.Error(c, http.StatusUnauthorized, constants.ErrorUnauthorized, constants.ErrInvalidAccessToken)
		}

		if headerParts[0] != constants.BearerName {
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

func (mw *AuthMiddleware) CookieValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(constants.AccessTokenKey)
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
