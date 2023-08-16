package delivery

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/chats/hub"
	"net/http"
)

type chatHandler struct {
	hub hub.IHub
}

func NewChatHandler(hub hub.IHub) ChatHandler {
	return &chatHandler{
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (ch *chatHandler) ChatConnect() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Request().Host, c.Request().RemoteAddr)

		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			c.Logger().Error(err)
		}

		client := &models.Client{Conn: ws}

		hub.Clients[client] = true

		// fmt.Println("clients", len(hub.Clients), hub.Clients, ws.RemoteAddr())

		ch.hub.Receiver(c.Request().Context(), client)

		// fmt.Println("existing", ws.RemoteAddr().String())
		defer delete(hub.Clients, client)

		return c.JSON(http.StatusOK, nil)
	}
}
