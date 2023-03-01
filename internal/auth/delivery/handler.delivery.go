package delivery

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"zoomer/utils"
	"zoomer/internal/auth"
	"zoomer/internal/auth/usecase"
	"zoomer/internal/auth/presenter"
)

type authHandler struct {
	useCase usecase.UseCase
}

func NewAuthHandler(useCase usecase.UseCase) Handler {
	return &authHandler {
		useCase: useCase,
	}
}

func (h *authHandler) SignUp() echo.HandlerFunc{
	return func(c echo.Context) error {
		input := &presenter.SignUpInput{}
		if err := utils.ReadRequest(c, input); err != nil {
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
	return func (c echo.Context) error {
		input := &presenter.LoginInput{}
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		token, err := h.useCase.SignIn(c.Request().Context(), input.Username, input.Password)

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

		return c.JSON(http.StatusOK, presenter.LogInResponse{Token: token})
	}
}

func (h *authHandler) SignOut() echo.HandlerFunc {
	return func (c echo.Context) error {
		usecase.WriteCookie(c, "jwt", "", -1, "", "", false, true)
		return c.NoContent(http.StatusNoContent)
	}
}