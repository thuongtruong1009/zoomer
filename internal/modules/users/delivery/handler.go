package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/users/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"net/http"
)

type userHandler struct {
	usecase usecase.IUserUseCase
	inter interceptor.IInterceptor
}

func NewUserHandler(usecase usecase.IUserUseCase, inter interceptor.IInterceptor) IUserHandler {
	return &userHandler{
		usecase: usecase,
		inter: inter,
	}
}

func (h *userHandler) GetUserByIdOrName() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := c.Param("idOrName")

		user, err := h.usecase.GetUserByIdOrName(c.Request().Context(), input)
		if err != nil {
			return h.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return h.inter.Data(c, http.StatusOK, user)
	}
}
