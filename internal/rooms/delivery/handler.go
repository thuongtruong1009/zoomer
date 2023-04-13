package delivery

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"zoomer/internal/auth/repository"
	"zoomer/internal/models"
	"zoomer/internal/rooms/presenter"
	"zoomer/internal/rooms/usecase"
	"zoomer/validators"
)

type roomHandler struct {
	roomUC usecase.UseCase
}

func NewRoomHandler(roomUC usecase.UseCase) *roomHandler {
	return &roomHandler{roomUC: roomUC}
}

func mapRoom(r *models.Room) *presenter.RoomResponse {
	return &presenter.RoomResponse{
		Id:        r.Id,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		CreatedBy: r.CreatedBy,
	}
}

func mapRooms(ro []*models.Room) []*presenter.RoomResponse {
	out := make([]*presenter.RoomResponse, len(ro))

	for i, b := range ro {
		out[i] = mapRoom(b)
	}
	return out
}

func (rh *roomHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		rooms, err := rh.roomUC.GetAllRooms(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, mapRooms(rooms))
	}
}

func (rh *roomHandler) GetUserRooms() echo.HandlerFunc {
	return func(c echo.Context) error {
		rawId := c.Param(repository.CtxUserKey)
		userId, err := uuid.Parse(rawId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		rooms, err := rh.roomUC.GetRoomsByUserId(c.Request().Context(), userId.String())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, mapRooms(rooms))
	}
}

func (rh *roomHandler) AddRoom() echo.HandlerFunc {
	return func(c echo.Context) error {

		userId := c.Get(repository.CtxUserKey)
		input := &presenter.RoomRequest{}

		if err := validators.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		err := rh.roomUC.CreateRoom(c.Request().Context(), fmt.Sprintf("%v", userId), input.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

// sync to redis
func (rh *roomHandler) ChatHistoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		u1 := c.QueryParam("u1")
		u2 := c.QueryParam("u2")

		fromTS, toTS := "0", "+inf"

		if c.QueryParam("from-ts") != "" && c.QueryParam("to-ts") != "" {
			fromTS = c.QueryParam("from-ts")
			toTS = c.QueryParam("to-ts")
		}

		chat := rh.roomUC.GetChatHistory(c.Request().Context(), u1, u2, fromTS, toTS)

		return c.JSON(http.StatusOK, chat)
	}
}

func (rh *roomHandler) ContactListHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.QueryParam("username")
		res := rh.roomUC.ContactList(c.Request().Context(), u)

		return c.JSON(http.StatusOK, res)
	}
}

func (rh *roomHandler) CreateFetchChatBetweenIndexHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		rh.roomUC.GetFetchChatBetweenIndex(c.Request().Context())
		return c.JSON(http.StatusOK, nil)
	}
}
