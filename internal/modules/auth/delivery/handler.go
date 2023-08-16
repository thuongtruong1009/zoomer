package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/validators"
	"net/http"
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
//	@Summary		Create a new user
//	@Description	Create a new user
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

		req := &models.User{}

		if req.IsUsernameInvalid() {
			return ah.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, constants.ErrUsernameInvalid)
		}

		if req.IsPasswordInvalid() {
			return ah.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, constants.ErrPasswordInvalid)
		}

		user, err := ah.useCase.SignUp(c.Request().Context(), input.Username, input.Password, input.Limit)
		if err != nil {
			return ah.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return ah.inter.Data(c, http.StatusCreated, user)
	}
}

// SignIn godoc
//
//	@Summary		Login to user account
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

		usecase.WriteCookie(c, constants.AccessTokenKey, user.Token, helpers.DurationSecond(ah.paramCfg.TokenTimeout), ah.paramCfg.CookiePath, ah.paramCfg.CookieDomain, ah.paramCfg.CookieSecure == true, ah.paramCfg.CookieHttpOnly == true)

		return ah.inter.Data(c, http.StatusOK, user)
	}
}

// SignOut godoc
//
//	@Summary		Logout user credentials
//	@Description	Logout user credentials
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
		usecase.WriteCookie(c, constants.AccessTokenKey, "", -1, "", "", ah.paramCfg.CookieSecure == true, ah.paramCfg.CookieHttpOnly == true)
		return c.NoContent(http.StatusNoContent)
	}
}
