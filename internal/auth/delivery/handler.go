package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/auth/presenter"
	"github.com/thuongtruong1009/zoomer/internal/auth/usecase"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/validators"
	"net/http"
)

type authHandler struct {
	useCase usecase.UseCase
	inter   interceptor.IInterceptor
}

func NewAuthHandler(useCase usecase.UseCase, inter interceptor.IInterceptor) AuthHandler {
	return &authHandler{
		useCase: useCase,
		inter:   inter,
	}
}

// SignUp godoc
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
func (h *authHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.SignUpInput{}
		if err := validators.ReadRequest(c, input); err != nil {
			return h.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		req := &models.User{}

		if req.IsUsernameInvalid() {
			return h.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, constants.ErrUsernameInvalid)
		}

		if req.IsPasswordInvalid() {
			return h.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, constants.ErrPasswordInvalid)
		}

		user, err := h.useCase.SignUp(c.Request().Context(), input.Username, input.Password, input.Limit)
		if err != nil {
			return h.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return h.inter.Data(c, http.StatusCreated, user)
	}
}

// SignIn godoc
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
func (h *authHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.SignInInput{}
		if err := validators.ReadRequest(c, input); err != nil {
			return h.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		user, err := h.useCase.SignIn(c.Request().Context(), input.Username, input.Password)

		if err != nil {
			if err == constants.ErrUserNotFound {
				return h.inter.Error(c, http.StatusNotFound, constants.ErrUserNotFound, err)
			}
			if err == constants.ErrWrongPassword {
				return h.inter.Error(c, http.StatusNotFound, constants.ErrWrongPassword, err)
			}
			return h.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		usecase.WriteCookie(c, constants.CookieKey, user.Token, 60*60*24, "/", "localhost", false, true)

		return h.inter.Data(c, http.StatusOK, user)
	}
}

// SignOut godoc
//	@Summary		Logout user credentials
//	@Description	Logout user credentials
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			userId	path	int	true	"User ID"
//	@Success		200		string	constants.Success
//	@Failure		400		error	constants.ErrorBadRequest.Error()
//	@Failure		401		error	constants.ErrorUnauthorized.Error()
//	@Failure		500		error	constants.ErrorInternalServer.Error()
//	@Security		bearerAuth
//	@Router			/auth/signout [post]
func (h *authHandler) SignOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase.WriteCookie(c, constants.CookieKey, "", -1, "", "", false, true)
		return c.NoContent(http.StatusNoContent)
	}
}
