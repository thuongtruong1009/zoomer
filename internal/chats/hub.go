package chats

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Rooms[client.RoomID]; ok {
				r := h.Rooms[client.RoomID]

				if _, ok := r.Clients[client.ID]; !ok {
					r.Clients[client.ID] = client
				}
			}
		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.RoomID]; ok {
				if _, ok := h.Rooms[client.RoomID].Clients[client.ID]; ok {
					if len(h.Rooms[client.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   client.RoomID,
							Username: client.Username,
						}
					}

					delete(h.Rooms[client.RoomID].Clients, client.ID)
					close(client.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, client := range h.Rooms[m.RoomID].Clients {
					client.Message <- m
				}
			}
		}
	}
}
