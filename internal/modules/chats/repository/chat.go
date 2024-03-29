package repository

import (
	"context"
	"encoding/json"
	// "fmt"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/chats/adapter"
	"log"
	"time"
)

type chatRepository struct {
	redisDB *redis.Client
}

func NewChatRepository(redisDB *redis.Client) *chatRepository {
	return &chatRepository{
		redisDB: redisDB,
	}
}

func (cr *chatRepository) UpdateContactList(ctx context.Context, username, contact string) error {
	zs := &redis.Z{Score: float64(time.Now().Unix()), Member: contact}

	err := cr.redisDB.ZAdd(context.Background(), adapter.ContactListZKey(username), zs).Err()
	if err != nil {
		log.Println("error while adding contact list. username: ", username, "contact: ", contact, err)
		return err
	}
	return nil
}

func (cr *chatRepository) CreateChat(ctx context.Context, c *models.Chat) (string, error) {
	chatKey := adapter.ChatKey()

	by, err := json.Marshal(c)
	if err != nil {
		log.Println("error while marshaling chat JSON", err)
		return "", err
	}

	res, err := cr.redisDB.Do(
		context.Background(),
		"json.set",
		chatKey,
		"$",
		string(by)).Result()

	if err != nil {
		log.Println("error while setting chat JSON in Redis", err)
		return "", err
	}

	log.Println("--> save chat successfully set", res)

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
