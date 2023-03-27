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
	"zoomer/utils"
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

		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		err := rh.roomUC.CreateRoom(c.Request().Context(), fmt.Sprintf("%v", userId), input.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, nil)
	}4
}

//sync to redis
func (rh *roomHandler) ChatHistoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		u1 := c.Query("u1")
		u2 := c.Query("u2")

		fromTS, toTS := "0" "+inf"

		if c.Query("from-ts") != ""  && c.Query("to-ts") != "" {
			fromTS = c.Query("from-ts")
			toTS = c.Query("to-ts")
		}

		chat := rh.roomUC.GetChatHistory(c.Request().Context(), u1, u2, fromTS, toTS)

		return c.JSON(http.StatusOK, chat)
	}
}
