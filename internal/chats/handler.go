package chats

import (
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	hub *Hub
}

func NewChatHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func (h *Handler) CreateRoom() echo.HandlerFunc {
	return func(c echo.Context) error{
		var req CreateRoomReq
		// if err := c.ShouldBindJSON(&req); err != nil {
		// 		c.Logger().Error(err)
		// }

		h.hub.Rooms[req.ID] = &Room{
			ID: req.ID,
			Name: req.Name,
			Clients: make(map[string]*Client),
		}

		return c.JSON(http.StatusCreated, req)
	}
}

func (h *Handler) JoinRoom() echo.HandlerFunc  {
	return func(c echo.Context) error {
		conn, err := upgrader.Upgrade(c.Response(), c.Request() ,nil)
		if err != nil {
			c.Logger().Error(err)
		}

		roomID := c.Param("roomId")
		clientID := c.FormValue("userId")
		username := c.FormValue("username")

		client := &Client{
			Conn: conn,
			Message: make(chan *Message, 10),
			ID: clientID,
			RoomID: roomID,
			Username: username,
		}

		m := &Message{
			Content: "A new user has joined the room",
			RoomID: roomID,
			Username: username,
		}

		h.hub.Register <- client
		h.hub.Broadcast <- m

		go client.writeMessage()
		client.readMessage(h.hub)

		return c.JSON(http.StatusOK, nil)
	}
}

type RoomRes struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms() echo.HandlerFunc {
	return func(c echo.Context) error {
		rooms := make([]RoomRes, 0)

		for _, r := range h.hub.Rooms {
			rooms = append(rooms, RoomRes{
				ID: r.ID,
				Name: r.Name,
			})
		}

		return c.JSON(http.StatusOK, rooms)
	}
}

type ClientRes struct {
	ID string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients() echo.HandlerFunc {
	return func(c echo.Context) error {
		var clients []ClientRes
		roomId := c.Param("roomId")

		if _, ok := h.hub.Rooms[roomId]; !ok {
			clients = make([]ClientRes, 0)
			c.JSON(http.StatusOK, clients)
		}

		for _, c := range h.hub.Rooms[roomId].Clients{
			clients = append(clients, ClientRes{
				ID: c.ID,
				Username: c.Username,
			})
		}

		return c.JSON(http.StatusOK, clients)
	}
}
