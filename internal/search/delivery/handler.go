package delivery

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"

	"zoomer/internal/search/presenter"
	"zoomer/internal/search/usecase"
	"zoomer/internal/search/views"
)

type searchHandler struct {
	usecase usecase.UseCase
}

func NewSearchHandler(useCase usecase.UseCase) Handler {
	return &searchHandler{
		useCase: useCase,
	}
}

func (h *searchHandler) SearchRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.SearchRoomInput{}
		input.Name = c.QueryParam("name")
		input.Description = c.QueryParam("description")
		input.Category = c.QueryParam("category")
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		rooms, err := h.useCase.SearchRoom(c.Request().Context(), input)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, presenter.SearchResponse{Rooms: rooms})
	}
}
