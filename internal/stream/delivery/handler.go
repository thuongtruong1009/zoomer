package delivery

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"zoomer/internal/stream/hub"
	"zoomer/internal/models"
	"zoomer/internal/stream/presenter"
	"zoomer/pkg/interceptor"
	"zoomer/pkg/constants"
)

type streamHandler struct {
	hub hub.IHub
	inter interceptor.IInterceptor
}

func Init(){
	hub.Mapper.Map =  make(map[string][]*models.Participant)
}

func NewStreamHandler(hub hub.IHub, inter interceptor.IInterceptor) StreamHandler {
	return &streamHandler{
		hub: hub,
		inter: inter,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (sh *streamHandler) CreateStream() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Request().Host, c.Request().RemoteAddr)

		roomID := sh.hub.CreateStream(c.Request().Context())

		return sh.inter.Data(c, http.StatusOK, presenter.StreamResponse{
			RoomID: roomID})
	}
}

func (sh *streamHandler) JoinStream() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomID := c.QueryParam("roomID")

		if roomID == "" {
			return sh.inter.Error(c, http.StatusBadRequest, constants.ErrStreamIDMissing, nil)
		}

		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			c.Logger().Error(err)
		}

		client := &models.Participant{
			Conn: ws,
			Host: false,
		}

		sh.hub.InsertIntoStream(c.Request().Context(), roomID, client)

		sh.hub.Receiver(c.Request().Context(), roomID, client)

		defer delete(hub.Mapper.Map, client.Conn.RemoteAddr().String())

		return sh.inter.Data(c, http.StatusOK, nil)
	}
}
