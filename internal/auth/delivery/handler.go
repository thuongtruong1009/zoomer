package delivery

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"zoomer/pkg/interceptor"
	"zoomer/validators"
	"zoomer/internal/models"
	"zoomer/pkg/constants"
	"zoomer/internal/auth/presenter"
	"zoomer/internal/auth/usecase"
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

		return h.inter.Data(c, http.StatusCreated, presenter.SignUpResponse{Id: user.Id, Username: user.Username, Limit: user.Limit})
	}
}

func (h *authHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.LoginInput{}
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

		return h.inter.Data(c, http.StatusOK, presenter.LogInResponse{UserId: user.UserId, Username: user.Username, Token: user.Token})
	}
}

func (h *authHandler) SignOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase.WriteCookie(c, constants.CookieKey, "", -1, "", "", false, true)
		return c.NoContent(http.StatusNoContent)
	}
}
