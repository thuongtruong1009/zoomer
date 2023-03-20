package delivery

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"zoomer/utils"
	"zoomer/internal/search/presenter"
	"zoomer/internal/search/usecase"
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
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		rooms := h.searchUC.SearchRooms(c.Request().Context(), input)

		return c.JSON(http.StatusOK, rooms)
	}
}
