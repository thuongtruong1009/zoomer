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

func NewChatHub(chatRepo repository.ChatRepository) IHub {
	return &Hub{
		chatRepo: chatRepo,
	}
}

func (h *Hub) Receiver(ctx context.Context, client *models.Client) {
	ctx := context.Background()
	deliveries, _ := RabbitMQAdapter(ctx, "test", []byte("Hello World!"))

	for delivery := range deliveries {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("Received a message: %s\n", string(delivery.Delivery.Body))
			delivery.Delivery.Ack(false)
		}
	}
}
