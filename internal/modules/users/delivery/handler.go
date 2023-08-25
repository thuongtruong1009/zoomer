package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/users/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"net/http"
	"strconv"
	"github.com/thuongtruong1009/zoomer/pkg/abstract"
)

type userHandler struct {
	usecase usecase.IUserUseCase
	inter   interceptor.IInterceptor
}

func NewUserHandler(usecase usecase.IUserUseCase, inter interceptor.IInterceptor) IUserHandler {
	return &userHandler{
		usecase: usecase,
		inter:   inter,
	}
}

// GetUserByIdOrName godoc
//
//	@Summary		Get user by id or name
//	@Description	Get user by id or name
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			idOrName	path		string	true	"id or name of user"
//	@Success		200			{object}	presenter.GetUserByIdOrNameResponse
//	@Failure		401			error		constants.ErrorBadRequest
//	@Failure		500			error		constants.ErrorInternalServer
//	@Router			/users/{idOrName} [get]
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

// SearchUser godoc
//
//	@Summary		Search user by name
//	@Description	Search user by name
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"name of user"
// @Param name query string false "name of user"
// @Param page query string false "page number"
// @Param size query string false "number of elements"
//	@Success		200			{object}	presenter.GetUserByIdOrNameResponse
//	@Failure		401			error		constants.ErrorBadRequest
//	@Failure		500			error		constants.ErrorInternalServer
//	@Router			/users/search [get]
func (h *userHandler) SearchUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return h.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}
		size, err := strconv.Atoi(c.QueryParam("size"))
		if err != nil {
			return h.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}

		pq := abstract.NewPagination(size, page, "")

		user, err := h.usecase.SearchUser(c.Request().Context(), c.QueryParam("name"), pq)
		if err != nil {
			return h.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}

		return h.inter.Data(c, http.StatusOK, user)
	}
}
