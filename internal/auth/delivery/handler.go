package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zoomer/internal/auth"
	"zoomer/internal/auth/presenter"
	"zoomer/internal/auth/usecase"
	"zoomer/validators"
)

type authHandler struct {
	useCase usecase.UseCase
}

func NewAuthHandler(useCase usecase.UseCase) Handler {
	return &authHandler{
		useCase: useCase,
	}
}

func (h *authHandler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.SignUpInput{}
		if err := validators.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		user, err := h.useCase.SignUp(c.Request().Context(), input.Username, input.Password, input.Limit)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, presenter.SignUpResponse{Id: user.Id, Username: user.Username, Limit: user.Limit})
	}
}

func (h *authHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.LoginInput{}
		if err := validators.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		userId, username, token, err := h.useCase.SignIn(c.Request().Context(), input.Username, input.Password)

		if err != nil {
			if err == auth.ErrUserNotFound {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			if err == auth.ErrWrongPassword {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		usecase.WriteCookie(c, "jwt", token, 60*60*24, "/", "localhost", false, true)

		return c.JSON(http.StatusOK, presenter.LogInResponse{UserId: userId, Username: username, Token: token})
	}
}

func (h *authHandler) SignOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		usecase.WriteCookie(c, "jwt", "", -1, "", "", false, true)
		return c.NoContent(http.StatusNoContent)
	}
}
