package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"zoomer/internal/chats/repository"
	"zoomer/internal/models"
)

var (
	Clients   = make(map[*models.Client]bool)
	Broadcast = make(chan *models.Chat)
)

type Hub struct {
	chatRepo repository.ChatRepository
}

// func NewHub(chatRepo repository.ChatRepository) *Hub {
// 	return &Hub{
// 		Clients: make(map[*models.Client]bool),
// 		Broadcast: make(chan *models.Chat),
// 		chatRepo: chatRepo,
// 	}
// }

func NewChatHub(chatRepo repository.ChatRepository) IHub {
	return &Hub{
		chatRepo: chatRepo,
	}
}

func (h *Hub) Receiver(ctx context.Context, client *models.Client) {
	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		m := &models.Message{}

		err = json.Unmarshal(p, m)
		if err != nil {
			log.Println("error while unmarshaling chat", err)
			continue
		}
		fmt.Println("host", client.Conn.RemoteAddr())
		if m.Type == "bootup" {
			client.Username = m.User
			fmt.Println("client succesfully mapped", &client, client, client.Username)
		} else {
			fmt.Println("received message", m.Type, m.Chat)
			c := m.Chat
			c.Timestamp = time.Now().Unix()

			id, err := h.chatRepo.CreateChat(ctx, &c)
			if err != nil {
				log.Println("error while saving chat in redis", err)
				return
			}

			c.ID = id
			Broadcast <- &c
		}
	}
}

func (h *Hub) Broadcaster() {
	for {
		message := <-Broadcast

		for client := range Clients {
			fmt.Println("username:", client.Username, " -  message: ", message, " - from:", message.From, " - to:", message.To)

			if client.Username == message.From || client.Username == message.To {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					log.Printf("websocket error: %s", err)
					client.Conn.Close()
					delete(Clients, client)
				}
			}
		}
	}
}
