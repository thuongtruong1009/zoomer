package delivery

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"zoomer/internal/chats/adapter"
	"zoomer/internal/chats/hub"
	"zoomer/internal/models"
)

type chatHandler struct {
	hub hub.IHub
}

func NewChatHandler(hub hub.IHub) ChatHandler {
	return &chatHandler{
		hub: hub,
	}
}

func (ch *chatHandler) ChatConnect() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Request().Host, c.Request().RemoteAddr)

		ws := adapter.HubUpgrader(c.Response(), c.Request())

		client := &models.Client{Conn: ws}

		hub.Clients[client] = true

		fmt.Println("clients", len(hub.Clients), hub.Clients, ws.RemoteAddr())

		// Receiver(client)
		go ch.hub.Broadcaster(c.Request().Context())
		go ch.hub.Receiver(c.Request().Context(), client)

		fmt.Println("existing", ws.RemoteAddr().String())
		delete(hub.Clients, client)

		return c.JSON(http.StatusOK, nil)
	}
}
