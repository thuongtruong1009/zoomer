package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/chats/repository"
	"log"
	"time"
)

var (
	Clients   = make(map[*models.Client]bool)
	Broadcast = make(chan *models.Chat)
)

type Hub struct {
	chatRepo repository.ChatRepository
}

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
		// fmt.Println("host", client.Conn.RemoteAddr())
		if m.Type == "bootup" {
			client.Username = m.User
			// fmt.Println("Created succesfully client mapped", &client, client, client.Username)
		} else {
			fmt.Println("Received message", m.Type, m.Chat)
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

		fmt.Println("---> New message: ", message, " - from:", message.From, " - to:", message.To)

		// fmt.Println("Clients: ", Clients)

		for client := range Clients {
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
