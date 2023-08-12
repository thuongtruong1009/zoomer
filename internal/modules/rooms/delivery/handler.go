package delivery

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/repository"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/rooms/presenter"
	"github.com/thuongtruong1009/zoomer/internal/modules/rooms/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
	"github.com/thuongtruong1009/zoomer/pkg/validators"
	"net/http"
)

type roomHandler struct {
	roomUC usecase.UseCase
	inter  interceptor.IInterceptor
}

func NewRoomHandler(roomUC usecase.UseCase, inter interceptor.IInterceptor) *roomHandler {
	return &roomHandler{roomUC: roomUC, inter: inter}
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
			return rh.inter.Error(c, http.StatusNotFound, constants.ErrorNotFound, err)
		}
		return rh.inter.Data(c, http.StatusOK, mapRooms(rooms))
	}
}

func (rh *roomHandler) GetUserRooms() echo.HandlerFunc {
	return func(c echo.Context) error {
		rawId := c.Param(repository.CtxUserKey)
		userId, err := uuid.Parse(rawId)
		if err != nil {
			return rh.inter.Error(c, http.StatusBadRequest, constants.ErrorBadRequest, err)
		}
		rooms, err := rh.roomUC.GetRoomsByUserId(c.Request().Context(), userId.String())
		if err != nil {
			return rh.inter.Error(c, http.StatusInternalServerError, constants.ErrorInternalServer, err)
		}
		return rh.inter.Data(c, http.StatusOK, mapRooms(rooms))
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
