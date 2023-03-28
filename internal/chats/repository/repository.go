package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"context"
	"time"
	"github.com/go-redis/redis/v8"
	"zoomer/db"
	"zoomer/internal/models"
	"zoomer/internal/chats/adapter"
)

func UpdateContactList(username, contact string) error {
	zs := &redis.Z{Score: float64(time.Now().Unix()), Member: contact}

	err := db.RedisClient.ZAdd(context.Background(), adapter.ContactListZKey(username), zs).Err()
	if err != nil {
		log.Println("error while adding contact list. username: ", username, "contact: ", contact, err)
		return err
	}
	return nil
}

func CreateChat(c *models.Chat) (string, error) {
	chatKey := adapter.ChatKey()
	fmt.Println("chat key", chatKey)

	by, _ := json.Marshal(c)

	res, err := db.RedisClient.Do(context.Background(), "JSON.SET", chatKey, "$", string(by)).Result()

	if err != nil {
		log.Println("error while creating setting chat json", err)
		return "", err
	}
	log.Println("chat succesfully set", res)

	err = UpdateContactList(c.From, c.To)
	if err != nil {
		log.Println("error while updating contact list of", c.From)
	}

	err = UpdateContactList(c.To, c.From)
	if err != nil {
		log.Println("error while updating contact list of", c.To)
	}

	return chatKey, nil
}