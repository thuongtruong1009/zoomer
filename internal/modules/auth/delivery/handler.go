package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/validators"
	"net/http"
	"strings"
	"time"
)

type authHandler struct {
	useCase  usecase.UseCase
	inter    interceptor.IInterceptor
	paramCfg *parameter.ParameterConfig
}

func NewAuthHandler(useCase usecase.UseCase, inter interceptor.IInterceptor, paramCfg *parameter.ParameterConfig) AuthHandler {
	return &authHandler{
		useCase:  useCase,
		inter:    inter,
		paramCfg: paramCfg,
	}
}

// SignUp godoc
//
//	@Summary		Register account
//	@Description	Create a new account for new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		presenter.SignUpInput	true	"Create User"
//	@Success		201		{object}	presenter.SignUpResponse
//	@Failure		400		error		constants.ErrorBadRequest
//	@Failure		500		error		constants.ErrorInternalServer
//	@Router			/auth/signup [post]
func (ah *authHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.SignUpInput{}
		if err := validators.ReadRequest(c, input); err != nil {
			return ah.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		err1 := input.IsUsernameValid()
		if err1 != nil {
			return ah.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err1)
		}

		err2 := input.IsPasswordValid()
		if err2 != nil {
			return ah.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err2)
		}

		input.IsLimitValid()

		user, err := ah.useCase.SignUp(c.Request().Context(), strings.ToLower(input.Username), input.Password, input.Limit)
		if err != nil {
			return ah.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return ah.inter.Data(c, http.StatusCreated, user)
	}
}

// SignIn godoc
//
//	@Summary		Login to account
//	@Description	Login to user account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		presenter.SignInInput	true	"Login user"
//	@Success		200		{object}	presenter.SignInResponse
//	@Failure		400		error		constants.ErrorBadRequest
//	@Failure		500		error		constants.ErrorInternalServer
//	@Router			/auth/signin [post]
func (ah *authHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.SignInInput{}
		if err := validators.ReadRequest(c, input); err != nil {
			return ah.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		user, err := ah.useCase.SignIn(c.Request().Context(), input.Username, input.Password)

		if err != nil {
			if err == constants.ErrUserNotFound {
				return ah.inter.Error(c, http.StatusNotFound, constants.ErrUserNotFound, err)
			}
			if err == constants.ErrWrongPassword {
				return ah.inter.Error(c, http.StatusNotFound, constants.ErrWrongPassword, err)
			}
			return ah.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}
		newCookie := &presenter.SetCookie{
			Name: constants.AccessTokenKey,
			Value: user.Token,
			Expires: helpers.DurationSecond(ah.paramCfg.TokenTimeout),
		}
		ah.writeCookie(c, newCookie)

		return ah.inter.Data(c, http.StatusOK, user)
	}
}

// SignOut godoc
//
//	@Summary		Logout user
//	@Description	Logout user with credentials or token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			userId	path	string	true	"User ID"
//	@Success		200		string	constants.Success
//	@Failure		400		error	constants.ErrorBadRequest.Error()
//	@Failure		401		error	constants.ErrorUnauthorized.Error()
//	@Failure		500		error	constants.ErrorInternalServer.Error()
//	@Security		bearerAuth
//	@Router			/auth/signout [post]
func (ah *authHandler) SignOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		newCookie := &presenter.SetCookie{
			Name: constants.AccessTokenKey,
			Value: "",
			Expires: -1,
		}
		ah.writeCookie(c, newCookie)
		return c.NoContent(http.StatusNoContent)
	}
}

func (ah *authHandler) writeCookie(c echo.Context, cookie *presenter.SetCookie) {
	newCookie := &http.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		Expires:  time.Now().Add(cookie.Expires),
		Path:     ah.paramCfg.CookiePath,
		Domain:   ah.paramCfg.CookieDomain,
		Secure:   (ah.paramCfg.CookieSecure == true),
		HttpOnly: (ah.paramCfg.CookieHttpOnly == true),
	}

	if cookie.Value == "" {
		newCookie.Path = ""
		newCookie.Domain = ""
	}

	c.SetCookie(newCookie)
}


// ResetPassword godoc
//
//	@Summary		Reset password
//	@Description	Reset or update change password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		presenter.SignInInput	true	"Reset password"
//	@Success		200		{object}	presenter.SignInResponse
//	@Failure		400		error		constants.ErrorBadRequest
//	@Failure		500		error		constants.ErrorInternalServer
//	@Router			/auth/reset-password [patch]
func (ah *authHandler) ResetPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		newEmail := &presenter.ResetPassword{}

		if err := validators.ReadRequest(c, newEmail); err != nil {
			return ah.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		return ah.useCase.ResetPassword(c.Request().Context(), newEmail)
	}
}
