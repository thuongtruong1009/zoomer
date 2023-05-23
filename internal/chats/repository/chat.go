package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"zoomer/db"
	"zoomer/internal/chats/adapter"
	"zoomer/internal/models"
)

type chatRepository struct{}

func NewChatRepository() *chatRepository {
	return &chatRepository{}
}

func (cr *chatRepository) UpdateContactList(ctx context.Context, username, contact string) error {
	zs := &redis.Z{Score: float64(time.Now().Unix()), Member: contact}

	err := db.GetRedisInstance().ZAdd(context.Background(), adapter.ContactListZKey(username), zs).Err()
	if err != nil {
		log.Println("error while adding contact list. username: ", username, "contact: ", contact, err)
		return err
	}
	return nil
}

func (cr *chatRepository) CreateChat(ctx context.Context, c *models.Chat) (string, error) {
	chatKey := adapter.ChatKey()
	fmt.Println("chat key", chatKey)

	by, err := json.Marshal(c)
	if err != nil {
		log.Println("error while marshaling chat JSON", err)
		return "", err
	}

	// Store chat JSON using HSET command
	err = db.GetRedisInstance().HSet(context.Background(), chatKey, "$", string(by)).Err()
	if err != nil {
		log.Println("error while setting chat JSON in Redis", err)
		return "", err
	}

	log.Println("chat successfully set")

	err = cr.UpdateContactList(ctx, c.From, c.To)
	if err != nil {
		log.Println("error while updating contact list of", c.From)
	}

	err = cr.UpdateContactList(ctx, c.To, c.From)
	if err != nil {
		log.Println("error while updating contact list of", c.To)
	}

	return chatKey, nil
}
