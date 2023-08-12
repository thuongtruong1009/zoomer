package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/search/presenter"
	"github.com/thuongtruong1009/zoomer/internal/modules/search/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/validators"
	"net/http"
)

type searchHandler struct {
	searchUC usecase.UseCase
}

func NewSearchHandler(searchUC usecase.UseCase) *searchHandler {
	return &searchHandler{
		searchUC: searchUC,
	}
}

func (h *searchHandler) SearchRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &presenter.RoomSearchParams{}
		input.Name = c.QueryParam("name")
		input.Description = c.QueryParam("description")
		input.Category = c.QueryParam("category")
		if err := validators.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		rooms := h.searchUC.SearchRooms(c.Request().Context(), input)

		return c.JSON(http.StatusOK, rooms)
	}
}
